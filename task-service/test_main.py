import unittest
from unittest.mock import MagicMock
from main import TaskServiceServicer
import task_service_pb2

class MockDb:
    def __init__(self):
        self.db = dict()

    def put(self, key, value):
        self.db[key] = value

    def Get(self, key):
        return self.db[key]

class TestMyGRPCService(unittest.TestCase):
    def test_get(self):
        mock_db = MockDb()
        mock_db.put('asd', '{"title":"1", "content":"2", "status":"3", "owner_id":"4"}'.encode())

        service = TaskServiceServicer(mock_db)

        result : task_service_pb2.GetResponse = service.Get(task_service_pb2.GetRequest(task_id="asd"), None)

        self.assertTrue(result.title == '1')
        self.assertTrue(result.content == '2')
        self.assertTrue(result.status == '3')
        self.assertTrue(result.owner_id == '4')

    def test_get_list(self):
        mock_db = MockDb()
        mock_db.put('asd', '{"tasks":["id1"]}'.encode())
        mock_db.put('id1', '{"title":"1", "content":"2", "status":"3"}')

        service = TaskServiceServicer(mock_db)

        result : task_service_pb2.GetListResponse = service.GetList(task_service_pb2.GetListRequest(owner_id="asd", from_pos=0, count=1), None)

        self.assertTrue(result.tasks[0].title == '1')
        self.assertTrue(result.tasks[0].content == '2')
        self.assertTrue(result.tasks[0].status == '3')

if __name__ == '__main__':
    unittest.main()
