package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

// Respond формирует и отправляет ответ в формате JSON
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Print(err)
	}
}

// Message формирует тело сообщения об ошибке
func Message(id string, message string) map[string]interface{} {
	return map[string]interface{}{
		"id":      id,
		"message": message,
	}
}
