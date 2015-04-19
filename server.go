package main

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"net/http"

	// Third party packages
	"github.com/gauthierbl/go-rest-service-example/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// GET a user
	r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		user := models.User{
			Name:   "Brandon G",
			Gender: "male",
			Age:    30,
			ID:     p.ByName("id"),
		}

		// Marshal to JSON
		userAsJSON, _ := json.Marshal(user)

		//Write content-type, statuscode and payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, string(userAsJSON))
	})

	r.GET("/blg", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "blg Test!\n")
	})

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}
