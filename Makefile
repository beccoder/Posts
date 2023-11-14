all: run

rebuild_db:
	sudo docker compose down
	sudo rm -rf .db-data
	sudo mkdir .db-data
	sudo docker compose up --build

run:
	go run cmd/main.go

#denis isaev linter golang