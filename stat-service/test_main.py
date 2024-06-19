import unittest
from unittest.mock import MagicMock
from main import TaskServiceServicer
import stat_service_pb2

class CHRet:
    def __init__(self, a):
        self.result_set = a

class MockDb:
    def __init__(self):
        self.db = dict()

    def put(self, key, value):
        self.db[key] = value

    def query(self, query, parameters = None):
        return self.db[(query, parameters)]

class TestMyGRPCService(unittest.TestCase):
    def test_get(self):
        mock_db = MockDb()
        mock_db.put(('SELECT count(distinct user_id) FROM likes_m WHERE task_id = %s', ('123',)), CHRet(a=[[2]]))
        mock_db.put(('SELECT count(distinct user_id) FROM views_m WHERE task_id = %s', ('123',)), CHRet(a=[[3]]))

        service = TaskServiceServicer(mock_db)

        result : stat_service_pb2.GetViewsAndLikesResponse = service.GetViewsAndLikes(stat_service_pb2.GetViewsAndLikesRequest(task_id="123"), None)

        self.assertTrue(result.task_id == '123')
        self.assertTrue(result.views == 3)
        self.assertTrue(result.likes == 2)

    def test_get_list(self):
        mock_db = MockDb()
        mock_db.put(("SELECT task_id, COUNT(DISTINCT user_id) AS unique_entries_count FROM likes_m GROUP BY task_id ORDER BY unique_entries_count DESC LIMIT 5", None), CHRet(a=[['asd', 123], ['dsa', 321]]))

        service = TaskServiceServicer(mock_db)

        result : stat_service_pb2.GetTop5LikedOrViewedTasksResponse = service.GetTop5LikedOrViewedTasks(stat_service_pb2.GetTop5LikedOrViewedTasksRequest(likes=True), None)

        self.assertTrue(result.tasks[0].task_id == 'asd')
        self.assertTrue(result.tasks[0].count == 123)
        self.assertTrue(result.tasks[1].task_id == 'dsa')
        self.assertTrue(result.tasks[1].count == 321)

if __name__ == '__main__':
    unittest.main()
