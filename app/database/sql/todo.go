package sql

import (
	"github.com/doug-martin/goqu/v9"
)

//CreateToDo A sql to create todo
func CreateToDo(userID int, text string) (sql string, params []interface{}, err error) {
	return pgBuilder().From("app.todo").Prepared(true).Insert().
			Cols("user_id", "text").
			Vals(
				goqu.Vals{ userID, text },
			).
			ToSQL()
}

//ListToDo A sql to list todo with  
func ListToDo(limit uint) (sql string, params []interface{}, err error) {
	return pgBuilder().From("app.todo").
			Order(goqu.C("id").Desc()).
			Limit(limit).
			ToSQL()
}