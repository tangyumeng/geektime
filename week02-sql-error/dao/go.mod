module example.com/dao

go 1.18

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/pkg/errors v0.9.1
)

replace example.com/dao => ./dao
