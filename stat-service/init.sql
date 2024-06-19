CREATE TABLE IF NOT EXISTS likes
(
    user_id TEXT,
    task_id TEXT
) ENGINE = Kafka()
SETTINGS kafka_broker_list = 'kafka:9092',
         kafka_topic_list = 'likes',
         kafka_group_name = 'likes',
         kafka_format = 'JSONEachRow',
         kafka_num_consumers = 1;

CREATE TABLE IF NOT EXISTS views
(
    user_id TEXT,
    task_id TEXT
) ENGINE = Kafka()
SETTINGS kafka_broker_list = 'kafka:9092',
         kafka_topic_list = 'views',
         kafka_group_name = 'views',
         kafka_format = 'JSONEachRow',
         kafka_num_consumers = 1;

CREATE TABLE IF NOT EXISTS likes_m (
    user_id TEXT,
    task_id TEXT
) ENGINE MergeTree ORDER BY user_id;

CREATE MATERIALIZED VIEW IF NOT EXISTS likes_mv TO likes_m AS SELECT * FROM likes;

CREATE TABLE IF NOT EXISTS views_m (
    user_id TEXT,
    task_id TEXT
) ENGINE MergeTree ORDER BY user_id;

CREATE MATERIALIZED VIEW IF NOT EXISTS views_mv TO views_m AS SELECT * FROM views;
