include .env

.PHONY: create-migrations
create-migrations:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))
	

.PHONY: migrations-up
migrations-up:
	@migrate -path=$(MIGRATIONS_PATH)  -database=$(DB_ADDR) up
	

.PHONY: migrations-down
migrations_down:
	@migrate -path=$(MIGRATIONS_PATH)  -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))
	

.PHONY: migrations-status
migrations-status:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) status
%:
	@:
