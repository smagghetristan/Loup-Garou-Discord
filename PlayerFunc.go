package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"Loup.Garou/config"
	Permissions "Loup.Garou/perm"

	"github.com/bwmarrin/discordgo"
)

func RollMessage(Type int, Roll int) (bool, *discordgo.MessageEmbed) {
	embed := &discordgo.MessageEmbed{}
	if Type == 1 {
		if Roll < 45 {
			embed = &discordgo.MessageEmbed{
				Title:       "Bravo !",
				Description: "Ca a marché ! (" + strconv.Itoa(Roll) + "/100)",
			}
			return true, embed
		} else {
			embed = &discordgo.MessageEmbed{
				Title:       "Dommage !",
				Description: "Ca n'a pas marché ! (" + strconv.Itoa(Roll) + "/100)",
			}
			return false, embed
		}
	} else if Type == 2 {
		if Roll < 35 {
			embed = &discordgo.MessageEmbed{
				Title:       "Bravo !",
				Description: "Ca a marché ! (" + strconv.Itoa(Roll) + "/100)",
			}
			return true, embed
		} else {
			embed = &discordgo.MessageEmbed{
				Title:       "Dommage !",
				Description: "Ca n'a pas marché ! (" + strconv.Itoa(Roll) + "/100)",
			}
			return false, embed
		}
	}
	return false, embed
}

func RollDice(Type int, Amount int) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 100
	roll := rand.Intn(max-min) + min
	var PFChan *discordgo.Channel
	exist := false
	for i := range config.CurrentGame.Channels {
		if config.CurrentGame.Channels[i].Name == "petite-fille" {
			PFChan = config.CurrentGame.Channels[i]
			exist = true
			break
		}
	}
	fmt.Println(PFChan)
	if exist {
		if Type == 1 {
			Won, embed := RollMessage(Type, roll)
			fmt.Println(embed)
			fmt.Println(PFChan.ID)
			if Won {
				_, err := dg.ChannelMessageSendEmbed(PFChan.ID, embed)
				if err != nil {
					fmt.Println(err)
				}
				for i := range config.CurrentGame.Channels {
					if config.CurrentGame.Channels[i].Name == "loup-garou" {
						MessageSlice, _ := dg.ChannelMessages(config.CurrentGame.Channels[i].ID, Amount, "", "", "")
						toSend := ""
						for k := range MessageSlice {
							toSend += "**Someone said** : " + MessageSlice[k].Content + `
`
						}
						dg.ChannelMessageSend(PFChan.ID, toSend)
						break
					}
				}
			} else {
				dg.ChannelMessageSendEmbed(PFChan.ID, embed)
			}
		} else if Type == 2 {
			Won, embed := RollMessage(Type, roll)
			if Won {
				dg.ChannelMessageSendEmbed(PFChan.ID, embed)
			} else {
				dg.ChannelMessageSendEmbed(PFChan.ID, embed)
			}
		}

	}
}

func ChienLoupChange() {
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].Role.Name == "Chien-Loup" {
			config.CurrentGame.Players[i].Role.Team.HasChannel = true
		}
	}
}

func MaitreMort() {
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].Role.Name == "Enfant Sauvage" {
			config.CurrentGame.Players[i].Role.Team.HasChannel = true
		}
	}
}

func MuteAll() {
	for i := range config.CurrentGame.Players {
		if !config.CurrentGame.Players[i].ManualMute {
			err := MuteSomeone(config.CurrentGame.Players[i].ID, true)
			if err != nil {
				//
			}
			config.CurrentGame.Players[i].ManualMute = true
			time.Sleep(10 * time.Millisecond)
		} else {
			err := MuteSomeone(config.CurrentGame.Players[i].ID, false)
			if err != nil {
				//
			}
			config.CurrentGame.Players[i].ManualMute = false
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func Mute(PlayerID string) {
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].ID == PlayerID {
			if !config.CurrentGame.Players[i].ManualMute {
				err := MuteSomeone(config.CurrentGame.Players[i].ID, true)
				if err != nil {
					//
				}
				config.CurrentGame.Players[i].ManualMute = true
				time.Sleep(10 * time.Millisecond)
			} else {
				err := MuteSomeone(config.CurrentGame.Players[i].ID, false)
				if err != nil {
					//
				}
				config.CurrentGame.Players[i].ManualMute = false
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}

func Infect(PlayerID string) {
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].ID == PlayerID {
			config.CurrentGame.Players[i].Infected = true
		}
	}
}

func Kill(PlayerID string) {
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].ID == PlayerID {
			dat, err := box.FindString("40x40/" + config.CurrentGame.Players[i].Role.Image)
			check(err)
			Image := &discordgo.File{
				Reader: strings.NewReader(dat),
				Name:   "role.png",
			}
			Params := &discordgo.MessageSend{
				Files: []*discordgo.File{Image},
				Embed: &discordgo.MessageEmbed{
					Title:       "__" + config.CurrentGame.Players[i].Username + "__ est mort(e) !",
					Description: "Il/Elle était : **" + config.CurrentGame.Players[i].Role.Name + "** !",
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: "attachment://role.png",
					},
					Color: 0xFFDD00,
				},
			}
			_, err = dg.ChannelMessageSendComplex(config.CurrentGame.GameStats.ID, Params)
			if err != nil {
				fmt.Println(err)
				return
			}

			config.CurrentGame.Players[i].Role.Name = "Mort"
			config.CurrentGame.Players[i].Role.Image = "mor.png"
			config.CurrentGame.Players[i].Role.ChannelName = "morts"

			err = MuteSomeone(config.CurrentGame.Players[i].ID, true)
			if err != nil {
				//
			}

			TextPerm := &discordgo.ChannelEdit{
				PermissionOverwrites: []*discordgo.PermissionOverwrite{{
					ID:    config.CurrentGame.Players[i].ID,
					Type:  "member",
					Allow: Permissions.VIEW_CHANNEL,
				}},
			}
			_, err = dg.ChannelEditComplex(config.CurrentGame.Deads.ID, TextPerm)
			if err != nil {
				//
			}
		}
	}
}
