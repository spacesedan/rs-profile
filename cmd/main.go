package main

import (
	"github.com/spacesedan/profile-tracker/internal/repo"
	"log"
)

func main() {
	db, err := repo.NewMongo()
	if err != nil {
		log.Fatalln(err)
	}

	app, err := inject(db)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(app.Run())

}
