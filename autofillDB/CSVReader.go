package autofillDB

import (
	"databaseManagement/utils"
	"databaseManagement/tableTemplate"
	"os"
	"github.com/gocarina/gocsv"
	"fmt"
)

func ReadFile(filepath string)(file *os.File){
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	utils.CheckError(err)
	//defer file.Close()
	return file
}

func UnmarshalStore(file *os.File)(stores []*tableTemplate.Store){
	err := gocsv.UnmarshalFile(file,&stores)

	utils.CheckError(err)
	for _, store := range stores {
		fmt.Println("Hello", store.StoreName)
	}
	return
}

func FillStruct(stores []*tableTemplate.Store)(){

	return
}

