package airqualitycollector

import (
	"os"

	"github.com/robfig/cron/v3"
	"github.com/withmandala/go-log"
)

var (
	c      *cron.Cron
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stderr)
	c = cron.New()
	c.AddFunc("@every 1m", doAirQualityUpdateRequest)
}

func Run() {
	c.Start()
}
