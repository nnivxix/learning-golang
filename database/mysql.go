package database

var connection string

func init() {
	// invoked automatically when package imported
	connection = "mySql"
}

func GetDatabase() string {
	return connection
}
