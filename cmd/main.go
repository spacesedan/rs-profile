package main

import (
	"github.com/spacesedan/profile-tracker/internal/datastores"
	"log"
)

func main() {
	if err := configEnv(); err != nil {
		log.Fatalln(err)
	}

	dao, err := datastores.NewDAO()
	if err != nil {
		log.Fatalln(err)
	}

	app, err := inject(dao)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("TODO: CHANGE ASSETS AND COLLECTIONS ENDPOINT TO USE RESERVOIR API")

	log.Fatalln(app.Run())

}
