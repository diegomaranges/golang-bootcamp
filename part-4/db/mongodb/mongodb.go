package mongodb

import (
	"errors"
	"fmt"
	"strings"

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
func CreateNewDBInstance(carID string, newDB bool) (*MongoStruct, error) {
	mongodb := &MongoStruct{}
	db := getSession().DB("CarApi")

	colections, err := db.CollectionNames()
	if err != nil {
		return mongodb, err
	}

	if newDB {
		for _, colection := range colections {
			if strings.Compare(colection, "Car"+carID) == 0 {
				return mongodb, errors.New("db is already exist")
			}
		}
		mongodb.collection = db.C("Car" + carID)
		return mongodb, nil
	}

	exist := false
	for _, colection := range colections {
		if strings.Compare(colection, "Car"+carID) == 0 {
			exist = true
			break
		}
	}

	if !exist {
		return mongodb, errors.New("db does not exist")
	}
	mongodb.collection = db.C("Car" + carID)
	return mongodb, nil
}

/*LoadBackUp algo
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
}*/

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
func GenerateBackUp() error {
	db := getSession().DB("CarApi")
	items := &fileinteraction.Items{}
	colections, err := db.CollectionNames()
	if err != nil {
		return err
	}
	for _, colection := range colections {
		file := fileinteraction.CreateNewFInstance("backup/", colection)
		if er := db.C(colection).Find(nil).All(items); err != nil {
			fmt.Println(colection)
			return er
		}
		if er := file.WriteFile(*items); er != nil {
			fmt.Println(colection)
			return er
		}
	}

	return nil
}

/*DeleteColection algo*/
func (m *MongoStruct) DeleteColection() error {
	return m.collection.DropCollection()
}
