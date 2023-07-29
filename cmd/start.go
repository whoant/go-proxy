package cmd

import (
	"github.com/spf13/cobra"
	"github.com/whoant/go-proxy/internal/proxy_server_service"
)

var port int
var username string
var password string

func init() {
	startCmd.PersistentFlags().IntVar(&port, "port", 8080, "port")
	startCmd.PersistentFlags().StringVar(&username, "username", "whoant", "username")
	startCmd.PersistentFlags().StringVar(&password, "password", "whoant", "password")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "Start proxy server",
	Long:    `Start proxy server`,
	Example: "start --port 8080 --username whoant --password whoant",
	Run: func(cmd *cobra.Command, args []string) {
		service := proxy_server_service.NewProxyServerService(port, username, password)
		service.Start()
	},
}
