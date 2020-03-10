package mongodb

import (
	"errors"

	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db/mongodb/fileinteraction"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*MongoDb interface*/
type MongoDb interface {
	LoadBackUp() error
	AddItem(item fileinteraction.Item) error
	ReturnItem(id string) (fileinteraction.Item, error)
	ReturnAllItems() (fileinteraction.Item, error)
	UpdateItem(id string, item fileinteraction.Item) error
	DeleteItem(id string) error
	GenerateBackUp() error
}

/*MongoStruct MongoDb Sctruct*/
type MongoStruct struct {
	file       *fileinteraction.DestinyFile
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
func CreateNewDBInstance(directory string, carID string) *MongoStruct {
	mongodb := &MongoStruct{}
	mongodb.file = fileinteraction.CreateNewFInstance(directory, carID)
	mongodb.collection = getSession().DB("CarApi").C("Car" + carID)
	return mongodb
}

/*LoadBackUp algo*/
func (m *MongoStruct) LoadBackUp() error {
	items := &fileinteraction.Items{}
	if err := m.file.ReadFile(items); err != nil {
		return err
	}

	if err := m.collection.DropCollection(); err != nil {
		return err
	}

	for _, item := range *items {
		if err := m.collection.Insert(item); err != nil {
			return err
		}
	}

	return nil
}

/*AddItem create new mongodb document, Return an error if:

function can not access to mongo db

does not exist the element*/
func (m *MongoStruct) AddItem(item fileinteraction.Item) error {
	numOfElements, err := m.collection.Find(bson.M{"_id": item.ID}).Count()
	if err != nil {
		return err
	}

	if numOfElements == 1 {
		item.Quantity++
		return m.collection.Update(bson.M{"_id": item.ID}, item)
	}

	return m.collection.Insert(item)
}

/*ReturnItem Return an error if:

function can not access to mongo db

does not exist the element*/
func (m *MongoStruct) ReturnItem(id string) (fileinteraction.Item, error) {
	result := &fileinteraction.Item{}
	err := m.collection.Find(bson.M{"_id": id}).One(&result)
	return *result, err
}

/*ReturnAllItems Return an error if:

function can not access to mongo db*/
func (m *MongoStruct) ReturnAllItems() (fileinteraction.Items, error) {
	var items fileinteraction.Items
	err := m.collection.Find(nil).All(&items)
	return items, err
}

/*UpdateItem Return an error if:

function can not access to mongo db

does not exist the element

or id is not the same to item.id*/
func (m *MongoStruct) UpdateItem(id string, quantity int) error {
	if quantity <= 0 {
		return errors.New("Wrong quantity imput")
	}

	myItem := &fileinteraction.Item{}
	err := m.collection.Find(bson.M{"_id": id}).One(&myItem)
	if err != nil {
		return err
	}

	if myItem == nil {
		return errors.New("not found")
	}

	myItem.Quantity = quantity
	return m.collection.Update(bson.M{"_id": id}, myItem)
}

/*DeleteItem Return an error if:

function can not access to mongo db

does not exist the element*/
func (m *MongoStruct) DeleteItem(id string) error {
	return m.collection.Remove(bson.M{"_id": id})
}

/*GenerateBackUp algo*/
func (m *MongoStruct) GenerateBackUp() error {
	items := &fileinteraction.Items{}

	if err := m.collection.Find(nil).All(&items); err != nil {
		return err
	}

	return m.file.WriteFile(*items)
}
