// 初期化処理
package db

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var connectionError error

// DB の接続情報
const (
	DRIVER_NAME = "mysql" // ドライバ名(mysql固定)
	// user:password@tcp(container-name:port)/dbname ※mysql はデフォルトで用意されているDB
	DATA_SOURCE_NAME = "root:password@tcp(mysql:3306)/mysql"
)

func init() {
	DB, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database: ", connectionError)
	}
	rand.Seed(time.Now().UnixNano()) //Seedで生成する乱数を固定しないように
}
