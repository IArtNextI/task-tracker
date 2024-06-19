from concurrent import futures
import grpc
import stat_service_pb2
import stat_service_pb2_grpc
import clickhouse_connect
import clickhouse_connect.driver.client
from grpc_reflection.v1alpha import reflection
import json

def getFreshClickHouse() -> clickhouse_connect.driver.client.Client:
    return clickhouse_connect.get_client(host='clickhouse', port=8123)

class CHDB:
    def query(self, str, parameters=None):
        ch = getFreshClickHouse()
        return ch.query(str, parameters)

class TaskServiceServicer(stat_service_pb2_grpc.StatServiceServicer):
    def __init__(self, db):
        super().__init__()
        self.db : CHDB = db

    def GetViewsAndLikes(self, request : stat_service_pb2.GetViewsAndLikesRequest, context : grpc.ServicerContext):
        task_id = request.task_id
        likes = self.db.query("SELECT count(distinct user_id) FROM likes_m WHERE task_id = %s", (task_id,)).result_set[0][0]
        views = self.db.query("SELECT count(distinct user_id) FROM views_m WHERE task_id = %s", (task_id,)).result_set[0][0]
        return stat_service_pb2.GetViewsAndLikesResponse(views=views, likes=likes, task_id=task_id)

    def GetTop5LikedOrViewedTasks(self, request : stat_service_pb2.GetTop5LikedOrViewedTasksRequest, context : grpc.ServicerContext):
        do_get_likes = request.likes
        tasks = []
        if do_get_likes:
            tasks = self.db.query("SELECT task_id, COUNT(DISTINCT user_id) AS unique_entries_count FROM likes_m GROUP BY task_id ORDER BY unique_entries_count DESC LIMIT 5").result_set
        else:
            tasks = self.db.query("SELECT task_id, COUNT(DISTINCT user_id) AS unique_entries_count FROM views_m GROUP BY task_id ORDER BY unique_entries_count DESC LIMIT 5").result_set
        response = stat_service_pb2.GetTop5LikedOrViewedTasksResponse(tasks=[
            stat_service_pb2.GetTop5LikedOrViewedTasksResponse.TaskIdAndCount(task_id=x[0], count=x[1])
            for x in tasks
        ])
        return response

    def GetTop3LikedUsers(self, request, context : grpc.ServicerContext):
        client = getFreshClickHouse()
        tasks = client.query("SELECT user_id, COUNT(DISTINCT task_id) AS unique_entries_count FROM likes_m GROUP BY user_id ORDER BY unique_entries_count DESC LIMIT 3").result_set
        response = stat_service_pb2.GetTop3LikedUsersResponse(user_ids=[
            stat_service_pb2.GetTop3LikedUsersResponse.UserIdAndLikeCount(user_id=x[0], like_count=x[1])
            for x in tasks
        ])
        return response


if __name__ == '__main__':
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    SERVICE_NAMES = (
        stat_service_pb2.DESCRIPTOR.services_by_name['StatService'].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    stat_service_pb2_grpc.add_StatServiceServicer_to_server(TaskServiceServicer(CHDB()), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()
