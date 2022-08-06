package lib
type DB interface{
	connect(configPath *string) error
}