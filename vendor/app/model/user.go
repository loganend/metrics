package model

import (
	"app/shared/database"
	"fmt"
)

type User struct {
	ID uint32 `db:"id" json:"id,string,omitempty"`
	Age uint16 `db:"age" json:"age,string,omitempty"`
	Sex string `db:"sex" json:"sex"`
}

func (u *User) UserID() string {
	return fmt.Sprintf("%v", u.ID)
}

func UserCreate(id, age, sex string) error {
	var err error

	_, err = database.SQL.Exec("INSERT INTO user (id, age, sex) VALUES (?,?,?)", id, age,
		sex)

	return standardizeError(err)
}