package tg

import (
	"github.com/IliaBelov/RPDB/music"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mus music.Music

const (
	commandStart        = "start"
	commandAddMusicList = "add_list"
	commandAddMusic     = "add_music"
	commandDelMusicList = "del_list"
	commandDelMusic     = "del_music"
	commandCheckCom     = "music_com"
	commandCheckText    = "music_text"
	commandCheckList    = "my_music"
	commandHelp         = "help"
	commandCancel       = "cancel"

	replyStart = "Привет!"
)

var (
	musicname   = false
	musictext   = false
	musicauthor = false
)
var (
	Start        = false
	AddMusicList = false
	AddMusic     = false
	DelMusicList = false
	DelMusic     = false
	CheckCom     = false
	CheckText    = false
	CheckList    = false
	Help         = false
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandAddMusicList:
		return b.handleAddMusicListCommand(message)
	case commandDelMusicList:
		return b.handleDelMusicListCommand(message)
	case commandAddMusic:
		return b.handleAddMusicCommand(message)
	case commandDelMusic:
		return b.handleDelMusicCommand(message)
	/*case commandCheckList:
	return b.handleCheckListCommand(message)*/
	case commandCheckText:
		return b.handleCheckTextCommand(message)
	case commandCheckCom:
		return b.handleCheckComCommand(message)
	case commandHelp:
		return b.handleHelpCommand(message)
	case commandCancel:
		return b.handleCancelCommand(message)

	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	switch {
	case AddMusic:
		switch {
		case musicname:
			mus.Music_name = msg.Text
			musicname = false
			musicauthor = true
			msg.Text = "Введите название автора:"
			b.bot.Send(msg)
			return nil

		case musicauthor:
			mus.Author = msg.Text
			musicauthor = false
			musictext = true
			msg.Text = "Введите текст песни:"
			b.bot.Send(msg)
			return nil

		case musictext:
			mus.Music_text = msg.Text
			musictext = false
			AddMusic = false
			//func
			err := b.store.AddMusic(&mus)
			if err != nil {
				msg.Text = err.Error()
				b.bot.Send(msg)
				return err
			}
			msg.Text = "Музыка добавлена"
			b.bot.Send(msg)
			mus.Author = ""
			mus.Id = -1
			mus.Music_name = ""
			mus.Music_text = ""
			return nil
		}
	case CheckText:
		mus.Music_name = msg.Text
		CheckText = false
		//func
		var err error
		msg.Text, err = b.store.CheckMusicText(&mus)
		if err != nil {
			return err
		}
		b.bot.Send(msg)
		return nil

	case CheckCom:
		CheckCom = false
		mus.Music_name = msg.Text
		//func

		return nil
	case DelMusic:
		DelMusic = false
		mus.Music_name = msg.Text
		//func
		err := b.store.DelMusic(&mus)
		if err != nil {
			msg.Text = err.Error()
			b.bot.Send(msg)
			return err
		}
		mus.Music_name = " "
		msg.Text = "Музыка удалена"
		b.bot.Send(msg)
		return nil
	case AddMusicList:
		AddMusicList = false
		mus.Music_name = msg.Text
		//func
		return nil
	case DelMusicList:
		DelMusicList = false
		mus.Music_name = msg.Text
		//func
		return nil
	}

	b.bot.Send(msg)
	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввёл команду /start")

	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCancelCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввёл команду /cancel")

	musicname = false
	musictext = false
	musicauthor = false

	Start = false
	AddMusicList = false
	AddMusic = false
	DelMusicList = false
	DelMusic = false
	CheckCom = false
	CheckText = false
	CheckList = false
	Help = false

	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Команды для поиска по: названию - ,автору")

	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleSearchNameCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handleSearchAuthorCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleAddMusicListCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввёл команду /add_list. Напиши название песни:")
	AddMusicList = true
	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handleDelMusicListCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввёл команду /del_list. Напиши название песни:")
	DelMusicList = true
	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleAddMusicCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввёл команду /add_music. Напиши название песни:")
	AddMusic = true
	musicname = true
	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handleDelMusicCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввёл команду /del_music. Напиши название песни:")
	DelMusic = true
	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}

/*
	func (b *Bot) handleCheckListCommand(message *tgbotapi.Message) error {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввёл команду /my_music")
		//func
		var list []music.Music
		list, err := b.store.CheckList()
		for i := range list {
			msg = tgbotapi.NewMessage(message.Chat.ID, list[i].Music_name)
			_, err = b.bot.Send(msg)
		}

		return err
	}
*/
func (b *Bot) handleCheckTextCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввёл команду /music_text. Напиши название песни:")
	CheckText = true
	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handleCheckComCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввёл команду /music_com. Напиши название песни:")
	CheckCom = true
	//msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")

	_, err := b.bot.Send(msg)
	return err
}
