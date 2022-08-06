package dbutils
type DB interface{
	connect(configPath *string) error
}