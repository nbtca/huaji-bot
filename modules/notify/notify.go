package notify

import (
	"log"
	"net/http"
	"net/rpc"
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/models"
	"github.com/Mrs4s/MiraiGo/message"
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
	msg := message.NewSendingMessage().
		Append(&message.TextElement{
			Content: req.Subject + "\n",
		}).
		Append(&message.TextElement{
			Content: "问题描述: " + req.Problem + "\n",
		}).
		Append(&message.TextElement{
			Content: "型号: " + req.Model + "\n",
		}).
		Append(&message.TextElement{
			Content: req.Link + "\n",
		}).
		Append(message.AtAll())
	n.bot.SendGroupMessage(915582432, msg)
	res.Success = true
	return nil
}

func (n *Notify) StartRPC() {
	rpc.Register(n)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
