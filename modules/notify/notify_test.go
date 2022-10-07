package notify_test

import (
	"net/rpc"
	"testing"

	"github.com/Logiase/MiraiGo-Template/models"
)

func Test_Notify(t *testing.T) {
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		t.Error(err)
	}
	req := models.EventActionNotifyRequest{
		Subject:   "test",
		Model:     "test",
		Problem:   "test",
		Link:      "https://repair.nbtca.space",
		GmtCreate: "2021-01-01 00:00:00",
	}
	res := models.EventActionNotifyResponse{}
	if err = conn.Call("Notify.EventCreate", req, &res); err != nil {
		t.Error(err)
	}
}
