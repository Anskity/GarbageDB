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

func (d *Database) CreateFile(path string) error {
    path = filepath.Join(d.RootPath, path)
    dirPath, _ := filepath.Split(path)

    _, err := os.Stat(dirPath);
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            err = os.MkdirAll(dirPath, os.ModePerm)
            if err != nil {
                return err
            }
        } else {
            return err
        }
    }

    _, err = os.Create(path)
    if err != nil {
        return err
    }

    return nil
}

func (d *Database) CreateFiles(paths... string) error {
    for _, path := range paths {
        err := d.CreateFile(path)

        if err != nil {
            return err
        }
    }
    return nil
}

type PathStatus int
const (
    PathDoesntExist PathStatus = iota
    PathIsFile
    PathIsDir
)
func (d *Database) GetPathStatus(path string) PathStatus {
    path = filepath.Join(d.RootPath, path)
    stat, err := os.Stat(path)
    if os.IsNotExist(err) {
        return PathDoesntExist
    }

    if stat.IsDir() {
        return PathIsDir
    }

    return PathIsFile
}

