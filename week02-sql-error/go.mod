module example.com/sqlerr

go 1.18

require example.com/dao v0.0.0-00010101000000-000000000000

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)

replace example.com/dao => ./dao
