// main.go
package main

import (
	// "time"
    "database/sql"
    "fmt"
	"log"
    _ "github.com/denisenkom/go-mssqldb"
)

type AccessRegion struct {
    ALM_TAGNAME string
    ALM_NATIVETIMEIN string
    ALM_NATIVETIMELAST string
    SL_CLOSETIME string
}

func main() {
    var server = "10.1.1.32"
    var port = 1433
    var user = "sa"
    var password = "Sealu2019"
    var database = "iFix_Data"

    //连接字符串
    connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", server, port, database, user, password)

    //建立连接
    db, err := sql.Open("mssql", connString)
    if err != nil {
        log.Fatal("Open Connection failed:", err.Error())
    }
    defer db.Close()
    _, err = db.Exec("exec [dbo].[inserttimein]")
    if err != nil {
        fmt.Println(err)
    }

    // 通过连接对象执行查询
    rows, err := db.Query(`exec dbo.valueselect`)
    if err != nil {
        log.Fatal("Query failed:", err.Error())
    }
    defer rows.Close()
    var rowsData []*AccessRegion
    //遍历每一行
    for rows.Next() {
        var row = new(AccessRegion)
        rows.Scan(&row.ALM_TAGNAME,&row.ALM_NATIVETIMEIN,&row.ALM_NATIVETIMELAST,&row.SL_CLOSETIME)
        rowsData = append(rowsData, row)
    }

    //打印数组
    for _, ar := range rowsData {
        if len(ar.SL_CLOSETIME)==0 {
            _, err = db.Exec("exec [dbo].[inserttimeout] @in=?",ar.ALM_NATIVETIMEIN)
        }
    }
}