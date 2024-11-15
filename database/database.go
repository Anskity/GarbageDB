package database

import (
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
    jsonPath := "db.json"
    _, err := os.Stat(filepath.Join(d.RootPath, jsonPath))

    if err != os.ErrNotExist {
        return nil
    }

    file, err := os.Create(jsonPath)
    if err != nil {
        return err
    }
    
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
