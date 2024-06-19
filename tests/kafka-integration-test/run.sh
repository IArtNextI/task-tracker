export COMPOSE_FILE=docker-compose.yaml

docker compose up -d --build zookeeper kafka clickhouse

sleep 15

docker compose run --build test-runner

docker compose down
