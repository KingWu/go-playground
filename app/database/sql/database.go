package sql

import (
	"github.com/doug-martin/goqu/v9"
  _ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func pgBuilder() goqu.DialectWrapper {
	return goqu.Dialect("postgres")
}