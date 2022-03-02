
build:
	docker-compose build proto
	docker-compose run --rm proto buf build -o server.pb --as-file-descriptor-set

proto: 
	docker-compose build proto
	docker-compose run --rm proto

server: 
	docker-compose build server
	docker-compose up -d server

run: proto clean server
	docker-compose up --build -d proxy
	docker-compose  logs -f

clean:
	docker-compose down --remove-orphans

