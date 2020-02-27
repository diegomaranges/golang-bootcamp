package fileinteraction

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

/*Items algo*/
type Items struct {
	Title    string `json:"title"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}

/*FileActions have all function for handle the interecctions with files*/
type FileActions interface {
	SetFile(string)
	CreateFile() error
	ReadFile(map[string]*Items) error
	WriteFile(map[string]*Items) error
	DeleteFile() error
	ReturnDestinyFile() string
}

/*DestinyFile save Destiny File for all information as a string*/
type DestinyFile struct {
	destinyFile string
}

const fileType = ".json"
const fileName = "cars/db"

/*CreateNewFInstance create new instance and save the destiny file*/
func CreateNewFInstance(dbID string) *DestinyFile {
	destiny := &DestinyFile{}
	destiny.destinyFile = fileName + dbID + fileType
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
Pre: externalMap different to nil*/
func (d *DestinyFile) ReadFile(externalMap map[string]*Items) error {
	if externalMap == nil {
		return errors.New("map is not initialized")
	}

	fileBytes, err := ioutil.ReadFile(d.destinyFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileBytes, &externalMap); err != nil {
		return err
	}

	return nil
}

/*WriteFile write information in file saved
Pre: externalMap different to nil*/
func (d *DestinyFile) WriteFile(externalMap map[string]*Items) error {
	if externalMap == nil {
		return errors.New("map is not initialized")
	}

	jsonString, err := json.Marshal(externalMap)
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
