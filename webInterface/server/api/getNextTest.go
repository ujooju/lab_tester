package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ujooju/lab_tester/webInterface/storage"
)

func NextTestHandler(w http.ResponseWriter, r *http.Request) {
	nextTest, err := storage.NextTest()
	if err != nil {
		http.Error(w, "failed to get next test", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	nextTestBytes, err := json.Marshal(nextTest)
	if err != nil {
		http.Error(w, "failed to Marshal response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write(nextTestBytes)
}
