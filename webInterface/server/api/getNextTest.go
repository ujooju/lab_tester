package api

import (
	"encoding/json"
	"net/http"

	"github.com/ujooju/lab_tester/webInterface/config"
	"github.com/ujooju/lab_tester/webInterface/storage"
)

func NextTestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	agentToken := r.FormValue("agent_token")
	if agentToken != config.AgentSecret {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	nextTest, err := storage.NextTest()
	if err != nil {
		http.Error(w, "failed to get next test", http.StatusInternalServerError)
		return
	}
	nextTestBytes, err := json.Marshal(nextTest)
	if err != nil {
		http.Error(w, "failed to Marshal response", http.StatusInternalServerError)
		return
	}
	w.Write(nextTestBytes)
}
