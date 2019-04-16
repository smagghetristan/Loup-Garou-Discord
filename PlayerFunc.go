package main

import (
	"Test/config"
	Permissions "Test/perm"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

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
	if exist {
		if Type == 1 {
			Won, embed := RollMessage(Type, roll)
			if Won {
				dg.ChannelMessageSendEmbed(PFChan.ID, embed)
				for i := range config.CurrentGame.Channels {
					if config.CurrentGame.Channels[i].Name == "loup-garou" {
						MessageSlice, _ := dg.ChannelMessages(config.CurrentGame.Channels[i].ID, Amount, "", "", "")
						toSend := ""
						for k := range MessageSlice {
							toSend += MessageSlice[k].Content + `
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
			VoicePerm := &discordgo.ChannelEdit{
				PermissionOverwrites: []*discordgo.PermissionOverwrite{{
					ID:   config.CurrentGame.Players[i].ID,
					Type: "member",
					Deny: Permissions.SPEAK,
				}},
			}
			_, err := dg.ChannelEditComplex(config.VoiceChannel, VoicePerm)
			if err != nil {
				//
			}
			time.Sleep(10 * time.Millisecond)
		} else {
			VoicePerm := &discordgo.ChannelEdit{
				PermissionOverwrites: []*discordgo.PermissionOverwrite{{
					ID:    config.CurrentGame.Players[i].ID,
					Type:  "member",
					Allow: Permissions.SPEAK,
				}},
			}
			_, err := dg.ChannelEditComplex(config.VoiceChannel, VoicePerm)
			if err != nil {
				//
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func Mute(PlayerID string) {
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].ID == PlayerID {
			if !config.CurrentGame.Players[i].ManualMute {
				VoicePerm := &discordgo.ChannelEdit{
					PermissionOverwrites: []*discordgo.PermissionOverwrite{{
						ID:   config.CurrentGame.Players[i].ID,
						Type: "member",
						Deny: Permissions.SPEAK,
					}},
				}
				_, err := dg.ChannelEditComplex(config.VoiceChannel, VoicePerm)
				if err != nil {
					//
				}
				time.Sleep(10 * time.Millisecond)
			} else {
				VoicePerm := &discordgo.ChannelEdit{
					PermissionOverwrites: []*discordgo.PermissionOverwrite{{
						ID:    config.CurrentGame.Players[i].ID,
						Type:  "member",
						Allow: Permissions.SPEAK,
					}},
				}
				_, err := dg.ChannelEditComplex(config.VoiceChannel, VoicePerm)
				if err != nil {
					//
				}
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
			config.CurrentGame.Players[i].Role.Name = "Mort"
			config.CurrentGame.Players[i].Role.Image = "mor.png"
			config.CurrentGame.Players[i].Role.ChannelName = "morts"
			emoji := strings.Replace(config.CurrentGame.Players[i].Role.Image, ".png", "", 1)
			_, err := dg.ChannelMessageSend(config.CurrentGame.GameStats.ID, "**"+config.CurrentGame.Players[i].Username+"** est mort(e) et était : :"+emoji+":")
			if err != nil {
				fmt.Println(err)
				return
			}

			VoicePerm := &discordgo.ChannelEdit{
				PermissionOverwrites: []*discordgo.PermissionOverwrite{{
					ID:   config.CurrentGame.Players[i].ID,
					Type: "member",
					Deny: Permissions.SPEAK,
				}},
			}
			_, err = dg.ChannelEditComplex(config.VoiceChannel, VoicePerm)
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
