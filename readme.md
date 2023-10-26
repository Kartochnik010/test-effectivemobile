# Тестовое задание на Junior Go-разработчика для *Effective Mobile*


## Developer notes

[*] hello world
[*] plan project structure
[*] setting up postgres with docker 
[*] .env config
[*] repository layer
[*] service layer 
[ ] trasport layer
[ ] web server
[ ] tests
[ ] swagger docs


## Endpoints
GET `/people`: Получение данных с фильтрами и пагинацией.

DELETE `/people/{id}`: Удаление записи по `id`.

PUT `/people/{id}`: Изменение сущности по `id`.

POST `/people`: Добавление людей в формате JSON.


## Makefile:
 help            Print this message
 run             run the cmd/api application
 migrate/up      migrations up
 migrate/down    migrations down
 postgres        start postgres conitainer
 postgres/rm     clear postgres container
 postgres/psql   enter postgres container

**Download** migrate tool:
```
go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
```

Create new migration file:
```
migrate create -ext sql -dir db/migrations -seq create_people_table
```

**Execute** migration:
```
migrate -path db/migrations -database $DB_DSN up
```