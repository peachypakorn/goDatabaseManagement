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
	Days []time.Weekday
}
// Convert the CSV string as internal date
func (weekdays *Weekday) UnmarshalCSV(csv string) (err error) {
	//fmt.Print(csv)
	dates := strings.Split(csv,"-")

	for _,date := range dates{
		if date != "" {
			numberDate,err := strconv.Atoi(date)
			utils.CheckError(err)
			if(numberDate>=7){
				numberDate = 0
			}
			weekdays.Days = append(weekdays.Days,time.Weekday(numberDate))

		}
	}
	//fmt.Print(weekdays.Days)
	return nil
}

type timesHHMM struct {
	time.Time
}
//Mon Jan 2 15:04:05 -0700 MST 2006
func (times *timesHHMM) UnmarshalCSV(csv string) (err error) {
	if len(csv) <=2 {
		csv += ".00"
	} else if len(csv)==3{
		csv+= "0"
	}
	value := strings.Split(csv,".")
	hour,err := strconv.Atoi(value[0])
	minute,err := strconv.Atoi(value[1])
	utils.CheckError(err)
	times.Time = times.Time.Add(time.Hour*time.Duration(hour)+time.Minute*time.Duration(minute))
		fmt.Print(csv+"  "+times.Time.String())
	return nil
}

