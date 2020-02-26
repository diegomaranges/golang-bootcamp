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
	ReadFile(map[string]*Items) error
	WriteFile(map[string]*Items) error
	ReturnDestinyFile() string
}

/*DestinyFile save Destiny File for all information as a string*/
type DestinyFile struct {
	destinyFile string
}

const fileType = ".json"

/*CreateNewFInstance create new instance of the object*/
func CreateNewFInstance() *DestinyFile {
	return &DestinyFile{}
}

/*SetFile save new route file without specific the type*/
func (d *DestinyFile) SetFile(destiny string) {
	d.destinyFile = destiny + fileType
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

	ioutil.WriteFile(d.destinyFile, jsonString, os.ModePerm)
	return nil
}

/*ReturnDestinyFile return a route string saved*/
func (d *DestinyFile) ReturnDestinyFile() string {
	return d.destinyFile
}
