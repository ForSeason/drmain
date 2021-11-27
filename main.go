package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	cron "github.com/robfig/cron/v3"
)

func main() {
	drmainPath := `C:\Drcom\DrUpdateClient\DrMain.exe`
	if len(os.Args) == 2 {
		drmainPath = os.Args[1]
	}
	crontab := cron.New()
	// crontab.AddFunc("* * * * *", func() {
	crontab.AddFunc("5 6 * * 1-5", func() {
		log.Println("开始重启……")
		cmd := exec.Command("taskkill", "/im", "DrClient.exe")
		cmd.Run()
		cmd = exec.Command("taskkill", "/im", "DrUpdate.exe")
		cmd.Run()
		cmd = exec.Command("taskkill", "/im", "DrMain.exe")
		cmd.Run()
		time.Sleep(time.Second * 10)
		cmd = exec.Command(drmainPath)
		cmd.Run()
		log.Println("重启完毕.")
	})
	crontab.Run()
}
