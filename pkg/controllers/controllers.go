package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adiiverma40/BookStore/pkg/models"
	"github.com/adiiverma40/BookStore/pkg/utils"
	"github.com/gorilla/mux"
)




var NewBook  models.Book


func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


func GetBookByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id , err := strconv.ParseInt(bookId, 0,0)
	if err != nil{
		fmt.Println("error", err)
	}

	bookdetails , _ := models.GetBookById(id)

	res, _ := json.Marshal(bookdetails)
	w.Header().Set("content-type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)


}

func CreateBook(w http.ResponseWriter, r *http.Request){
	createBook := &models.Book{}

	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res,_ := json.Marshal(b)
	w.Header().Set("content-type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)


}


func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookid := vars["bookId"]
	id, err := strconv.ParseInt(bookid, 0,0)
	if err != nil{
		fmt.Println("error", err)
	}
	book := models.DeleteBook(id)
	res,_ := json.Marshal(book)
	w.Header().Set("content-type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
