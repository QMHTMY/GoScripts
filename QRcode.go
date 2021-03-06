// 将网址制作成二维码并保存再本地
package main

import (
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
	"os"
)

// init提前检测参数值
func init() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: " + os.Args[0] + " https://xxx.xxx.xxx")
	}
}

// main 实现url只作二维码
func main() {
	qr, err := qrcode.New(os.Args[1], qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}
	qr.BackgroundColor = color.RGBA{255, 99, 71, 255}
	qr.ForegroundColor = color.White
	qr.WriteFile(256, "./urlQrcode.png")
}
