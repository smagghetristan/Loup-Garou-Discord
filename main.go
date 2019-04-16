package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"Loup.Garou/config"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"github.com/gobuffalo/packr"
	"github.com/gorilla/websocket"
)

var dg *discordgo.Session
var box packr.Box

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func vocalUpdate(s *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	// Find the guild for that channel.
	if event.ChannelID == config.VoiceChannel || event.ChannelID == "" {
		scan(s)
	}
}
func scan(s *discordgo.Session) {
	if !config.CurrentGame.Started {
		config.Players = []config.Player{}
		g, err := s.Guild(config.GuildID)
		if err != nil {
			return
		}
		// Look for the message sender in that guild's current voice states.
		for _, vs := range g.VoiceStates {
			if vs.ChannelID == config.VoiceChannel {
				user, err := s.User(vs.UserID)
				if err != nil {
					return
				}
				config.Players = append(config.Players, config.Player{
					ID:        vs.UserID,
					Username:  user.Username,
					AvatarURL: user.AvatarURL("512"),
				})
			}
		}
		config.SendPlayerUpdate()
	}
}

func botReady(s *discordgo.Session, event *discordgo.Ready) {
	color.Green("Bot is ready !")
	scan(s)
}

func MuteSomeone(userID string, State bool) (err error) {
	data := struct {
		Muted bool `json:"mute"`
	}{State}

	_, err = dg.RequestWithBucketID("PATCH", discordgo.EndpointGuildMember(config.GuildID, userID), data, discordgo.EndpointGuildMember(config.GuildID, ""))
	if err != nil {
		return
	}

	return
}

func main() {
	box = packr.NewBox("./www")
	var err error
	// Create a new Discord session using the provided bot token.
	dg, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	//Setups
	config.SetupRole()
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(vocalUpdate)
	dg.AddHandler(botReady)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	ChannelDelete()
	http.HandleFunc("/ws", ws)
	http.HandleFunc("/", handler)
	http.Handle("/b/", http.StripPrefix("/b/", http.FileServer(box)))
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
