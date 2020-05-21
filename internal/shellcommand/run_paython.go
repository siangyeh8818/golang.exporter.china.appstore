package shellcommand

import (
	"log"
	"time"

	configs "github.com/siangyeh8818/golang.exporter.china.appstore/internal/config"
)

func Callpython() {
	log.Println("---------------------Callpython()---------------------")
	config := configs.BaseConfig{}
	(&config).InitConfig()
	log.Println("---------------------InitConfig()---------------------")
	//interval := config.CRAWLER_INTERNAL_TIME
	//intervaltime, _ := strconv.Atoi(config.CRAWLER_INTERNAL_TIME)
	//d, _ := time.ParseDuration(interval)
	//for range time.Tick(d) {
	for {
		log.Println("------------------------------------------")
		log.Println("[Routine] Running apple_download.py to getting latest downloadcount to csv")
		log.Println("------------------------------------------")
		//go func() {
		// Run python
		RunCommand("python3 /root/apple_download.py")
		//}()
		time.Sleep(300 * time.Second)
	}

}
