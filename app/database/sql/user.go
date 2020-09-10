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

func GetUsers(ids [] string) (sql string, params []interface{}, err error) {
	tempIds := make([]interface{}, len(ids))
	for i, v := range ids {
		tempIds[i] = v
	}
	return pgBuilder().From("app.user").
			Select("id", "name").
			Where(goqu.C("id").In(tempIds...)).
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