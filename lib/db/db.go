package db
type DB interface{
	connect(configPath *string) error
}