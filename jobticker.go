package main

import (
	"fmt"
	"time"
)


type jobTicker struct {
		timer *time.Timer
		good  *goodHandler
}

const INTERVAL_PERIOD time.Duration = 24 * time.Hour

const HOUR 	 int = 00
const MINUTE int = 00
const SECOND int= 00



func (j *jobTicker) run() {
	var msg_array []map[string]interface{}
	j.updateTimer()
	for {
		select {
		case msg := <-j.good.msg:
			msg_array = append(msg_array, msg)
			fmt.Println(msg_array)
			//j.outputMessage(&msg_array)
		case <-j.timer.C:
			fmt.Println(time.Now(), "- time_ticker 発火")
			j.outputMessage(&msg_array)
			j.updateTimer()
		}
	}
}



// Tickerを更新
func (j *jobTicker) updateTimer() {
	location,_ := time.LoadLocation("Asia/Tokyo")
	now := time.Now()
	nextTick := time.Date(time.Now().Year(), time.Now().Month(),time.Now().Day(), HOUR, MINUTE, SECOND, 0, location)

	fmt.Println("nextTick: ", nextTick, "now: ", now)

	// nextTick < now  日をまたぐときだけ、24時間足す
	if nextTick.Before(now) {
		fmt.Println("ok")
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	fmt.Println(nextTick, "次のアラート")

	// nextTickまでの秒数
	diff := nextTick.Sub(now)
	if j.timer == nil {
		// 初めの処理のみここでtickerを作成する
		j.timer = time.NewTimer(diff)
	} else {
		if !j.timer.Stop() {
			fmt.Println("止まってない")
			select {
			case <-j.timer.C:
				fmt.Println("止めた")
			default:
			}
		}
		// diff秒後に期限切れになるタイマーをセット
		j.timer.Reset(diff)
	}

}
