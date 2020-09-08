package sql

import (
	"github.com/doug-martin/goqu/v9"
)

//GetUser A sql to get user
func GetUser(id string) (sql string, params []interface{}, err error) {
	return pgBuilder().From("app.user").
			Select("id", "name").
			Where(goqu.C("id").Eq(id)).
			ToSQL()
}

//CreateUser A sql to create todo
func CreateUser(name string) (sql string, params []interface{}, err error) {
	return pgBuilder().From("app.user").Prepared(true).Insert().
			Cols("name").
			Vals(
				goqu.Vals{ name },
			).
			Returning("id").
			ToSQL()
}