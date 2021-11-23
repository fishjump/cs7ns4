package main

import (
	"flag"
	"fmt"
	"sync"

	airqualitycollector "github.com/fishjump/cs7ns4/modules/air-quality-collector"
	"github.com/fishjump/cs7ns4/modules/server"
)

type (
	fn func()

	config struct {
		ip         string
		port       int
		launchList map[string]fn
	}
)

var (
	wg   sync.WaitGroup
	conf config
)

func showBanner() {
	fmt.Print(`
	██╗   ██╗ ██████╗    ██╗  ██╗    ███████╗███████╗██████╗ ██╗   ██╗███████╗██████╗ 
	██║   ██║██╔════╝    ██║  ██║    ██╔════╝██╔════╝██╔══██╗██║   ██║██╔════╝██╔══██╗
	██║   ██║██║         ███████║    ███████╗█████╗  ██████╔╝██║   ██║█████╗  ██████╔╝
	██║   ██║██║         ╚════██║    ╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██╔══╝  ██╔══██╗
	╚██████╔╝╚██████╗         ██║    ███████║███████╗██║  ██║ ╚████╔╝ ███████╗██║  ██║
	 ╚═════╝  ╚═════╝         ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝                                                                         													
`)
}

func parseParams() {
	flag.StringVar(&conf.ip, "ip", "0.0.0.0", "server ip")
	flag.IntVar(&conf.port, "port", 8888, "server port")
}

func launchWrapper(f fn) fn {
	return func() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}
}

func init() {
	showBanner()

	parseParams()

	conf.launchList = make(map[string]fn)

	conf.launchList["server"] = launchWrapper(func() { server.Run(conf.ip, conf.port) })
	conf.launchList["collector"] = launchWrapper(func() { airqualitycollector.Run() })
}

func main() {
	for _, f := range conf.launchList {
		f()
	}

	wg.Wait()
}
