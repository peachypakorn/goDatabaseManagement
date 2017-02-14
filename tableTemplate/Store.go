package tableTemplate

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Store struct {
	StoreID  	bson.ObjectId `bson:"_id,omitempty"`
	StoreName 	string
	Branch    	string
	Phone     	string
	City		string
	Province	string
	EnteringDay	[]time.Weekday
	DayOff		[]time.Weekday
	StartTime	time.Time
	Duration 	int
	SellUnza	bool
	SellBio		bool
	LastAddStock	time.Time
}