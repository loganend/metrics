package model

import (
	"time"
	"app/shared/database"
)

type Stat struct {
	ID uint32 `db:"id" json:"id,string,omitempty"`
	UID       uint32  `db:"user_id" json:"user,string,omitepty"`
	Action string `db:"user_id" json:"action"`
	Datetime time.Time `db:"datetime" json:"ts,string,omitepty"`
}

func StatCreate(user uint32, action string, datetime time.Time) error{
	var err error

	_, err = database.SQL.Exec("INSERT INTO stat (user_id, action, datetime) VALUES (?,?,?)", user, action,
		datetime)

	return standardizeError(err)
}

