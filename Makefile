include .env

DB_STORAGE_FOLDER=.db-data

all: prod_var create_db_storage_folder build_image compose_up run

run:
	go run cmd/main.go

build_image:
	docker compose build

compose_up:
	docker compose up -d

compose_down:
	docker compose down

stop_container:
	docker stop ${DB_CONTAINER_NAME}

remove_container:
	docker remove ${DB_CONTAINER_NAME}

remove_db_storage_folder:
	sudo rm -rf ${DB_STORAGE_FOLDER}

create_db_storage_folder:
	mkdir ${DB_STORAGE_FOLDER}

test_var:
	sed -i '/^RUN_MODE=/d' .env
	echo "RUN_MODE=test" >> .env


prod_var:
	sed -i '/^RUN_MODE=/d' .env
	echo "RUN_MODE=prod" >> .env

swag_init:
	swag init -g internal/app.go

add_admin:
	go run cmd/main.go admin admin admin

test: test_var build_image compose_up run_test

run_test:
	go test ./...

clean: compose_down

rebuild: build_image compose_up run

f_clean: compose_down remove_db_storage_folder

f_rebuild: f_clean create_db_storage_folder build_image compose_up run

.PHONY: f_clean run build_image compose_up compose_down stop_container remove_container remove_db_storage_folder create_db_storage_folder clean re_run swag-init
