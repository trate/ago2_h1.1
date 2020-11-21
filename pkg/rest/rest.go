package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteAsJSON write data in a JSON format
func WriteAsJSON(writer http.ResponseWriter, items interface{}) error {
	data, err := json.Marshal(items)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)
	if err != nil {
		log.Print(err)
	}
	return err
}