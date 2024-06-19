export COMPOSE_FILE=docker-compose.yaml

docker compose up -d --build redis-tasks task-service

sleep 15

python3 -m unittest test.py || exit 1

docker compose down
