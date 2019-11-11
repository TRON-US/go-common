package constant

const (
	// -- ERRORS --
	// connection
	DBWriteConnectionError  = "Cannot connect to the write database!"
	DBReadConnectionError   = "Cannot connect to the read database!"
	DBMasterConnectionError = "Cannot connect to the master database!"
	DBSlaveConnectionError  = "Cannot connect to the slave database!"
	DBURLParseError         = "Cannot parse database connection URL!"
	RDConnectionError 		= "Cannot connect to the RD database!"

	// CRUD
	DBQueryError         = "Cannot query the database!"
	DBUpsertError        = "Cannot upsert into the database!"
	DBInsertError        = "Cannot insert into the database!"
	DBDeleteError        = "Cannot delete from the database!"
	DBCountEstimateError = "Cannot estimate count the database!"
	PostgresTxContext    = "Postgres Transaction"

	// migration
	DBMigrationError = "Cannot migrate to latest database schema!"

	// -- DEBUG --
	DBConnectionHealthy = "Connection to database is healthy."
	RDConnectionHealthy = "Connection to redis is healthy."

	// write
	DBWriteAvailability = "DBWriteAvailability"
	DBWriteError        = "DBWriteError"
	DBWriteUser         = "DBWriteUser"
	DBWriteAddress      = "DBWriteAddress"
	DBWriteName         = "DBWriteName"
	DBWriteConnection   = "Connected to Write Database."
	// read
	DBReadAvailability = "DBReadAvailability"
	DBReadError        = "DBReadError"
	DBReadUser         = "DBReadUser"
	DBReadAddress      = "DBReadAddress"
	DBReadName         = "DBReadName"
	DBReadConnection   = "Connected to Read Database."

	// -- ORM --
	PGQueryBuilderError = "Cannot build query filter for PG driver!"
)
