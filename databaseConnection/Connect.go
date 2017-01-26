package databaseConnection

import "gopkg.in/mgo.v2"

var session *mgo.Session

func StartConnection()  {
	var err error
	session, err = mgo.Dial("127.0.0.1");
	if err != nil {
		panic(err)
	}
	//defer session.Close()
}

func DropOldDatabase(databaseName string)  {
	err := session.DB(databaseName).DropDatabase()
		if err != nil {
			panic(err)
		}
	defer session.Close()

}
