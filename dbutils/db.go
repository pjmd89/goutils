package dbutils

import "fmt"
type DBInterface interface{
	Connect() (err error)
	Close() (err error)
	GetClient() (client interface{})
}
type DBCommandsInterface interface{
	Create(inputs interface{}, opts interface{}, results []*interface{})( err  error)
	Read(where interface{}, opts interface{}, results []*interface{})( err  error)
	Update(inputs interface{}, where interface {}, opts interface{}, results []*interface{}) ( err  error)
	Delete(where interface {}, opts interface{}, results []*interface{}) ( err  error)
	Count(where interface {}, opts interface{}, results []*interface{}) ( err  error)
	SetDatabase(db string)
	SetCollection(collection string)
	GetDatabase() string
	GetCollection() string
}
type DBCommands struct{
	DBCommandsInterface
	database string
	collection string
}
type DB struct {
	DBInterface
	Host 			string		`json:"host"`
	Port 			string		`json:"port"`
	User 			string		`json:"user"`
	Pass 			string		`json:"pass"`
	DataBase   		string		`json:"db"`
	Reconnect		bool		`json:"reconnect"`
	SkipCollection	[]string	`json:"skipCollection"`
	Commands		DBCommands
}
func(o *DBCommands)Create(inputs interface{}, opts interface{}, results []*interface{})( err  error){
	err = fmt.Errorf("No declared method")
	return err;
}
func(o *DBCommands)Read(where interface{}, opts interface{}, results []*interface{})( err  error){
	err = fmt.Errorf("No declared method")
	return err;
}
func(o *DBCommands)Update(inputs interface{}, where interface {}, opts interface{}, results []*interface{}) ( err  error){
	err = fmt.Errorf("No declared method")
	return err;
}
func(o *DBCommands)Delete(where interface {}, opts interface{}, results []*interface{}) ( err  error){
	err = fmt.Errorf("No declared method")
	return err;
}
func(o *DBCommands)Count(where interface {}, opts interface{}, results []*interface{}) ( err  error){
	err = fmt.Errorf("No declared method")
	return err;
}
func(o *DBCommands)SetDatabase(db string){
	o.database = db
}
func(o *DBCommands)GetDatabase() string{
	return o.database
}
func(o *DBCommands)SetCollection(collection string){
	o.collection = collection
}
func(o *DBCommands)GetCollection() string{
	return o.collection
}