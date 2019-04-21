package main

import (
	"time"

	Permissions "Loup.Garou/perm"

	"Loup.Garou/config"

	"github.com/bwmarrin/discordgo"
)

func ChamDayPerm(userID string) *discordgo.ChannelEdit {
	Perm := &discordgo.ChannelEdit{
		PermissionOverwrites: []*discordgo.PermissionOverwrite{{
			ID:   config.GuildID,
			Type: "role",
			Deny: Permissions.VIEW_CHANNEL,
		}, {
			ID:    dg.State.User.ID,
			Type:  "member",
			Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
		}, {
			ID:   userID,
			Type: "member",
			Deny: Permissions.VIEW_CHANNEL,
		}},
	}
	return Perm
}

func NormalDayPerm(userID string) *discordgo.ChannelEdit {
	Perm := &discordgo.ChannelEdit{
		PermissionOverwrites: []*discordgo.PermissionOverwrite{{
			ID:   config.GuildID,
			Type: "role",
			Deny: Permissions.VIEW_CHANNEL,
		}, {
			ID:    dg.State.User.ID,
			Type:  "member",
			Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
		}, {
			ID:   userID,
			Type: "member",
			Deny: Permissions.VIEW_CHANNEL,
		}},
	}
	return Perm
}

func ChamNightPerm(userID string) *discordgo.ChannelEdit {
	Perm := &discordgo.ChannelEdit{
		PermissionOverwrites: []*discordgo.PermissionOverwrite{{
			ID:   config.GuildID,
			Type: "role",
			Deny: Permissions.VIEW_CHANNEL,
		}, {
			ID:    dg.State.User.ID,
			Type:  "member",
			Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
		}, {
			ID:    userID,
			Type:  "member",
			Allow: Permissions.VIEW_CHANNEL,
			Deny:  Permissions.SEND_MESSAGES + Permissions.READ_MESSAGE_HISTORY,
		}},
	}
	return Perm
}

func NormalNightPerm(userID string) *discordgo.ChannelEdit {
	Perm := &discordgo.ChannelEdit{
		PermissionOverwrites: []*discordgo.PermissionOverwrite{{
			ID:   config.GuildID,
			Type: "role",
			Deny: Permissions.VIEW_CHANNEL,
		}, {
			ID:    dg.State.User.ID,
			Type:  "member",
			Allow: Permissions.VIEW_CHANNEL + Permissions.SEND_MESSAGES,
		}, {
			ID:    userID,
			Type:  "member",
			Allow: Permissions.VIEW_CHANNEL,
		}},
	}
	return Perm
}

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
			_, err := dg.ChannelEditComplex(config.CurrentGame.Votes.ID, NormalDayPerm(config.CurrentGame.Players[i].ID))
			if err != nil {
				//
			}
			time.Sleep(10 * time.Millisecond)

			if config.CurrentGame.Players[i].Role.Name == "Chaman" {
				Clear(config.CurrentGame.Deads.ID)
				_, err := dg.ChannelEditComplex(config.CurrentGame.Deads.ID, ChamDayPerm(config.CurrentGame.Players[i].ID))
				if err != nil {
					//
				}
				time.Sleep(10 * time.Millisecond)
			}

			if config.CurrentGame.Players[i].Role.Name != "Mort" {
				err := MuteSomeone(config.CurrentGame.Players[i].ID, false)
				if err != nil {
					//
				}
				time.Sleep(10 * time.Millisecond)
			}

			for k := range config.CurrentGame.Channels {

				if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalDayPerm(config.CurrentGame.Players[i].ID))
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}

				if config.CurrentGame.Players[i].Role.ChannelName == config.CurrentGame.Channels[k].Name {
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalDayPerm(config.CurrentGame.Players[i].ID))
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}
				if config.CurrentGame.Players[i].Role.Team.HasChannel && config.CurrentGame.Players[i].Role.Team.ChannelName == config.CurrentGame.Channels[k].Name {
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalDayPerm(config.CurrentGame.Players[i].ID))
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
			_, err := dg.ChannelEditComplex(config.CurrentGame.Votes.ID, NormalNightPerm(config.CurrentGame.Players[i].ID))
			if err != nil {
				//
			}
			time.Sleep(10 * time.Millisecond)

			if config.CurrentGame.Players[i].Role.Name == "Chaman" {
				Clear(config.CurrentGame.Deads.ID)
				_, err := dg.ChannelEditComplex(config.CurrentGame.Deads.ID, ChamNightPerm(config.CurrentGame.Players[i].ID))
				if err != nil {
					//
				}
				time.Sleep(10 * time.Millisecond)
			}

			if config.CurrentGame.Players[i].Role.Name != "Mort" {
				err := MuteSomeone(config.CurrentGame.Players[i].ID, true)
				if err != nil {
					//
				}
				time.Sleep(10 * time.Millisecond)
			}
			time.Sleep(10 * time.Millisecond)

			for k := range config.CurrentGame.Channels {

				if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalNightPerm(config.CurrentGame.Players[i].ID))
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}
				if config.CurrentGame.Players[i].Role.ChannelName == config.CurrentGame.Channels[k].Name {
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalNightPerm(config.CurrentGame.Players[i].ID))
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}
				if config.CurrentGame.Players[i].Role.Team.HasChannel && config.CurrentGame.Players[i].Role.Team.ChannelName == config.CurrentGame.Channels[k].Name {
					_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalNightPerm(config.CurrentGame.Players[i].ID))
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
				_, err := dg.ChannelEditComplex(config.CurrentGame.Votes.ID, NormalDayPerm(config.CurrentGame.Players[i].ID))
				if err != nil {
					//
				}
				time.Sleep(10 * time.Millisecond)

				if config.CurrentGame.Players[i].Role.Name == "Chaman" {
					Clear(config.CurrentGame.Deads.ID)
					_, err := dg.ChannelEditComplex(config.CurrentGame.Deads.ID, ChamDayPerm(config.CurrentGame.Players[i].ID))
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}

				if config.CurrentGame.Players[i].Role.Name != "Mort" {
					err := MuteSomeone(config.CurrentGame.Players[i].ID, false)
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}

				for k := range config.CurrentGame.Channels {

					if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalDayPerm(config.CurrentGame.Players[i].ID))
						if err != nil {
							//
						}
						time.Sleep(10 * time.Millisecond)
					}

					if config.CurrentGame.Players[i].Role.ChannelName == config.CurrentGame.Channels[k].Name {
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalDayPerm(config.CurrentGame.Players[i].ID))
						if err != nil {
							//
						}
						time.Sleep(10 * time.Millisecond)
					}
					if config.CurrentGame.Players[i].Role.Team.HasChannel && config.CurrentGame.Players[i].Role.Team.ChannelName == config.CurrentGame.Channels[k].Name {
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalDayPerm(config.CurrentGame.Players[i].ID))
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
				_, err := dg.ChannelEditComplex(config.CurrentGame.Votes.ID, NormalNightPerm(config.CurrentGame.Players[i].ID))
				if err != nil {
					//
				}
				time.Sleep(10 * time.Millisecond)

				if config.CurrentGame.Players[i].Role.Name == "Chaman" {
					Clear(config.CurrentGame.Deads.ID)
					_, err := dg.ChannelEditComplex(config.CurrentGame.Deads.ID, ChamNightPerm(config.CurrentGame.Players[i].ID))
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}

				if config.CurrentGame.Players[i].Role.Name != "Mort" {
					err := MuteSomeone(config.CurrentGame.Players[i].ID, true)
					if err != nil {
						//
					}
					time.Sleep(10 * time.Millisecond)
				}

				for k := range config.CurrentGame.Channels {

					if config.CurrentGame.Players[i].Infected && config.CurrentGame.Channels[k].Name == "loup-garou" {
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalNightPerm(config.CurrentGame.Players[i].ID))
						if err != nil {
							//
						}
						time.Sleep(10 * time.Millisecond)
					}

					if config.CurrentGame.Players[i].Role.ChannelName == config.CurrentGame.Channels[k].Name {
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalNightPerm(config.CurrentGame.Players[i].ID))
						if err != nil {
							//
						}
						time.Sleep(10 * time.Millisecond)
					}
					if config.CurrentGame.Players[i].Role.Team.HasChannel && config.CurrentGame.Players[i].Role.Team.ChannelName == config.CurrentGame.Channels[k].Name {
						_, err := dg.ChannelEditComplex(config.CurrentGame.Channels[k].ID, NormalNightPerm(config.CurrentGame.Players[i].ID))
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
