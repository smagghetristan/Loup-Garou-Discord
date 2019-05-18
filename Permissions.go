package main

import (
	"time"

	Permissions "Loup.Garou/perm"

	"Loup.Garou/config"

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
	Clear(config.CurrentGame.Votes.ID)
	for i := range config.CurrentGame.Channels {
		CurrentPerm := &discordgo.ChannelEdit{
			PermissionOverwrites: []*discordgo.PermissionOverwrite{{
				ID:   config.GuildID,
				Type: "role",
				Deny: Permissions.VIEW_CHANNEL,
			}, {
				ID:    dg.State.User.ID,
				Type:  "member",
				Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
			}},
		}
		for k := range config.CurrentGame.Players {
			if config.CurrentGame.Players[k].Role.Name != "Mort" && config.CurrentGame.Players[k].Role.Name != "MDJ" {
				err := MuteSomeone(config.CurrentGame.Players[k].ID, false)
				if err != nil {
					//
				}
			}
			time.Sleep(config.SleepTime * time.Millisecond)
			if config.CurrentGame.Channels[i].Name == config.CurrentGame.Players[k].Role.ChannelName || config.CurrentGame.Channels[i].Name == config.CurrentGame.Players[k].Role.Team.ChannelName {
				CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:   config.CurrentGame.Players[k].ID,
					Type: "member",
					Deny: Permissions.VIEW_CHANNEL,
				})
			} else if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
				CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:   config.CurrentGame.Players[k].ID,
					Type: "member",
					Deny: Permissions.VIEW_CHANNEL,
				})
			}
		}
		_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[i].ID, CurrentPerm)
		if err != nil {
			//
		}
		time.Sleep(config.SleepTime * time.Millisecond)
	}

	VotePerm := &discordgo.ChannelEdit{
		PermissionOverwrites: []*discordgo.PermissionOverwrite{{
			ID:   config.GuildID,
			Type: "role",
			Deny: Permissions.VIEW_CHANNEL,
		}, {
			ID:    dg.State.User.ID,
			Type:  "member",
			Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
		}},
	}

	CurrentPerm := &discordgo.ChannelEdit{
		PermissionOverwrites: []*discordgo.PermissionOverwrite{{
			ID:   config.GuildID,
			Type: "role",
			Deny: Permissions.VIEW_CHANNEL,
		}, {
			ID:    dg.State.User.ID,
			Type:  "member",
			Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
		}},
	}
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].Role.Name != "MDJ" {
			if config.CurrentGame.Players[i].Role.Name == "Mort" {
				CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:    config.CurrentGame.Players[i].ID,
					Type:  "member",
					Allow: Permissions.VIEW_CHANNEL,
				})
				VotePerm.PermissionOverwrites = append(VotePerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:    config.CurrentGame.Players[i].ID,
					Type:  "member",
					Allow: Permissions.VIEW_CHANNEL,
					Deny:  Permissions.SEND_MESSAGES,
				})
			} else {
				CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:   config.CurrentGame.Players[i].ID,
					Type: "member",
					Deny: Permissions.VIEW_CHANNEL,
				})
			}
		}
	}
	_, err := dg.ChannelEditComplex(config.CurrentGame.Deads.ID, CurrentPerm)
	if err != nil {
		//
	}
	time.Sleep(config.SleepTime * time.Millisecond)
	_, err = dg.ChannelEditComplex(config.CurrentGame.Votes.ID, VotePerm)
	if err != nil {
		//
	}
}

func NightPerm() {
	Clear(config.CurrentGame.Votes.ID)
	for i := range config.CurrentGame.Channels {
		CurrentPerm := &discordgo.ChannelEdit{
			PermissionOverwrites: []*discordgo.PermissionOverwrite{{
				ID:   config.GuildID,
				Type: "role",
				Deny: Permissions.VIEW_CHANNEL,
			}, {
				ID:    dg.State.User.ID,
				Type:  "member",
				Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
			}},
		}
		for k := range config.CurrentGame.Players {
			if config.CurrentGame.Players[k].Role.Name != "Mort" && config.CurrentGame.Players[k].Role.Name != "MDJ" {
				err := MuteSomeone(config.CurrentGame.Players[k].ID, true)
				if err != nil {
					//
				}
			}
			time.Sleep(config.SleepTime * time.Millisecond)
			if config.CurrentGame.Players[k].Role.Name == "Chaman" {
				CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:    config.CurrentGame.Players[k].ID,
					Type:  "member",
					Allow: Permissions.VIEW_CHANNEL,
					Deny:  Permissions.SEND_MESSAGES + Permissions.READ_MESSAGE_HISTORY,
				})
			} else if config.CurrentGame.Channels[i].Name == config.CurrentGame.Players[k].Role.ChannelName || config.CurrentGame.Channels[i].Name == config.CurrentGame.Players[k].Role.Team.ChannelName {
				CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:    config.CurrentGame.Players[k].ID,
					Type:  "member",
					Allow: Permissions.VIEW_CHANNEL,
				})
			} else if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
				CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:    config.CurrentGame.Players[k].ID,
					Type:  "member",
					Allow: Permissions.VIEW_CHANNEL,
				})
			}
		}
		_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[i].ID, CurrentPerm)
		if err != nil {
			//
		}
		time.Sleep(config.SleepTime * time.Millisecond)
	}

	VotePerm := &discordgo.ChannelEdit{
		PermissionOverwrites: []*discordgo.PermissionOverwrite{{
			ID:   config.GuildID,
			Type: "role",
			Deny: Permissions.VIEW_CHANNEL,
		}, {
			ID:    dg.State.User.ID,
			Type:  "member",
			Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
		}},
	}

	CurrentPerm := &discordgo.ChannelEdit{
		PermissionOverwrites: []*discordgo.PermissionOverwrite{{
			ID:   config.GuildID,
			Type: "role",
			Deny: Permissions.VIEW_CHANNEL,
		}, {
			ID:    dg.State.User.ID,
			Type:  "member",
			Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
		}},
	}
	for i := range config.CurrentGame.Players {
		if config.CurrentGame.Players[i].Role.Name != "MDJ" {
			if config.CurrentGame.Players[i].Role.Name == "Mort" {
				CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:    config.CurrentGame.Players[i].ID,
					Type:  "member",
					Allow: Permissions.VIEW_CHANNEL,
				})
				VotePerm.PermissionOverwrites = append(VotePerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:    config.CurrentGame.Players[i].ID,
					Type:  "member",
					Allow: Permissions.VIEW_CHANNEL,
					Deny:  Permissions.SEND_MESSAGES,
				})
			} else {
				CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
					ID:   config.CurrentGame.Players[i].ID,
					Type: "member",
					Deny: Permissions.VIEW_CHANNEL,
				})
			}
		}
	}
	_, err := dg.ChannelEditComplex(config.CurrentGame.Deads.ID, CurrentPerm)
	if err != nil {
		//
	}
	time.Sleep(config.SleepTime * time.Millisecond)
	_, err = dg.ChannelEditComplex(config.CurrentGame.Votes.ID, VotePerm)
	if err != nil {
		//
	}
}

func DayPermPlayer(PlayerID string) {
	Clear(config.CurrentGame.Votes.ID)
	for i := range config.CurrentGame.Channels {
		CurrentPerm := &discordgo.ChannelEdit{
			PermissionOverwrites: []*discordgo.PermissionOverwrite{{
				ID:   config.GuildID,
				Type: "role",
				Deny: Permissions.VIEW_CHANNEL,
			}, {
				ID:    dg.State.User.ID,
				Type:  "member",
				Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
			}},
		}
		for k := range config.CurrentGame.Players {
			if config.CurrentGame.Players[k].ID == PlayerID {
				if config.CurrentGame.Players[k].Role.Name != "Mort" && config.CurrentGame.Players[k].Role.Name != "MDJ" {
					err := MuteSomeone(config.CurrentGame.Players[k].ID, false)
					if err != nil {
						//
					}
				}
				time.Sleep(config.SleepTime * time.Millisecond)
				if config.CurrentGame.Channels[i].Name == config.CurrentGame.Players[k].Role.ChannelName || config.CurrentGame.Channels[i].Name == config.CurrentGame.Players[k].Role.Team.ChannelName {
					CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
						ID:   config.CurrentGame.Players[k].ID,
						Type: "member",
						Deny: Permissions.VIEW_CHANNEL,
					})
				} else if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
					CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
						ID:   config.CurrentGame.Players[k].ID,
						Type: "member",
						Deny: Permissions.VIEW_CHANNEL,
					})
				}
			}
		}
		_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[i].ID, CurrentPerm)
		if err != nil {
			//
		}
		time.Sleep(config.SleepTime * time.Millisecond)
	}
}

func NightPermPlayer(PlayerID string) {
	Clear(config.CurrentGame.Votes.ID)
	for i := range config.CurrentGame.Channels {
		CurrentPerm := &discordgo.ChannelEdit{
			PermissionOverwrites: []*discordgo.PermissionOverwrite{{
				ID:   config.GuildID,
				Type: "role",
				Deny: Permissions.VIEW_CHANNEL,
			}, {
				ID:    dg.State.User.ID,
				Type:  "member",
				Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
			}},
		}
		for k := range config.CurrentGame.Players {
			if config.CurrentGame.Players[k].ID == PlayerID {
				if config.CurrentGame.Players[k].Role.Name != "Mort" && config.CurrentGame.Players[k].Role.Name != "MDJ" {
					err := MuteSomeone(config.CurrentGame.Players[k].ID, true)
					if err != nil {
						//
					}
				}
				time.Sleep(config.SleepTime * time.Millisecond)
				if config.CurrentGame.Players[k].Role.Name == "Chaman" {
					CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
						ID:    config.CurrentGame.Players[k].ID,
						Type:  "member",
						Allow: Permissions.VIEW_CHANNEL,
						Deny:  Permissions.SEND_MESSAGES + Permissions.READ_MESSAGE_HISTORY,
					})
				} else if config.CurrentGame.Channels[i].Name == config.CurrentGame.Players[k].Role.ChannelName || config.CurrentGame.Channels[i].Name == config.CurrentGame.Players[k].Role.Team.ChannelName {
					CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
						ID:    config.CurrentGame.Players[k].ID,
						Type:  "member",
						Allow: Permissions.VIEW_CHANNEL,
					})
				} else if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
					CurrentPerm.PermissionOverwrites = append(CurrentPerm.PermissionOverwrites, &discordgo.PermissionOverwrite{
						ID:    config.CurrentGame.Players[k].ID,
						Type:  "member",
						Allow: Permissions.VIEW_CHANNEL,
					})
				}
			}
		}
		_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[i].ID, CurrentPerm)
		if err != nil {
			//
		}
		time.Sleep(config.SleepTime * time.Millisecond)
	}
}
