package main
import (
   "fmt"
   "net/http"
   "html/template"
   "github.com/gorilla/mux"
)
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

// все функции/////////////////////////////////////////////////
func handleFunc(){
   rtr:=mux.NewRouter()
   rtr.HandleFunc("/",index).Methods("GET")
   http.Handle("/",rtr)
   http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("./static/"))))
   http.ListenAndServe(":5555",nil)
}
func main(){
   handleFunc()
}
