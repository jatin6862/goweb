package main
import (
  "fmt"
  "net/http"
)
func index_handler(w http.ResponseWriter,r * http.Request){
fmt.Fprintf(w, "Whoooo Go is Amazing")
}
func main(){
  http.HandleFunc("/",index_handler)
  http.ListenAndServe(":8080",nil)
}
