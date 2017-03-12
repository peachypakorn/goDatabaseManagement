package main

import (
	"databaseManagement/autofillDB"
	"databaseManagement/databaseConnection"
)

func main(){
	databaseConnection.DropAllDatabase();
	//databaseConnection.DropOldDatabase("test2");
	file := autofillDB.ReadFile("/home/peachy/gopath/src/databaseManagement/asset/store.csv")
	stores := autofillDB.UnmarshalStore(file)
	//stores = autofillDB.
	databaseConnection.AddDataToDatabase("pcwutl","stores",[]string{"StoreID"},stores)
	//autofillDB.FillStruct(store)
	//fmt.Print(time.Weekday(1))
}