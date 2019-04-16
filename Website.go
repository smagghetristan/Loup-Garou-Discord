package main

import (
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
			fmt.Println(err)
			return
		}
		if rec.Action == "start" {
			GameStart()
		} else if rec.Action == "begin" {
			GameBegin(rec)
		} else if rec.Action == "delete" {
			ChannelDelete()
		} else if rec.Action == "day" {
			DayPerm()
		} else if rec.Action == "night" {
			NightPerm()
		} else if rec.Action == "dayP" {
			DayPermPlayer(rec.ID)
		} else if rec.Action == "nightP" {
			NightPermPlayer(rec.ID)
		} else if rec.Action == "kill" {
			GameBegin(rec)
		} else if rec.Action == "infect" {
			Infect(rec.ID)
		} else if rec.Action == "mute" {
			Mute(rec.ID)
		} else if rec.Action == "muteall" {
			MuteAll()
		} else if rec.Action == "reloadchan" {
			ChannelReload()
		} else if rec.Action == "dice" {
			RollDice(rec.Type, rec.Amount)
		} else if rec.Action == "chlg" {
			ChienLoupChange()
		} else if rec.Action == "mmort" {
			MaitreMort()
		} else if rec.Action == "vwin" {
			VillageWin()
		} else if rec.Action == "wwin" {
			WolvesWin()
		} else if rec.Action == "awin" {
			AngelWin()
		} else if rec.Action == "cwin" {
			CoupleWin()
		} else if rec.Action == "lgbwin" {
			LoupBlancWin()
		}
	}
}
