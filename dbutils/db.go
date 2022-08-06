package dbutils
type DB interface{
	Connect(configPath *string) error
}