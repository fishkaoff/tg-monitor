package main

import (
	"context"
	"log"

	config "github.com/fishkaoff/tg-monitor/configs"
	Bot "github.com/fishkaoff/tg-monitor/internal/bot/services"
	"github.com/fishkaoff/tg-monitor/internal/metric"
	db "github.com/fishkaoff/tg-monitor/internal/repository/postgres"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {
	// init config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to init config")
	}

	// logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// init database
	conn, err := pgxpool.New(context.Background(), config.DBURL)
	if err != nil {
		sugar.Fatal("Unable to connect to database: %v\n", err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		sugar.Fatal("Database connection lost")
	}
	defer conn.Close()

	// start bot
	bot, err := tgbotapi.NewBotAPI(config.TGTOKEN)
	if err != nil {
		sugar.Fatal("Error while auth")
	}

	// IoC
	db := db.NewDB(conn)
	mtr := metric.NewMetric()
	tg := Bot.NewBot(bot, mtr, db, sugar)

	sugar.Info("Authentificated with token %v\n", config.TGTOKEN)
	sugar.Info("Starting Bot........")
	tg.Start()
}
