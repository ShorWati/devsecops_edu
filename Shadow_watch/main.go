package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type AlertLog struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Attacker  string `json:"attacker_ip"`
	Message   string `json:"message"`
}

func LogSecurityAlert(ip string) {
	alert := AlertLog{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     "Critical",
		Attacker:  ip,
		Message:   "[Alert] Обнаружена попытка несанкционированного подключения",
	}
	jsonData, _ := json.Marshal(alert)
	fmt.Println(string(jsonData))
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("Ошибка запуска: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("INFO: Shadow_watch запущен, активно на порту 9999")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		attackerAddr := conn.RemoteAddr().String()
		attackerIP, _, _ := net.SplitHostPort(attackerAddr)

		LogSecurityAlert(attackerIP)

		conn.Close()

	}
}
