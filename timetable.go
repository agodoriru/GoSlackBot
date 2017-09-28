package main

import (
	// "log"
	"strings"
	// "os"
	"strconv"
	// "fmt"
	// "bufio"
	// "time"

	//APIのimport
	"github.com/nlopes/slack"
)


//時間割用の配列を作成
var table [7][7]string

func timeTable(ev *slack.MessageEvent, rtm *slack.RTM) {

	//初期設定的な何か・生配列に直で書き込み
	if strings.HasPrefix(ev.Text, "set "){

		var tmp []string=strings.Fields(ev.Text)
		
		// var k int
		// for k=0;k<len(tmp);k++ {
		// 	rtm.SendMessage(rtm.NewOutgoingMessage(tmp[k],ev.Channel))
		// }

		var weekday_period string = tmp[1]
		var class_name string =tmp[2]

		//確認出力
		// rtm.SendMessage(rtm.NewOutgoingMessage("曜日・時限:"+weekday_period+" 授業名:"+class_name,ev.Channel))

		var arr []string
		arr = strings.Split(weekday_period, "")
		
		//何曜日に授業があるか
		var weekday_tmp string =arr[0]
		var weekday int

		if weekday_tmp=="月"{
			weekday=0
		}else if weekday_tmp=="火"{
			weekday=1
		}else if weekday_tmp=="水"{
			weekday=2
		}else if weekday_tmp=="木"{
			weekday=3
		}else if weekday_tmp=="金"{
			weekday=4
		}else if weekday_tmp=="土" {
			weekday=5
		}

		//何時限目に授業があるか
		var period_tmp string =arr[1]
		var period int =0
		hoge, _:= strconv.Atoi(period_tmp)
		period=period+hoge


		// 確認出力
		// rtm.SendMessage(rtm.NewOutgoingMessage("時限:"+period, ev.Channel))


		//出力
		rtm.SendMessage(rtm.NewOutgoingMessage(weekday_tmp+"曜日の"+period_tmp+"限目に"+class_name+"を設定しました。",ev.Channel))
		
		//時間割配列に代入		
		table[weekday][period-1]=class_name

	}

	// 時間割を表示/消すかもしれない
	// if ev.Text=="時間割表示"{
	// 	rtm.SendMessage(rtm.NewOutgoingMessage("月 火 水 木 金",ev.Channel))
	// 	for i:=0;i<5;i++{
	// 		rtm.SendMessage(rtm.NewOutgoingMessage(table[0][i]+" "+table[1][i]+" "+table[2][i]+" "+table[3][i]+" "+table[4][i], ev.Channel))
	// 	}
	// }

	
	//現在の時間を取る関数




	//次の時間割の教室を通知する関数
	//次の時間はーで教室はーですみたいな

}