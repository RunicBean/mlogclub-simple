package sqls

import (
	"database/sql"

	"github.com/RunicBean/mlogclub-simple/common/strs"
)

func SqlNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  len(value) > 0,
	}
}

func KeywordWrap(keyword string) string {
	if strs.IsBlank(keyword) {
		return keyword
	}
	return "\"" + keyword + "\""
}
