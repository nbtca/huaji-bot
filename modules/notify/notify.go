package notify

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/models"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/spf13/viper"
)

var instance *Notify

type Notify struct {
	bot *bot.Bot
}

func init() {
	instance = &Notify{}
	bot.RegisterModule(instance)
}

func (n *Notify) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "internal.Notify",
		Instance: instance,
	}
}

func (n *Notify) Init() {

}

func (n *Notify) PostInit() {

}

func (n *Notify) Serve(b *bot.Bot) {
	n.bot = b
}

func (n *Notify) Start(b *bot.Bot) {
	go n.StartRPC()
}

func (n *Notify) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

func (n *Notify) EventActionNotify(req *models.EventActionNotifyRequest, res *models.EventActionNotifyResponse) error {
	groups := viper.GetIntSlice("notifyGroup")
	req.Link = "https://repair.nbtca.space"
	msg := message.NewSendingMessage()
	if req.ActorAlias != "" {
		msg.Append(message.NewText(fmt.Sprintf("%v(%v)\n", req.Subject, req.ActorAlias)))
	} else {
		msg.Append(message.NewText(fmt.Sprintf("%v\n", req.Subject)))
	}
	if req.Problem != "" {
		msg.Append(message.NewText(fmt.Sprintf("问题: %v\n", req.Model)))
	}
	if req.Model != "" {
		msg.Append(message.NewText(fmt.Sprintf("型号: %v\n", req.Model)))
	}
	msg.Append(message.NewText(fmt.Sprintf("%v", req.Link)))

	for _, group := range groups {
		n.bot.SendGroupMessage(int64(group), msg)
	}
	log.Println("EventCreate", req)
	res.Success = true
	return nil
}

func (n *Notify) StartRPC() {
	rpc.Register(n)
	rpc.HandleHTTP()
	port := viper.GetString("rpcPort")
	log.Println("StartRPC", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
