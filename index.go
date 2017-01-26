package main

import "databaseManagement/databaseConnection"

func main(){
	databaseConnection.StartConnection();
	databaseConnection.DropOldDatabase("test2");
}