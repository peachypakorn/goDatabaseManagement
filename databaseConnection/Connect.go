package databaseConnection

import (
	"gopkg.in/mgo.v2"
	"databaseManagement/utils"
	"fmt"
	//"databaseManagement/tableTemplate"
	"databaseManagement/tableTemplate"
)

//var session *mgo.Session

func connect() (session *mgo.Session , err error)  {
	//TODO change this to configure file
	session, err = mgo.Dial("127.0.0.1");
	if err != nil {
		return nil , err
	}
	session.SetMode(mgo.Monotonic, true)
	return session,nil
}


func DropDatabase( databaseName string)  {
	session , err := connect()
	utils.CheckError(err)
	defer session.Close()

	err = session.DB(databaseName).DropDatabase()
	utils.CheckError(err)
}

func DropAllDatabase(){
	session , err := connect()
	utils.CheckError(err)
	defer session.Close()

	names , err := session.DatabaseNames()
	utils.CheckError(err)

	for _, name := range names {
		err := session.DB(name).DropDatabase()
		utils.CheckError(err)
	}
}
func CreateNewDatabase(databaseName string){
	session , err := connect()
	utils.CheckError(err)
	defer session.Close()

	names , err := session.DatabaseNames()
	utils.CheckError(err)

	for _,name := range names{
		if name == databaseName{
			return
		}
	}
	database := session.DB(databaseName)
	fmt.Print(database.Name)
}

func AddDataToDatabase(databaseName string , collectionName string, keys []string, data []*tableTemplate.Store){
	session , err := connect()
	utils.CheckError(err)
	defer session.Close()

	collection := session.DB(databaseName).C(collectionName);

	index := mgo.Index{
		Key:        keys,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = collection.EnsureIndex(index)
	utils.CheckError(err)

	if data != nil{
		for _,store := range data{
			err = collection.Insert(store)
			utils.CheckError(err)
		}

	}


}



func ConnectToDatabase(session *mgo.Session, databaseName string)(database *mgo.Database){
	database  = session.DB(databaseName)
	return

}
