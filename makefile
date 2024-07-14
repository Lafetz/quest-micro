.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./db/migrations ${name}
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./db/migrations -database postgresql://user:password@localhost:5432/quest?sslmode=disable up

proto-knight:
	protoc --go_out=. --go_opt=paths=source_relative \
  	  --go-grpc_out=. --go-grpc_opt paths=source_relative \
	  proto/knight/knight.proto