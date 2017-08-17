package model

import (
	"time"
	"app/shared/database"
	"log"
	"strings"
)


type Row struct{
	ID uint32 `json:"id,string,omitepty"`
	Age uint16 `json:"age,string,omitepty"`
	Sex string `json:"sex"`
	Count uint32 `json:"count,string,omitepty"`
}


func GetTopUsersByAction(action, fromDate, toDate, limit string) (map[string][]Row, error){
	var err error

	from, err := time.Parse("2006-01-02", fromDate)
	to, err := time.Parse("2006-01-02", toDate)
	mapItems := make(map[string][]Row)
	setIds := make(map[uint32]uint32)
	var sliceIds []uint32
	var dates []string
	var udates []string


	for !from.Equal(to) {
		dates = append(dates, from.Format("2006.01.02"))
		from = from.AddDate(0, 0, 1)
	}
	dates = append(dates, to.Format("2006.01.02"))


	for i := 0; i < len(dates); i++ {
		rows, err := database.SQL.Query("Select user_id, count(*) as count from stat " +
			"WHERE action = ? AND DATE(datetime) = ? " +
			"GROUP BY user_id " +
			"ORDER BY count(*) desc " +
			"LIMIT ?", action, dates[i], limit)

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		flag:=true
		for rows.Next() {
			if(flag) {
				udates = append(udates, dates[i])
				flag = false
			}

			var user_id uint32
			var count uint32

			if err := rows.Scan(&user_id, &count); err != nil {
				log.Fatal(err)
			}

			mapItems[dates[i]] = append(mapItems[dates[i]], Row{user_id, 0, "", count})
			setIds[user_id] = user_id
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}

	if(len(setIds) == 0){
		return mapItems, err
	}

	for k := range setIds {
			sliceIds = append(sliceIds, k)
	}

	args := []interface{}{}
	for _, id := range sliceIds{
		args = append(args, id)
	}

	users, err := database.SQL.Query("SELECT id, age, sex FROM USER " +
		"WHERE id in (?" + strings.Repeat(",?", len(sliceIds)-1) + ") ", args...)

	if err != nil {
		log.Fatal(err)
	}

	defer users.Close()

	mapUsers := make(map[uint32]Row)

	for users.Next() {
		var id uint32
		var age uint16
		var sex string


		if err := users.Scan(&id, &age, &sex); err != nil {
			log.Fatal(err)
		}

		mapUsers[id] = Row{id, age, sex, 0}
	}

	for i := 0; i < len(udates); i++{
		rows:= mapItems[udates[i]]
		for j:=0; j <len(rows);j++{
			rows[j].Age = mapUsers[rows[j].ID].Age
			rows[j].Sex = mapUsers[rows[j].ID].Sex
		}
	}
	return mapItems, err
}

//if err := users.Err(); err != nil {

//ids, err := database.SQL.Query("SELECT st.user_id FROM STAT st " +
//	"WHERE st.action = ? AND st.datetime > ?  and st.datetime < ? " +
//	"GROUP BY st.user_id " +
//	"ORDER BY count(*) DESC LIMIT ?", action, from, to, num)
//
//if err != nil {
//	log.Fatal(err)
//}
//
//defer ids.Close()
//
//var listIds []uint32
//
//for ids.Next() {
//	var user_id uint32
//
//
//	if err := ids.Scan(&user_id); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("id: %d\n", user_id)
//
//	listIds = append(listIds, user_id)
//}
//if err := ids.Err(); err != nil {
//	log.Fatal(err)
//}
//
//fmt.Println(err)
//fmt.Println(len(listIds)-1)
//
//args := []interface{}{}
//args = append(args, action)
//args = append(args, from)
//args = append(args, to)
//for _, id := range listIds{
//	args = append(args, id)
//}
//
//fmt.Println("args")
//fmt.Println(args)
//
//
//rows, err := database.SQL.Query("SELECT u.id, u.age, u.sex, count(*), DATE(st.datetime) FROM STAT st " +
//	"JOIN USER u ON st.user_id = u.id WHERE action = ? " +
//	"AND  st.datetime > ?  and st.datetime < ? " +
//	"AND st.user_id in (?" + strings.Repeat(",?", len(listIds)-1) + ") " +
//	"GROUP BY DATE(st.datetime), st.user_id " +
//	"ORDER BY DATE(st.datetime)", args...)
//
//
//if err != nil {
//	log.Fatal(err)
//}
//defer rows.Close()
//
//
//mapItems := make(map[string][]Row)
//
//for rows.Next() {
//	var id uint32
//	var age uint16
//	var sex string
//	var count uint32
//	var datetime time.Time
//
//	if err := rows.Scan(&id, &age, &sex, &count, &datetime); err != nil {
//		log.Fatal(err)
//	}
//
//	mapItems[datetime.Format("2006.01.02")] = append(mapItems[datetime.Format("2006.01.02")], Row{id, age, sex, count})
//
//}
//if err := rows.Err(); err != nil {
//	log.Fatal(err)
//}

