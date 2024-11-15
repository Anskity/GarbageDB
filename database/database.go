package database

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const VERSION uint = 1

type Database struct {
	RootPath string
}

func New(rootPath string) *Database {
	return &Database{
		RootPath: rootPath,
	}
}

func (d *Database) Setup() error {
    jsonPath := filepath.Join(d.RootPath, "db.json")
    _, err := os.Stat(jsonPath)

    if !errors.Is(err, os.ErrNotExist)  {
        return nil
    }

    if err = os.Mkdir(d.RootPath, os.ModePerm); err != nil {
        return err
    }
    file, err := os.Create(jsonPath)
    if err != nil {
        return err
    }
    defer file.Close()
    
    json :=
`{
    "version": %d,
}
`

    _, err = file.WriteString(fmt.Sprintf(json, VERSION))
    if err != nil {
        return err
    }

    return nil
}
