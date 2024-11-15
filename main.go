package main

import (
	"github.com/Anskity/GarbageDB/database"
)

func main() {
    db := database.New("./test")
    err := db.Setup()
    if err != nil {
        panic(err)
    }
}
