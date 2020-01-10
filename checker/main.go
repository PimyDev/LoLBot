package checker

import (
	"../file_utils"
	"../lolparser"
	"github.com/alexbyk/panicif"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func postMessage(discord discordgo.Session, channelID string, update lolparser.Update) {
	embedMessage := &discordgo.MessageEmbed{Title: update.Title,
		Description: update.Text,
		Image: &discordgo.MessageEmbedImage{
			URL:      update.ImageUrl,
			ProxyURL: "",
			Width:    640,
			Height:   360,
		},
		URL: update.Url,}
	_, err := discord.ChannelMessageSendEmbed(channelID, embedMessage)
	panicif.Err(err)
}

func readLastUpdateUrl() string {
	var updateUrl string
	filename := "temp.txt"
	if file_utils.FileExists(filename) {
		data, err := ioutil.ReadFile(filename)
		panicif.Err(err)
		updateUrl = string(data)
	} else {
		_, err := os.Create(filename)
		panicif.Err(err)
	}

	return updateUrl
}

func writeLastUpdateUrl(url string) {
	filename := "temp.txt"
	err := ioutil.WriteFile(filename, []byte(url), 0644)
	panicif.Err(err)
}

func Start(discord discordgo.Session, channelID string)  {
	for {
		updates := lolparser.GetLastUpdates()
		update := lolparser.GetUpdateInfo(updates[0])
		lastUpdateUrl := readLastUpdateUrl()
		if update.Url != lastUpdateUrl {
			postMessage(discord, channelID, update)
			writeLastUpdateUrl(update.Url)
			log.Println("Posted new update to channel: " + channelID)
		}
		time.Sleep(5*time.Second)
	}
}

