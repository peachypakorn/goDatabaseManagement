package autofillDB

import (
	"encoding/csv"
	"io"
	"log"
	"fmt"
	"io/ioutil"
	"databaseManagement/utils"
	"strings"
)

func ReadFile(filepath string)(readedFile string){
	file, err := ioutil.ReadFile(filepath)
	utils.CheckError(err)
	//fmt.Print(string(file))
	readedFile = string(file)
	return
}

func ReadStore(readedFile string){
	reader := csv.NewReader(strings.NewReader(readedFile))
	reader.Comma = ';'
	reader.Comment = '#'
	reader.Read()
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}

