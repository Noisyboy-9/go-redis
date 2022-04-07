package database

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"regexp"
)

// Database holds every database's data
type Database struct {
	Name       string
	StoredData map[string]string
}

func New(name string) *Database {
	return &Database{
		Name:       name,
		StoredData: make(map[string]string),
	}
}

// SetValue persists a key-value pair to database
func (db *Database) SetValue(key string, value string) {
	db.StoredData[key] = value
}

// GetValueByKey gets value of given key if exist
func (db *Database) GetValueByKey(key string) (value string, err error) {
	if value, exist := db.StoredData[key]; exist {
		return value, nil
	}

	return "", errors.New(fmt.Sprintf("Given key:%s doesn't exist", key))
}

// DeleteByKey deletes the value of the given key if exit
func (db *Database) DeleteByKey(key string) error {
	if _, exist := db.StoredData[key]; exist {
		delete(db.StoredData, key)
		return nil
	}

	return errors.New(fmt.Sprintf("Given key:%s doesn't exist", key))
}

// KeysMatchPattern returns all keys matching the given regex pattern
func (db *Database) KeysMatchPattern(pattern string) (keys []string, compileError error) {
	for key, _ := range db.StoredData {
		matchString, err := regexp.MatchString(pattern, key)
		if err != nil {
			return nil, err
		}

		if matchString {
			keys = append(keys, key)
		}
	}

	return keys, nil
}

// SaveToFile saves the database to given io.Writer
func SaveToFile(db *Database, writer io.Writer) error {
	encoder := gob.NewEncoder(writer)
	return encoder.Encode(*db)
}

// ReadFromFile reads the database from the given reader
func ReadFromFile(reader io.Reader) (*Database, error) {
	db := New("")
	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(db); err != nil {
		return nil, err
	}

	return db, nil
}
