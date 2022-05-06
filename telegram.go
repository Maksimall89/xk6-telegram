package telegram

import (
	"errors"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.k6.io/k6/js/modules"
)

// Register the extension on module initialization, available to
func init() {
	modules.Register("k6/x/telegram", new(Telegram))
}

// Telegram is the k6 extension.

type Telegram struct{}

// Connect to telegram bot API
func (t *Telegram) Connect(token string, isDebug bool) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = isDebug
	return bot, nil
}

// Send a text message to chat
func (t *Telegram) Send(client *tgbotapi.BotAPI, chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "html"
	_, err := client.Send(msg)
	return err
}

// SendReplay message to chat
func (t *Telegram) SendReplay(client *tgbotapi.BotAPI, chatID int64, replayID int, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "html"
	msg.ReplyToMessageID = replayID
	_, err := client.Send(msg)
	return err
}

// ShareImage exist a photo like http://myphoto.com/photo.png
func (t *Telegram) ShareImage(client *tgbotapi.BotAPI, chatID int64, fileID string) error {
	msg := tgbotapi.NewPhotoShare(chatID, fileID)
	_, err := client.Send(msg)
	return err
}

// UploadImageByte a photo like bytes array
func (t *Telegram) UploadImageByte(client *tgbotapi.BotAPI, chatID int64, photoBytes []byte) error {
	if len(photoBytes) == 0 {
		return errors.New("photo bytes is empty")
	}
	fileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}

	msg := tgbotapi.NewPhotoUpload(chatID, fileBytes)
	_, err := client.Send(msg)
	return err
}

// UploadImagePath a photo use local path to photo
func (t *Telegram) UploadImagePath(client *tgbotapi.BotAPI, chatID int64, photoPath string) error {
	photoBytes, err := os.ReadFile(photoPath)
	if err != nil {
		return err
	}
	return t.UploadImageByte(client, chatID, photoBytes)
}

// GetUpdate last message from chat
func (t *Telegram) GetUpdate(client *tgbotapi.BotAPI) (tgbotapi.Update, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := client.GetUpdatesChan(u)
	if err != nil {
		return tgbotapi.Update{}, err
	}

	return <-updates, nil
}
