package database

var connection string

func init() {
	connection = "My SQL"
}

func GetDatabase() string {
	return connection	
}