package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/josepmdc/goboilerplate/logger"

	"github.com/josepmdc/goboilerplate/api"
	"github.com/josepmdc/goboilerplate/app"
	"github.com/josepmdc/goboilerplate/conf"

	"github.com/spf13/cobra"
)

// RootCommand is the main command for starting the server. It reads the command flags port and config
func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "[options]",
		Run: run,
	}

	rootCmd.Flags().IntP("port", "p", 0, "--port [number] | -p [number]")
	rootCmd.PersistentFlags().StringP("config", "c", "", "--config [path] | -c [path]")

	return &rootCmd
}

func run(cmd *cobra.Command, _ []string) {
	config, err := conf.LoadConfig(cmd)
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	err = logger.ConfigureLogger(&config.LogConfig)
	if err != nil {
		panic("Failed to configure logger: " + err.Error())
	}

	s, err := app.InitServices(config)
	if err != nil {
		logger.Logger.Fatalf("Couldn't initialize services: %s", err.Error())
	}
	startServer(config, s)
}

func startServer(config *conf.Config, services *app.Services) {
	handler := api.NewRouter(services)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Domain, config.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	logger.Logger.Errorf("Error starting the server: %s", srv.ListenAndServe().Error())
}
