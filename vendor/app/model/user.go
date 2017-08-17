package model

import (
	"app/shared/database"
	"fmt"
	"log"
	"encoding/json"
)

type User struct {
	ID uint32 `db:"id" json:"id,string,omitempty"`
	Age uint16 `db:"age" json:"age,string,omitempty"`
	Sex string `db:"sex" json:"sex"`
}

func (u *User) UserID() string {
	return fmt.Sprintf("%v", u.ID)
}

func (u *User) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &u)
}

func UserExist(id string) (User, error){
	var err error

	row, err := database.SQL.Query("SELECT * FROM USER WHERE id = ?", id)

	defer row.Close()

	if err != nil {
		log.Fatal(err)
	}
	var user User
	if row.Next(){
		var id uint32
		var age uint16
		var sex string

		if err := row.Scan(&id, &age, &sex); err != nil {
			log.Fatal(err)
		}
		user.ID = id
		user.Age = age
		user.Sex = sex

		return user, err
	}
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}


	return user, err
}

func UserCreate(id, age, sex string) error {
	var err error

	_, err = database.SQL.Exec("INSERT INTO user (id, age, sex) VALUES (?,?,?)", id, age,
		sex)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func UserRemove(id uint32) (int64, error){
	var err error
	stmt, err := database.SQL.Prepare("delete from user where id=?")
	if err != nil {
		log.Println(err)
		return 0, err
	}

	res, err := stmt.Exec(id)

	affect, err := res.RowsAffected()

	fmt.Println(affect)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return affect, err
}
