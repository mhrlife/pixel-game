package cmd

import (
	"context"
	"github.com/centrifugal/gocent/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/mhrlife/tonference/internal/app/endpoint"
	"github.com/mhrlife/tonference/internal/app/service"
	"github.com/mhrlife/tonference/internal/ent"
	"github.com/mhrlife/tonference/internal/telegram"
	"github.com/mhrlife/tonference/pkg/framework"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
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

		liveKitRoomService := lksdk.NewRoomServiceClient(
			os.Getenv("LIVEKIT_HOST"),
			os.Getenv("LIVEKIT_API_KEY"),
			os.Getenv("LIVEKIT_API_SECRET"),
		)

		tonferenceService := service.NewService(client, app, liveKitRoomService)

		app.RegisterEndpoints(
			endpoint.NewUsers(tonferenceService),
			endpoint.NewRooms(tonferenceService),
		)

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
