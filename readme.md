# Тестовое задание на Junior Go-разработчика для *Effective Mobile*

## Endpoints
GET `/people`: Получение данных с фильтрами и пагинацией.

DELETE `/people/{id}`: Удаление записи по `id`.

PUT `/people/{id}`: Изменение сущности по `id`.

POST `/people`: Добавление людей в формате JSON.


## Developer notes

> Some steps will be included in Makefile

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