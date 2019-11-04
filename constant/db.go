package constant

const (
	// connection
	DBWriteConnectionError = "Cannot connect to the write database!"
	DBReadConnectionError  = "Cannot connect to the read database!"
	DBURLParseError        = "Cannot parse database URL!"
	DBConnectionHealthy    = "DB Connection Healthy"
	// CRUD
	DBUpsertError    = "DbUpsertError!"
	DBEmptyResultSet = "pg: no rows in result set"
	DBQueryError     = "Cannot query the database!"
	// migration
	DBMigrationError = "Cannot migrate to latest database schema!"
)
