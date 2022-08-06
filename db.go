package main
type DB interface{
	connect(configPath *string) error
}