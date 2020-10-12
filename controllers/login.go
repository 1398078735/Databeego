package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/skip2/go-qrcode"
	"math/rand"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	rand.Seed(time.Now().UnixNano())
	time := rand.Intn(1000)
	err := qrcode.WriteFile(" 扫码 "+string(time), qrcode.Medium, 256, "./static/img/qrcode.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	l.TplName = "login.html"
}

