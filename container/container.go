package container

import (
	"github.com/Noisyboy-9/go-redis/database"
)

// Container is a dependency injector for the app
type Container struct {
	AllDatabases    []*database.Database
	CurrentDatabase *database.Database
}

func New() *Container {
	return &Container{
		AllDatabases:    nil,
		CurrentDatabase: nil,
	}
}

// AddDatabase adds the given list of database
func (c *Container) AddDatabase(name string) {
	db := database.New(name)
	c.AllDatabases = append(c.AllDatabases, db)
}

// DatabaseExist checks if the given database exist or not
func (c *Container) DatabaseExist(name string) (bool, *database.Database) {
	for _, db := range c.AllDatabases {
		if db.Name == name {
			return true, db
		}
	}

	return false, nil
}

func (c *Container) GetOrCreateDatabaseByName(name string) *database.Database {
	if exist, db := c.DatabaseExist(name); exist {
		return db
	}

	// 	 database doesn't exist, create a new one
	newDb := database.New(name)
	c.AllDatabases = append(c.AllDatabases, newDb)
	return newDb
}

func (c *Container) GetAllDatabases() []*database.Database {
	return c.AllDatabases
}

func (c *Container) UpdateDatabaseContents(name string, db *database.Database) {
	
}
