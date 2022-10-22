package dbutils

import "fmt"

type ModelInterface interface {
	Create(inputs map[string]interface{}, opts interface{}) (r interface{}, err error)
	Read(where interface{}, opts interface{}) (r interface{}, err error)
	Update(inputs map[string]interface{}, where interface{}, opts interface{}) (r interface{}, err error)
	Delete(where interface{}, opts interface{}) (r interface{}, err error)
	Count(where interface{}, opts interface{}) (r int64, err error)
	Init(self interface{}, db DBInterface)
	RawRead(where interface{}, opts interface{}) (r interface{}, err error)
	RawUpdate(inputs map[string]interface{}, where interface{}, opts interface{}) (r interface{}, err error)
	Replace(inputs map[string]interface{}, where interface{}, opts interface{}) (r interface{}, err error)
	InterfaceUpdate(inputs interface{}, where interface{}, opts interface{}) (r interface{}, err error)
	InterfaceReplace(data interface{}, where interface{}, opts interface{}) (r interface{}, err error)
	Repare() (r bool, err error)
	RenameCollection(from, to string)
}
type DBInterface interface {
	Connect() (err error)
	Close() (err error)
	Create(inputs interface{}, collection string, opts interface{}) (results interface{}, err error)
	Read(where interface{}, collection string, opts interface{}) (results interface{}, err error)
	Update(inputs interface{}, where interface{}, collection string, opts interface{}) (results interface{}, err error)
	Replace(inputs interface{}, where interface{}, collection string, opts interface{}) (results interface{}, err error)
	Delete(where interface{}, collection string, opts interface{}) (results interface{}, err error)
	Count(where interface{}, collection string, opts interface{}) (results interface{}, err error)
	SetDatabase(db string)
	SetCollection(collection string)
	GetClient() (client interface{})
	GetDatabase() string
}
type DB struct {
	DBInterface
	Host       string `json:"host"`
	Port       string `json:"port"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	DataBase   string `json:"db"`
	Reconnect  bool   `json:"reconnect"`
	database   string
	collection string
	OnDatabase func(currentDB string, currentCollection string) string
}

func (o *DB) Create(inputs interface{}, collection string, opts interface{}) (results interface{}, err error) {
	err = fmt.Errorf("No declared method")
	return results, err
}
func (o *DB) Read(where interface{}, collection string, opts interface{}, results interface{}) (err error) {
	err = fmt.Errorf("No declared method")
	return err
}
func (o *DB) Update(inputs interface{}, where interface{}, collection string, opts interface{}) (results interface{}, err error) {
	err = fmt.Errorf("No declared method")
	return results, err
}
func (o *DB) Delete(where interface{}, collection string, opts interface{}) (results interface{}, err error) {
	err = fmt.Errorf("No declared method")
	return results, err
}
func (o *DB) Count(where interface{}, collection string, opts interface{}) (results interface{}, err error) {
	err = fmt.Errorf("No declared method")
	return results, err
}
func (o *DB) SetDatabase(db string) {
	o.database = db
}
func (o *DB) GetDatabase() string {
	return o.database
}
func (o *DB) SetCollection(collection string) {
	o.collection = collection
}
func (o *DB) GetCollection() string {
	return o.collection
}
