package telegram

import (
	"context"
	"database/sql"
	"log"
	"telebot/config"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	// models
	Bot *tgbotapi.BotAPI
}

func NewBot(cfg *config.Config) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}
	if cfg.Log.Level == "DEBUG" {
		bot.Debug = true
	}
	return &Bot{
		Bot: bot,
	}, nil
}

func (b *Bot) Run() error {
	log.Printf("Authorized on account %s", b.Bot.Self.UserName)

	updates := b.InitUpdatesChannel()

	b.listenUpdates(updates)
	return nil
}

func (b *Bot) InitUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.Bot.GetUpdatesChan(u)
}

func (b *Bot) listenUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Handle commands
		if update.Message.IsCommand() {
			// if err := a.Bot.handleCommand(update.Message); err != nil {
			// 	b.handleError(update.Message.Chat.ID, err)
			// }
			b.Bot.Send(tgbotapi.NewMessage(int64(update.Message.Chat.ID), update.Message.Text))
			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			// b.handleError(update.Message.Chat.ID, err)
		}
	}
}

type WeatherResponse struct {
	Location struct {
		Name           string  `json:"name,omitempty"`
		Region         string  `json:"region,omitempty"`
		Country        string  `json:"country,omitempty"`
		Lat            float64 `json:"lat,omitempty"`
		Lon            float64 `json:"lon,omitempty"`
		TzID           string  `json:"tz_id,omitempty"`
		LocaltimeEpoch int     `json:"localtime_epoch,omitempty"`
		Localtime      string  `json:"localtime,omitempty"`
	} `json:"location,omitempty"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch,omitempty"`
		LastUpdated      string  `json:"last_updated,omitempty"`
		TempC            float64 `json:"temp_c,omitempty"`
		TempF            float64 `json:"temp_f,omitempty"`
		IsDay            int     `json:"is_day,omitempty"`
		Condition        struct {
			Text string `json:"text,omitempty"`
			Icon string `json:"icon,omitempty"`
			Code int    `json:"code,omitempty"`
		} `json:"condition,omitempty"`
		WindMph    float64 `json:"wind_mph,omitempty"`
		WindKph    float64 `json:"wind_kph,omitempty"`
		WindDegree int     `json:"wind_degree,omitempty"`
		WindDir    string  `json:"wind_dir,omitempty"`
		PressureMb float64 `json:"pressure_mb,omitempty"`
		PressureIn float64 `json:"pressure_in,omitempty"`
		PrecipMm   float64 `json:"precip_mm,omitempty"`
		PrecipIn   float64 `json:"precip_in,omitempty"`
		Humidity   int     `json:"humidity,omitempty"`
		Cloud      int     `json:"cloud,omitempty"`
		FeelslikeC float64 `json:"feelslike_c,omitempty"`
		FeelslikeF float64 `json:"feelslike_f,omitempty"`
		VisKm      float64 `json:"vis_km,omitempty"`
		VisMiles   float64 `json:"vis_miles,omitempty"`
		Uv         float64 `json:"uv,omitempty"`
		GustMph    float64 `json:"gust_mph,omitempty"`
		GustKph    float64 `json:"gust_kph,omitempty"`
	} `json:"current,omitempty"`
}

func openDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)

	// time.ParseDuration() function to convert the idle timeout duration string
	// to a time.Duration type.
	duration, err := time.ParseDuration(cfg.DB.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
