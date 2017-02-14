package main

import (

	"databaseManagement/autofillDB"
	"time"
	"fmt"
)

func main(){
	//databaseConnection.StartConnection();
	//databaseConnection.DropOldDatabase("test2");
	reader := autofillDB.ReadFile("/home/peachy/gopath/src/databaseManagement/asset/store.csv")
	data := autofillDB.ReadStore(reader)
	autofillDB.FillStruct(data)
	fmt.Print(time.Weekday(1))
}