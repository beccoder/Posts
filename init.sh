#!/bin/sh
cp .env ./internal/repository/

sed -i '/^APPLICATION_MODE=/d' .env
echo "APPLICATION_MODE=TESTING" >> .env

go test ./... -v

sed -i '/^APPLICATION_MODE=/d' .env
echo "APPLICATION_MODE=PRODUCTION" >> .env

go run cmd/main.go admin admin admin
./blogs-app