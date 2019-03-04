package main

import (
	"EvaluationServer/pb"
	"log"
	"time"
)

//Waiting 是否开始等待
var waiting = false
var waittime = &pb.Time{}

//设置时间
func setTime(h int32, m int32, s int32) {
	waittime.Hour = h
	waittime.Minute = m
	waittime.Second = s
}

//时间转换
func changeToS(time *pb.Time) int32 {
	return ((time.Hour*60+time.Minute)*60 + time.Second)
}

//时间转换
func changeFrom(time int32) {
	waittime.Second = time % 60
	waittime.Minute = time / 60 % 60
	waittime.Hour = time / 3600
}

//开始等待倒计时
func timeStart() {
	waiting = true
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		time := changeToS(waittime)
		if time <= 0 {
			return
		}
		time--
		changeFrom(time)
		log.Println("S:", time, "time:", waittime)
	}
}
