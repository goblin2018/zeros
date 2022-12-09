package model

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

type Executable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var UserAuthTypeSystem string = "system"
var UserAuthTypeSmallWX string = "wxMini"
