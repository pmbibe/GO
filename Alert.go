package main
import(
	"log"
	"database/sql"
	"os/exec"
	_ "github.com/go-sql-driver/mysql"
)

func checkServiceRunning(service string, server string, lastStatusService uint) {
	serviceName := "./exitCode.sh " + service + " " + server + " ;echo $?"
	StatusCode := exec.Command("sh", "-c", serviceName)
	statusCode, _ := StatusCode.Output()
	sttCode := string(statusCode)
	if sttCode == "0\n" && (lastStatusService != 0) {
		log.Print("Service is running")
		log.Print(changeStatustoOK(service,server))
	} else if sttCode != "0\n" && (lastStatusService != 1) {
		log.Print("Service is Dead")
		log.Print(changeStatustoFail(service,server))
	}

}

func main(){
	db, _ := sql.Open("mysql", "root:minhduc7b@tcp(192.168.141.204:3306)/monitor_byGo")

defer db.Close()

slect, _ := db.Query("SELECT * FROM monitor")

defer slect.Close()
type Tag struct {
	HOSTNAME string
	SERVICE string
	STATUSCODE uint
}
for slect.Next(){
	var tag Tag
	_ = slect.Scan(&tag.HOSTNAME,&tag.SERVICE,&tag.STATUSCODE)
	checkServiceRunning(tag.SERVICE, tag.HOSTNAME, tag.STATUSCODE)
	
}
}
func changeStatustoOK(service string, server string) string {
	return "UPDATE monitor SET statusCode = 0 where Hostname = '" + server + "' AND Service = '" + service +"'"
	
}
func changeStatustoFail(service string, server string) string {
	return "UPDATE monitor SET statusCode = 1 where Hostname = '" + server + "' AND Service = '" + service +"'"
	
}
