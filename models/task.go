package models
import(
  "fmt"
  "time"
  "net/http"
	"strconv"
"github.com/jinzhu/gorm"
//"github.com/gin-contrib/sessions"
"github.com/gin-gonic/gin"
"github.com/bradfitz/slice"
)

//var Modeltest int =5
type (
	// TodoModel describes a TodoModel type
	TodoModel struct {
		gorm.Model
    Title2     string `json:"title2"`
		Title1     string `json:"title1"`
		Title     string `json:"title"`
		Completed int    `json:"completed"`

	}

	// transformedTodo represents a formatted Todo
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

  Tasks struct {
    //gorm.Model this has set id ,cautous!!!http://jinzhu.me/gorm/models.html
		//ID        uint   `json:"id"`
    ID        uint    `gorm:"primary_key"`
    Task      string   `json:"task"`
		User     string `json:"user"`
		Email    string   `json:"email"`
    Place    string   `json:"place"`
    Status  string   `json:"status"`
    Project  string  `json:"project"`
    Plantime  string  `json:"plantime"`
    Finishtime  string `json:"finishtime"`
    Note        string `json:"note"`
    Parentproject  string `json:"parentproject"`
    Ifdissect  string     `json:"ifdissect"`
    AccurateFinishtime  string `json:"AccurateFinishtime"`
    Longitude string `json:"longitude"`
    Latitude  string `json:"latitude"`
	}

  Person struct {
          Name   string
          Emails []string
     }


//{"name":'yangming','children':[]}
     Thinkmapofreview  struct{
       Name                    string
       Children             []Thinkmapofreview
       }


  Projects  struct{
    Name                    string
    Alltasksinproject       []Tasks
    }

    Everyday  struct{
      Name                    string
      Alldays       []Tasks
      }

      Place  struct{
        Name                    string
        Allplaces              []Tasks
        }

 )
















 // createTodo add a new todo
 func Createtaskfromios(c *gin.Context) {
   //emailcookie,_:=c.Request.Cookie("email")
   //fmt.Println(emailcookie.Value)
   //email:=emailcookie.Value
   email := c.PostForm("email")
   inbox := c.PostForm("inbox")
   fmt.Println(inbox)
   project := c.PostForm("project")
   place := c.PostForm("place")
   plantime := c.PostForm("plantime")
   if plantime =="today"{
     loc, _ := time.LoadLocation("Asia/Shanghai")
     //plantimeofanotherforamt :=  time.Now().In(loc)
     //
     plantime =  time.Now().In(loc).Format("060102")
   }
   if plantime  =="tomorrow"{
     loc, _ := time.LoadLocation("Asia/Shanghai")
    plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
   }
   status := c.PostForm("taskstatus")
   parentproject := c.PostForm("parentproject")
   note := c.PostForm("note")
   ifdissect := c.PostForm("ifdissect")

 if status!="unfinished"{
   clientfinishtime:=  c.PostForm("finishtime")
   fmt.Println("=================")
   fmt.Println(clientfinishtime)
   loc, _ := time.LoadLocation("Asia/Shanghai")
   finishtime :=  time.Now().In(loc)
   if clientfinishtime!="unspecified"{
   task := Tasks{Note:note,Ifdissect:ifdissect,Parentproject:parentproject,Task:inbox,User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
   db.Create(&task)
   }else{
   task := Tasks{Note:note,Ifdissect:ifdissect,Parentproject:parentproject,Task:inbox,User:email,Finishtime:finishtime.Format("060102"),Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
   db.Create(&task)
    }


 }else{
   task := Tasks{Task:inbox,User:email,Finishtime:"unfinished",Status:status,Email:email,Place:place,Project:project, Plantime:plantime}
   db.Create(&task)
 }
 c.JSON(200, gin.H{
     "status":  "posted",
     "message": "u have uploaded info,please come on!",
   })
 	}


























// createTodo add a new todo
func Createtask(c *gin.Context) {

  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value
  inbox := c.PostForm("inbox")
  fmt.Println(inbox)
  project := c.PostForm("project")
  place := c.PostForm("place")
  plantime := c.PostForm("plantime")
  if plantime =="today"{
    loc, _ := time.LoadLocation("Asia/Shanghai")
    //plantimeofanotherforamt :=  time.Now().In(loc)
    //
    plantime =  time.Now().In(loc).Format("060102")
  }
  if plantime  =="tomorrow"{
    loc, _ := time.LoadLocation("Asia/Shanghai")
   plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
  }
  status := c.PostForm("taskstatus")
  parentproject := c.PostForm("parentproject")
  note := c.PostForm("note")
  ifdissect := c.PostForm("ifdissect")

if status!="unfinished"{
  clientfinishtime:=  c.PostForm("finishtime")
  fmt.Println("=================")
  fmt.Println(clientfinishtime)
  loc, _ := time.LoadLocation("Asia/Shanghai")
  finishtime :=  time.Now().In(loc)
  if clientfinishtime!="unspecified"{
  task := Tasks{Note:note,Ifdissect:ifdissect,Parentproject:parentproject,Task:inbox,User:email,Finishtime:clientfinishtime,Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
  db.Create(&task)
  }else{
  task := Tasks{Note:note,Ifdissect:ifdissect,Parentproject:parentproject,Task:inbox,User:email,Finishtime:finishtime.Format("060102"),Status:status,Email:email,Place:place, Project:project, Plantime:plantime}
  db.Create(&task)
   }


}else{
  task := Tasks{Task:inbox,User:email,Finishtime:"unfinished",Status:status,Email:email,Place:place,Project:project, Plantime:plantime}
  db.Create(&task)
}
c.JSON(200, gin.H{
    "status":  "posted",
    "message": "u have uploaded info,please come on!",
  })
	}


  // createTodo add a new todo
  func Update(c *gin.Context) {
    emailcookie,_:=c.Request.Cookie("email")
    fmt.Println(emailcookie.Value)
    email:=emailcookie.Value
    inbox := c.PostForm("inbox")
    place := c.PostForm("place")
    fmt.Println(inbox)
    id := c.PostForm("id")
    project := c.PostForm("project")
    finishtime := c.PostForm("finishtime")
    plantime := c.PostForm("plantime")
    if plantime =="today"{
      loc, _ := time.LoadLocation("Asia/Shanghai")
      //plantimeofanotherforamt :=  time.Now().In(loc)
      //
      plantime =  time.Now().In(loc).Format("060102")
    }
    if plantime  =="tomorrow"{
      loc, _ := time.LoadLocation("Asia/Shanghai")
     plantime =  time.Now().In(loc).AddDate(0, 0, 1).Format("060102")
    }



    status := c.PostForm("taskstatus")
    parentproject := c.PostForm("parentproject")
    note := c.PostForm("note")
    //status := c.PostForm("taskstatus")
    ifdissect := c.PostForm("ifdissect")
    fmt.Println(status,plantime,project,id,inbox,email)
    var task Tasks
    db.Where("Email= ?", email).First(&task, id)
    fmt.Println(task)
    fmt.Println(task.Email)
    if task.Email!=email{
      c.JSON(200, gin.H{
          "status":  "posted",
          "message": "updated id not exsit",
        })
        //using python design method to return none
      return
    }
    if place!="unspecified"{db.Model(&task).Update("Place", place)}
    if project!="inbox"{db.Model(&task).Update("Project", project)}
    if inbox!="nocontent"{db.Model(&task).Update("Task", inbox)}
    if plantime!="unspecified"{db.Model(&task).Update("Plantime", plantime)}
    if parentproject!="unspecified"{db.Model(&task).Update("Parentproject", parentproject)}
    if ifdissect!="no"{db.Model(&task).Update("Ifdissect", ifdissect)}
    if note!="unspecified"{db.Model(&task).Update("Note", note)}
    //using it to format time https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
    if status!="unfinished"{
     //locate timezone https://stackoverflow.com/questions/27991671/how-to-get-the-current-timestamp-in-other-timezones-in-golang
      loc, _ := time.LoadLocation("Asia/Shanghai")
      now :=  time.Now().In(loc)

     db.Model(&task).Update("Finishtime",now.Format("060102"))
      //now1 :=  time.Now().In(loc)
      //db.Model(&task).Update("AccurateFinishtime",now1.String()）
      db.Model(&task).Update("Status", status)}

    if finishtime!="unspecified"{db.Model(&task).Update("Finishtime", finishtime)}

 c.JSON(200, gin.H{
                        "status":  "posted",
                        "message": "123",
                        "nick": "234",
                })
        }



    // createTodo add a new todo
    func Test(c *gin.Context) {
      c.JSON(200, gin.H{
    			"status":  "posted",
    			"message": "123",
    			"nick": "234",
    		})
    	}

      func Blockchain(c *gin.Context) {
        c.JSON(200, gin.H{
      			"status":  "posted",
      			"message": "123",
      			"nick": "234",
      		})
      	}










      func  Mainboard(c *gin.Context) {
      /*  c.HTML(http.StatusOK, "inbox.html",gin.H{
          "task":"ha",
        })
*/
       fmt.Println("hahhhahhah============")
          		c.HTML(http.StatusOK, "index.html",nil)
       }






func Inbox(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  db.Where("Email= ?", email).Order("id desc").Find(&tasks)
  //fmt.Println(tasks)
  //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
  //  looptest := "string"
  //fmt.Println(looptest)
  //try to STATISTICS for gtd
  //http://doc.gorm.io/crud.html#query
  /*
  db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Find(&users).Count(&count)
//// SELECT * from USERS WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (users)
//// SELECT count(*) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (count)

db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
//// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)

db.Table("deleted_users").Count(&count)
//// SELECT count(*) FROM deleted_users;


  */
/*
not reference

db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
//// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");



*/



var countofalltasks  int
var countofunfinishedtasks  int
var finishedrate   float64
//reference http://doc.gorm.io/crud.html#query  query with condition
//db.Table("tasks").Where("status = ?", "finish").Count(&countofalltasks)
//db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
//// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
//i use SELECT COUNT(CustomerID), Country FROM Customers GROUP BY Country; to verify which status  items are there?

db.Table("tasks").Not("status", []string{"unfinished"}).Count(&countofunfinishedtasks)//reference not keyword
db.Table("tasks").Count(&countofalltasks)

fmt.Println("+++++++++++++")
fmt.Println(countofalltasks)
fmt.Println(countofunfinishedtasks)
//https://stackoverflow.com/questions/32815400/how-to-perform-division-in-go
finishedrate = float64(countofunfinishedtasks)/float64(countofalltasks)
//strconv.FormatFloat(finishedrate, 'f', -1, 64)
fmt.Println("%.6f",finishedrate)
var finishedratebyend string
//https://gobyexample.com/string-formatting
finishedratebyend = fmt.Sprintf("%.6f", finishedrate)
fmt.Println("+++++++++++++")

 c.HTML(http.StatusOK, "inbox.html",gin.H{
   "task":tasks,"finishedrate":finishedratebyend,
  })

/*
  c.JSON(200, gin.H{
      "task":tasks,
    })
*/
}








func Inboxjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
    db.Where("Email= ?", email).Order("id desc").Find(&tasks)

  c.JSON(200, gin.H{
      "task":tasks,

    })

}



func Unfinishedtaskjson(c *gin.Context) {
  //i use email as identifier
//https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

  //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  //http://doc.gorm.io/crud.html#query to desc
  //db.Where("Email= ?", email).Order("id desc").Find(&tasks)


  //Query Chains http://doc.gorm.io/crud.html#query
  db.Where("Email= ?", email).Where("status in (?)", []string{"unfinish", "unfinished"}).Order("id desc").Find(&tasks)
  c.JSON(200, gin.H{
      "task":tasks,
    })

}
















    func Review(c *gin.Context) {
      //i use email as identifier
    //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
      emailcookie,_:=c.Request.Cookie("email")
      fmt.Println(emailcookie.Value)
      email:=emailcookie.Value
      fmt.Println(email)
      //build search algorithm to get project relationship
      /*
      1.set root project be dm
      2.select datastucture to store
      3.fetch every line to add --------



      */

/*
      var tasks []Tasks
      //email:="yangming1"
      db.Where("Email= ?", email).Find(&tasks)
      alldays:=make(map[string] []Tasks)
      make(map[string]  []string)//{"na
      for _,item :=range tasks{
         alldays[item.Plantime]=append(alldays[item.Plantime],item)
         //alldays[item.Finishtime]=append(alldays[item.Finishtime],item)
      }
      var alleverydays []Everyday
      for k,v := range alldays{
         alleverydays =append(alleverydays,Everyday{k,v})
      }
*/






      c.HTML(http.StatusOK, "review.html", nil)
      	}




func Everydays(c *gin.Context) {
      //i use email as identifier
    //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
      emailcookie,_:=c.Request.Cookie("email")
      fmt.Println(emailcookie.Value)
      email:=emailcookie.Value

      //fmt.Println(cookie1.Value)
      var tasks []Tasks
      //email:="yangming1"
      db.Where("Email= ?", email).Find(&tasks)
      alldays:=make(map[string] []Tasks)
      for _,item :=range tasks{
         alldays[item.Plantime]=append(alldays[item.Plantime],item)
         //alldays[item.Finishtime]=append(alldays[item.Finishtime],item)
      }
      var alleverydays []Everyday
      var unspecifiedday  Everyday
      for k,v := range alldays{
        if k!="unspecified"{
         alleverydays =append(alleverydays,Everyday{k,v})
      }
      if k=="unspecified"{
       //alleverydays =append(alleverydays,Everyday{k,v})
       unspecifiedday = Everyday{k,v}
    }

    }

      slice.Sort(alleverydays, func(i, j int) bool {
return alleverydays[i].Name > alleverydays[j].Name
})
  alleverydays =append(alleverydays,unspecifiedday)


      fmt.Println("====================")
      k:=alleverydays[0].Alldays
      fmt.Println(k[0].ID)
      fmt.Println(k[0])
      fmt.Println("=====================")

      fmt.Println(alleverydays)
      //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
      //  looptest := "string"
      //fmt.Println(looptest)
      c.HTML(http.StatusOK, "everyday.html",gin.H{
       "alldays":alleverydays,
      })
      	}






        func Finished(c *gin.Context) {
              //i use email as identifier
            //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
              emailcookie,_:=c.Request.Cookie("email")
              fmt.Println(emailcookie.Value)
              email:=emailcookie.Value

              //fmt.Println(cookie1.Value)
              var tasks []Tasks
              //email:="yangming1"
              db.Where("Email= ?", email).Find(&tasks)
              alldays:=make(map[string] []Tasks)
              for _,item :=range tasks{
                // alldays[item.Plantime]=append(alldays[item.Plantime],item)
          if item.Status!="unfinished"{
            if item.Status!="unfinish"{
            alldays[item.Finishtime]=append(alldays[item.Finishtime],item)}}
              }
              var alleverydays []Everyday
              var  daybefore180119 Everyday
              for k,v := range alldays{
                if k!="180119before"{
                 alleverydays =append(alleverydays,Everyday{k,v})}
                 if k=="180119before"{
                   daybefore180119 = Everyday{k,v}
                 }


              }


              slice.Sort(alleverydays, func(i, j int) bool {
       return alleverydays[i].Name > alleverydays[j].Name
   })
             alleverydays =append(alleverydays,daybefore180119)

            /*  fmt.Println("====================")
              k:=alleverydays[0].Alldays
              fmt.Println(k[0].ID)
              fmt.Println(k[0])
              fmt.Println("=====================")
             */
              fmt.Println(alleverydays)
              //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
              //  looptest := "string"
              //fmt.Println(looptest)
              c.HTML(http.StatusOK, "pride.html",gin.H{
               "alldays":alleverydays,
              })
              	}





  func Placebased(c *gin.Context) {
  //i use email as identifier
  //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
  emailcookie,_:=c.Request.Cookie("email")
  fmt.Println(emailcookie.Value)
  email:=emailcookie.Value

 //fmt.Println(cookie1.Value)
  var tasks []Tasks
  //email:="yangming1"
  db.Where("Email= ?", email).Find(&tasks)
  allplaces:=make(map[string] []Tasks)
                              for _,item :=range tasks{
                                // alldays[item.Plantime]=append(alldays[item.Plantime],item)
                            //  if item.Place!="unspecified"{
                        //  if item.Status!="unfinish"{
                          //  alldays[item.Finishtime]=append(alldays[item.Finishtime],item)}}
                             allplaces[item.Place]=append(allplaces[item.Place],item)
                             //}
                           }
                             /*var alleverydays []Everyday
                             for k,v := range alldays{
                                alleverydays =append(alleverydays,Everyday{k,v})
                             }
*/
                              var places []Place
                              for k,v := range allplaces{
                                 places =append(places,Place{k,v})
                              }

                              slice.Sort(places, func(i, j int) bool {
                       return places[i].Name < places[j].Name
                   })


                            /*  fmt.Println("====================")
                              k:=alleverydays[0].Alldays
                              fmt.Println(k[0].ID)
                              fmt.Println(k[0])
                              fmt.Println("=====================")
                             */
                             fmt.Println(places)
                              //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
                              //  looptest := "string"
                              //fmt.Println(looptest)
                              c.HTML(http.StatusOK, "place.html",gin.H{
                               "places":places,
                              })
                              	}










func Project(c *gin.Context) {

  //the algorithm can be upgrade
              //i use email as identifier
            //https://github.com/gin-gonic/gin/issues/165 use it to set cookie
      emailcookie,_:=c.Request.Cookie("email")
      fmt.Println(emailcookie.Value)
      email:=emailcookie.Value
      var tasks []Tasks
      //fmt.Println(cookie1.Value)
              //email:="yangming1"
      db.Where("Email= ?", email).Find(&tasks)
      var projects []string
      for _, item := range tasks {

        projects = append(projects,item.Project)
       }
    //get only project
     var onlyprojects []string
     onlyprojects=append(onlyprojects,projects[0])
     for _,item :=range projects{
         piot:="no"
         for _,item1 :=range onlyprojects{
           if item == item1{piot="yes"}
         }
         if piot=="no"{onlyprojects=append(onlyprojects,item)}
     }
     fmt.Println("--------------")
     fmt.Println(onlyprojects)

    //use maps to aviod to design complex algorithm
     allclassproject:=make(map[string] []Tasks)
     for _,item :=range tasks{
        allclassproject[item.Project]=append(allclassproject[item.Project],item)
     }
     var allprojects []Projects
     for k,v := range allclassproject{
        allprojects =append(allprojects,Projects{k,v})

     }


     slice.Sort(allprojects, func(i, j int) bool {
return allprojects[i].Name < allprojects[j].Name
})


     fmt.Println(allclassproject["gtd1"])


              //fmt.Println(tasks)
              //html  render https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
              //  looptest := "string"
              //fmt.Println(looptest)
      c.HTML(http.StatusOK, "project.html",gin.H{
               "projects":allprojects,
        })
      }














// createTodo add a new todo
func CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := TodoModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

// fetchAllTodo fetch all todos
func FetchAllTodo(c *gin.Context) {
	var todos []TodoModel
	var _todos []TransformedTodo

	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

// fetchSingleTodo fetch a single todo
func FetchSingleTodo(c *gin.Context) {
	var todo TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

// updateTodo update a todo
func UpdateTodo(c *gin.Context) {
	var todo TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// deleteTodo remove a todo
func DeleteTodo(c *gin.Context) {
	var todo TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
