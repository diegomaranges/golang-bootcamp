package mongodb

import (
	"gopkg.in/mgo.v2"
)

/*Item is a struct used for read element from the Json requests*/
type Item struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}

/*MongoDb interface*/
type MongoDb interface {
	AddItem(item Item) error
}

/*MongoStruct MongoDb Sctruct*/
type MongoStruct struct {
	collection *mgo.Collection
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

/*CreateNewDBInstance create a new mongo struncture*/
func CreateNewDBInstance(id string) *MongoStruct {
	mongodb := &MongoStruct{}
	mongodb.collection = getSession().DB("CarApi").C("Car" + id)
	return mongodb
}

/*AddItem create new mongodb document*/
func (m *MongoStruct) AddItem(item Item) error {

	return m.collection.Insert(item)

}
