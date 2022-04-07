package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Noisyboy-9/go-redis/container"
	"github.com/Noisyboy-9/go-redis/database"
)

// Parser is responsible for handling program loop and also getting cli commands and delegating to corresponding entity
type Parser struct {
	app    *container.Container
	reader *bufio.Reader
}

func New(c *container.Container) *Parser {
	return &Parser{
		reader: bufio.NewReader(os.Stdin),
		app:    c,
	}
}

func (p *Parser) StartProgramLoop() error {
	for {
		input, err := p.reader.ReadString('\n')
		if err != nil {
			return err
		}

		if err := p.parse(strings.TrimSuffix(input, "\n")); err != nil {
			return err
		}
	}
}

func (p *Parser) parse(input string) error {
	cmd := strings.Split(input, " ")[0]

	switch cmd {
	case "set":
		if p.app.CurrentDatabase == nil {
			return errors.New("no database selected")
		}
		key := strings.Split(input, " ")[1]
		value := strings.Split(input, " ")[2]
		p.app.CurrentDatabase.SetValue(key, value)

	case "get":
		if p.app.CurrentDatabase == nil {
			return errors.New("no database selected")
		}
		key := strings.Split(input, " ")[1]
		value, err := p.app.CurrentDatabase.GetValueByKey(key)
		if err != nil {
			return err
		}
		fmt.Println(value)

	case "del":
		if p.app.CurrentDatabase == nil {
			return errors.New("no database selected")
		}
		key := strings.Split(input, " ")[1]
		err := p.app.CurrentDatabase.DeleteByKey(key)
		if err != nil {
			return err
		}

	case "keys":
		if p.app.CurrentDatabase == nil {
			return errors.New("no database selected")
		}
		pattern := strings.Split(input, " ")[1]
		keys, err := p.app.CurrentDatabase.KeysMatchPattern(pattern)
		if err != nil {
			return err
		}
		fmt.Println(keys)

	case "use":
		dbName := strings.Split(input, " ")[1]

		p.app.CurrentDatabase = p.app.GetOrCreateDatabaseByName(dbName)

		fmt.Println("")

	case "list":
		databases := p.app.GetAllDatabases()
		for _, db := range databases {
			fmt.Println(db.Name)
		}

	case "dump":
		if p.app.CurrentDatabase == nil {
			return errors.New("no database selected")
		}
		filePath := strings.Split(input, " ")[1]
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		err = database.SaveToFile(p.app.CurrentDatabase, file)
		if err != nil {
			return err
		}

	case "load":
		filePath := strings.Split(input, " ")[1]
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		db, err := database.ReadFromFile(file)
		if err != nil {
			return err
		}

		exist, oldDb := p.app.DatabaseExist(db.Name)
		if exist {
			oldDb.StoredData = db.StoredData
		}

		p.app.CurrentDatabase = db

	case "exit":
		fmt.Println("hope you have enjoyed!")
		os.Exit(0)

	default:
		return errors.New("invalid command")
	}

	return nil
}
