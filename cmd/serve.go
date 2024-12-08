/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/centrifugal/gocent/v3"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"nevissGo/app/endpoint"
	"nevissGo/app/service"
	"nevissGo/ent"
	"nevissGo/framework"
	"nevissGo/telegram"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve",
	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load()

		client, err := ent.Open("mysql", os.Getenv("MYSQL_DSN"))
		if err != nil {
			logrus.WithError(err).Fatal("failed opening connection to mysql")
		}

		defer client.Close()

		if err := client.Schema.Create(context.Background()); err != nil {
			logrus.WithError(err).Fatal("failed creating schema resources")
		}

		logrus.Info("mysql connection established")

		// SETUP APP
		app := framework.NewApp(
			client,
			framework.NewCentrifugoClient(
				gocent.New(gocent.Config{
					Addr: os.Getenv("CENTRIFUGO_ADDR_API"),
					Key:  os.Getenv("CENTRIFUGO_SECRET_KEY"),
				}),
			),
			framework.Config{
				Addr: ":8001",
			},
		)

		hypeService := service.NewHype(app)

		bridge := service.Bridge{
			Hype: hypeService,
		}

		app.RegisterEndpoints(
			endpoint.NewUsers(service.NewUsers(app)),
			endpoint.NewPixels(service.NewPixels(app, bridge, time.Microsecond, 40, 40, 1)),
			endpoint.NewHype(hypeService),
			endpoint.NewOnlineUsers(service.NewOnlineUsers(app)),
		)

		// SETUP BOT
		bot, err := telegram.NewTelegram()
		if err != nil {
			logrus.WithError(err).Fatal("failed creating telegram bot")
		}

		go func() {
			logrus.Info("starting telegram bot")
			bot.Start()
		}()

		logrus.WithError(app.ServeEndpoints()).Error("failed serving endpoints")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
