package sql

import (
	"github.com/doug-martin/goqu/v9"
)

//CreateUser A sql to create todo
func CreateUser(name string) (sql string, params []interface{}, err error) {
	return pgBuilder().From("app.user").Prepared(true).Insert().
			Cols("name").
			Vals(
				goqu.Vals{ name },
			).
			ToSQL()
}