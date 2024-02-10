
POSTGRES_PASSWORD ?= "localpass"

PG_PORT ?= 5086
PG_DATA_DIR ?= "$(HOME)/postgres-data"
PG_NAME ?= "postgres-worble"

DB_NAME ?= worble
DB_USER ?= worble

DATABASE_URL ?= "postgres://$(DB_USER):$(POSTGRES_PASSWORD)@localhost:$(PG_PORT)/$(DB_NAME)"

test:
	echo $(DATABASE_URL)

pg-run: pg-clean
	docker run --name $(PG_NAME) -v $(PG_DATA_DIR):/var/lib/postgresql/data -p $(PG_PORT):5432 -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres

pg-clean: pg-stop
	@docker rm $(PG_NAME) 2>/dev/null || true

pg-stop:
	@docker stop $(PG_NAME) 2>/dev/null || true

pg-psql:
	docker exec -it $(PG_NAME) psql -h localhost -p 5432 -U postgres -d postgres

psql:
	psql $(DATABASE_URL)

app-run:
	DATABASE_URL=$(DATABASE_URL) go run ./app/main.go
