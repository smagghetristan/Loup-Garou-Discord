package main

import (
	"Test/config"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func VillageWin() {
	dat, err := box.FindString("static/vil.png")
	check(err)
	Image := &discordgo.File{
		Reader: strings.NewReader(dat),
		Name:   "role.png",
	}
	Params := &discordgo.MessageSend{
		Files: []*discordgo.File{Image},
		Embed: &discordgo.MessageEmbed{
			Title:       "Les Villageois ont gagnés la partie !",
			Description: "Félicitations les Villageois !",
			Image: &discordgo.MessageEmbedImage{
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
	config.CurrentGame = config.Game{}
}

func WolvesWin() {
	dat, err := box.FindString("static/lg.png")
	check(err)
	Image := &discordgo.File{
		Reader: strings.NewReader(dat),
		Name:   "role.png",
	}
	Params := &discordgo.MessageSend{
		Files: []*discordgo.File{Image},
		Embed: &discordgo.MessageEmbed{
			Title:       "Les Loups-Garous ont gagnés la partie !",
			Description: "Félicitations les Loups-Garous !",
			Image: &discordgo.MessageEmbedImage{
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
	config.CurrentGame = config.Game{}
}

func AngelWin() {
	dat, err := box.FindString("static/ang.png")
	check(err)
	Image := &discordgo.File{
		Reader: strings.NewReader(dat),
		Name:   "role.png",
	}
	Params := &discordgo.MessageSend{
		Files: []*discordgo.File{Image},
		Embed: &discordgo.MessageEmbed{
			Title:       "L'ange a gagné la partie !",
			Description: "Félicitations !",
			Image: &discordgo.MessageEmbedImage{
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
	config.CurrentGame = config.Game{}
}

func CoupleWin() {
	dat, err := box.FindString("static/cup.png")
	check(err)
	Image := &discordgo.File{
		Reader: strings.NewReader(dat),
		Name:   "role.png",
	}
	Params := &discordgo.MessageSend{
		Files: []*discordgo.File{Image},
		Embed: &discordgo.MessageEmbed{
			Title:       "Le couple a gagné la partie !",
			Description: "Félicitations le couple !",
			Image: &discordgo.MessageEmbedImage{
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
	config.CurrentGame = config.Game{}
}

func LoupBlancWin() {
	dat, err := box.FindString("static/lgb.png")
	check(err)
	Image := &discordgo.File{
		Reader: strings.NewReader(dat),
		Name:   "role.png",
	}
	Params := &discordgo.MessageSend{
		Files: []*discordgo.File{Image},
		Embed: &discordgo.MessageEmbed{
			Title:       "Le Loup-Garou Blanc a gagné la partie !",
			Description: "Félicitations !",
			Image: &discordgo.MessageEmbedImage{
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
	config.CurrentGame = config.Game{}
}
