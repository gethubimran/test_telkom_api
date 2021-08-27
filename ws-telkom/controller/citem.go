package controller

import(
	"fmt"
	"net/http"
    "log"
    //"context"
    /*"encoding/base64"
    "encoding/json"
	"github.com/gorilla/mux"*/
    mdl "ws-telkom/model"
    "encoding/json"
    //qry "gocintaid/query"
    //Config "gocintaid/config"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)
	



func GetItem(w http.ResponseWriter, r *http.Request) {
    

    db, err := sql.Open("mysql", "root:moza1210@tcp(127.0.0.1:3306)/opt_telkom")

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
    }

    results, err := db.Query("select id, urgency, due_interval, due_unit, description from m_item;")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    defer db.Close()
    
    ListDtItem := []mdl.Item{}
    for results.Next() {
        var tag mdl.Item
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.Id, &tag.Urgency, &tag.Due_interval, &tag.Due_unit, &tag.Description)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        ListDtItem = append(
            ListDtItem, mdl.Item{
                Id : tag.Id,
                Urgency : tag.Urgency,
				Due_interval : tag.Due_interval,
				Due_unit : tag.Due_unit,
				Description : tag.Description,
            },
        )

    }

    b, err := json.Marshal(ListDtItem)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Fprintf(w, string(b))
    
    
}