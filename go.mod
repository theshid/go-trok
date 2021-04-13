module github.com/theshid/go-trok

go 1.16

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/jackc/pgx/v4 v4.11.0
	github.com/theshid/go-trok/src/models v1.2.3
	github.com/theshid/go-trok/src/routes v1.2.3
)

replace github.com/theshid/go-trok/src/models => ./src/models

replace github.com/theshid/go-trok/src/routes => ./src/routes
