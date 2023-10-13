package controllers

import(
"encoding/json"
"fmt"
"github.com/gorilla/mux"
"net/http"
"strconv"
"github.com/vadiraj/pro3/pkg/utils"
"github.com/vadiraj/pro3/pkg/models"
)
var NewBook models.Book
func GetBooks(w http.ResponseWriter,r *http.Request){
newBooks:=models.GetAllBooks()
res,_:=json.Marshal(newBooks)
w.Header().Set("Content-Type","pkglication/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}
func GetBookById(w http.ResponseWriter,r *http.Request){
vars:=mux.Vars(r)
bookId:=vars["bookId"]
ID,err:=strconv.ParseInt(bookId,0,0)
if err!=nil{
fmt.Println("error while parsing")
}
bookDetails,_:=models.GetBookById(ID)
res,_:=json.Marshal(bookDetails)
w.Header().Set("Content-Type","pkglication/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}
func CreateBook(w http.ResponseWriter,r *http.Request){
newBook:=&models.Book{}
utils.ParseBody(r,newBook)
b:=newBook.CreateBook()
res,_:=json.Marshal(b)
w.Header().Set("Content-Type","pkglication/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}
func DeleteBook(w http.ResponseWriter,r *http.Request){
vars:=mux.Vars(r)
bookId:=vars["bookId"]
ID,err:=strconv.ParseInt(bookId,0,0)
if err!=nil{
fmt.Println("Error in parsing")
}
book:=models.DeleteBook(ID)
res,_:=json.Marshal(book)
w.Header().Set("Content-Type","pkglication/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}
func UpdateBook(w http.ResponseWriter,r *http.Request){
var updateBook=&models.Book{}
utils.ParseBody(r,updateBook)
vars:=mux.Vars(r)
bookId:=vars["bookId"]
ID,err:=strconv.ParseInt(bookId,0,0)
if err!=nil{
fmt.Println("error while parsing")
}
bookDetails,db:=models.GetBookById(ID)
if updateBook.Name!=""{
bookDetails.Name=updateBook.Name
}
if updateBook.Author!=""{
bookDetails.Author=updateBook.Author
}
if updateBook.Publication!=""{
bookDetails.Publication=updateBook.Publication
}
db.Save(&bookDetails)
res,_:=json.Marshal(bookDetails)
w.Header().Set("Content-Type","pkglication/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}