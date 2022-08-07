package dbutils
type DBInterface interface{
	Connect(configPath *string) error
	Close() error
}
type DBCommandsInterface interface{
	Create(inputs interface{}, opts interface{}, results []*interface{})( err  error)
	Read(where interface{}, opts interface{}, results []*interface{})( err  error)
	Update(inputs interface{}, where interface {}, opts interface{}, results []*interface{}) ( err  error)
	Delete(where interface {}, opts interface{}, results []*interface{}) ( err  error)
	Count(where interface {}, opts interface{}, results []*interface{}) ( err  error)
	SetDatabase(db string)
	SetCollection(collection string)
	
}
type DBCommands struct{
	DBCommandsInterface
	db string
	collection string
}
type DB struct {
	DBInterface
	Host 			string		`json:"host"`
	Port 			string		`json:"port"`
	User 			string		`json:"user"`
	Pass 			string		`json:"pass"`
	DB   			string		`json:"db"`
	Reconnect		bool		`json:"reconnect"`
	SkipCollection	[]string	`json:"skipCollection"`
	Commands		DBCommands
}