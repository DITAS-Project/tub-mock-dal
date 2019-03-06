package main

import (
	"flag"

	"github.com/DITAS-Project/tub-mock-dal/dal"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {

}

func cliSetup() {
	viper.SetDefault("port", 8080)

	flag.Int("port", 8080, "set the port to listen to")
	flag.Bool("strict", false, "validate access token before responding to requests")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

}

func main() {
	cliSetup()
	setupServer()
}

func setupServer() {
	server := dal.NewDalServer(viper.GetBool("strict"), viper.GetInt("port"))
	server.StartAndListen()
}
