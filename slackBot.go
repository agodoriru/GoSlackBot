package main

import (
    "log"
    "strings"
    // "os"

    "github.com/nlopes/slack"
)


var date string
var classroom string

func setClassroom(ev *slack.MessageEvent, rtm *slack.RTM) {
    
        var arr []string=strings.Fields(ev.Text)

        date = arr[1]
        classroom = arr[2]

        rtm.SendMessage(rtm.NewOutgoingMessage("設定完了",ev.Channel))

}

func deletation(ev *slack.MessageEvent, rtm *slack.RTM) {

    date = ""
    classroom = ""

    rtm.SendMessage(rtm.NewOutgoingMessage("削除完了",ev.Channel))

}

func showClassroom(ev *slack.MessageEvent, rtm *slack.RTM) {
    
    rtm.SendMessage(rtm.NewOutgoingMessage("次の活動日は"+date+"で、活動場所は"+classroom+"です。",ev.Channel))

}

func main() {
    // API Clientを作成する
    api := slack.New(getApiToken())

    // WebSocketでSlack RTM APIに接続する
    rtm := api.NewRTM()
    // goroutineで並列化する
    go rtm.ManageConnection()

    // イベントを取得する
    for msg := range rtm.IncomingEvents {
        // 型swtichで型を比較する
        switch ev := msg.Data.(type) {

        case *slack.HelloEvent:
                log.Print("activated")
        


        case *slack.MessageEvent:

            // test case
            if strings.HasPrefix(ev.Text,"test") {
            
                rtm.SendMessage(rtm.NewOutgoingMessage("Test Success!",ev.Channel))
            
            }else if strings.HasPrefix(ev.Text,"del") {
            
                deletation(ev, rtm)
            
            }else if strings.HasPrefix(ev.Text,"set "){
            
                setClassroom(ev, rtm)

            }else{
                showClassroom(ev, rtm)

            }


        
        case *slack.InvalidAuthEvent:
                log.Print("Invalid credentials")
                return 




        }
    }
}