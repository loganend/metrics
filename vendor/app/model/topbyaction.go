package model

import (
	"time"
	"app/shared/database"
	"log"
	"strconv"
)


type Row struct{
	ID uint32 `json:"id,string,omitepty"`
	Age uint16 `json:"age,string,omitepty"`
	Sex string `json:"sex"`
	Count uint32 `json:"count,string,omitepty"`
}


func GetTopUsersByAction(action, fromDate, toDate, limit string) (map[string][]Row, error){
	var err error

	mapItems := make(map[string][]Row)
	lim, err := strconv.Atoi(limit)

	hell, err := database.SQL.Query("SELECT u.id, u.age, u.sex, count(*) as count, DATE(st.datetime) as date " +
		"FROM STAT st JOIN USER u ON st.user_id = u.id  " +
		"WHERE st.action = ? AND  st.datetime >= ?  and st.datetime <= ?  " +
		"GROUP BY DATE(st.datetime), st.user_id " +
		"ORDER BY DATE(st.datetime), count(*) desc", action, fromDate , toDate)

	if err != nil {
		log.Fatal(err)
	}
	defer hell.Close()


	for hell.Next() {
		var user_id uint32
		var count uint32
		var date time.Time
		var sex string
		var age uint16

		if err := hell.Scan(&user_id, &age, &sex, &count, &date); err != nil {
			log.Fatal(err)
		}

		mapItems[date.Format("2006.01.02")] = append(mapItems[date.Format("2006.01.02")], Row{user_id, age, sex, count})
	}
	if err := hell.Err(); err != nil {
		log.Fatal(err)
	}


	for i, rows:= range mapItems{
		if len(rows) >= lim {
			mapItems[i] = rows[:lim]
		}
	}

	return mapItems, err
}



////////////// Version 2 ////////

//
//
//
//hell, err := database.SQL.Query("Select user_id, count(*) as count,  DATE(datetime) as date, " +
//	"@num := if(@type = user_id, @num + 1, 1) as row_number, " +
//	"@type := user_id as dummy  from stat " +
//	"WHERE action = ? AND  DATE(datetime) >= ?  and DATE(datetime) <= ? " +
//	"GROUP BY DATE(datetime), user_id " +
//	"having row_number <= ? " +
//	"ORDER BY DATE(datetime), count(*) desc", action, fromDate , toDate, limit)
//
//if err != nil {
//	log.Fatal(err)
//}
//defer hell.Close()
//
//
//for hell.Next() {
//	var user_id uint32
//	var count uint32
//	var date time.Time
//	var row_num uint32
//	var dummy uint32
//
//
//	if err := hell.Scan(&user_id, &count, &date, &row_num, &dummy); err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(user_id)
//	fmt.Println(count)
//	fmt.Println(date)
//	fmt.Println("\n\n\n\\n\n")
//
//	mapItems[date.Format("2006.01.02")] = append(mapItems[date.Format("2006.01.02")], Row{user_id, 0, "", count})
//	setIds[user_id] = user_id
//}
//if err := hell.Err(); err != nil {
//	log.Fatal(err)
//}
//
//
//
////for i := 0; i < len(dates); i++ {
////	rows, err := database.SQL.Query("Select user_id, count(*) as count from stat " +
////		"WHERE action = ? AND DATE(datetime) = ? " +
////		"GROUP BY user_id " +
////		"ORDER BY count(*) desc " +
////		"LIMIT ?", action, dates[i], limit)
////
////	if err != nil {
////		log.Fatal(err)
////	}
////	defer rows.Close()
////
////	flag:=true
////	for rows.Next() {
////		if(flag) {
////			udates = append(udates, dates[i])
////			flag = false
////		}
////
////		var user_id uint32
////		var count uint32
////
////		if err := rows.Scan(&user_id, &count); err != nil {
////			log.Fatal(err)
////		}
////
////		mapItems[dates[i]] = append(mapItems[dates[i]], Row{user_id, 0, "", count})
////		setIds[user_id] = user_id
////	}
////	if err := rows.Err(); err != nil {
////		log.Fatal(err)
////	}
////}
//
//if(len(setIds) == 0){
//	return mapItems, err
//}
//
//fmt.Println(setIds)
//
//for k := range setIds {
//		sliceIds = append(sliceIds, k)
//}
//
//fmt.Println("\n\n\n")
//fmt.Println(sliceIds)
//
//args := []interface{}{}
//for _, id := range sliceIds{
//	args = append(args, id)
//}
//
//users, err := database.SQL.Query("SELECT id, age, sex FROM USER " +
//	"WHERE id in (?" + strings.Repeat(",?", len(sliceIds)-1) + ") ", args...)
//
//if err != nil {
//	log.Fatal(err)
//}
//
//defer users.Close()
//
//mapUsers := make(map[uint32]Row)
//
//for users.Next() {
//	var id uint32
//	var age uint16
//	var sex string
//
//	if err := users.Scan(&id, &age, &sex); err != nil {
//		log.Fatal(err)
//	}
//	mapUsers[id] = Row{id, age, sex, 0}
//}
//
//fmt.Println(mapUsers)
//fmt.Println(len(mapItems))
//
//
//for _, rows:= range mapItems{
//	for i := range rows{
//		rows[i].Age = mapUsers[rows[i].ID].Age
//		rows[i].Sex = mapUsers[rows[i].ID].Sex
//	}
//}
//
//
//
////for i := 0; i < len(udates); i++{
////	rows:= mapItems[udates[i]]
////	for j:=0; j <len(rows);j++{
////		rows[j].Age = mapUsers[rows[j].ID].Age
////		rows[j].Sex = mapUsers[rows[j].ID].Sex
////	}
////}
//



////////////// Version 1 ////////


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

