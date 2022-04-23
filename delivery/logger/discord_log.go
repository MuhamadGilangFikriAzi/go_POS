package logger

import (
	"fmt"
	"github.com/gtuk/discordwebhook"
	"github.com/rs/zerolog"
	"gopos.com/m/utility"
	"os"
)

var (
	Log *zerolog.Logger
)

func NewLoger(logLevel string) {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if logLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	loger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = &loger
}

func SendLogToDiscord(serve string, errSend error) {
	var username = "Log Api Error"
	var url = "https://discord.com/api/webhooks/961186062586236958/4-EXW7G7eVjZbqsslxH-sVyEOBqc05xOlJiVKqS5xFk96xb8qjOy9HJV_c_OPqeqx93C"

	currentTime := utility.ThisTimeStamp()
	content := fmt.Sprintf("(%s)\n Serve: %s\n Error: %s", currentTime, serve, errSend.Error())
	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	discordwebhook.SendMessage(url, message)
	//if err != nil {
	//	log.Fatal(err)
	//}
	Log.Err(errSend).Msg(serve)
}
