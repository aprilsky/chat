package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	r "github.com/revel/revel"
	"log"
	"os"
)

var (
	Dbm *gorp.DbMap
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func createTable() *gorp.DbMap {
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "/tmp/post_db.db")
	checkErr(err, "sql.Open failed")

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}
	setColumnUnique := func(t *gorp.TableMap, colUnique map[string]bool) {
		for col, unique := range colUnique {
			t.ColMap(col).Unique = unique
		}
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	t_user := dbmap.AddTableWithName(User{}, "t_user").SetKeys(true, "Id")
	t_chat_log := dbmap.AddTableWithName(ChatLog{}, "t_chat_log").SetKeys(true, "Id")
	setColumnSizes(t_user, map[string]int{
		"Id":       10,
		"NickName": 20,
		"Password": 32,
		"Email":    50,
	})
	setColumnUnique(t_user, map[string]bool{
		"NickName": true,
		"Email":    true,
	})
	setColumnSizes(t_chat_log, map[string]int{
		"Id":           10,
		"FromUserId":   10,
		"ToUserId":     10,
		"ToUserName":   20,
		"FromUserName": 20,
		"MsgContent":   255,
		"ReadTime":     10,
	})

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	//todo check alert table(add column)
	r.INFO.Println("Create tables success ")
	return dbmap
}

func InitDB() {
	Dbm = createTable()
	//defer Dbm.Db.Close()
	Dbm.TraceOn("[gorp]", log.New(os.Stdout, "chat.sql.:", log.Lmicroseconds))
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
