package databaseConnection

import "gopkg.in/mgo.v2"

//var session *mgo.Session

func StartConnection() (session *mgo.Session)  {
	var err error
	session, err = mgo.Dial("127.0.0.1");
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	return
}


func DropOldDatabase(session *mgo.Session, databaseName string)  {
	err := session.DB(databaseName).DropDatabase()
		if err != nil {
			panic(err)
		}
	defer session.Close()

}

func ConnectToDatabase(session *mgo.Session, databaseName string)(database *mgo.Database){
	database  = session.DB(databaseName)
	return
}
