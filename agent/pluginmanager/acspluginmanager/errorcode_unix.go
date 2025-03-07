// +build linux freebsd

package acspluginmanager

import (
	"errors"
	"fmt"
	"syscall"
	"github.com/aliyun/aliyun_assist_client/agent/util/errnoutil"
)

func errProcess(function string, exitCode int, err error, tip string) (int, string) {
	var errorCode string
	var ok bool
	if errorCode, ok = ErrorStrMap[exitCode]; !ok {
		errorCode = "UNKNOWN"
	}
	var errno syscall.Errno
	if errors.As(err, &errno) {
		if errnoPhrase, ok := errnoutil.ErrnoPhrases[errno]; ok {
			errorCode += "." + errnoPhrase
		}
	}
	fmt.Printf("%s %s: %s\n", function, errorCode, tip)
	return exitCode, errorCode
}
