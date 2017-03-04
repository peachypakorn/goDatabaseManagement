package tableTemplate

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"databaseManagement/utils"
	"fmt"

	"strings"
	"strconv"

)
type Store struct {
	StoreID  	bson.ObjectId `csv:"-",bson:"_id,omitempty"`
	StoreName 	string `csv:"StoreID"`
	Branch    	string `csv:"StoreName"`
	Phone     	string `csv:"Branch"`
	City		string `csv:"City"`
	Province	string `csv:"Province"`
	EnteringDay	Weekday `csv:"EnteringDay"`
	DayOff		Weekday `csv:"DayOff"`
	StartTime	timesHHMM `csv:"StartTime"`
	EndTime 	timesHHMM `csv:"EndTime"`
	SellUnza	bool `csv:"sellUnza"`
	SellBio		bool `csv:"sellBio"`
	LastAddStock	time.Time `csv:"-"`
}
type Weekday struct {
	days []time.Weekday
}
// Convert the CSV string as internal date
func (weekdays Weekday) UnmarshalCSV(csv string) (err error) {
	//fmt.Print(csv)
	dates := strings.Split(csv,"-")
	for _,date := range dates{
		numberDate,err := strconv.Atoi(date)
		utils.CheckError(err)
		if(numberDate>7){
			numberDate = 7
		}
		weekdays.days = append(weekdays.days,time.Weekday(numberDate))
	}
	return nil
}

type timesHHMM struct {
	time.Time
}
//Mon Jan 2 15:04:05 -0700 MST 2006
func (times timesHHMM) UnmarshalCSV(csv string) (err error) {
	if len(csv) <=2 {
		csv += ".00"
	} else if len(csv)==3{
		csv+= "0"
	}
	tempTime,err := strconv.ParseFloat(csv,32)
	utils.CheckError(err)
	if tempTime >= 13{
		csv+="pm"
	} else {
		csv+="am"
	}
	times.Time, err = time.Parse("11:04pm", csv)
	fmt.Print(csv+"  "+times.Time.String())
	return nil
}