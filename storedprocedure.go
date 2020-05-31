// main.go
package main

import (
    "database/sql"
    "fmt"
	"log"

    _ "github.com/denisenkom/go-mssqldb"
)

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
    _, err = db.Exec("CREATE PROCEDURE [dbo].[valueselect] AS BEGIN select ALM_TAGNAME,ALM_NATIVETIMEIN,ALM_NATIVETIMELAST,SL_CLOSETIME from [iFix_Data].[dbo].[FIXALARMS] a where not exists(select 1 from [iFix_Data].[dbo].[FIXALARMS] where ALM_NATIVETIMEIN=a.ALM_NATIVETIMEIN and ALM_TAGNAME=a.ALM_TAGNAME and ALM_NATIVETIMELAST>a.ALM_NATIVETIMELAST) and ALM_ALMSTATUS='OK' and SL_TAG_TIMEIN is NULL END")
    _, err = db.Exec("CREATE PROCEDURE [dbo].[inserttimein] AS BEGIN update [iFix_Data].[dbo].[FIXALARMS] set SL_TAG_TIMEIN=Replace((concat(concat(ALM_TAGNAME,'_'),concat(ALM_DATEIN,'_'),ALM_TIMEIN)),' ','') where SL_TAG_TIMEIN is NULL END")
    _, err = db.Exec("CREATE PROCEDURE [dbo].[inserttimeout] @in varchar(64) AS BEGIN update [iFix_Data].[dbo].[FIXALARMS] set SL_CLOSETIME=ALM_NATIVETIMELAST where ALM_ALMSTATUS='OK' and ALM_NATIVETIMEIN=@in END")
}