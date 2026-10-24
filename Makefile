dc_run:
	docker-compose -f ./docker/docker-compose.yml up -d
	sleep 5
	migrate -path ./schema -database 'postgres://postgres_user:postgres_password@localhost/postgres_db?sslmode=disable' up

dc_stop:
	docker-compose -f ./docker/docker-compose.yml stop

dc_down:
	docker-compose -f ./docker/docker-compose.yml down -v
