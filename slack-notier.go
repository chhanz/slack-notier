package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
    "github.com/ashwanthkumar/slack-go-webhook"
    "time"
)

func main() {
    // time 설정
    t := time.Now()
    timef := t.Format(time.ANSIC)

    db, err := sql.Open("mysql", "root:<pwd>@tcp(localhost:3306)/TESTDB")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
 
    var idx int
    var name string
    var content string 
  
    rows, err := db.Query("SELECT idx,name,content FROM contact where notification in('N')")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close() 

    data := map[int]string{}
    var msgfull string

    for rows.Next() {
        err := rows.Scan(&idx, &name, &content)
        if err != nil {
            log.Fatal(err)
        	}
		data[idx] = name+"-  ```"+content+"```\n"
		msgfull += data[idx]
	}
    
		fmt.Println(msgfull)
		log.Println(idx)    

    if idx != 0 {
    webhookUrl := "https://hooks.slack.com/services/<webhook>"

    payload := slack.Payload {
      Text: timef+"\n신규 문의가 등록 되었습니다.\n"+msgfull ,
      Username: "robot",
      Channel: "#chhanz-webhook",
      IconEmoji: ":monkey_face:",
    }
    slack.Send(webhookUrl, "", payload)
	}

}
