package main

import (
	"Test/config"
	Permissions "Test/perm"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Clear(channelID string) {
	number := 99
	if number < 100 {
		MessageSlice, err := dg.ChannelMessages(channelID, number, "", "", "")
		if err != nil {
			dg.ChannelMessageSend(channelID, "I had some problem deleting the messages.")
		} else {
			var Params []string
			for i := 0; i < len(MessageSlice); i++ {
				Params = append(Params, MessageSlice[i].ID)
			}
			dg.ChannelMessagesBulkDelete(channelID, Params)
		}
	} else {
		dg.ChannelMessageSend(channelID, "I cannot delete more than 100 messages.")
	}
}

func DayPerm() {
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].Role.Name != "MDJ" {
			if config.CurrentGame.Players[i].Role.Name != "Mort" {
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

			for k := range config.CurrentGame.Channels {

				if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
					TextPerm := &discordgo.ChannelEdit{
						PermissionOverwrites: []*discordgo.PermissionOverwrite{{
							ID:   config.CurrentGame.Players[i].ID,
							Type: "member",
							Deny: Permissions.VIEW_CHANNEL,
						}},
					}
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}

				if config.CurrentGame.Players[i].Role.ChannelName == config.CurrentGame.Channels[k].Name {
					TextPerm := &discordgo.ChannelEdit{
						PermissionOverwrites: []*discordgo.PermissionOverwrite{{
							ID:   config.CurrentGame.Players[i].ID,
							Type: "member",
							Deny: Permissions.VIEW_CHANNEL,
						}},
					}
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}
				if config.CurrentGame.Players[i].Role.Team.HasChannel && config.CurrentGame.Players[i].Role.Team.ChannelName == config.CurrentGame.Channels[k].Name {
					TextPerm := &discordgo.ChannelEdit{
						PermissionOverwrites: []*discordgo.PermissionOverwrite{{
							ID:   config.CurrentGame.Players[i].ID,
							Type: "member",
							Deny: Permissions.VIEW_CHANNEL,
						}},
					}
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}
			}
		}
	}
}

func NightPerm() {
	Clear(config.CurrentGame.Votes.ID)
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].Role.Name != "MDJ" {
			if config.CurrentGame.Players[i].Role.Name != "Mort" {
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
			}
			time.Sleep(10 * time.Millisecond)

			for k := range config.CurrentGame.Channels {

				if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
					TextPerm := &discordgo.ChannelEdit{
						PermissionOverwrites: []*discordgo.PermissionOverwrite{{
							ID:    config.CurrentGame.Players[i].ID,
							Type:  "member",
							Allow: Permissions.VIEW_CHANNEL,
						}},
					}
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}

				if config.CurrentGame.Players[i].Role.ChannelName == config.CurrentGame.Channels[k].Name {
					TextPerm := &discordgo.ChannelEdit{}
					if config.CurrentGame.Players[i].Role.Name != "Chaman" {
						TextPerm = &discordgo.ChannelEdit{
							PermissionOverwrites: []*discordgo.PermissionOverwrite{{
								ID:    config.CurrentGame.Players[i].ID,
								Type:  "member",
								Allow: Permissions.VIEW_CHANNEL,
							}},
						}
					} else {
						Clear(config.CurrentGame.Deads.ID)
						TextPerm = &discordgo.ChannelEdit{
							PermissionOverwrites: []*discordgo.PermissionOverwrite{{
								ID:    config.CurrentGame.Players[i].ID,
								Type:  "member",
								Allow: Permissions.VIEW_CHANNEL,
								Deny:  Permissions.SEND_MESSAGES + Permissions.READ_MESSAGE_HISTORY,
							}},
						}
					}
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}
				if config.CurrentGame.Players[i].Role.Team.HasChannel && config.CurrentGame.Players[i].Role.Team.ChannelName == config.CurrentGame.Channels[k].Name {
					TextPerm := &discordgo.ChannelEdit{
						PermissionOverwrites: []*discordgo.PermissionOverwrite{{
							ID:    config.CurrentGame.Players[i].ID,
							Type:  "member",
							Allow: Permissions.VIEW_CHANNEL,
						}},
					}
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}
			}
		}
	}
}

func DayPermPlayer(PlayerID string) {
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].ID == PlayerID {
			if config.CurrentGame.Players[i].Role.Name != "MDJ" {
				if config.CurrentGame.Players[i].Role.Name != "Mort" {
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

				for k := range config.CurrentGame.Channels {

					if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
						TextPerm := &discordgo.ChannelEdit{
							PermissionOverwrites: []*discordgo.PermissionOverwrite{{
								ID:   config.CurrentGame.Players[i].ID,
								Type: "member",
								Deny: Permissions.VIEW_CHANNEL,
							}},
						}
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
						if err != nil {
							//
						}
						time.Sleep(10 * time.Millisecond)
					}

					if config.CurrentGame.Players[i].Role.ChannelName == config.CurrentGame.Channels[k].Name {
						TextPerm := &discordgo.ChannelEdit{
							PermissionOverwrites: []*discordgo.PermissionOverwrite{{
								ID:   config.CurrentGame.Players[i].ID,
								Type: "member",
								Deny: Permissions.VIEW_CHANNEL,
							}},
						}
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
						if err != nil {
							//
						}
						time.Sleep(10 * time.Millisecond)
					}
					if config.CurrentGame.Players[i].Role.Team.HasChannel && config.CurrentGame.Players[i].Role.Team.ChannelName == config.CurrentGame.Channels[k].Name {
						TextPerm := &discordgo.ChannelEdit{
							PermissionOverwrites: []*discordgo.PermissionOverwrite{{
								ID:   config.CurrentGame.Players[i].ID,
								Type: "member",
								Deny: Permissions.VIEW_CHANNEL,
							}},
						}
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
						if err != nil {
							//
						}
						time.Sleep(10 * time.Millisecond)
					}
				}
			}
		}
	}
}

func NightPermPlayer(PlayerID string) {
	Clear(config.CurrentGame.Votes.ID)
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].ID == PlayerID {
			if config.CurrentGame.Players[i].Role.Name != "MDJ" {
				if config.CurrentGame.Players[i].Role.Name != "Mort" {
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
				}

				for k := range config.CurrentGame.Channels {

					if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
						TextPerm := &discordgo.ChannelEdit{
							PermissionOverwrites: []*discordgo.PermissionOverwrite{{
								ID:    config.CurrentGame.Players[i].ID,
								Type:  "member",
								Allow: Permissions.VIEW_CHANNEL,
							}},
						}
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
						if err != nil {
							//
						}
						time.Sleep(10 * time.Millisecond)
					}

					if config.CurrentGame.Players[i].Role.ChannelName == config.CurrentGame.Channels[k].Name {
						TextPerm := &discordgo.ChannelEdit{}
						if config.CurrentGame.Players[i].Role.Name != "Chaman" {
							TextPerm = &discordgo.ChannelEdit{
								PermissionOverwrites: []*discordgo.PermissionOverwrite{{
									ID:    config.CurrentGame.Players[i].ID,
									Type:  "member",
									Allow: Permissions.VIEW_CHANNEL,
								}},
							}
						} else {
							Clear(config.CurrentGame.Deads.ID)
							TextPerm = &discordgo.ChannelEdit{
								PermissionOverwrites: []*discordgo.PermissionOverwrite{{
									ID:    config.CurrentGame.Players[i].ID,
									Type:  "member",
									Allow: Permissions.VIEW_CHANNEL,
									Deny:  Permissions.SEND_MESSAGES + Permissions.READ_MESSAGE_HISTORY,
								}},
							}
						}
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
						if err != nil {
							//
						}
						time.Sleep(10 * time.Millisecond)
					}
					if config.CurrentGame.Players[i].Role.Team.HasChannel && config.CurrentGame.Players[i].Role.Team.ChannelName == config.CurrentGame.Channels[k].Name {
						TextPerm := &discordgo.ChannelEdit{
							PermissionOverwrites: []*discordgo.PermissionOverwrite{{
								ID:    config.CurrentGame.Players[i].ID,
								Type:  "member",
								Allow: Permissions.VIEW_CHANNEL,
							}},
						}
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, TextPerm)
						if err != nil {
							//
						}
						time.Sleep(10 * time.Millisecond)
					}
				}
			}
		}
	}
}
