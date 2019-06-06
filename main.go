/*
 * TestAPI
 *
 * Testing viability of OpenAPI2.0 for new URY API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"
	"github.com/BurntSushi/toml"
	ut "github.com/Skwunk/PrototypeAPI/utils"
)

func main() {
	log.Printf("Server started")

	config := &ut.Config{}
	_, err := toml.DecodeFile("config.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	ut.NewDatabaseConnection(config)
	ut.NewMemcacheClient(config)
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
