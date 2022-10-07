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

func (n *Notify) EventCreate(req *models.EventActionNotifyRequest, res *models.EventActionNotifyResponse) error {
	groups := viper.GetIntSlice("notifyGroup")
	// content := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?> <msg templateID="12345" action="web" brief="new event" serviceID="1" url="http://repair.nbtca.space">
	// 		<item layout="0">
	// 			<title>%v</title>
	// 			<summary>问题描述: %v</summary>
	// 			<summary>型号: %v</summary>
	// 		</item>
	// 	</msg>`, req.Subject, req.Problem, req.Model)
	// msg := message.NewSendingMessage().
	// 	Append(message.NewRichXml(content, 0))
	req.Link = "https://repair.nbtca.space"
	msg := message.NewSendingMessage().
		Append(message.NewText(fmt.Sprintf("新事件: %v\n问题描述: %v\n型号: %v\n时间: %v\n链接: %v", req.Subject, req.Problem, req.Model, req.GmtCreate, req.Link)))
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
