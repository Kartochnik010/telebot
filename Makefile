
include .env

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


# ==================================================================================== # 
# DEVELOPMENT
# ==================================================================================== #

## dev: run in development mode
.PHONY: dev
dev:
	go run cmd/bot/main.go

## db: access db
.PHONY: db
db:
	docker exec -it e7049d04f0420d9aa26be28f03c9ecd263b8797a74be90522559a989896f2602 psql ${BD_DSN} 