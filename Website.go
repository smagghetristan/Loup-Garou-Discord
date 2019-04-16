package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Loup.Garou/config"
)

// ReceivedPlayer : fsdq
type ReceivedPlayer struct {
	ID       string `json:"id"`
	RoleName string `json:"role"`
}

// ReceivedGame : fsdq
type ReceivedGame struct {
	Players   []ReceivedPlayer `json:"players"`
	RoleNames []string         `json:"roles"`
}

// Received : fsdq
type Received struct {
	Action string       `json:"act"`
	Game   ReceivedGame `json:"roles"`
	ID     string       `json:"id"`
	Type   int          `json:"vari"`
	Amount int          `json:"amount"`
}

type SendGameWebsocket struct {
	Players []config.Player `json:"players"`
}

func SendCards() string {
	Map := map[string]string{}
	for i := range config.AllRoles {
		Map[config.AllRoles[i].Name] = config.AllRoles[i].Image
	}
	StringMap, _ := json.Marshal(Map)
	return string(StringMap)
}

func SendGame() string {
	ToSend := SendGameWebsocket{
		Players: config.CurrentGame.Players,
	}
	StringJSON, _ := json.Marshal(ToSend)
	return string(StringJSON)
}

func handler(w http.ResponseWriter, r *http.Request) {
	Path := r.URL.Path[1:]
	if Path == "" && !config.CurrentGame.Started {
		dat, err := box.FindString("index.html")
		check(err)
		fmt.Fprintf(w, dat)
	} else if Path == "" && config.CurrentGame.Started {
		dat, err := box.FindString("game.html")
		check(err)
		fmt.Fprintf(w, dat)
	} else if Path == "control" && config.CurrentGame.Started {
		dat, err := box.FindString("control.html")
		check(err)
		dat += "<script> var Cards = " + SendCards() + "</script>"
		dat += "<script> var Game = " + SendGame() + "</script>"
		fmt.Fprintf(w, dat)
	} else {
		dat, err := box.FindString("index.html")
		check(err)
		fmt.Fprintf(w, dat)
	}
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	if err != nil {
		return
	}
	conn.SetCloseHandler(func(code int, text string) error {
		for i := range config.Connections {
			if config.Connections[i] == conn {
				config.Connections = append(config.Connections[:i], config.Connections[i+1:]...)
			}
		}
		return nil
	})
	config.Connections = append(config.Connections, conn)
	config.SendPlayerUpdate()

	for {
		// Read message from browser
		rec := Received{}
		err := conn.ReadJSON(&rec)
		if err != nil {
			return
		}
		if rec.Action == "start" {
			go GameStart()
		} else if rec.Action == "begin" {
			go GameBegin(rec)
		} else if rec.Action == "delete" {
			go ChannelDelete()
		} else if rec.Action == "day" {
			go DayPerm()
		} else if rec.Action == "night" {
			go NightPerm()
		} else if rec.Action == "dayP" {
			go DayPermPlayer(rec.ID)
		} else if rec.Action == "nightP" {
			go NightPermPlayer(rec.ID)
		} else if rec.Action == "kill" {
			go Kill(rec.ID)
		} else if rec.Action == "infect" {
			go Infect(rec.ID)
		} else if rec.Action == "mute" {
			go Mute(rec.ID)
		} else if rec.Action == "muteall" {
			go MuteAll()
		} else if rec.Action == "reloadchan" {
			go ChannelReload()
		} else if rec.Action == "dice" {
			go RollDice(rec.Type, rec.Amount)
		} else if rec.Action == "chlg" {
			go ChienLoupChange()
		} else if rec.Action == "mmort" {
			go MaitreMort()
		} else if rec.Action == "vwin" {
			go VillageWin()
		} else if rec.Action == "wwin" {
			go WolvesWin()
		} else if rec.Action == "awin" {
			go AngelWin()
		} else if rec.Action == "cwin" {
			go CoupleWin()
		} else if rec.Action == "lgbwin" {
			go LoupBlancWin()
		}
	}
}
