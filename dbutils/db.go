package dbutils

import "fmt"
type DBInterface interface{
	Connect() (err error)
	Close() (err error)
	Create(inputs interface{}, opts interface{}, results []*interface{})( err  error)
	Read(where interface{}, opts interface{}, results []*interface{})( err  error)
	Update(inputs interface{}, where interface {}, opts interface{}, results []*interface{}) ( err  error)
	Delete(where interface {}, opts interface{}, results []*interface{}) ( err  error)
	Count(where interface {}, opts interface{}, results []*interface{}) ( err  error)
	SetDatabase(db string)
	SetCollection(collection string)
	GetClient() (client interface{})
	GetDatabase() string
	GetCollection() string
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
	database 		string
	collection 		string
}
func(o *DB)Create(inputs interface{}, opts interface{}, results []*interface{})( err  error){
	err = fmt.Errorf("No declared method")
	return err;
}
func(o *DB)Read(where interface{}, opts interface{}, results []*interface{})( err  error){
	err = fmt.Errorf("No declared method")
	return err;
}
func(o *DB)Update(inputs interface{}, where interface {}, opts interface{}, results []*interface{}) ( err  error){
	err = fmt.Errorf("No declared method")
	return err;
}
func(o *DB)Delete(where interface {}, opts interface{}, results []*interface{}) ( err  error){
	err = fmt.Errorf("No declared method")
	return err;
}
func(o *DB)Count(where interface {}, opts interface{}, results []*interface{}) ( err  error){
	err = fmt.Errorf("No declared method")
	return err;
}
func(o *DB)SetDatabase(db string){
	o.database = db
}
func(o *DB)GetDatabase() string{
	return o.database
}
func(o *DB)SetCollection(collection string){
	o.collection = collection
}
func(o *DB)GetCollection() string{
	return o.collection
}