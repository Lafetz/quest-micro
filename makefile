.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./db/migrations ${name}
.PHONY: db/migrations/up
db/migrations/up: 
	@echo 'Running up migrations...'
	migrate -path ./db/migrations -database postgresql://user:password@localhost:5432/quest?sslmode=disable up

proto-knight:
	protoc --go_out=. --go_opt=paths=source_relative \
  	  --go-grpc_out=. --go-grpc_opt paths=source_relative \
	  proto/knight/knight.proto
run-knight:
	@echo "Loading environment variables from .env file"
	@set -o allexport; source ./load_env.sh; set +o allexport; \
	echo "Running knight service application"; \
	go run ./cmd/knight/main.go
run-quest:
	@echo "Loading environment variables from .env file"
	@set -o allexport; source ./load_env.sh; set +o allexport; \
	echo "Running quest service application"; \
	go run ./cmd/quest/main.go