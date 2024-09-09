module firm.com/function

go 1.23.0

replace firm.com/connectDB => ../connectDB

require (
	firm.com/connectDB v0.0.0-00010101000000-000000000000
	firm.com/models v0.0.0-00010101000000-000000000000
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
)

replace firm.com/models => ../models
