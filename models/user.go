package models

import (
"fmt"
"net/http"
"github.com/jinzhu/gorm"
"github.com/gin-gonic/gin"
//"github.com/gin-contrib/sessions"
_ "github.com/jinzhu/gorm/dialects/mysql"

)
var Yangming int
var db *gorm.DB
type (
	//when username use lowcase,the db will not include the items
	Accounts struct {
		gorm.Model
		Email     string `json:"email"`
		Username     string `json:"username"`
		Password    string    `json:"password"`
                Uid         string    `json:"uid"` 
	}

   Students  struct {
                gorm.Model
                Email     string `json:"email"`
                Wechatid     string `json:"wechatid"`
                Password    string    `json:"password"`
                Hadboughtcoursesurls     string    `json:"hadboughtcoursesurls"`

        }



             Teachers  struct {
                gorm.Model
                Email     string `json:"email"`
                Teacherid     string `json:"teacherid"`
                Password    string    `json:"password"`
                Description    string     `json:"description"`
                Suppliedcoursesurls     string    `json:"suppliedcoursesurls"`
        }




             Courses  struct {
                gorm.Model
                Courseid     string `json:"Courseid"`
                Courseurls     string    `json:"coursesurls"`
                Description    string     `json:"description"`
                Teacherid    string     `json:"teacherid"`
        }









)

func init() {
	//open a db connection
	//var a =m add(2,3)
	//fmt.Println(a)
	var err error
        fmt.Println("-------i am here to change web --------------")
	//mysql://dt_admin:dt2016@localhost/dreamteam_db
	db, err = gorm.Open("mysql", "root:123456@/yongxinxue?charset=utf8&parseTime=True&loc=Local")
  //db, err := gorm.Open("sqlite3", "./yangming.sqlite")
	//defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
     //db.AutoMigrate(&Account{},&TodoModel{},&Tasks{})
       db.AutoMigrate(&Students{},&Teachers{},&Courses{})
//http://jinzhu.me/gorm/database.html#migration delete database table column
 //db.Model(&Tasks{}).DropColumn("Uer")
}





func  User(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html",nil)
 }

 func Register(c *gin.Context) {

	 fmt.Println("-----------------")
	 fmt.Println(Yangming)
 	 Email := c.PostForm("email")
 	 Passowrd:= c.DefaultPostForm("password", "anonymous")
 	 Username:= c.PostForm("username")
 	 User1 := Accounts{Email: Email,Username:Username,Password:Passowrd}
 	 fmt.Println(Email,Passowrd,Username)
 	 fmt.Println(User1)
	// db, _ = gorm.Open("mysql", "dt_admin:dt2016@/dreamteam_db?charset=utf8&parseTime=True&loc=Local")
 	 db.Save(&User1)
 	 c.HTML(http.StatusOK, "user.html", nil)
  }


func  Login(c *gin.Context) {
	   //cookie set
	  //store := sessions.NewCookieStore([]byte("secret"))
	  //router.Use(sessions.Sessions("mysession", store))
	    email := c.PostForm("email")
      passowrd:= c.PostForm("password")
      client:=c.PostForm("client")
			fmt.Println(client)
			//session := sessions.Default(c)
			//session.Set("count", "yangming")
			//session.Save()
			fmt.Println(email,passowrd,client)
			cookie := &http.Cookie{
							Name:  "username",
							Value: email,
					}
			http.SetCookie(c.Writer, cookie)
			cookie1 := &http.Cookie{
							Name:  "email",
							Value: email,
					}
			http.SetCookie(c.Writer, cookie1)
			cookie2 := &http.Cookie{
							Name:  "logintime",
							Value: "now-nounspecify",
					}
			http.SetCookie(c.Writer, cookie2)
			cookie3 := &http.Cookie{
							Name:  "client",
							Value: client,
					}
			http.SetCookie(c.Writer, cookie3)
					//c.String(http.StatusOK, "0")
			if client == "web"{
				//https://github.com/gin-gonic/gin to redirect
				c.Redirect(http.StatusMovedPermanently, "/inbox")
				//c.Redirect(http.StatusMovedPermanently, "/mainboard")
  		 //c.HTML(http.StatusOK, "user.html", nil)
		 }else{
			 c.JSON(http.StatusOK,  gin.H{
			"status":  "logined",
		})
		 }
  }


func  checkcookie() bool{
	return true
}
