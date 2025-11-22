package runner

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	httpcurl "github.com/ujooju/http-curl/lib"
	"github.com/ujooju/lab_tester/testRunner/cage"
	"github.com/ujooju/lab_tester/testRunner/config"
	"github.com/ujooju/lab_tester/testRunner/models"
)

func GetNextTest() (*models.TestRecord, error) {
	reqUrl := config.LTURL + "/agent/next-test?agent_token=" + config.AgentSecret
	response, err := httpcurl.HttpCurl(httpcurl.CurlOption{
		"-X":         httpcurl.CurlValue{"GET"},
		"--location": httpcurl.CurlValue{reqUrl},
		"-H":         httpcurl.CurlValue{"Content-Type: application/json"},
		"--tls-max":  httpcurl.CurlValue{"1.2"},
	}, time.Second*10)
	if err != nil {
		return nil, err
	}

	testInfo := models.TestRecord{}
	err = json.Unmarshal(response, &testInfo)
	if err != nil {
		return nil, err
	}
	return &testInfo, nil
}

func Run() {
	for {
		test, err := GetNextTest()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 15)
			continue
		}
		if test.ID == 0 {
			time.Sleep(time.Second * 15)
			continue
		}
		log.Printf("started test %d %s %s %s\n", test.ID, test.Owner, test.RepoName, test.Branch)
		report, err := cage.StartTest(test, time.Minute*10)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 15)
			continue
		}
		err = SubmitTest(*test, &report)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 15)
			continue
		}

	}
}

func SubmitTest(testRecord models.TestRecord, report *models.Report) error {
	testRecord.Status = strconv.Itoa(report.Score)
	testRecord.Report = report.Text
	testRecordBytes, err := json.Marshal(testRecord)
	if err != nil {
		return err
	}
	reqUrl := config.LTURL + "/agent/report?agent_token=" + config.AgentSecret
	response, err := httpcurl.HttpCurl(httpcurl.CurlOption{
		"-X":         httpcurl.CurlValue{"POST"},
		"--location": httpcurl.CurlValue{reqUrl},
		"-H":         httpcurl.CurlValue{"Content-Type: application/json"},
		"--tls-max":  httpcurl.CurlValue{"1.2"},
		"-d":         httpcurl.CurlValue{string(testRecordBytes)},
	}, time.Second*10)
	if err != nil {
		return err
	}
	log.Println(string(response))
	return nil
}
