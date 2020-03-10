package mongodb

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*Item is a struct used for read element from the Json requests*/
type Item struct {
	ID       string `bson:"_id,omitempty"`
	Title    string `json:"title"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}

/*MongoDb interface*/
type MongoDb interface {
	//loadbackup
	AddItem(item Item) error
	ReturnItem(id string) (Item, error)
	//returnall
	UpdateItem(id string, item Item) error
	DeleteItem(id string) error
	//savedata
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

/*AddItem create new mongodb document, Return an error if:

function can not access to mongo db

does not exist the element*/
func (m *MongoStruct) AddItem(item Item) error {
	numOfElements, err := m.collection.Find(bson.M{"_id": item.ID}).Count()
	if err != nil {
		return err
	}

	if numOfElements == 1 {
		return m.collection.Update(bson.M{"_id": item.ID}, item)
	}

	return m.collection.Insert(item)
}

/*ReturnItem Return an error if:

function can not access to mongo db

does not exist the element*/
func (m *MongoStruct) ReturnItem(id string) (Item, error) {
	result := &Item{}
	err := m.collection.Find(bson.M{"_id": id}).One(&result)
	return *result, err
}

/*UpdateItem Return an error if:

function can not access to mongo db

does not exist the element

or id is not the same to item.id*/
func (m *MongoStruct) UpdateItem(id string, item Item) error {
	numOfElements, err := m.collection.Find(bson.M{"_id": id}).Count()
	if err != nil {
		return err
	}

	if numOfElements == 0 {
		return errors.New("not found")
	}

	return m.collection.Update(bson.M{"_id": id}, item)
}

/*DeleteItem Return an error if:

function can not access to mongo db

does not exist the element*/
func (m *MongoStruct) DeleteItem(id string) error {
	return m.collection.Remove(bson.M{"_id": id})
}
