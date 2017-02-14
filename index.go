package main

import (

	"databaseManagement/autofillDB"
)

func main(){
	//databaseConnection.StartConnection();
	//databaseConnection.DropOldDatabase("test2");
	reader := autofillDB.ReadFile("/home/peachy/gopath/src/databaseManagement/asset/store.csv")
	autofillDB.ReadStore(reader)

}