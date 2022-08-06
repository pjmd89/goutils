package goutils
type DB interface{
	Connect(configPath *string) error
}