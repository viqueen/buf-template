harness-up:
	@echo "Starting harness..."
	@docker-compose -f _harness/docker-compose.yaml up -d

harness-down:
	@echo "Stopping harness..."
	@docker-compose -f _harness/docker-compose.yaml down

migrations:
	@echo "flyway..."
	@docker run --rm \
		--network="todo_network" \
		--volume $(PWD)/backend/data/schema:/flyway/backend/data/schema \
		flyway/flyway:7.2.0-alpine \
			-url="jdbc:postgresql://database:5432/todo" \
			-user="todo" -password="todo" \
			-schemas="public" -defaultSchema="public" \
			-locations=filesystem:/flyway/backend/data/schema \
			migrate info

run-backend-server:
	@echo "Running backend server..."
	@go run ./backend/cmd/server/main.go