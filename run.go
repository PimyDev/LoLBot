package main

import (
	"./checker"
	"./config"
	"github.com/alexbyk/panicif"
	"github.com/bwmarrin/discordgo"
)

func main()  {
	cfg := config.LoadConfig()
	discord, err := discordgo.New("Bot " + cfg.Token)
	panicif.Err(err)
	checker.Start(*discord, cfg.ChannelID)
}
