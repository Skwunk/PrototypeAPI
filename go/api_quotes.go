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
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func AddQuote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var quote Quote
	err := decoder.Decode(&quote)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = Database.Query(`INSERT INTO quotes(text, source, suspended) VALUES ($1, $2, $3)`, quote.Text, quote.Source, quote.Suspended)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	rows, err := Database.Query("SELECT * FROM quotes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var quotes []Quote
	for rows.Next() {
		var quote Quote
		err := rows.Scan(&quote.QuoteId, &quote.Text, &quote.Source, &quote.Date, &quote.Suspended)
		if err != nil {
			log.Fatal(err)
		}
		quotes = append(quotes, quote)
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(quotes)
	w.WriteHeader(http.StatusOK)
}

func GetQuote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rows, err := Database.Query("SELECT * FROM quotes WHERE quote_id = $1", params["quote_id"])
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var quote Quote
	rows.Next()
	err = rows.Scan(&quote.QuoteId, &quote.Text, &quote.Source, &quote.Date, &quote.Suspended)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(quote)
	w.WriteHeader(http.StatusOK)
}

func UpdateQuote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var quote Quote
	err := decoder.Decode(&quote)
	if err != nil {
		log.Fatal(err)
	}

	_, err = Database.Query(`UPDATE quotes SET text = $1, source = $2, suspended = $3 WHERE quote_id = $4`, quote.Text, quote.Source, quote.Suspended, params["quote_id"])
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateQuoteDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateQuoteSource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var quote Quote
	err := decoder.Decode(&quote)
	if err != nil {
		log.Fatal(err)
	}

	_, err = Database.Query(`UPDATE quotes SET source = $1 WHERE quote_id = $2`, quote.Source, params["quote_id"])
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateQuoteSuspended(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateQuoteText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
