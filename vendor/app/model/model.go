package model

import (
	"database/sql"
	"errors"

	"gopkg.in/mgo.v2"
)

var (
	ErrNoResult = errors.New("Result not found.")
)

func standardizeError(err error) error {
	if err == sql.ErrNoRows || err == mgo.ErrNotFound {
		return ErrNoResult
	}

	return err
}
