package main

import (
	// Standard library packages
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	// Third party packages
	"github.com/gauthierbl/go-rest-service-example/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	// User resource
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.RemoveUser)

	// test url
	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Test!\n")
	})

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}

// Connects to mongodb
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
