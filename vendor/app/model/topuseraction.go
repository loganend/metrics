package model

import (
	"time"
	"app/shared/database"
	"log"
	"fmt"
	"strings"
)


type Row struct{
	ID uint32 `json:"id,string,omitepty"`
	Age uint16 `json:"age,string,omitepty"`
	Sex string `json:"sex,string,omitepty"`
	Count uint32 `json:"count,string,omitepty"`
}


func GetTopUsersByAction(action, from, to, num string) (map[string][]Row, error){
	var err error

	//numi, err := strconv.Atoi(num)
	ids, err := database.SQL.Query("SELECT st.user_id FROM STAT st " +
		"WHERE st.action = ? AND st.datetime > ?  and st.datetime < ? " +
		"GROUP BY st.user_id " +
		"ORDER BY count(*) DESC LIMIT ?", action, from, to, num)

	if err != nil {
		log.Fatal(err)
	}

	defer ids.Close()

	var listIds []uint32

	for ids.Next() {
		var user_id uint32


		if err := ids.Scan(&user_id); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d\n", user_id)

		listIds = append(listIds, user_id)
	}
	if err := ids.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(err)
	fmt.Println(len(listIds)-1)

	args := []interface{}{}
	args = append(args, action)
	args = append(args, from)
	args = append(args, to)
	for _, id := range listIds{
		args = append(args, id)
	}

	fmt.Println("args")
	fmt.Println(args)


	rows, err := database.SQL.Query("SELECT u.id, u.age, u.sex, count(*), DATE(st.datetime) FROM STAT st " +
		"JOIN USER u ON st.user_id = u.id WHERE action = ? " +
		"AND  st.datetime > ?  and st.datetime < ? " +
		"AND st.user_id in (?" + strings.Repeat(",?", len(listIds)-1) + ") " +
		"GROUP BY DATE(st.datetime), st.user_id " +
		"ORDER BY DATE(st.datetime)", args...)


	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	mapItems := make(map[string][]Row)

	for rows.Next() {
		var id uint32
		var age uint16
		var sex string
		var count uint32
		var datetime time.Time

		if err := rows.Scan(&id, &age, &sex, &count, &datetime); err != nil {
			log.Fatal(err)
		}

		mapItems[datetime.Format("2006.01.02")] = append(mapItems[datetime.Format("2006.01.02")], Row{id, age, sex, count})

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}


	return mapItems, standardizeError(err)
}
