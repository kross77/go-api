package mysql

import (
	"os"
	"log"
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
)

var DB *sql.DB
var IsConnected bool

func ConnectToMySQL() {
	database:="golang"
	user:=os.Getenv("adminzFhGeYC")
	password:=os.Getenv("qzeFQCbJ7mT_")
	var err error
	DB, err = sql.Open("golang", "tcp:"+os.Getenv("OPENSHIFT_MYSQL_DB_HOST")+":"+os.Getenv("OPENSHIFT_MYSQL_DB_PORT")+"*"+database+"/"+user+"/"+password)
	if err != nil {
		log.Fatal(err)
	}else{
		IsConnected = true;
	}
}

