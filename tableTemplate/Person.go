package tableTemplate

import "gopkg.in/mgo.v2/bson"

type PC struct {
	PcID  	bson.ObjectId `bson:"_id,omitempty"`
	Name 		string
	Surname 	string
	Username	string
	Password 	string
	PhoneNumber	string
}
