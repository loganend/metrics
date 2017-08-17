package model

import (
	"time"
	"app/shared/database"
)

type Stat struct {
	ID uint32 `db:"id" json:"id,string,omitempty"`
	UID       uint32  `db:"user_id" json:"user,string,omitepty"`
	Action string `db:"action" json:"action"`
	Datetime time.Time `db:"datetime" json:"ts,string,omitepty"`
}

func StatGet(user, action, datetime string) error{

	//fmt.Println(datetime)
	//date, err := time.Parse("2006-01-02 15:04:05", datetime)

	result := Stat{}
	err := database.SQL.Get(&result, "select distinct * from stat where user_id = ? and action = ? and datetime = ?", user, action, datetime);

	return err;
}

func StatCreate(user uint32, action string, datetime time.Time) error{
	var err error

	_, err = database.SQL.Exec("INSERT INTO stat (user_id, action, datetime) VALUES (?,?,?)", user, action,
		datetime)

	return standardizeError(err)
}

func StatRemove(user uint32, action string, datetime string) error{
	var err error

	_, err = database.SQL.Exec("DELETE FROM stat WHERE user_id = ? and action = ? and datetime = ?", user, action,
		datetime)

	return standardizeError(err)
}