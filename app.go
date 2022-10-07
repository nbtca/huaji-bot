package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/spf13/viper"

	_ "github.com/Logiase/MiraiGo-Template/modules/logging"
	_ "github.com/Logiase/MiraiGo-Template/modules/notify"
)

func init() {
	utils.WriteLogToFS(utils.LogInfoLevel, utils.LogWithStack)
	config.Init()
}

func main() {

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 快速初始化
	bot.Init()

	// 初始化 Modules
	bot.StartService()

	// 使用协议
	// 不同协议可能会有部分功能无法使用
	// 在登陆前切换协议
	bot.UseProtocol(bot.AndroidPhone)

	// 登录
	bot.Login()

	bot.SaveToken()
	// 刷新好友列表，群列表
	bot.RefreshList()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)
	<-ch
	bot.Stop()
}
