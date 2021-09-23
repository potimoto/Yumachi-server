package repository

import (
	"fmt"
	"log"
	"net/http"
	"work/repository/db"
)

func InsertByUserData(appID string, gender string) {
	//insert
	ins, err := db.DB.Prepare("INSERT INTO `appDB`.`pairing_user`(`app_id`,`gender`)VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(appID, gender)
}

func UpdateBySetPass(appID string, strpass string) {
	//upd
	upd, err := db.DB.Prepare("UPDATE `appDB`.`pairing_user` set pairing_id = ? where `app_id`= ?")
	if err != nil {
		log.Fatal(err)
	}
	upd.Exec(strpass, appID)
}

func UpdateByRegisterPassChild(PairingID string, appID string) {
	upd, err := db.DB.Prepare("UPDATE `appDB`.`pairing_user` set pairing_id = ? where `app_id`= ?")
	if err != nil {
		log.Fatal(err)
	}
	upd.Exec(PairingID, appID)
}

var m string

func SelectByCheckMark(appID string, PairingID string, setmark string, w http.ResponseWriter) {
	rows, err := db.DB.Query("select `pairing_id`, ifnull(`mark_str`,'') from `appDB`.`pairing_user` where `app_id` = ?", appID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var mark_str string
		var pairing_id string
		if err := rows.Scan(&pairing_id, &mark_str); err != nil {
			log.Fatal(err)
		}
		log.Println("before")
		log.Println("1回目は虚無", mark_str)
		log.Println(PairingID)
		log.Println(pairing_id)
		if pairing_id == PairingID {
			if mark_str == "" {
				UpdateBySetMark(setmark, PairingID)
				fmt.Fprintf(w, "%v\n", setmark)
				log.Println("afterA")
				log.Println(setmark)
			} else {
				fmt.Fprintf(w, "%v\n", mark_str)
				log.Println("afterB")
				log.Println(mark_str)
			}
		}
		/*if mark_str == "" { //mark_strが空なら更新
			UpdateBySetMark(setmark, PairingID)
		} else if pairing_id == PairingID { //そうでなければセレクトできたものを送信
			fmt.Fprintf(w, "%v\n", mark_str)
			//log.Printf("%v\n", i)
		}*/
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func UpdateBySetMark(setmark string, PairingID string) {
	upd, err := db.DB.Prepare("UPDATE `appDB`.`pairing_user` set mark_str = ? where `pairing_id` = ?")
	if err != nil {
		log.Fatal(err)
	}
	upd.Exec(setmark, PairingID)
}

func UpdateBySetBeaconID(UUID string, appID string) {
	upd, err := db.DB.Prepare("UPDATE `appDB`.`pairing_user` set beacon_id = ? where `app_id`= ?")
	if err != nil {
		log.Fatal(err)
	}
	upd.Exec(UUID, appID)
}

//今は使ってないSelectByCheckMarkに変更
func SelectByMark(PairingID string, appID string, w http.ResponseWriter) {
	rows, err := db.DB.Query("select `pairing_id`, ifnull(`mark_str`,'') from `appDB`.`pairing_user` where `app_id` = ?", appID) //ifnullが正しいかが分からん
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var pairing_id string
		var mark_str string
		if err := rows.Scan(&pairing_id, &mark_str); err != nil {
			log.Fatal(err)
		}
		if pairing_id == PairingID {
			fmt.Fprintf(w, "%v", mark_str)
			_, err := db.DB.Prepare("delete from `appDB`.`pairing_user`") //デモ用
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

//デバッグ
func SelectAll(w http.ResponseWriter) {
	rows, err := db.DB.Query("select * from `appDB`.`pairing_user`") //stringはnullを受け付けない使用のため工夫が必要
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var app_id string
		var pairing_id string
		var gender string
		var beacon_id string
		var mark_str string
		var RasPi_id string
		if err := rows.Scan(&app_id, &pairing_id, &gender, &beacon_id, &mark_str, &RasPi_id); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%v,%v,%v,%v,%v,%v\n", app_id, pairing_id, gender, beacon_id, mark_str, RasPi_id)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func UpdateByGetRasPiIDandSetMark(RasPiID string, UUID string) {
	upd, err := db.DB.Prepare("UPDATE `appDB`.`pairing_user` set RasPi_id = ? where `beacon_id`= ?") //ここで格納するテーブルを変更する？？
	if err != nil {
		log.Fatal(err)
	}
	upd.Exec(RasPiID, UUID)
}

func SelectByChangeMark(RasPiID string, UUID string, w http.ResponseWriter) {
	rows, err := db.DB.Query("select `mark_str` from `appDB`.`pairing_user` where `beacon_id`=?", UUID) //stringはnullを受け付けない使用のため工夫が必要
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() { //rowからデータを持ってこれないと動かないやつ
		var mark_str string
		if err := rows.Scan(&mark_str); err != nil {
			log.Fatal(err)
		}
		if RasPiID == "00000000a4ec6eaa" { //大浴場で受け取ったやつ
			fmt.Fprintf(w, "%v,blue", mark_str)
		} else if RasPiID == "oyaji" { //脱衣所で受け取ったやつ
			fmt.Fprintf(w, "%v,yelloew", mark_str)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func Manji(w http.ResponseWriter) {
	if w != nil {
		_, err := db.DB.Exec("delete from `appDB`.`pairing_user`") //デモ用
		log.Println(&w)
		if err != nil {
			log.Fatal(err)
		}
	}
}
