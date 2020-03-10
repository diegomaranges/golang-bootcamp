package fileinteraction

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

/*Items algo*/
type Items []Item

/*Item is a struct used for read element from the Json requests*/
type Item struct {
	ID       string `bson:"_id,omitempty"`
	Title    string `json:"title"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}

/*FileActions have all function for handle the interecctions with files*/
type FileActions interface {
	CreateFile() error
	ReadFile(Items) error
	WriteFile(Items) error
	DeleteFile() error
	ReturnDestinyFile() string
}

/*DestinyFile save Destiny File for all information as a string*/
type DestinyFile struct {
	destinyFile string
}

const fileType = ".json"
const fileName = "db"

/*CreateNewFInstance create new instance and save the destiny file

destinyFile = directory + db name + ID + Default type*/
func CreateNewFInstance(directory string, dbID string) *DestinyFile {
	destiny := &DestinyFile{}
	destiny.destinyFile = directory + fileName + dbID + fileType
	return destiny
}

/*CreateFile reed information from file saved

Warning: If the map element have some information it will delete

Pre: externalMap different to nil*/
func (d *DestinyFile) CreateFile() error {
	_, err := ioutil.ReadFile(d.destinyFile)
	if err == nil {
		return errors.New("File directory already exist")
	}

	return ioutil.WriteFile(d.destinyFile, nil, os.ModePerm)
}

/*ReadFile reed information from file saved

Warning: If the map element have some information it will delete

Pre: externalStruct different to nil*/
func (d *DestinyFile) ReadFile(externalStruct *Items) error {
	if externalStruct == nil {
		return errors.New("map is not initialized")
	}

	fileBytes, err := ioutil.ReadFile(d.destinyFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileBytes, &externalStruct); err != nil {
		return err
	}

	return nil
}

/*WriteFile write information in file saved

Pre: externalStruct different to nil*/
func (d *DestinyFile) WriteFile(externalStruct Items) error {
	if externalStruct == nil {
		return errors.New("map is not initialized")
	}

	jsonString, err := json.Marshal(externalStruct)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(d.destinyFile, jsonString, os.ModePerm)
}

/*DeleteFile return a error if the dile does not exist*/
func (d *DestinyFile) DeleteFile() error {
	return os.Remove(d.destinyFile)
}

/*ReturnDestinyFile return a route string saved*/
func (d *DestinyFile) ReturnDestinyFile() string {
	return d.destinyFile
}
