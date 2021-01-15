package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

var ConfigFile string

func main() {
	args := os.Args

	fs := pflag.NewFlagSet("bdk-k8s-daemon", pflag.ExitOnError)
	fs.StringVar(&ConfigFile, "config", "config.yaml", "The path to configuration file (default \"/etc/bdk-k8s/config.yaml\")")
	fs.BoolP("help", "h", false, "help and exit")

	fs.Parse(args)

	fmt.Println(ConfigFile)
}
