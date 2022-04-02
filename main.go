package main

import (
	"log"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	// ブラウザはChromeを指定して起動
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	// ログインページに遷移
	if err := page.Navigate("https://elms.u-aizu.ac.jp/login/index.php"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	// ID, Passの要素を取得し、値を設定
	identity := page.FindByID("username")
	password := page.FindByID("password")
	identity.Fill("own id")
	password.Fill("own password")
	// formをサブミット
	/*
		if err := page.FindByClass("loginbtn.btn.btn-primary.btn-block.mt-3").Submit(); err != nil {
			log.Fatalf("Failed to login:%v", err)
		}
	*/
	btn := page.FindByButton("Log in")
	if err = btn.Click(); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	// 処理完了後、10秒間ブラウザを表示しておく
	time.Sleep(10 * time.Second)

}
