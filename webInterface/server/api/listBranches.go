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

func ListBranchesHandler(w http.ResponseWriter, r *http.Request) {
	branches := []*gitea.Branch{}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse url values", http.StatusBadRequest)
		log.Println(err)
		return
	}
	forkOwner := r.FormValue("owner")
	forkName := r.FormValue("name")
	reqUrl := config.GiteaURL + "/api/v1/repos/" + forkOwner + "/" + forkName + "/branches?access_token=" + r.Context().Value("token").(string)
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
	err = json.Unmarshal(response, &branches)
	if err != nil {
		http.Error(w, "failed to unmarshal repsonse", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	result := []BranchInfo{}
	var branchInfo BranchInfo
	for _, branch := range branches {
		branchInfo = BranchInfo{
			Name:      branch.Name,
			SubmitURL: "/api/submit?owner=" + forkOwner + "&name=" + forkName + "&branch=" + branch.Name,
		}
		result = append(result, branchInfo)
	}
	resultBytes, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write(resultBytes)
}

type BranchInfo struct {
	Name      string `json:"name"`
	SubmitURL string `json:"submit_url"`
}
