/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)
var TeleToken  = os.Getenv("TELE_TOKEN")

var tbotCmd = &cobra.Command{
	Use:   "tbot",
	Aliases: []string{"go"},
	Short: "Start tg bot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tbot called")
		tbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}
		
		tbot.Handle(telebot.OnText,  func(m telebot.Context) error {
			log.Print(m.Text())
			if strings.HasPrefix(strings.ToLower(m.Text()), "/addtimezone") {
				parts := strings.Fields(m.Text())
				if len(parts) < 2 {
					tbot.Send(m.Sender(), "Usage: /addtimezone <timezone>")
					return nil
				}
				timezone := parts[1]
				loc, err := time.LoadLocation(timezone)
				if err != nil {
					tbot.Send(m.Sender(), fmt.Sprintf("Failed to load timezone %s: %s", timezone, err))
					return nil
				}
				timezones = append(timezones, timezone)
				timezoneMap[timezone] = loc
				tbot.Send(m.Sender(), fmt.Sprintf("Timezone %s added successfully!", timezone))
			} else if strings.ToLower(m.Text()) == "/time" {
				utcTime := time.Now().UTC()
				timeFormat := "15:04:05 (MST)"
				var timeStrings []string
				for _, loc := range timezones {
					timeStrings = append(timeStrings, fmt.Sprintf("%s: %s", loc, utcTime.In(timezoneMap[loc]).Format(timeFormat)))
				}
				tbot.Send(m.Sender(), strings.Join(timeStrings, "\n"))
			} else if strings.ToLower(m.Text()) =="/start"{
				 m.Send("Hello I'm tbot!\nUse /time for display time\n/addtimezone <timezone> for adding timezone")
			}
			return nil
		})
	
		tbot.Start()

	},
}

var timezones = []string{
	"America/New_York",
	"Europe/Kiev",
	"Australia/Sydney",
}

var timezoneMap = map[string]*time.Location{}

func init() {
	for _, tz := range timezones {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			log.Fatalf("Failed to load location %s: %s", tz, err)
		}
		timezoneMap[tz] = loc
	}
	rootCmd.AddCommand(tbotCmd)
}
