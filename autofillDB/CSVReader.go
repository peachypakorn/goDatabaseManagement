package autofillDB

import (
	"encoding/csv"
	"io/ioutil"
	"databaseManagement/utils"
	"strings"
	"databaseManagement/tableTemplate"

)

func ReadFile(filepath string)(readedFile string){
	file, err := ioutil.ReadFile(filepath)
	utils.CheckError(err)
	//fmt.Print(string(file))
	readedFile = string(file)
	return
}

func ReadStore(readedFile string)(data *csv.Reader){
	data = csv.NewReader(strings.NewReader(readedFile))
	data.Comma = ';'
	data.Comment = '#'
	data.TrimLeadingSpace = true
	//just read header
	data.Read()
	return
}

func FillStruct(data *csv.Reader)(stores []tableTemplate.Store){
	//store := new(tableTemplate.Store)

	//rs, _ := csv.(data, store)
	//for rs.Get() {
	//	if date.Year == p.Joined.Year {
	//		fmt.Println(p)
	//	}
	//}
	return
}

