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

test: test_var create_db_storage_folder build_image compose_up run

clean: compose_down remove_db_storage_folder

re_run: clean create_db_storage_folder build_image compose_up run

.PHONY: run build_image compose_up compose_down stop_container remove_container remove_db_storage_folder create_db_storage_folder clean re_run
