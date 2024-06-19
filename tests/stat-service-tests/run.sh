export COMPOSE_FILE=docker-compose.yaml

docker compose up -d --build stat-service clickhouse

sleep 15

python3 -m unittest test.py || exit 1

docker compose down
