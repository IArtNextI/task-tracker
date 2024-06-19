import unittest
from task_service_pb2_grpc import TaskServiceStub
from task_service_pb2 import CreateRequest, GetRequest, CreateResponse, GetResponse, GetListRequest, GetListResponse, UpdateRequest, UpdateResponse
import grpc

class TaskServiceTests(unittest.TestCase):
    def test_create_task_and_get(self):
        with grpc.insecure_channel('localhost:50051') as channel:
            stub = TaskServiceStub(channel)
            response1 : CreateResponse = stub.Create(CreateRequest(title="test", content="testc", owner_id="testo"))
            task_id = response1.task_id
            self.assertTrue(response1.task_id)
            response2 : GetResponse = stub.Get(GetRequest(task_id=task_id))
            self.assertEqual((response2.title, response2.content, response2.owner_id, response2.status), ("test", "testc", "testo", "Unknown"))
            response3 : UpdateResponse = stub.Update(UpdateRequest(task_id=task_id, owner_id="testo", title="test2", content="testc2", status="done"))
            response4 : GetListResponse = stub.GetList(GetListRequest(owner_id="testo", from_pos=-1, count=1))
            self.assertEqual([(response4.tasks[0].title, response4.tasks[0].content, response4.tasks[0].status)], [("test2", "testc2",  "done")])
            response5 : GetResponse = stub.Get(GetRequest(task_id=task_id))
            self.assertEqual((response5.title, response5.content, response5.owner_id, response5.status), ("test2", "testc2", "testo", "done"))


if __name__ == '__main__':
    unittest.main()
