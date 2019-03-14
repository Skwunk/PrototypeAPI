/*
 * TestAPI
 *
 * Testing viability of OpenAPI2.0 for new URY API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package testapi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/Skwunk/TestAPI/1.0.0/",
		Index,
	},

	Route{
		"AddQuote",
		strings.ToUpper("Post"),
		"/Skwunk/TestAPI/1.0.0/quotes",
		AddQuote,
	},

	Route{
		"GetAllQuotes",
		strings.ToUpper("Get"),
		"/Skwunk/TestAPI/1.0.0/quotes",
		GetAllQuotes,
	},

	Route{
		"GetQuote",
		strings.ToUpper("Get"),
		"/Skwunk/TestAPI/1.0.0/quotes/{quote_id}",
		GetQuote,
	},

	Route{
		"UpdateQuote",
		strings.ToUpper("Put"),
		"/Skwunk/TestAPI/1.0.0/quotes/{quote_id}",
		UpdateQuote,
	},

	Route{
		"UpdateQuoteDate",
		strings.ToUpper("Put"),
		"/Skwunk/TestAPI/1.0.0/quotes/{quote_id}/date",
		UpdateQuoteDate,
	},

	Route{
		"UpdateQuoteSource",
		strings.ToUpper("Put"),
		"/Skwunk/TestAPI/1.0.0/quotes/{quote_id}/source",
		UpdateQuoteSource,
	},

	Route{
		"UpdateQuoteSuspended",
		strings.ToUpper("Put"),
		"/Skwunk/TestAPI/1.0.0/quotes/{quote_id}/suspended",
		UpdateQuoteSuspended,
	},

	Route{
		"UpdateQuoteText",
		strings.ToUpper("Put"),
		"/Skwunk/TestAPI/1.0.0/quotes/{quote_id}/text",
		UpdateQuoteText,
	},
}