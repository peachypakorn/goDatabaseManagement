package autofillDB

import (
	"databaseManagement/utils"
	"databaseManagement/tableTemplate"
	"os"
	"github.com/gocarina/gocsv"
	"strings"
	"gopkg.in/mgo.v2/bson"
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
		store.StoreName = strings.TrimSpace(store.StoreName)
		store.StoreID = bson.NewObjectId()
		//res2B, _ := json.Marshal(store)
		//fmt.Println(string(res2B))
		//fmt.Println(store.EnteringDay)
	}
	return
}

func FillStruct(stores []*tableTemplate.Store)(){

	return
}

