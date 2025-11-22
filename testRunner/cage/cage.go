package cage

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/ujooju/lab_tester/testRunner/config"
	"github.com/ujooju/lab_tester/testRunner/models"
)

func RunScript(ctx context.Context, testInfo *models.TestRecord) (models.Report, error) {
	//в скрипт первым параметром передаётся url для клонирования репозитория
	//url сразу с токеном
	//вторым параметром ветка
	//script.sh <url> <branch>

	buf := bytes.NewBufferString("")
	cmd := exec.CommandContext(ctx,
		//1 arg. script name
		config.ScriptLocation+"/"+config.ScriptName,
		//2 arg. clone url
		fmt.Sprintf( //making clone url
			"%s://%s:%s@%s/%s/%s.git",
			config.GitURLProtoName,
			config.CheckerName,
			config.CheckerToken,
			config.GitURLHostName,
			testInfo.Owner,
			testInfo.RepoName,
		),
		//3 arg. branch name
		testInfo.Branch,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	if err != nil {
		return models.Report{}, err
	}
	cmd.Wait()

	report := models.Report{}
	reportFile, err := os.Open("report.txt")
	reportFileBuf := bufio.NewReader(reportFile)
	if err != nil {
		return models.Report{}, err
	}

	statusString, err := reportFileBuf.ReadString('\n')
	statusString = strings.TrimSpace(statusString)
	if err != nil {
		return models.Report{}, err
	}

	score, err := strconv.Atoi(statusString)
	if err != nil {
		return models.Report{}, err
	}

	reportText, err := io.ReadAll(reportFileBuf)
	if err != nil {
		return models.Report{}, err
	}

	report.Text = string(reportText)
	report.Score = score

	err = os.RemoveAll(testInfo.RepoName)
	if err != nil {
		fmt.Println(err)
	}

	return report, nil
}

func StartTest(testInfo *models.TestRecord, timeout time.Duration) (models.Report, error) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	report, err := RunScript(ctx, testInfo)
	if err != nil {
		return models.Report{}, err
	}
	return report, nil
}
