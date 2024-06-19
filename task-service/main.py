from concurrent import futures
import grpc
import task_service_pb2
import task_service_pb2_grpc
import redis
from grpc_reflection.v1alpha import reflection
import json

from uuid import uuid4

HOST = "redis-tasks"

def getFreshRedis() -> redis.Redis:
    return redis.Redis(host=HOST, port=6380, db=0, password='e64a1e3065c49e781aa2721ce86e2725')


class RedisDb:
    def Get(self, key):
        r = getFreshRedis()
        return r.get(key)

class TaskServiceServicer(task_service_pb2_grpc.TaskServiceServicer):
    def __init__(self, db):
        super().__init__()
        self.db : RedisDb = db

    def Create(self, request : task_service_pb2.CreateRequest, context : grpc.ServicerContext):
        owner_id = request.owner_id
        title = request.title
        content = request.content

        if not owner_id:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            return

        r = getFreshRedis()
        newId = str(uuid4())

        task_desc = {'owner_id' : owner_id, 'title' : title, 'content' : content, 'status' : "Unknown"}

        r.set(newId, json.dumps(task_desc))

        llst = r.get(owner_id)
        if llst is not None:
            lst = json.loads(llst.decode())["tasks"]
            lst.append(newId)
            r.set(owner_id, json.dumps({"tasks":lst}))
        else:
            r.set(owner_id, json.dumps({"tasks":[newId]}))

        return task_service_pb2.CreateResponse(task_id=newId)

    def Update(self, request : task_service_pb2.UpdateRequest, context : grpc.ServicerContext):
        task_id = request.task_id
        owner_id = request.owner_id
        title = request.title
        content = request.content
        status = request.status

        r = getFreshRedis()

        val = json.loads(r.get(task_id).decode())

        if owner_id != val['owner_id']:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            return

        if title:
            val['title'] = title
        if content:
            val['content'] = content
        if status:
            val['status'] = status

        r.set(task_id, json.dumps(val))

        return task_service_pb2.UpdateResponse()

    def Delete(self, request : task_service_pb2.DeleteRequest, context : grpc.ServicerContext):
        task_id = request.task_id
        owner_id = request.owner_id

        r = getFreshRedis()

        val = json.loads(r.get(task_id).decode())

        if owner_id != val["owner_id"]:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            return

        lst = json.loads(r.get(owner_id).decode())["tasks"]

        lst = set(lst).difference([task_id])

        lst = list(lst)

        r.set(owner_id, json.dumps({"tasks":lst}))
        r.delete(task_id)

        return task_service_pb2.DeleteResponse()

    def Get(self, request : task_service_pb2.GetRequest, context : grpc.ServicerContext):
        task_id = request.task_id

        val = json.loads(self.db.Get(task_id).decode())

        return task_service_pb2.GetResponse(title=val['title'], content=val['content'], status=val['status'], owner_id=val['owner_id'])

    def GetList(self, request : task_service_pb2.GetListRequest, context : grpc.ServicerContext):
        owner_id = request.owner_id
        from_pos = request.from_pos
        count = request.count

        lst = json.loads(self.db.Get(owner_id).decode())['tasks'][from_pos:][:count]

        response = task_service_pb2.GetListResponse()
        for task_id in lst:
            val = json.loads(self.db.Get(task_id))
            response.tasks.append(task_service_pb2.Task(task_id=task_id, title=val['title'], content=val['content'], status=val['status']))

        return response

if __name__ == '__main__':
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    SERVICE_NAMES = (
        task_service_pb2.DESCRIPTOR.services_by_name['TaskService'].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    task_service_pb2_grpc.add_TaskServiceServicer_to_server(TaskServiceServicer(RedisDb()), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()
