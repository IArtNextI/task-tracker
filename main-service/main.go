package main

import (
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	stat_service "main-service/stat-service"
	"main-service/task_service"
	"net/http"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type registerRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type userData struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Birthday    string `json:"birthday"`
	Email       string `json:"email"`
	Mobilephone string `json:"mobilephone"`
}

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func GetRegisterHandler(db Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request registerRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		val, err := db.Exists(request.Login)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		if val != 0 {
			c.Status(http.StatusForbidden)
			return
		}

		var data userData

		data.Login = request.Login

		digest := sha256.Sum256([]byte(request.Password))
		hexdigest := hex.EncodeToString(digest[:])

		data.Password = hexdigest

		json_data, err := json.Marshal(data)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		added, err := db.SetNX(request.Login, json_data, time.Duration(time.Duration.Seconds(1)))

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		if !added {
			c.Status(http.StatusForbidden)
			return
		}

		c.Status(http.StatusOK)
	}
}

func update(c *gin.Context) {
	var data userData
	err := c.BindJSON(&data)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	jwt_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	claims := jwt_token.Claims.(jwt.MapClaims)

	login := claims["login"]

	fmt.Println(login)
	fmt.Println(data.Login)
	if login != data.Login {
		fmt.Println("How?")
		c.Status(http.StatusForbidden)
		return
	}

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "e64a1e3065c49e781aa2721ce86e2725",
		DB:       0,
	})

	json_data, err := json.Marshal(data)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	_, err = rdb.Set(ctx, data.Login, json_data, time.Duration(time.Duration.Seconds(1))).Result()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func login(c *gin.Context) {
	var request registerRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "e64a1e3065c49e781aa2721ce86e2725",
		DB:       0,
	})
	resp, err := rdb.Get(ctx, request.Login).Result()
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	var data userData
	json.Unmarshal([]byte(resp), &data)
	digest := sha256.Sum256([]byte(request.Password))
	hexdigest := hex.EncodeToString(digest[:])

	if hexdigest != data.Password {
		c.Status(http.StatusUnauthorized)
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = request.Login
	signed_token, err := token.SignedString(privateKey)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "token",
		Value: signed_token,
	})
	c.Status(http.StatusOK)
}

type createTaskRequst struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func createTask(c *gin.Context) {
	var request createTaskRequst
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	jwt_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	claims := jwt_token.Claims.(jwt.MapClaims)

	log.Println(claims)

	login := claims["login"]

	conn, err := grpc.Dial("task-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := task_service.NewTaskServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Create(ctx, &task_service.CreateRequest{Title: request.Title, Content: request.Content, OwnerId: login.(string)})
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, task_service.CreateResponse{TaskId: r.GetTaskId()})
}

type updateTaskRequst struct {
	TaskId  string `json:"task_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

func updateTask(c *gin.Context) {
	var request updateTaskRequst
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	jwt_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	claims := jwt_token.Claims.(jwt.MapClaims)

	login := claims["login"]

	conn, err := grpc.Dial("task-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := task_service.NewTaskServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.Update(ctx, &task_service.UpdateRequest{Title: request.Title, Content: request.Content, OwnerId: login.(string), TaskId: request.TaskId, Status: request.Status})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

type deleteTaskRequst struct {
	TaskId string `json:"task_id"`
}

func deleteTask(c *gin.Context) {
	var request deleteTaskRequst
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	jwt_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	claims := jwt_token.Claims.(jwt.MapClaims)

	login := claims["login"]

	conn, err := grpc.Dial("task-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := task_service.NewTaskServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.Delete(ctx, &task_service.DeleteRequest{OwnerId: login.(string), TaskId: request.TaskId})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

type getTaskRequst struct {
	TaskId string `json:"task_id"`
}

func getTask(c *gin.Context) {
	var request getTaskRequst
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	conn, err := grpc.Dial("task-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := task_service.NewTaskServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Get(ctx, &task_service.GetRequest{TaskId: request.TaskId})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, r)
}

type getListRequst struct {
	OwnerId string `json:"owner_id"`
	FromPos int32  `json:"from_pos"`
	Count   int32  `json:"count"`
}

func getList(c *gin.Context) {
	var request getListRequst
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	conn, err := grpc.Dial("task-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := task_service.NewTaskServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetList(ctx, &task_service.GetListRequest{OwnerId: request.OwnerId, FromPos: request.FromPos, Count: request.Count})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, r)
}

type addLikeRequst struct {
	TaskId string `json:"task_id"`
}

type addLikeKafkaRequst struct {
	UserId string `json:"user_id"`
	TaskId string `json:"task_id"`
}

func addLike(c *gin.Context) {
	var request addLikeRequst
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	token, err := c.Cookie("token")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	jwt_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	claims := jwt_token.Claims.(jwt.MapClaims)

	login := claims["login"] // who wanted to like

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to: %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	add_like_kafka_request := addLikeKafkaRequst{UserId: login.(string), TaskId: request.TaskId}

	add_like_kafka_request_json, err := json.Marshal(add_like_kafka_request)
	if err != nil {
		panic(err)
	}

	topic := "likes"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          add_like_kafka_request_json,
	}, nil)

	p.Flush(15 * 1000)

	c.Status(http.StatusOK)
}

type addViewRequst struct {
	TaskId string `json:"task_id"`
}

type addViewKafkaRequst struct {
	UserId string `json:"user_id"`
	TaskId string `json:"task_id"`
}

func addView(c *gin.Context) {
	var request addViewRequst
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	token, err := c.Cookie("token")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	jwt_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	claims := jwt_token.Claims.(jwt.MapClaims)

	login := claims["login"] // who wanted to add a view

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to: %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	add_view_kafka_request := addViewKafkaRequst{UserId: login.(string), TaskId: request.TaskId}

	add_view_kafka_request_json, err := json.Marshal(add_view_kafka_request)
	if err != nil {
		panic(err)
	}

	topic := "views"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          add_view_kafka_request_json,
	}, nil)

	p.Flush(15 * 1000)

	c.Status(http.StatusOK)
}

type getByTaskRequest struct {
	TaskId string `json:"task_id"`
}

func getByTask(c *gin.Context) {
	var request getByTaskRequest
	err := c.BindJSON(&request)
	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	conn, err := grpc.Dial("stat-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := stat_service.NewStatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetViewsAndLikes(ctx, &stat_service.GetViewsAndLikesRequest{TaskId: request.TaskId})
	if err != nil {
		panic(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, r)
}

type getTop5TasksRequest struct {
	ByLikes bool `json:"by_likes"`
}

type getTop5TasksResponseUnit struct {
	TaskId string `json:"task_id"`
	UserId string `json:"user_id"`
	Count  int    `json:"count"`
}

type getTop5TasksResponse struct {
	Tasks []getTop5TasksResponseUnit `json:"tasks"`
}

func getTop5Tasks(c *gin.Context) {
	var request getTop5TasksRequest
	err := c.BindJSON(&request)
	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	conn, err := grpc.Dial("stat-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := stat_service.NewStatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetTop5LikedOrViewedTasks(ctx, &stat_service.GetTop5LikedOrViewedTasksRequest{Likes: request.ByLikes})
	if err != nil {
		panic(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	conn2, err := grpc.Dial("task-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer conn2.Close()

	client2 := task_service.NewTaskServiceClient(conn2)

	cnt := 5
	if cnt > len(r.GetTasks()) {
		cnt = len(r.GetTasks())
	}

	resp := getTop5TasksResponse{Tasks: make([]getTop5TasksResponseUnit, cnt)}

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()

	for i := 0; i < cnt; i++ {
		r2, err := client2.Get(ctx2, &task_service.GetRequest{TaskId: r.Tasks[i].GetTaskId()})
		if err != nil {
			panic(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		resp.Tasks[i] = getTop5TasksResponseUnit{TaskId: r.Tasks[i].GetTaskId(), UserId: r2.GetOwnerId(), Count: int(r.Tasks[i].GetCount())}
	}

	c.JSON(http.StatusOK, resp)
}

type getTop3UsersRequest struct {
	ByLikes bool `json:"by_likes"`
}

type getTop3UsersResponseUnit struct {
	TaskId string `json:"task_id"`
	UserId string `json:"user_id"`
	Count  int    `json:"count"`
}

type getTop3UsersResponse struct {
	Tasks []getTop5TasksResponseUnit `json:"tasks"`
}

func getTop3Users(c *gin.Context) {
	conn, err := grpc.Dial("stat-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := stat_service.NewStatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetTop3LikedUsers(ctx, &emptypb.Empty{})
	if err != nil {
		panic(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, r)
}

func main() {
	privateKeyPath := flag.String("private", "", "Path to private key")
	publicKeyPath := flag.String("public", "", "Path to public key")
	flag.Parse()

	privateKeyFileContent, err := os.ReadFile(*privateKeyPath)
	if err != nil {
		panic(err)
	}

	publicKeyFileContent, err := os.ReadFile(*publicKeyPath)
	if err != nil {
		panic(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyFileContent)
	if err != nil {
		panic(err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyFileContent)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	var redisDb RedisDb
	r.POST("/register", GetRegisterHandler(&redisDb))
	r.POST("/update", update)
	r.POST("/login", login)

	r.POST("/tasks/create", createTask)
	r.POST("/tasks/update", updateTask)
	r.DELETE("/tasks/delete", deleteTask)
	r.GET("/tasks/get", getTask)
	r.GET("/tasks/getlist", getList)
	r.POST("/tasks/addLike", addLike)
	r.POST("/tasks/addView", addView)

	r.POST("/stats/getByTask", getByTask)
	r.POST("/stats/getTop5Tasks", getTop5Tasks)
	r.GET("/stats/getTop3Users", getTop3Users)

	r.Run("0.0.0.0:7777")
}
