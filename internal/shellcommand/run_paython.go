package shellcommand

import (
	"log"
	"os"
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

		internal_time, _ := time.ParseDuration(os.Getenv("CRAWLER_INTERNAL_TIME"))

		RunCommand("python3 /root/apple_download.py")
		//}()
		time.Sleep(time.Duration(internal_time))
	}

}
