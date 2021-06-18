package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	outfile, _ = os.OpenFile("/var/log/begeek.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	LogFile    = log.New(outfile, "", 0)
)

func ForError(err error) {
	if err != nil {
		text := fmt.Sprintf("%s: %s\n", time.Now().Format("2006.01.02 15:04:05"), err)
		LogFile.Println(text)
	}
}
