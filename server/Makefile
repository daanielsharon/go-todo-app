postgresinit: 
	docker run --rm --name postgres -p 1234:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

postgres:
	docker exec -it postgres psql

check_connection:
	docker exec -it postgres psql -c "SELECT count(*) FROM pg_stat_activity;"

createdb:
	docker exec -it postgres createdb --username=root --owner=root todoapp

dbup:
	migrate -path migrations -database "postgresql://root:root@localhost:1234/todoapp?sslmode=disable" -verbose up

dbdown:
	migrate -path migrations -database "postgresql://root:root@localhost:1234/todoapp?sslmode=disable" -verbose down

createdbtest:
	docker exec -it postgres createdb --username=root --owner=root todoapp_test

dbtestup:
	migrate -path migrations -database "postgresql://root:root@localhost:1234/todoapp_test?sslmode=disable" -verbose up

dbtestdown:
	migrate -path migrations -database "postgresql://root:root@localhost:1234/todoapp_test?sslmode=disable" -verbose down

dropdb:
	docker exec -it postgres dropdb todoapp

dropdbtest:
	docker exec -it postgres dropdb todoapp_test
