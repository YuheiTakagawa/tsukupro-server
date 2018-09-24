//package main
package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v1"
)

type UserInfo struct {
	Userid int32  `db:"userid"`
	Name   string `db:"name"`
	Birth  string `db:"birth"`
}

/*
type Transaction struct {
	ID        int    `db:"id"`
	TxHash    string `db:"TxHash"`
	BlockID   int    `db:"BlockID"`
	Input     string `db:"Input"`
	Output    string `db:"Output"`
	Amount    int    `db:"Amount"`
	Timestamp string `db:"Timestamp"`
	Sign      string `db:"Sign"`
	Pubkey    string `db:"Pubkey"`
}
*/
/*
func main() {
	dbmap := InitDb()
	defer dbmap.Db.Close()

	//insert用のテストデータ
	tx1 := &UserInfo{1, "Yuhei", "2018"}
	err := dbmap.Insert(tx1)
	CheckErr(err, "Insert failed")

	var transactions []UserInfo
	_, err = dbmap.Select(&transactions, "select * from userinfo order by userid")
	CheckErr(err, "Select failed")
	log.Println("All rows:")
	for x, p := range transactions {
		log.Printf("    %d: %v\n", x, p)
	}
}
*/
func InitDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/db")
	CheckErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(UserInfo{}, "userinfo").SetKeys(false, "userid")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")

	return dbmap
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
