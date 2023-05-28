package utils_test

import (
	"testing"

	"github.com/Logiase/MiraiGo-Template/utils"
)

func TestCallPushDeer(t *testing.T) {
	err := utils.CallPushDeer("pushKey", "text")
	if err != nil {
		t.Error(err)
	}
}
