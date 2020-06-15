package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var mapdb = make(map[string]string)

func connectToDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "test02:test1234@tcp(192.168.193.102:3306)/test")
	if db != nil {
		db.SetMaxOpenConns(100) // 최대 커넥션 개수를 제한</span>
		db.SetMaxIdleConns(10)  // 대기 커넥션 최대 개수를 설정</span>
	}

	if err != nil {
		return nil, err
	}

	return db, err
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func currentTime(c *gin.Context) {
	con, err := connectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	var time string
	err = con.QueryRow("select now()").Scan(&time)

	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"CurrentTime": time,
	})
}

func findUser(c *gin.Context) {

	user := c.Params.ByName("name")

	con, err := connectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	var password string
	var dept string
	row := con.QueryRow("SELECT USER_NAME, USER_ACCOUNT_PASSWORD, DEPT_NAME FROM BLOCK_USER WHERE USER_ID = ?", user)
	err = row.Scan(&name, &password, &dept)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"name": name, "password": password, "dept": dept})
}

func findAdminUser(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)

	fmt.Println("user : " + user)

	con, err := connectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	var password string
	var dept string
	var keyFile string
	row := con.QueryRow("SELECT USER_NAME, USER_ACCOUNT_PASSWORD, DEPT_NAME, USER_KEY_FILE FROM BLOCK_USER WHERE USER_ID = ?", user)
	err = row.Scan(&name, &password, &dept, &keyFile)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"name": name, "password": password, "dept": dept, "keyFile": keyFile})
}

func myRouter() *gin.Engine {
	// Default Router
	r := gin.Default()

	// Ping test
	r.GET("/ping", pong)

	// Time test
	r.GET("/time", currentTime)

	// Get user value
	r.GET("/user/:name", findUser)

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"ether1": "ether1", // user:ether1, password:bether1ar
		"mall":   "mall",   // user:mall, password:mall
	}))

	authorized.GET("/user", findAdminUser)

	return r
}

func main() {
	router := myRouter()

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
