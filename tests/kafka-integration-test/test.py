import unittest
import clickhouse_connect
from uuid import uuid4
from confluent_kafka import Producer
import time

conf = {
    'bootstrap.servers': 'kafka:9092',
    'client.id': 'python-producer'
}

delivered = 0
def delivery_report(err, msg):
    global delivered
    if err is None:
        delivered += 1
    else:
        raise RuntimeError

class TaskServiceTests(unittest.TestCase):
    def test_clickhouse_pulls(self):
        ch_client = clickhouse_connect.create_client(host="clickhouse", port=8123)
        producer = Producer(conf)
        producer.produce("likes", '{"task_id":"123", "user_id":"321"}', callback=delivery_report)
        producer.produce("likes", '{"task_id":"124", "user_id":"321"}', callback=delivery_report)
        producer.produce("likes", '{"task_id":"125", "user_id":"321"}', callback=delivery_report)
        producer.flush()
        ch_client.query("SELECT user_id, task_id FROM likes_m").result_set
        time.sleep(10)
        vals = set(ch_client.query("SELECT * FROM likes_m").result_set)
        self.assertEqual(vals, set([('321', '123'), ('321', '124'), ('321', '125')]))

if __name__ == '__main__':
    unittest.main()
