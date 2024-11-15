package main

import (
	"github.com/Anskity/GarbageDB/database"
)

func main() {
	db := database.New("./test")
    if err := db.Setup(); err != nil {
		panic(err)
	}

    err := db.CreateFiles(
        "users/theprimeagen/data.json",
        "users/fireship/data.json",
        "users/TJ/data.json",
        "users/Theo/data.json",
        "users/Thor/data.json",
        "users/Dreams of Code/data.json",
    )

    if err != nil {
        panic(err)
    }
}
