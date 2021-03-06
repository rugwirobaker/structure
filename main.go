//Package structure is a a map for web application structure
package main

import (
	"fmt"
	"log"

	"github.com/rugwirobaker/structure/app"
)

func main() {
	fmt.Println("*** Loading application configuration...")
	file, err := app.LoadConfig(app.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	conf := app.NewAppConfig(file)

	//log.Println(conf)

	a := &app.App{Conf: conf}
	a.Initialize()
	defer a.DB.Close()
	a.Run(":8080")

}

//TODO: move handlers to own package --> DONE
//TODO: improve error reporting
//--JSONIFY handler errors
//--Errors as middleware
//--Handler model/database errors
//TODO: add caching middleware
//TODO: add users and authentication

//LEARNED: function wrapping
//CREATE: a basic logging package
