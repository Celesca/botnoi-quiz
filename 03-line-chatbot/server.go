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
		responseMessage := linebot.NewTextMessage("สวัสดีครับ บอท Folk พร้อมให้บริการครับ")
		_, err := bot.ReplyMessage(replyToken, responseMessage).Do()
		if err != nil {
			log.Print(err)
		}

	case "Button":
		responseButtonTemplate(replyToken)
	default:
	}

}

func responseCarousel(replyToken string) {
	imgURL := "https://example.com/bot/images/image.jpg"
	template := linebot.NewCarouselTemplate()
}

func responseButtonTemplate(replyToken string) {
	template := linebot.NewButtonsTemplate(
		"https://botnoi.ai/assets/etc/botnoi.png", // Thumbnail image URL
		"Menu",          // Title
		"Please select", // Text
		linebot.NewURIAction("View detail", "http://example.com/page/123"),
		linebot.NewPostbackAction("Buy", "action=buy&itemid=123", "", "", "", ""),
		linebot.NewPostbackAction("Add to cart", "action=add&itemid=123", "", "", "", ""),
		linebot.NewURIAction("View detail", "http://example.com/page/123"),
	)
	template.ImageAspectRatio = "rectangle"
	template.ImageSize = "cover"
	template.ImageBackgroundColor = "#FFFFFF"

	message := linebot.NewTemplateMessage("This is a buttons template", template)

	if _, err := bot.ReplyMessage(replyToken, message).Do(); err != nil {
		log.Print(err)
	}
}
