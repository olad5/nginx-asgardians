docker.start.nginx:
	docker compose --file ./docker-compose.yml  up -d

docker.stop.nginx:
	docker compose --file ./docker-compose.yml  down

run.loki: 
		go run ./loki/main.go

run.odin:
		go run ./odin/main.go

run.thor: 
		go run ./thor/main.go
