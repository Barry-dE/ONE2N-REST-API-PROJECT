include .env

.PHONY: create-migrations
create-migrations:
	@echo "Creating SQL migration files"
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))
	

.PHONY: migrations-up
migrations-up:
	@echo "Migrating..."
	@migrate -path=$(MIGRATIONS_PATH)  -database=$(DB_ADDR) up
	

.PHONY: migrations-down
migrations_down:
	@echo "Rolling back..."
	@migrate -path=$(MIGRATIONS_PATH)  -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))
	

%:
	@:
