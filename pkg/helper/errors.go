package helper

import (
	"errors"

	"github.com/lib/pq"
)

var ErrInternal = errors.New("服务器开小差")
var ErrUserNotFound = errors.New("用户不存在")

// 是否存在唯一索引错误
func IsPGUniqueIndexErr(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		return pgErr.Code == "23505"
	}
	return false
}
