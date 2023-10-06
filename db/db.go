package db

// Conn // don't use global exported variables if it can lead to unexpected results when changed
var Conn string

func Open(dbName string) {
	Conn = dbName
}
