package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ujooju/lab_tester/testRunner/cage"
	"github.com/ujooju/lab_tester/testRunner/models"
)

func RunHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	dataBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	testInfo := models.TestInfo{}
	err = json.Unmarshal(dataBytes, &testInfo)
	if err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	report, err := cage.StartTest(&testInfo, time.Second*100)
	if err != nil {
		http.Error(w, "internal error while running tests", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(report)
	if err != nil {
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(responseJSON))
}
