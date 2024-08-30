package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	initEnv()

	discordBot, err := discordgo.New("Bot " + ENVIRONMENT.BotToken)
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

	dsn := "host=" + ENVIRONMENT.DatabaseHost + " port=" + ENVIRONMENT.DatabasePort + " user=" + ENVIRONMENT.DatabaseUser + " dbname=" + ENVIRONMENT.DatabaseName + " password=" + ENVIRONMENT.DatabasePassword + " sslmode=disable"
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
