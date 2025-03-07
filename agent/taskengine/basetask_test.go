package taskengine

import (
	"encoding/base64"
	"math/rand"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/aliyun/aliyun_assist_client/agent/taskengine/models"
	"github.com/aliyun/aliyun_assist_client/agent/util"
	"github.com/aliyun/aliyun_assist_client/agent/util/osutil"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func addMockServer() {
	httpmock.Activate()

	httpmock.RegisterResponder("GET", "http://100.100.100.200/latest/meta-data/region-id",
		httpmock.NewStringResponder(200, `cn-test`))

	httpmock.RegisterResponder("GET", "https://cn-test.axt.aliyun.com/luban/api/connection_detect",
		httpmock.NewStringResponder(200, `ok`))

	httpmock.RegisterResponder("POST", "https://cn-test.axt.aliyun.com/luban/api/v1/task/finish",
		httpmock.NewStringResponder(200, `ok`))

	httpmock.RegisterResponder("POST", "https://cn-test.axt.aliyun.com/luban/api/v1/task/running",
		httpmock.NewStringResponder(200, `ok`))
}

func removeMockServer() {
	httpmock.DeactivateAndReset()
}

func TestRunTask(t *testing.T) {
	util.NilRequest.Set()
	defer util.NilRequest.Clear()
	addMockServer()
	defer removeMockServer()
	guard := gomonkey.ApplyFunc(util.GetServerHost, func() string{
		return "cn-test.axt.aliyun.com"
	})
	defer guard.Reset()

	var commandType string
	var content string
	var workingDir string
	if osutil.GetOsType() == osutil.OSLinux || osutil.GetOsType() == osutil.OSFreebsd {
		commandType = "RunShellScript"
		content = base64.StdEncoding.EncodeToString([]byte("pwd"))
		workingDir = "/tmp"
	} else if runtime.GOOS == "windows" {
		commandType = "RunBatScript"
		content = base64.StdEncoding.EncodeToString([]byte("chdir"))
		workingDir = "C:\\Users"
	}

	rand.Seed(time.Now().UnixNano())
	rand_num := rand.Intn(10000000)
	rand_str := strconv.Itoa(rand_num)

	info := models.RunTaskInfo{
		InstanceId:  "i-test",
		CommandType: commandType,
		TaskId:      "t-test" + rand_str,
		CommandId:   "c-test",
		TimeOut:     "120",
		WorkingDir:  workingDir,
		Content:     content,
	}
	task, _ := NewTask(info, nil, nil, nil)

	errcode, err := task.Run()

	/*if runtime.GOOS == "windows" {
		assert.Contains(t, output, "Users")
	}

	if runtime.GOOS == "linux" {
		assert.Contains(t, output, "tmp")
	}*/

	assert.Equal(t, nil , err)
	assert.Equal(t, 0 , int(errcode))
}
