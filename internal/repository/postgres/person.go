package postgres

import (
	"context"
	"fmt"
	"reflect"

	"github.com/Kartochnik010/test-effectivemobile/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

const PersonTable = "person"

type PersonRepo struct {
	dbClient *pgx.Conn
}

func NewPersonRepo(dbConn *pgx.Conn) *PersonRepo {
	return &PersonRepo{
		dbClient: dbConn,
	}
}

// Throws pgx.ErrNoRows if no rows are returned
func (r *PersonRepo) FindPersonById(id int) (models.Person, error) {
	var person models.Person
	query := "SELECT id, name, surname, patronymic, age, gender, nationality FROM person WHERE id = $1"

	err := r.dbClient.QueryRow(context.Background(), query, id).Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality)
	if err != nil {
		return models.Person{}, err
	}
	return person, nil
}

func (r *PersonRepo) UpdatePersonById(id int, person models.Person) (models.Person, error) {
	// conditional update
	query, args := buildPersonUpdateQuery(person)
	log.Debug().Msg(fmt.Sprint("Preparing query: "+query, args))

	err := r.dbClient.QueryRow(context.Background(), query, append([]interface{}{id}, args...)).Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality)
	if err != nil {
		return models.Person{}, err
	}

	return person, nil
}

func (r *PersonRepo) SearchPerson(filters models.Filters) ([]models.Person, models.Metadata, error) {

	return nil, models.Metadata{}, nil
}

func (r *PersonRepo) InsertPerson(person models.Person) (models.Person, error) {
	// conditional insert
	query, args := buildPersonInsertQuery(person)
	log.Debug().Msg(fmt.Sprint("Preparing query: "+query, args))

	err := r.dbClient.QueryRow(context.Background(), query).Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality)
	if err != nil {
		return models.Person{}, err
	}
	return person, err
}
func (r *PersonRepo) DeletePersonById(id int) error {
	_, err := r.dbClient.Exec(context.Background(), "DELETE FROM person WHERE id = $1", id)
	return err
}

func buildPersonInsertQuery(person models.Person) (string, []interface{}) {
	arr, args := buildArgs(person)
	query := "INSERT INTO person ("
	for _, e := range arr {
		query += fmt.Sprintf("%s,", e)
	}

	query = query[:len(query)-1] + ") VALUES ("
	for _, e := range args {
		if reflect.TypeOf(e).Kind() == reflect.Int {
			query += fmt.Sprintf("%s,", e)
			continue
		}
		query += fmt.Sprintf("'%s',", e)
	}
	query = query[:len(query)-1] + ") returning id, name, surname, patronymic, age, gender, nationality;"
	return query, args
}

func buildPersonUpdateQuery(person models.Person) (string, []interface{}) {
	arr, args := buildArgs(person)
	query := "UPDATE person SET "
	for i, e := range arr {
		query += fmt.Sprintf("%s = $%v, ", e, i+2)
	}
	return query[:len(query)-1] + " FROM person WHERE id = $1 returning id, name, surname, patronymic, age, gender, nationality;", args
}

// retrieves arguments from Person struct for building sql query
func buildArgs(person models.Person) ([]string, []interface{}) {
	arr := []string{}
	args := []interface{}{}
	if person.Name != "" {
		arr = append(arr, "name")
		args = append(args, person.Name)
	}
	if person.Surname != "" {
		arr = append(arr, "surname")
		args = append(args, person.Surname)
	}
	if person.Patronymic.String != "" {
		arr = append(arr, "patronymic")
		args = append(args, person.Patronymic.String)
	}
	if person.Age.Int32 != 0 {
		arr = append(arr, "age")
		a := fmt.Sprintf("%v", person.Age.Int32)
		args = append(args, a)

	}
	if person.Gender.String == "" {
		arr = append(arr, "gender")
		args = append(args, person.Gender.String)
	}
	if person.Nationality.String != "" {
		arr = append(arr, "nationality")
		args = append(args, person.Nationality.String)
	}
	return arr, args
}
