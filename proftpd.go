package main

import (
	"encoding/json"
	"log"
	"os/exec"
)

type ftpwho struct {
	Server      ftpwhoServer       `json:"server"`
	Connections []ftpwhoConnection `json:"connections"`
}

type ftpwhoServer struct {
	PID        int    `json:"pid"`
	ServerType string `json:"server_type"`
	StartedMs  int    `json:"started_ms"`
}

type ftpwhoConnection struct {
	ConnectedSinceMs int    `json:"connected_since_ms"`
	Idling           bool   `json:"idling"`
	IdleSinceMs      int    `json:"idle_since_ms"`
	LocalAddress     string `json:"local_address"`
	LocalPort        int    `json:"local_port"`
	Location         string `json:"location"`
	PID              int    `json:"pid"`
	Protocol         string `json:"protocol"`
	RemoteName       string `json:"remote_name"`
	RemoteAddress    string `json:"remote_address"`
	Username         string `json:"user"`
}

type ftpwhoUniqueConnection struct {
	Count      int
	Connection ftpwhoConnection
}

func parseFtpwho() ftpwho {
	cmd := exec.Command("/usr/bin/ftpwho", "-v", "-o", "json")
	ftpwhoOutput, err := cmd.Output()
	if err != nil {
		log.Printf("[ERROR] Unable to call `ftpwho`: %s", err)
	}
	// parse ftpwhoOutput
	var ftpwhoParsed ftpwho
	err = json.Unmarshal(ftpwhoOutput, &ftpwhoParsed)
	if err != nil {
		log.Printf("[ERROR] Unable to Parse `ftpwho` output: %s", err)
	}

	return ftpwhoParsed
}

func (fw *ftpwhoConnection) DetermineState() string {
	if fw.Idling {
		return "idle"
	}
	return "active"
}
