package dao

import (
	"Go-000/Week02/model"
	"database/sql"
	"github.com/pkg/errors"
)

var db *sql.DB

func QueryList() ([]model.User, error) {
	data, err := db.Query("select * from user where id = 1")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "query users errors")
		}
	}
	var user []model.User
	if err = data.Scan(&user); err != nil {
		return nil, errors.Wrap(err, "scan find error")
	}
	return user, nil
}
