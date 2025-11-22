package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ujooju/lab_tester/webInterface/models"
	"github.com/ujooju/lab_tester/webInterface/storage"
)

func PostReportHandler(w http.ResponseWriter, r *http.Request) {
	record := models.TestRecord{}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = json.Unmarshal(reqBytes, &record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	err = storage.UpdateRecord(&record)
	if err != nil {
		http.Error(w, "failed to update record", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write([]byte("record updated"))
}
