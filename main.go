package main
import (
   "fmt"
   "net/http"
   "html/template"
   "github.com/gorilla/mux"
   "database/sql"
  _ "github.com/go-sql-driver/mysql"
)
import "github.com/go-martini/martini"
type Article struct{
   Id uint16
   Title,Anons,FullText,Description string
}
var posts=[]Article{}
var showPost=Article{}
func index(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/index.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   t.ExecuteTemplate(w,"index",nil)
}
func createWindows(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/createWindows.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   t.ExecuteTemplate(w,"createWindows",nil)
}
func createPython(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/createPython.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   t.ExecuteTemplate(w,"createPython",nil)
}
func createJava(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/createJava.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   t.ExecuteTemplate(w,"createJava",nil)
}
func createKotlin(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/createKotlin.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   t.ExecuteTemplate(w,"createKotlin",nil)
}
func createHtml_css(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/createHtml_css.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   t.ExecuteTemplate(w,"createHtml_css",nil)
}
func createJavaScript(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/createJavaScript.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   t.ExecuteTemplate(w,"createJavaScript",nil)
}
func createPhp(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/createPhp.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   t.ExecuteTemplate(w,"createPhp",nil)
}
func createRuby(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/createRuby.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   t.ExecuteTemplate(w,"createRuby",nil)
}
//c++windows language//////////////////////////////////////////////////////////////////
func windows(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/windows.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
   if err!=nil{
      panic(err)
   }
   defer db.Close()
   //choose data
     res,err:=db.Query("SELECT * FROM `windows`")
   if err!=nil{
      panic(err)
   }
   posts=[]Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      posts=append(posts,post)
   }
   t.ExecuteTemplate(w,"windows",posts)
}
//python//////////////////////////////////////////////////////////////////
func python(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/python.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
   if err!=nil{
      panic(err)
   }
   defer db.Close()
   //choose data
     res,err:=db.Query("SELECT * FROM `python`")
   if err!=nil{
      panic(err)
   }
   posts=[]Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      posts=append(posts,post)
   }
   t.ExecuteTemplate(w,"python",posts)
}
//java//////////////////////////////////////////////////////
func java(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/java.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
    db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
   if err!=nil{
      panic(err)
   }
   defer db.Close()
   //choose data
     res,err:=db.Query("SELECT * FROM `java`")
   if err!=nil{
      panic(err)
   }
   posts=[]Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      posts=append(posts,post)
   }
   t.ExecuteTemplate(w,"java",posts)
}
//kotlin//////////////////////////////////////////////////////
func kotlin(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/kotlin.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
   if err!=nil{
      panic(err)
   }
   defer db.Close()
   //choose data
     res,err:=db.Query("SELECT * FROM `kotlin`")
   if err!=nil{
      panic(err)
   }
   posts=[]Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      posts=append(posts,post)
   }
   t.ExecuteTemplate(w,"kotlin",posts)
}
//html-css//////////////////////////////////////////////////////
func html_css(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/html_css.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
   if err!=nil{
      panic(err)
   }
   defer db.Close()
   //choose data
     res,err:=db.Query("SELECT * FROM `html_css`")
   if err!=nil{
      panic(err)
   }
   posts=[]Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      posts=append(posts,post)
   }
   t.ExecuteTemplate(w,"html_css",posts)
}
//javaScript//////////////////////////////////////////////////////
func javaScript(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/javaScript.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
   if err!=nil{
      panic(err)
   }
   defer db.Close()
   //choose data
     res,err:=db.Query("SELECT * FROM `javaScript`")
   if err!=nil{
      panic(err)
   }
   posts=[]Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      posts=append(posts,post)
   }
   t.ExecuteTemplate(w,"javaScript",posts)
}
//php//////////////////////////////////////////////////////
func php(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/php.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
   if err!=nil{
      panic(err)
   }
   defer db.Close()
   //choose data
     res,err:=db.Query("SELECT * FROM `php`")
   if err!=nil{
      panic(err)
   }
   posts=[]Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      posts=append(posts,post)
   }
   t.ExecuteTemplate(w,"php",posts)
}
/////////////ruby///////////////////////////////////////////
func ruby(w http.ResponseWriter,r *http.Request){
   t,err := template.ParseFiles("templates/ruby.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
   if err!=nil{
      panic(err)
   }
   defer db.Close()
   //choose data
     res,err:=db.Query("SELECT * FROM `ruby`")
   if err!=nil{
      panic(err)
   }
   posts=[]Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      posts=append(posts,post)
   }
   t.ExecuteTemplate(w,"ruby",posts)
}
//сохранить C++ windows//////////////////////////////////////////////////////////////////
func save_articleWindows(w http.ResponseWriter,r *http.Request){
   title:=r.FormValue("title")
   anons:=r.FormValue("anons")
   full_text:=r.FormValue("full_text")
   description:=r.FormValue("description")
   //Проверка
   if title== "" || anons=="" || full_text=="" || description==""{
       fmt.Fprintf(w,"Не все данные заполнены")
   }else{
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()
      //Send data
      insert,err := db.Query(fmt.Sprintf("INSERT INTO `windows`(`title`,`anons`,`full_text`,`description`)VALUES ('%s' ,'%s','%s','%s')",title,anons,full_text,description))
      if err!=nil{
         panic(err)
      }
      defer insert.Close()
      http.Redirect(w ,r ,"/windows",http.StatusSeeOther)
  }
}
//сохранить данные Python//////////////////////////////////////////////////////////////////
func save_articlePython(w http.ResponseWriter,r *http.Request){
   title:=r.FormValue("title")
   anons:=r.FormValue("anons")
   full_text:=r.FormValue("full_text")
   description:=r.FormValue("description")
   //Проверка
   if title== "" || anons=="" || full_text=="" || description==""{
       fmt.Fprintf(w,"Не все данные заполнены")
   }else{
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()
      //Send data
      insert,err := db.Query(fmt.Sprintf("INSERT INTO `python`(`title`,`anons`,`full_text`,`description`)VALUES ('%s' ,'%s','%s','%s')",title,anons,full_text,description))
      if err!=nil{
         panic(err)
      }
      defer insert.Close()
      http.Redirect(w ,r ,"/python",http.StatusSeeOther)
  }
}
//сохранить данные Java//////////////////////////////////////////////////////////////////
func save_articleJava(w http.ResponseWriter,r *http.Request){
   title:=r.FormValue("title")
   anons:=r.FormValue("anons")
   full_text:=r.FormValue("full_text")
   description:=r.FormValue("description")
   //Проверка
   if title== "" || anons=="" || full_text=="" || description==""{
       fmt.Fprintf(w,"Не все данные заполнены")
   }else{
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()
      //Send data
      insert,err := db.Query(fmt.Sprintf("INSERT INTO `java`(`title`,`anons`,`full_text`,`description`)VALUES ('%s' ,'%s','%s','%s')",title,anons,full_text,description))
      if err!=nil{
         panic(err)
      }
      defer insert.Close()
      http.Redirect(w ,r ,"/java",http.StatusSeeOther)
  }
}
//сохранить данные Kotlin//////////////////////////////////////////////////////////////////
func save_articleKotlin(w http.ResponseWriter,r *http.Request){
   title:=r.FormValue("title")
   anons:=r.FormValue("anons")
   full_text:=r.FormValue("full_text")
   description:=r.FormValue("description")
   //Проверка
   if title== "" || anons=="" || full_text=="" || description==""{
       fmt.Fprintf(w,"Не все данные заполнены")
   }else{
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()
      //Send data
      insert,err := db.Query(fmt.Sprintf("INSERT INTO `kotlin`(`title`,`anons`,`full_text`,`description`)VALUES ('%s' ,'%s','%s','%s')",title,anons,full_text,description))
      if err!=nil{
         panic(err)
      }
      defer insert.Close()
      http.Redirect(w ,r ,"/createKotlin",http.StatusSeeOther)
  }
}
//сохранить данные Html-css//////////////////////////////////////////////////////////////////
func save_articleHtml_css(w http.ResponseWriter,r *http.Request){
   title:=r.FormValue("title")
   anons:=r.FormValue("anons")
   full_text:=r.FormValue("full_text")
   description:=r.FormValue("description")
   //Проверка
   if title== "" || anons=="" || full_text=="" || description==""{
       fmt.Fprintf(w,"Не все данные заполнены")
   }else{
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()
      //Send data
      insert,err := db.Query(fmt.Sprintf("INSERT INTO `html_css`(`title`,`anons`,`full_text`,`description`)VALUES ('%s' ,'%s','%s','%s')",title,anons,full_text,description))
      if err!=nil{
         panic(err)
      }
      defer insert.Close()
      http.Redirect(w ,r ,"/createHtml_css",http.StatusSeeOther)
  }
}
//сохранить данные javaScript//////////////////////////////////////////////////////////////////
func save_articleJavaScript(w http.ResponseWriter,r *http.Request){
   title:=r.FormValue("title")
   anons:=r.FormValue("anons")
   full_text:=r.FormValue("full_text")
   description:=r.FormValue("description")
   //Проверка
   if title== "" || anons=="" || full_text=="" || description==""{
       fmt.Fprintf(w,"Не все данные заполнены")
   }else{
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()
      //Send data
      insert,err := db.Query(fmt.Sprintf("INSERT INTO `javaScript`(`title`,`anons`,`full_text`,`description`)VALUES ('%s' ,'%s','%s','%s')",title,anons,full_text,description))
      if err!=nil{
         panic(err)
      }
      defer insert.Close()
      http.Redirect(w ,r ,"/createJavaScript",http.StatusSeeOther)
  }
}
//сохранить данные php//////////////////////////////////////////////////////////////////
func save_articlePhp(w http.ResponseWriter,r *http.Request){
   title:=r.FormValue("title")
   anons:=r.FormValue("anons")
   full_text:=r.FormValue("full_text")
   description:=r.FormValue("description")
   //Проверка
   if title== "" || anons=="" || full_text=="" || description==""{
       fmt.Fprintf(w,"Не все данные заполнены")
   }else{
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()
      //Send data
      insert,err := db.Query(fmt.Sprintf("INSERT INTO `php`(`title`,`anons`,`full_text`,`description`)VALUES ('%s' ,'%s','%s','%s')",title,anons,full_text,description))
      if err!=nil{
         panic(err)
      }
      defer insert.Close()
      http.Redirect(w ,r ,"/createPhp",http.StatusSeeOther)
  }
}
//сохранить данные Ruby//////////////////////////////////////////////////////////////////
func save_articleRuby(w http.ResponseWriter,r *http.Request){
   title:=r.FormValue("title")
   anons:=r.FormValue("anons")
   full_text:=r.FormValue("full_text")
   description:=r.FormValue("description")
   //Проверка
   if title== "" || anons=="" || full_text=="" || description==""{
       fmt.Fprintf(w,"Не все данные заполнены")
   }else{
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()
      //Send data
      insert,err := db.Query(fmt.Sprintf("INSERT INTO `ruby`(`title`,`anons`,`full_text`,`description`)VALUES ('%s' ,'%s','%s','%s')",title,anons,full_text,description))
      if err!=nil{
         panic(err)
      }
      defer insert.Close()
      http.Redirect(w ,r ,"/createRuby",http.StatusSeeOther)
  }
}
//Показать данные Windows//////////////////////////////////////////////////////////////////
func show_post(w http.ResponseWriter,r *http.Request){
   vars := mux.Vars(r)
   t,err := template.ParseFiles("templates/showWindows.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
    db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()

      //choose data
     res,err:=db.Query(fmt.Sprintf("SELECT * FROM `windows` WHERE `id`='%s'",vars["id"]))
   if err!=nil{
      panic(err)
   }
   showPost=Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      showPost=post
   }
   t.ExecuteTemplate(w,"showWindows",showPost)

}
/////показать данные в Питоне
func show_post_Python(w http.ResponseWriter,r *http.Request){
   vars := mux.Vars(r)
   t,err := template.ParseFiles("templates/showPython.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()

      //choose data
     res,err:=db.Query(fmt.Sprintf("SELECT * FROM `python` WHERE `id`='%s'",vars["id"]))
   if err!=nil{
      panic(err)
   }
   showPost=Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      showPost=post
   }
   t.ExecuteTemplate(w,"showPython",showPost)
}
/////показать данные в Java
func show_post_Java(w http.ResponseWriter,r *http.Request){
   vars := mux.Vars(r)
   t,err := template.ParseFiles("templates/showJava.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()

      //choose data
     res,err:=db.Query(fmt.Sprintf("SELECT * FROM `java` WHERE `id`='%s'",vars["id"]))
   if err!=nil{
      panic(err)
   }
   showPost=Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      showPost=post
   }
   t.ExecuteTemplate(w,"showJava",showPost)
}
/////показать данные в kotlin
func show_post_Kotlin(w http.ResponseWriter,r *http.Request){
   vars := mux.Vars(r)
   t,err := template.ParseFiles("templates/showKotlin.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()

      //choose data
     res,err:=db.Query(fmt.Sprintf("SELECT * FROM `kotlin` WHERE `id`='%s'",vars["id"]))
   if err!=nil{
      panic(err)
   }
   showPost=Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      showPost=post
   }
   t.ExecuteTemplate(w,"showKotlin",showPost)
}
/////показать данные в html_css
func show_post_Html_css(w http.ResponseWriter,r *http.Request){
   vars := mux.Vars(r)
   t,err := template.ParseFiles("templates/showHtml_css.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()

      //choose data
     res,err:=db.Query(fmt.Sprintf("SELECT * FROM `html_css` WHERE `id`='%s'",vars["id"]))
   if err!=nil{
      panic(err)
   }
   showPost=Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      showPost=post
   }
   t.ExecuteTemplate(w,"showHtml_css",showPost)
}
/////показать данные в javaScript
func show_post_JavaScript(w http.ResponseWriter,r *http.Request){
   vars := mux.Vars(r)
   t,err := template.ParseFiles("templates/showJavaScript.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()

      //choose data
     res,err:=db.Query(fmt.Sprintf("SELECT * FROM `javaScript` WHERE `id`='%s'",vars["id"]))
   if err!=nil{
      panic(err)
   }
   showPost=Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      showPost=post
   }
   t.ExecuteTemplate(w,"showJavaScript",showPost)
}
/////показать данные в php
func show_post_php(w http.ResponseWriter,r *http.Request){
   vars := mux.Vars(r)
   t,err := template.ParseFiles("templates/showPhp.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()

      //choose data
     res,err:=db.Query(fmt.Sprintf("SELECT * FROM `php` WHERE `id`='%s'",vars["id"]))
   if err!=nil{
      panic(err)
   }
   showPost=Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      showPost=post
   }
   t.ExecuteTemplate(w,"showPhp",showPost)
}
/////показать данные в Ruby
func show_post_ruby(w http.ResponseWriter,r *http.Request){
   vars := mux.Vars(r)
   t,err := template.ParseFiles("templates/showRuby.html","templates/header.html","templates/footer.html")
   if err!=nil{
      fmt.Fprintf(w,err.Error())
   }
   db ,err:= sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/nanta-code" )
      if err!=nil{
         panic(err)
      }
      defer db.Close()

      //choose data
     res,err:=db.Query(fmt.Sprintf("SELECT * FROM `ruby` WHERE `id`='%s'",vars["id"]))
   if err!=nil{
      panic(err)
   }
   showPost=Article{}
   for res.Next(){
      var post Article
      err=res.Scan(&post.Id,&post.Title,&post.Anons,&post.FullText,&post.Description)
      if err!=nil{
         panic(err)
      }
      showPost=post
   }
   t.ExecuteTemplate(w,"showRuby",showPost)
}
// все функции/////////////////////////////////////////////////
func handleFunc(){
   rtr:=mux.NewRouter()
   rtr.HandleFunc("/",index).Methods("GET")
   rtr.HandleFunc("/windows",windows).Methods("GET")
   rtr.HandleFunc("/python",python).Methods("GET")
   rtr.HandleFunc("/java",java).Methods("GET")
   rtr.HandleFunc("/kotlin",kotlin).Methods("GET")
   rtr.HandleFunc("/html_css",html_css).Methods("GET")
   rtr.HandleFunc("/javaScript",javaScript).Methods("GET")
    rtr.HandleFunc("/php",php).Methods("GET")
   rtr.HandleFunc("/ruby",ruby).Methods("GET")
   rtr.HandleFunc("/createWindows",createWindows).Methods("GET")
   rtr.HandleFunc("/createPython",createPython).Methods("GET")
   rtr.HandleFunc("/createJava",createJava).Methods("GET")
   rtr.HandleFunc("/createKotlin",createKotlin).Methods("GET")
   rtr.HandleFunc("/createHtml_css",createHtml_css).Methods("GET")
   rtr.HandleFunc("/createJavaScript",createJavaScript).Methods("GET")
    rtr.HandleFunc("/createPhp",createPhp).Methods("GET")
   rtr.HandleFunc("/createRuby",createRuby).Methods("GET")
   rtr.HandleFunc("/save_articleWindows",save_articleWindows).Methods("POST")
   rtr.HandleFunc("/save_articlePython",save_articlePython).Methods("POST")
   rtr.HandleFunc("/save_articleJava",save_articleJava).Methods("POST")
   rtr.HandleFunc("/save_articleKotlin",save_articleKotlin).Methods("POST")
   rtr.HandleFunc("/save_articleHtml_css",save_articleHtml_css).Methods("POST")
   rtr.HandleFunc("/save_articleJavaScript",save_articleJavaScript).Methods("POST")
   rtr.HandleFunc("/save_articlePhp",save_articlePhp).Methods("POST")
   rtr.HandleFunc("/save_articleRuby",save_articleRuby).Methods("POST")
   rtr.HandleFunc("/post/{id:[0-9]+}",show_post).Methods("GET","POST")
   rtr.HandleFunc("/post/python/{id:[0-9]+}",show_post_Python).Methods("GET")
   rtr.HandleFunc("/post/java/{id:[0-9]+}",show_post_Java).Methods("GET")
   rtr.HandleFunc("/post/kotlin/{id:[0-9]+}",show_post_Kotlin).Methods("GET")
   rtr.HandleFunc("/post/html_css/{id:[0-9]+}",show_post_Html_css).Methods("GET")
   rtr.HandleFunc("/post/javaScript/{id:[0-9]+}",show_post_JavaScript).Methods("GET")
    rtr.HandleFunc("/post/php/{id:[0-9]+}",show_post_php).Methods("GET")
   rtr.HandleFunc("/post/ruby/{id:[0-9]+}",show_post_ruby).Methods("GET")
   http.Handle("/",rtr)
   http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("./static/"))))
   http.ListenAndServe(":5555",nil)
}
func main(){
   handleFunc()
   m := martini.Classic()
	m.Get("/", func() string {
		return "Hello World"
	})
	m.Run()
}