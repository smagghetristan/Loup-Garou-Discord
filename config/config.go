package config

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/websocket"
)

var ImagePath = "static/"

//Different IDs
var VoiceChannel = "398084144862724098"
var GuildID = "248880737518878720"
var CategoryID = "398083936745553920"
var Token = ""
var CurrentGame Game

//Slices

var AllTeams []Team
var AllRoles []Role
var Players []Player

var Channels []string
var TeamChannels []string
var SpecialChannels []string
var Connections []*websocket.Conn

type Team struct {
	Name        string
	ChannelName string
	HasChannel  bool
}

type Role struct {
	Name        string `json:"name"`
	Image       string `json:"card"`
	ChannelName string `json:"-"`
	Description string `json:"-"`
	Team        Team   `json:"-"`
}

type RoleChannel struct {
	Role Role
	ID   string
}

type TeamChannel struct {
	Team Team
	ID   string
}

type Player struct {
	Username   string `json:"username"`
	ID         string `json:"id"`
	AvatarURL  string `json:"picture"`
	Infected   bool   `json:"-"`
	ManualMute bool   `json:"-"`
	Role       Role   `json:"role"`
}

type Game struct {
	Started   bool
	Finished  bool
	SendComp  bool
	Players   []Player
	Roles     []Role
	GuildID   string
	Channels  []*discordgo.Channel
	GameStats *discordgo.Channel
	Votes     *discordgo.Channel
	Deads     *discordgo.Channel
}
