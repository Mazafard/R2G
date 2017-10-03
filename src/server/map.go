package server

var CacheDbMap  struct{
	Tables []*Table
}

//func Init()  {
//	dbmap := ConnectMysql("test")
//	CacheDbMap.Tables = dbmap.GetTables()
//}