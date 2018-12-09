package main

/*GO*/
//https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3
//https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
//the design mode learn from  https://github.com/beego/samples/blob/master/todo/models/task.go
//the import package learn from https://golang.org/doc/code.html
//please attention use  the things u donnot famialr with
//gopath https://github.com/golang/go/wiki/SettingGOPATH


/*database*/
//in order to keep db from losing ,i using db backup https://www.eversql.com/how-to-transfer-a-mysql-database-between-two-servers/
//https://www.wikihow.com/Delete-a-MySQL-Database




import (
	//"net/http"
	//"github.com/yangming/stringutil"
	//"fmt"
	//"./math"
	"github.com/gin-gonic/gin"
	//"gopkg.in/olahol/melody.v1"
	//"github.com/jinzhu/gorm"
  //"github.com/gin-contrib/sessions"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// the  . https://www.golang-book.com/books/intro/11
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/yongxinxue/models"
)

//var db *gorm.DB
//about init https://stackoverflow.com/questions/24790175/when-is-the-init-function-run

//about init https://stackoverflow.com/questions/24790175/when-is-the-init-function-run




func main() {
	router := gin.Default()
  router.GET("/test", Test)

	router.Run(":8080")

}






