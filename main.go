// package main
//
// import (
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"fmt"
//
// 	"github.com/go-telegram-bot-api/telegram-bot-api"
// )
//
// var numericKeyboard = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("تماس با ما"),
// 		tgbotapi.NewKeyboardButton("نشریه ها"),
// 	),
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("5"),
// 		tgbotapi.NewKeyboardButton("6"),
// 	),
// )
// var backbutton = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("بازگشت"),
// 	),
// )
//
// type userdata struct {
// 	Name  string
// 	Phone string
// 	id    string
// 	course string
// }
//
// var userMap = make(map[int64]userdata)
//
// func main() {
// 	bot, err := tgbotapi.NewBotAPI("6377789715:AAGggk8Aj1w2wVHkZ7SQIWhMOi09kq8IHQc")
// 	if err != nil {
// 		log.Panic(err)
// 	}
//
// 	bot.Debug = true
//
// 	log.Printf(" %s was started ...", bot.Self.UserName)
//
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60
//
// 	updates, _ := bot.GetUpdatesChan(u)
//
// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}
//
// 		if update.Message.IsCommand() {
// 			switch update.Message.Command() {
// 			case "start":
// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "به ربات خوش آمدید")
// 				msg.ReplyMarkup = numericKeyboard
// 				bot.Send(msg)
// 			case "product1":
// 				sendPhotoToChat(bot,update.Message.Chat.ID,"./15.png","اولین محصول تستی")
// 				requestname(bot,update.Message.Chat.ID)
// 				userdata:=userMap[update.Message.Chat.ID]
// 				userdata.course="کارگاه انتگرال"
// 			}
// 		} else {
// 			noncommandhandler(bot, update.Message)
// 		}
//
// 		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
//
// 	}
// }
//
// func saveToFile(filename string, userData userdata) error {
// 	// باز کردن یا ایجاد فایل برای نوشتن
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
//
// 	// نوشتن اطلاعات به فایل به صورت خطی
// 	_, err = file.WriteString(fmt.Sprintf("%s|%s|%s|%s\n",userData.course,userData.Name,userData.Phone,userData.id ))
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println("اطلاعات با موفقیت در فایل ذخیره شد.")
// 	return nil
// }
//
//
// func noncommandhandler(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
// 	if userData, ok := userMap[message.Chat.ID]; ok {
// 		if userData.Name == "" {
// 			userData.Name = message.Text
// 			msg := tgbotapi.NewMessage(message.Chat.ID, "لطفاً شماره تماس خود را وارد کنید:")
// 			bot.Send(msg)
// 		} else if userData.Phone == "" {
// 			userData.Phone = message.Text
// 			msg := tgbotapi.NewMessage(message.Chat.ID, "لطفا شماره دانشجویی خود را وارد نمایید")
// 			bot.Send(msg)
// 			}else if userData.id==""{
// 				userData.id=message.Text
// 				msg := tgbotapi.NewMessage(message.Chat.ID, "اطلاعات شما ذخیره شد. ممنون!")
// 				bot.Send(msg)
// 			// اینجا می‌توانید اطلاعات ذخیره شده را به دیگر سرویس‌ها یا پایگاه داده ارسال کنید
// 				saveToFile(userData.course,userData)
// 		}
// 	}
// }
//
// func sendPhotoToChat(bot *tgbotapi.BotAPI, chatID int64, photoPath, caption string) {
// 	// Open the photo file
// 	photo, err := os.Open(photoPath)
// 	if err != nil {
// 		log.Println("Error opening photo:", err)
// 		return
// 	}
// 	defer photo.Close()
//
// 	photobyte, _ := ioutil.ReadAll(photo)
//
// 	photoConfig := tgbotapi.NewPhotoUpload(chatID, tgbotapi.FileBytes{
// 		Name:  "photo",
// 		Bytes: photobyte,
// 	})
// 	photoConfig.Caption = caption
//
// 	// Send the photo to the chat
// 	_, err = bot.Send(photoConfig)
// 	if err != nil {
// 		log.Println("Error sending photo:", err)
// 		return
// 	}
//
// 	log.Println("Photo sent successfully.")
// }
//
// func sendLocation(bot *tgbotapi.BotAPI, chatID int64) {
// 	latitude := 36.3585433433271
// 	longitude := 59.48413923382759
// 	location := tgbotapi.NewLocation(chatID, latitude, longitude)
// 	_, err := bot.Send(location)
// 	if err != nil {
// 		log.Println("Error sending location with caption:", err)
// 		return
// 	}
// 	log.Println("Location with caption sent successfully.")
// }
//
// func requestname(bot *tgbotapi.BotAPI, chatID int64) {
// 	if userData, ok := userMap[chatID]; ok {
// 		// اگر یک کاربر با این chatID وجود داشته باشد
// 		userData.Name = "" // یا دستورات دیگر بر اساس نیاز
// 		userData.Phone = ""
// 		userData.id = ""
// 		userData.course = ""
// 	} else {
// 		// اگر کاربر با این chatID وجود نداشته باشد
// 		userMap[chatID] = userdata{}
// 	}
//
// 	msg := tgbotapi.NewMessage(chatID, "لطفاً نام و نام خانوادگی خود را وارد کنید:")
// 	bot.Send(msg)
// }

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var eventsbutton = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("کارگاه انتگرال"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("کارگاه برنامه نویسی HTML و CSS"),
	),
)

// var backbutton = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("بازگشت"),
// 	),
// )

var userMap = make(map[int64]*userdata)

type userdata struct {
	Name   string
	Phone  string
	ID     string
	Course string
}

type Event struct {
	Name    string
	Photoid string
	Caption string
}

var events []Event

func main() {

	events = append(events, Event{Name: "کارگاه انتگرال",
		Photoid: "AgACAgQAAxkBAAO5ZU2XHOq54v7PakbnywmFJYUURqcAAum6MRvnkHFSWv3tMh0BVPoBAAMCAAN5AAMzBA",
		Caption: `🎗 انجمن علمی کامپیوتر دانشگاه خاوران برگزار میکند :


		▫️برگزاری کارگاه انتگرال
		
		
		📅زمان: شنبه 1402/08/13
		⏱ ساعت حرکت از دانشگاه: 17:45
		👤 ظرفیت: 20 نفر
		
		💰 هزینه شرکت در بازدید: رایگان
		
		📮 جهت ثبت نام از طریق زیر وارد ربات انجمن علمی کامپیوتر دانشگاه خاوران مشهد شوید👇
		
		📣 @khavaran_uni_bot
		`})
	events = append(events, Event{Name: "کارگاه برنامه نویسی HTML و CSS",
		Photoid: "AgACAgQAAxkBAAEni79lTkNkSQ74B5ltbyH5x-dsliPD6wACMb8xG-2NcFJXtR_SIg6MaQEAAwIAA3kAAzME",
		Caption: `🎗 انجمن علمی کامپیوتر دانشگاه خاوران برگزار میکند :


		▫️برگزاری کارگاه HTML , CSS
		
		
		📅زمان: شنبه 1402/08/13
		⏱ ساعت حرکت از دانشگاه: 17:45
		👤 ظرفیت: 20 نفر
		
		💰 هزینه شرکت در بازدید: رایگان
		
		📮 جهت ثبت نام از طریق زیر وارد ربات انجمن علمی کامپیوتر دانشگاه خاوران مشهد شوید👇
		
		📣 @khavaran_uni_bot
		`})

	bot, err := tgbotapi.NewBotAPI("6377789715:AAEVoRj4TOMkJXQl0hOjTLjK3Hm8OV4FZs0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("%s was started ...", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			handleCommand(bot, update.Message)
		} else {
			nonCommandHandler(bot, update.Message)
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func handleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		welcomeMessage := "سلام! ✋ \n به ربات انجمن علمی کامپیوتر دانشگاه خاوران خوش آمدید. 🌷\n"
		welcomeMessage += "شما می‌توانید با استفاده از من درخواست‌های خود را به اعضای انجمن علمی کامپیوتر اعلام کنید تا من آنها را به دستشان برسانم. 🙏\n"
		welcomeMessage += "برای استفاده از ربات، از منوی زیر استفاده کنید. 😉"
		msg := tgbotapi.NewMessage(message.Chat.ID, welcomeMessage)
		msg.ReplyMarkup=tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
	case "events":
		msg := tgbotapi.NewMessage(message.Chat.ID, "یک کارگاه را انتخاب کنید")
		msg.ReplyMarkup = eventsbutton
		bot.Send(msg)
	case "journals":
		caption := `اولین نشریه پاییز 1402 🍁📝
		با موضوعات زیر منتشر شد :
		📌 هوش مصنوعی و دیجیتال مارکتینگ
		📌 کاربرد هوش مصنوعی در دیجیتال مارکتینگ
		📌 بهبود UX سایت`
		sendPDFToChat(bot, message.Chat.ID, "BQACAgQAAxkBAAIBemVOJMT_WTxzlMYlqQwGWf3SpCNlAAKUEAACxfJxUiLb9KQC9T76MwQ",caption)
	case "about":
		aboutText := `🔹 انجمن علمی کامپیوتر دانشگاه خاوران مشهد در سال 1397 و با هدف ارتقای سطح علمی دانشجویان این دانشگاه تاسیس شد. این انجمن تا کنون با برگزاری سمینارهای مختلف و اردو های متعدد فعالیت های خود را ادامه داده است.
	
	🔸 این انجمن در سال 1402 با اعضای جدید در صدد برگزاری دوره های مختلف آموزشی و برگزاری مسابقات مهارتی همراه با جوایز نقدی و غیر نقدی است.  
	
	👥 اعضای کنونی انجمن علمی کامپیوتر:
	
	👨🏻‍🎓 استاد مشاور: دکتر الهام مهدی پور
	👨🏻‍🎓 دبیر انجمن: علیرضا خوش صحبت
	👨🏻‍🎓 نائب دبیر: پرنیا براتی
	👨🏻‍🎓 گرافیست: حسین سروی
	👨🏻‍🎓 تیم تحریر: حسین اسماعیلی (شکیبا شاکری، علیرضا نیکنیان)
	👨🏻‍🎓 برگزاری کلاس ها: پرنیا براتی (لیلا مروج، ریحانه حسینی پور)
	👨🏻‍🎓 تولید و محتوای دیجیتال: غزاله اصغری (نیاز به عضو)
	👨🏻‍🎓 ارتباط با سایر انجمن ها: علیرضا خوش صحبت (نیاز به عضو)
	
	◾️ آدرس دانشگاه:
	بلوار شهید فلاحی، شهید فلاحی ۲، بین برادران شهید حسینی ۱۰ و ۱۲`
	
		msg := tgbotapi.NewMessage(message.Chat.ID, aboutText)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
		sendLocation(bot, message.Chat.ID)
	
	}
}

func sendphotoToChat(bot *tgbotapi.BotAPI, chatID int64, telegramFileID, caption string) {
	photoConfig := tgbotapi.NewPhotoShare(chatID, telegramFileID)
	photoConfig.Caption = caption
	photoConfig.ReplyMarkup=tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.Send(photoConfig)
	if err != nil {
		log.Println("Error sending photo:", err)
		return
	}

	log.Println("Photo sent successfully.")
}

func requestName(bot *tgbotapi.BotAPI, chatID int64) {
	if userData := getUserData(chatID); userData != nil {
		userData.Name = ""
		userData.Phone = ""
		userData.ID = ""
		userData.Course = ""
	} else {
		userMap[chatID] = &userdata{}
	}

	msg := tgbotapi.NewMessage(chatID, "لطفاً نام و نام خانوادگی خود را وارد کنید:")
	bot.Send(msg)
}

func getUserData(chatID int64) *userdata {
	if userData, ok := userMap[chatID]; ok {
		return userData
	}
	return nil
}

func nonCommandHandler(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	for i := 0; i < len(events); i++ {
		if message.Text == events[i].Name {
			sendphotoToChat(bot, message.Chat.ID, events[i].Photoid, events[i].Caption)
			requestName(bot, message.Chat.ID)
			userdata := getUserData(message.Chat.ID)
			userdata.Course = events[i].Name

			break
		} else if userData := getUserData(message.Chat.ID); userData != nil {
			if userData.Name == "" && isText(message) {
				userData.Name = message.Text
				msg := tgbotapi.NewMessage(message.Chat.ID, "لطفاً شماره تماس خود را وارد کنید:")
				bot.Send(msg)
				break
			} else if userData.Phone == "" && isText(message) {
				userData.Phone = message.Text
				msg := tgbotapi.NewMessage(message.Chat.ID, "لطفا شماره دانشجویی خود را وارد نمایید")
				bot.Send(msg)
				break
			} else if userData.ID == "" && isText(message) {
				userData.ID = message.Text
				msg := tgbotapi.NewMessage(message.Chat.ID, "اطلاعات شما ذخیره شد. ممنون!")
				msg.ReplyMarkup=tgbotapi.NewRemoveKeyboard(true)
				bot.Send(msg)
				saveToFile(userData.Course, userData)
				delete(userMap, message.Chat.ID)
				break
			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, "لطفا مجددا وارد نمایید")
				bot.Send(msg)
			}
		}
	}

}

func saveToFile(filename string, userData *userdata) error {
	var file *os.File
	var err error
	if _, statErr := os.Stat(filename); os.IsNotExist(statErr) {
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
	} else {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return err
		}
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s|%s|%s|%s\n", userData.Course, userData.Name, userData.Phone, userData.ID))
	if err != nil {
		return err
	}
	fmt.Println("اطلاعات با موفقیت در فایل ذخیره شد.")
	return nil
}

func isText(message *tgbotapi.Message) bool {
	return message.Text != "" || message.Caption != ""
}

func sendPDFToChat(bot *tgbotapi.BotAPI, chatID int64, fileID, caption string) {
	file := tgbotapi.NewDocumentShare(chatID, fileID)
	file.Caption = caption
	bot.Send(file)
}

func sendLocation(bot *tgbotapi.BotAPI, chatID int64) {
	latitude := 36.3585433433271
	longitude := 59.48413923382759
	location := tgbotapi.NewLocation(chatID, latitude, longitude)
	_, err := bot.Send(location)
	if err != nil {
		log.Println("Error sending location with caption:", err)
		return
	}
	log.Println("Location with caption sent successfully.")
}

// sendphotoToChat(bot, message.Chat.ID, "AgACAgQAAxkBAAO5ZU2XHOq54v7PakbnywmFJYUURqcAAum6MRvnkHFSWv3tMh0BVPoBAAMCAAN5AAMzBA", "اولین محصول تستی")
