package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type environmentVariables struct {
	BotToken         string `env:"BOT_TOKEN,required"`
	LogLevel         string `env:"LOG_LEVEL" envDefault:"info"`
	DatabaseHost     string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort     string `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseName     string `env:"DATABASE_NAME" envDefault:"legbot"`
	DatabaseUser     string `env:"DATABASE_USER" envDefault:"postgres"`
	DatabasePassword string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	CronApiUrl       string `env:"CRON_API_URL" envDefault:"https://natural-cron-api.fly.dev"`
}

var ENVIRONMENT *environmentVariables

func main() {
	envVars := &environmentVariables{}
	err := env.Parse(envVars)
	if err != nil {
		log.Fatalf("Error parsing environment variables: %v", err)
	}
	ENVIRONMENT = envVars

	discordBot, err := discordgo.New("Bot " + envVars.BotToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	err = discordBot.Open()
	if err != nil {
		log.Fatalf("Error opening connection to Discord: %v", err)
	}

	defer discordBot.Close()

	discordBot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		//if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		//	handler(s, i)
		//}
	})

	dsn := "host=" + envVars.DatabaseHost + " port=" + envVars.DatabasePort + " user=" + envVars.DatabaseUser + " dbname=" + envVars.DatabaseName + " password=" + envVars.DatabasePassword + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	loadSchedules(db, cron.New())

	log.Print("Bot is running. Press CTRL-C to exit.")
}

func loadSchedules(db *gorm.DB, scheduler *cron.Cron) {
	// Load schedules from the database
	var schedules []Schedule
	result := []map[string]interface{}{}
	db.Model(&schedules).Select("cron").Find(&result)

	// Iterate over schedules and schedule them
	for _, schedule := range schedules {
		scheduler.AddFunc(schedule.Cron, func() {
			remindUser(db, schedule.DiscordID, schedule.ID)
		})
	}
}

func remindUser(db *gorm.DB, discordID string, scheduleID uint) {
	// Send a reminder to the user
}
