package helper

import (
	"fmt"
	"strings"

	"github.com/mirobidjon/go_dynamic_service/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendTelegramMessage(message string) error {
	cfg := config.Load()

	bot, err := tgbotapi.NewBotAPI(cfg.TGBotToken)
	if err != nil {
		return fmt.Errorf("failed to create bot: %v", err)
	}

	serviceName := fmt.Sprintf("[%s] - [%s]\n", strings.ToUpper(cfg.ServiceName), strings.ToUpper(cfg.Environment))
	msg := tgbotapi.NewMessage(cfg.TGChatId, serviceName+message)

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	return nil
}

func SendFileTelegram(cfg config.Config, fileName string, data []byte) error {
	bot, err := tgbotapi.NewBotAPI(cfg.TGBotToken)
	if err != nil {
		return fmt.Errorf("failed to create bot: %v", err)
	}

	fileBytes := tgbotapi.FileBytes{
		Name:  fileName,
		Bytes: data,
	}

	msg := tgbotapi.NewDocumentUpload(-825653122, fileBytes)

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	return nil
}
