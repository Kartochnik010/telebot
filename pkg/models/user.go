package models

import (
	"database/sql"
	"log"
)

// Статистика
// - показать юзеру когда был его первый запрос,
// - сколько всего запросов было
// - можно добавить ещё показателей, которые мы можем получить из хранимых данных.

type User struct {
	FirstName string
	LastName  string
	ID        int64
	CreatedAt string

	FirstQueryTimeStamp string
	QueryCounter        int
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(user User) error {
	log.Printf("Insert " + user.FirstName)
	return nil
}

func (u *UserModel) Get(id string) (*User, error) {
	log.Printf("Get user with id " + id)
	return nil, nil
}
