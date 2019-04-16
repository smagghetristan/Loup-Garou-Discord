package main

import (
	"Test/config"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GameStart() {
	config.CurrentGame = config.Game{
		Started:  true,
		Finished: false,
		Players:  config.Players,
	}
}

func ChannelGenerator() {
	toCreate := []string{}
	for i := range config.CurrentGame.Players {
		existNormal := false
		existTeam := false
		for k := range toCreate {
			if toCreate[k] == config.CurrentGame.Players[i].Role.ChannelName {
				existNormal = true
			}
			if toCreate[k] == config.CurrentGame.Players[i].Role.ChannelName {
				existTeam = true
				break
			} else {
				break
			}
		}
		if !existNormal {
			toCreate = append(toCreate, config.CurrentGame.Players[i].Role.ChannelName)
		}
		if !existTeam {
			toCreate = append(toCreate, config.CurrentGame.Players[i].Role.Team.ChannelName)
		}
	}

	for i := range toCreate {
		data := discordgo.GuildChannelCreateData{
			Name:     toCreate[i],
			Type:     discordgo.ChannelTypeGuildText,
			ParentID: config.CategoryID,
		}
		channel, err := dg.GuildChannelCreateComplex(config.GuildID, data)
		if err != nil {
			//
		}
		config.CurrentGame.Channels = append(config.CurrentGame.Channels, channel)
		time.Sleep(10 * time.Millisecond)
	}
}

func DMSender() {
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].Role.Name == "MDJ" {
			dm, err := dg.UserChannelCreate(config.CurrentGame.Players[i].ID)
			if err != nil {
				fmt.Println(err)
				return
			}
			embed := &discordgo.MessageEmbed{
				Title:       "Loup-Garou",
				Description: "Les rôles ont étés distribués",
				Color:       0xFFDD00,
			}
			_, err = dg.ChannelMessageSendEmbed(dm.ID, embed)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			dm, err := dg.UserChannelCreate(config.CurrentGame.Players[i].ID)
			if err != nil {
				fmt.Println(err)
				return
			}
			dat, err := box.FindString("static/" + config.CurrentGame.Players[i].Role.Image)
			check(err)
			Image := &discordgo.File{
				Reader: strings.NewReader(dat),
				Name:   "role.png",
			}
			Params := &discordgo.MessageSend{
				Files: []*discordgo.File{Image},
				Embed: &discordgo.MessageEmbed{
					Title:       config.CurrentGame.Players[i].Role.Name,
					Description: config.CurrentGame.Players[i].Role.Description,
					Image: &discordgo.MessageEmbedImage{
						URL: "attachment://role.png",
					},
					Color: 0xFFDD00,
				},
			}
			_, err = dg.ChannelMessageSendComplex(dm.ID, Params)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
}

func GameBegin(Data Received) {
	for i := range Data.Game.Players {
		for k := range config.CurrentGame.Players {
			if config.CurrentGame.Players[k].ID == Data.Game.Players[i].ID {
				rl, error := config.GetRoleByName(Data.Game.Players[i].RoleName)
				if error {
					return
				}
				config.CurrentGame.Players[k].Role = rl
			}
		}
	}
	//go ChannelGenerator()
	go DMSender()
}
