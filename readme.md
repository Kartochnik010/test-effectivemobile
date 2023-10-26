# Тестовое задание на Junior Go-разработчика для *Effective Mobile*

Всегда можем обсудить любые вопросы у меня в телеграмме: [@ilyas_amantaev](https://t.me/ilyas_amantaev)

## Developer notes

- [x] hello world
- [x] plan project structure
- [x] setting up postgres with docker 
- [x] .env config
- [x] repository layer
- [x] service layer 
- [x] trasport layer
- [x] logging
- [x] web server
- [ ] pagination, filtering and validation
- [ ] tests
- [ ] swagger docs

*You can check commits with comments*

## Endpoints
GET `/people`: Получение данных с фильтрами и пагинацией.

GET `/people/{id}`: Получение данных с фильтрами и пагинацией.

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