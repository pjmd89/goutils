package goutils
type DB interface{
	connect(configPath *string) error
}