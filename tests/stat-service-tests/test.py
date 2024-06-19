import unittest
from stat_service_pb2_grpc import StatServiceStub
from stat_service_pb2 import GetViewsAndLikesRequest, GetViewsAndLikesResponse
import clickhouse_connect
import grpc
from uuid import uuid4


class TaskServiceTests(unittest.TestCase):
    def test_get_views_and_likes(self):
        with grpc.insecure_channel('localhost:50051') as channel:
            ch_client = clickhouse_connect.create_client(host="localhost", port=8123)
            task_id = str(uuid4())
            user_id = str(uuid4())
            ch_client.query("INSERT INTO likes_m (task_id, user_id) VALUES (%s, %s), (%s, %s)", (task_id, user_id, task_id, user_id,))
            ch_client.query("INSERT INTO views_m (task_id, user_id) VALUES (%s, %s), (%s, %s)", (task_id, user_id, task_id, user_id,))

            stub = StatServiceStub(channel)
            response1 : GetViewsAndLikesResponse = stub.GetViewsAndLikes(GetViewsAndLikesRequest(task_id=task_id))
            self.assertEqual(response1.likes, 1)
            self.assertEqual(response1.views, 1)


if __name__ == '__main__':
    unittest.main()
