package main

 import (
     "log"
     "net/http"
     "os"

     "github.com/line/line-bot-sdk-go/linebot"
     "github.com/line/line-bot-sdk-go/linebot/httphandler"
 )

 func main() {

     // CHANNEL_SECRETとCHANNEL_TOKENはLine Developersから取得できる
     handler, err := httphandler.New(
         os.Getenv("CHANNEL_SECRET"),
         os.Getenv("CHANNEL_TOKEN"),
     )
     if err != nil {
         log.Fatal(err)
     }

     // Setup HTTP Server for receiving requests from LINE platform
     handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
         bot, err := handler.NewClient()
         if err != nil {
             log.Print(err)
             return
         }
         for _, event := range events {
             if event.Type == linebot.EventTypeMessage {
                 switch message := event.Message.(type) {
                 case *linebot.TextMessage:
                     if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
                         log.Print(err)
                     }
                 }
             }
         }
     })
     http.Handle("/callback", handler)

     //if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
     //  log.Fatal(err)
     //}

     // HTTPではLine APIを利用できないため、HTTPSに変更する
     // CERT_FILEとKEY_FILEはSSL証明書を発行した後に書き換える
     if err := http.ListenAndServeTLS(":443", "CERT_FILE", "KEY_FILE", nil); err != nil {
         log.Fatal("ListenAndServe: ", err)
     }
 }