package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strings"

	"github.com/cirocosta/syslog-send/priority"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Address   string `long:"address" description:"address of the syslog server" default:"127.0.0.1:514"`
	Transport string `long:"transport" description:"transport to use (tcp|udp)" default:"udp"`
	Facility  string `long:"facility" description:"name of the syslog facility to send msgs to" default:"local0"`
	Severity  string `long:"severity" description:"severity of the message" default:"emerg"`
	Args      struct {
		Message []string
	} `positional-args:"yes" required:"yes"`
}

func must(err error) {
	if err == nil {
		return
	}

	log.Fatal(err)
}

func main() {
	var syslogPriority syslog.Priority

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	switch opts.Transport {
	case "udp":
	case "tcp":
	default:
		must(fmt.Errorf("Unknown transport %s", opts.Transport))
	}

	syslogPriority, err = priority.NewPriority(opts.Facility, opts.Severity)
	must(err)

	logger, err := syslog.Dial(opts.Transport, opts.Address,
		syslogPriority, "custom-tag")
	must(err)

	defer logger.Close()

	fmt.Fprintf(logger, strings.Join(opts.Args.Message, " "))
}
