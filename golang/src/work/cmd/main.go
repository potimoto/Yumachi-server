package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"work/repository"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

var mark = []string{"magro", "ebi", "tamago"}

//ユーザデータ格納
func UserData(w http.ResponseWriter, r *http.Request) {
	appID := r.FormValue("appID")
	gender := r.FormValue("gender")
	repository.InsertByUserData(appID, gender)
}

//ペアリング用パスワードを生成し、親となるスマホに格納
func SetPass(w http.ResponseWriter, r *http.Request) {
	var pass []int
	var strpass string
	for i := 0; i < 4; i++ {
		pass = append(pass, rand.Intn(10))
		strpass = strpass + strconv.Itoa(pass[i])
	}
	fmt.Fprintf(w, "%v\n", strpass) //4桁のパスワードをリクエストに対してレスポンス
	appID := r.FormValue("appID")
	repository.UpdateBySetPass(appID, strpass)
}

//子となるスマホのパスを登録
func RegisterPassChild(w http.ResponseWriter, r *http.Request) {
	appID := r.FormValue("appID")
	PairingID := r.FormValue("PairingID")
	repository.UpdateByRegisterPassChild(PairingID, appID)
}

//beaconのUUIDを格納
func SetUUIDandGetMark(w http.ResponseWriter, r *http.Request) {
	//Post受け取り
	appID := r.FormValue("appID")
	UUID := r.FormValue("UUID")
	PairingID := r.FormValue("PairingID")
	//認証番号につきマークを設定
	setmark := mark[rand.Intn(3)]

	repository.UpdateBySetBeaconID(UUID, appID)
	repository.SelectByCheckMark(appID, PairingID, setmark, w)
	//repository.SelectByMark(PairingID, appID, w)
}

//片方のbeaconidを受け取る処理を書く

//ラズパイから固有IDを受け取りとマーク送信
func ChangeMark(w http.ResponseWriter, r *http.Request) {
	//Post受け取り
	RasPiID := r.FormValue("RasPiID") //ラズパイ固有ID
	UUID := r.FormValue("UUID")       //ラズパイで受信したbeaconのUUID

	repository.UpdateByGetRasPiIDandSetMark(RasPiID, UUID)
	//ラズパイごとにIDで検知しようと思ったけど処理としては各々別のリンクに飛ばして処理を書いた方が楽かもしれない
	repository.SelectByChangeMark(RasPiID, UUID, w)
}

//できてない
func Debug(w http.ResponseWriter, r *http.Request) {
	repository.Reset(w)
}

func main() {
	// gorilla mux でルーティングする
	router := mux.NewRouter()
	//router.HandleFunc("/", getDBInfo).Methods("GET")
	router.HandleFunc("/UserData", UserData)                   //appIDとgenderを格納するときのあれこれ
	router.HandleFunc("/SetPass", SetPass)                     //パスワード生成と親スマホへの格納
	router.HandleFunc("/RegisterPass", RegisterPassChild)      //子となるスマホに登録するときのあれこれ
	router.HandleFunc("/SetUUIDandGetMark", SetUUIDandGetMark) //uuid登録とmarkに当たる文字列送信を同時に行うあれこれ
	router.HandleFunc("/ChangeMark", ChangeMark)               //ラズパイ間とのあれこれ
	router.HandleFunc("/Debug", Debug)
	//defer db.DB.Close()
	fmt.Println("Server Start...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
