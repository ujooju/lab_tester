package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"code.gitea.io/sdk/gitea"
	httpcurl "github.com/ujooju/http-curl/lib"
	"github.com/ujooju/lab_tester/webInterface/config"
)

func ListForksHandler(w http.ResponseWriter, r *http.Request) {
	forks := []*gitea.Repository{}
	reqUrl := config.GiteaURL + "/api/v1/repos/" + config.CurrentTaskOwner + "/" + config.CurrentTaskName + "/forks?access_token=" + r.Context().Value("token").(string)
	response, err := httpcurl.HttpCurl(httpcurl.CurlOption{
		"-X":         httpcurl.CurlValue{"GET"},
		"--location": httpcurl.CurlValue{reqUrl},
		"-H":         httpcurl.CurlValue{"Content-Type: application/json"},
		"--tls-max":  httpcurl.CurlValue{"1.2"},
	}, time.Second*10)
	if err != nil {
		http.Error(w, "failed to get forks", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = json.Unmarshal(response, &forks)
	if err != nil {
		http.Error(w, "failed to unmarshal forks\n"+string(response), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	resp := []TaskFork{}
	for _, fork := range forks {
		taskFork := TaskFork{
			Owner:         fork.Owner.UserName,
			Name:          fork.Name,
			URL:           fork.HTMLURL,
			Status:        "to be done",
			Result:        "to be done",
			SubmitsCnt:    "tbd",
			ReportURL:     "to be done",
			ForkStatusURL: "/home/fork?owner=" + fork.Owner.UserName + "&name=" + fork.Name,
		}
		resp = append(resp, taskFork)
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write(respBytes)
}

type TaskFork struct {
	Owner         string `json:"owner"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	Result        string `json:"result"`
	URL           string `json:"url"`
	SubmitsCnt    string `json:"cnt"`
	ReportURL     string `json:"report_url"`
	ForkStatusURL string `json:"fork_status_url"`
}
