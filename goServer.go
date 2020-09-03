package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type indexHandler struct {
	content string
}

var db *sql.DB

func init() {
	DB, err := sql.Open("mysql", "数据库用户名:密码@tcp(数据库ip:端口)/数据库名")
	if err != nil {
		fmt.Println("open mysql failed", err)
	}
	db = DB
}
func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("insert into my_test (temp) values(123)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("插入成功")
	//defer row.Close()
}

func main() {
	http.Handle("/", &indexHandler{content: "hello world"})
	http.ListenAndServe(":8001", nil)
}
