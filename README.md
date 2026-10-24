# POST SERVICE GO

Пример Rest API для создания, получения, изменения и удаления постов с прикрепленными файлами.

## RUN

`migrate -path ./schema -database 'postgres://postgres_user:postgres_password@localhost/postgres_db?sslmode=disable' up`
`go run ./cmd/main.go`
