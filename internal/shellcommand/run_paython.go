package shellcommand

import (
	"log"
	"time"

	config "github.com/siangyeh8818/golang.exporter.china.appstore/internal/config"
)

func callpython() {

	config := config.BaseConfig
	config.InitConfig()

	interval := config.CRAWLER_INTERNAL_TIME
	d, _ := time.ParseDuration(interval)
	for range time.Tick(d) {
		log.Println("------------------------------------------")
		log.Println("[Routine] Running apple_download.py to getting latest downloadcount to csv")
		log.Println("------------------------------------------")
		go func() {
			// Run python
			ExecShell("python3 /root/apple_download.py")
		}()
	}

}
