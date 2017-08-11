package main

import (
	"databaseManagement/databaseConnection"
	"databaseManagement/databaseUtil"
)

func main(){
	//databaseConnection.DropAllDatabase();
	databaseConnection.DropDatabase("test2");
	file := databaseUtil.ReadFile("asset/store.csv")
	//stores := databaseUtil.UnmarshalStore(file)
	databaseConnection.AddDataToDatabase("pcwutl","stores",[]string{"StoreID"},stores)
	//databaseUtil.FillStruct(store)
	//fmt.Print(time.Weekday(1))
}