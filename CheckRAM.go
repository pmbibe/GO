package main

import (
	"database/sql"

	_ "fmt"

	"fmt"

	"os/exec"
	_ "os/exec"

	_ "time"

	_ "github.com/go-sql-driver/mysql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var updates tgbotapi.UpdatesChannel

var bot *tgbotapi.BotAPI

var db = connect()

var last_status uint

type Tag struct {
	HOSTNAME string

	SERVICE string

	STATUSCODE uint
}

func main() {

	last_status = 200

	// updates, bot = startBOT()

	// defer db.Close()

	// go func() {

	// checkServiceAllServer()

	// }()

	// for {

	// 	check_server_JIRA()

	// }

	// for {

	// checkRAMServer()
	checkRAMServer()

	// }

}

func getValueRAM(server string) (string, string, string) {
	command := "./getTotalRAM.sh " + server
	result := exec.Command("sh", "-c", command)
	totalRAM, _ := result.Output()
	command = "./getFreeRAM.sh " + server
	result = exec.Command("sh", "-c", command)
	FreeRAM, _ := result.Output()
	command = "./getAvailableRAM.sh " + server
	result = exec.Command("sh", "-c", command)
	AvailableRAM, _ := result.Output()
	return string(totalRAM), string(FreeRAM), string(AvailableRAM)
}
func checkRAMServer() {
	slect, _ := db.Query("SELECT Hostname FROM monitor")

	defer slect.Close()

	for slect.Next() {

		var tag Tag

		_ = slect.Scan(&tag.HOSTNAME)

		fmt.Println(getValueRAM(tag.HOSTNAME))

	}
}

func startBOT() (tgbotapi.UpdatesChannel, *tgbotapi.BotAPI) {

	bot, _ := tgbotapi.NewBotAPI("1153037633:AAHp5oGyFvTncdN_9hkhoNyEQpuM4cwYnns")

	u := tgbotapi.NewUpdate(0)

	updates, _ := bot.GetUpdatesChan(u)

	return updates, bot

}

func sendMsg(Msg string) {

	msg := tgbotapi.NewMessage(817269876, "")

	msg.Text = Msg

	bot.Send(msg)

}

func connecttoDB(user, password, hostname, port, database string) string {

	return user + ":" + password + "@" + "tcp(" + hostname + ":" + port + ")/" + database

}

func connect() *sql.DB {

	db, _ := sql.Open("mysql", connecttoDB("root", "minhduc7b", "192.168.141.204", "3306", "monitor_byGo"))

	return db

}
