package autofillDB

import (
	"path/filepath"
	"os"
	"databaseManagement/tableTemplate"
)

func ReadCSVFromFile(filepath string)(readedFile *os.File){
	readedFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer readedFile.Close()
	return
}

func ConvertToStoreStruct(readedFile *os.File)(stores []*tableTemplate.Store){
	stores = []*tableTemplate.Store{}
	err := UnmarshalFile(clientsFile, &clients)
}

if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
panic(err)
}
for _, client := range clients {
fmt.Println("Hello", client.Name)
}
