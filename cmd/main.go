package main

import (
	"flag"
	"fmt"
	"sync"

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

func init() {
	showBanner()

	parseParams()

	conf.launchList = make(map[string]fn)

	conf.launchList["server"] = func() { defer wg.Done(); server.Run(conf.ip, conf.port) }
	// launchList["collector"] = func() { defer wg.Done(); }
}

func main() {
	for _, fn := range conf.launchList {
		wg.Add(1)
		go fn()
	}

	wg.Wait()
}
