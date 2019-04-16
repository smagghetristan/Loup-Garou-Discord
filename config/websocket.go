package config

import (
	"encoding/json"
)

type PlayerUpdate struct {
	Action string   `json:"act"`
	IDs    []Player `json:"players"`
}

type RoleUpdate struct {
	Action string `json:"act"`
	Roles  []Role `json:"roles"`
}

func BroadcastString(msg string) {
	for i := range Connections {
		err := Connections[i].WriteMessage(1, []byte(msg))
		if err != nil {
			Connections = append(Connections[:i], Connections[i+1:]...)
		}
	}
}

func SendPlayerUpdate() {
	b, err := json.Marshal(PlayerUpdate{
		Action: "players",
		IDs:    Players,
	})
	if err != nil {
		return
	}
	c, err := json.Marshal(RoleUpdate{
		Action: "roles",
		Roles:  AllRoles,
	})
	if err != nil {
		return
	}
	BroadcastString(string(b))
	BroadcastString(string(c))
}
