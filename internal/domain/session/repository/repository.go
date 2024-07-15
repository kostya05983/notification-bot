package repository

import (
	"github.com/ydb-platform/ydb-go-sdk/v3"
)

type Repository struct {
	driver *ydb.Driver
}
