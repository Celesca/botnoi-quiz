package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func main() {
	var err error

	channelSecret := "539cf9da750ff5f5ab1e93d5f677d5f4"
	channelToken := "lvfhMH23GeRN88NQ75GpQb5/k2V9xYs8SDKrsXJlFDAiyR1nt10yuVFtHWZAszy6mSajXcDIwKinoiF6LypjIDZpnpsqO5CzPkW7UvbPMibyf6lb6erMS8L9v7F4X3Fy28YJ7NaCnOGZRnxu4I9VRQdB04t89/1O/w1cDnyilFU="

	bot, err = linebot.New(
		channelSecret,
		channelToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.POST("/callback", callbackHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func callbackHandler(c echo.Context) error {
	events, err := bot.ParseRequest(c.Request())
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.NoContent(http.StatusInternalServerError)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			handleMessageEvent(event)
		}
	}

	return c.NoContent(http.StatusOK)
}

func handleMessageEvent(event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		handleTextMessage(event.ReplyToken, message)
	default:
		log.Printf("Unsupported message: %v", message)
	}
}

func handleTextMessage(replyToken string, message *linebot.TextMessage) {
	switch message.Text {
	case "Text":
		responseTextMessage(replyToken, "สวัสดีครับนี่คือข้อความตอบกลับครับ")
	case "Button":
		responseButtonTemplate(replyToken)
	case "Carousel":
		responseCarouselTemplate(replyToken)
	default:
		responseTextMessage(replyToken, "ขออภัยครับ ผมไม่เข้าใจคำถามของคุณ")
	}
}

func getQuickReplyItems() *linebot.QuickReplyItems {
	quickReplyItems := []*linebot.QuickReplyButton{
		linebot.NewQuickReplyButton(
			"",
			linebot.NewMessageAction("ดู Text", "Text"),
		),
		linebot.NewQuickReplyButton(
			"",
			linebot.NewMessageAction("ดู Button", "Button"),
		),
		linebot.NewQuickReplyButton(
			"",
			linebot.NewMessageAction("ดู Carousel", "Carousel"),
		),
	}

	return linebot.NewQuickReplyItems(quickReplyItems...)
}

func responseTextMessage(replyToken string, message string) {
	responseMessage := linebot.NewTextMessage(message).WithQuickReplies(getQuickReplyItems())
	_, err := bot.ReplyMessage(replyToken, responseMessage).Do()
	if err != nil {
		log.Print(err)
	}
}

func responseCarouselTemplate(replyToken string) {
	var actions []linebot.TemplateAction

	actions = append(actions, linebot.NewMessageAction("ดูเพิ่มเติม", "วันนี้อากาศดี"))
	actions = append(actions, linebot.NewURIAction("ติดต่อสอบถาม", "https://botnoi.ai/"))

	imgURI := "https://cdn.pixabay.com/photo/2024/04/13/18/22/barberry-8694277_1280.jpg"
	secondImgURI := "https://cdn.pixabay.com/photo/2022/12/02/21/20/blue-7631674_960_720.jpg"

	var columns []*linebot.CarouselColumn

	columns = append(columns, linebot.NewCarouselColumn(imgURI, "Garden", "lorem ipsum magnito", actions...))
	columns = append(columns, linebot.NewCarouselColumn(secondImgURI, "Diamond", "lorem ipsum mana", actions...))
	carousel := linebot.NewCarouselTemplate(columns...)
	template := linebot.NewTemplateMessage("Carousel", carousel).WithQuickReplies(getQuickReplyItems())

	_, err := bot.ReplyMessage(replyToken, template).Do()
	if err != nil {
		log.Print(err)
	}
}

func responseButtonTemplate(replyToken string) {
	template := linebot.NewButtonsTemplate(
		"https://cdn.pixabay.com/photo/2024/04/13/18/22/barberry-8694277_1280.jpg",
		"บริการของบอทน้อย",
		"โปรดเลือกเมนูที่ต้องการครับ",
		linebot.NewURIAction("ดูรายละเอียด", "https://botnoi.ai/"),
		linebot.NewMessageAction("ติดต่อเรา", "ฉันต้องการติดต่อคุณ"),
	)
	template.ImageAspectRatio = "rectangle"
	template.ImageSize = "cover"
	template.ImageBackgroundColor = "#FFFFFF"

	message := linebot.NewTemplateMessage("This is a buttons template", template).WithQuickReplies(getQuickReplyItems())

	_, err := bot.ReplyMessage(replyToken, message).Do()
	if err != nil {
		log.Print(err)
	}
}
