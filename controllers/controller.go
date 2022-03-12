package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/octavio-luna/bookstore/methods"
)

var BookController bookControllerInterface = &bookController{}

type bookControllerInterface interface {
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	InsertBook(w http.ResponseWriter, r *http.Request)
	GetAvailableBooks(w http.ResponseWriter, r *http.Request)
	DeleteBookByTitle(w http.ResponseWriter, r *http.Request)
	GetBooksByTitle(w http.ResponseWriter, r *http.Request)
	GetBooksByAuthor(w http.ResponseWriter, r *http.Request)

	GetMembersByName(w http.ResponseWriter, r *http.Request)
	GetAllMembers(w http.ResponseWriter, r *http.Request)
	GetMembersByLastName(w http.ResponseWriter, r *http.Request)
	InsertMember(w http.ResponseWriter, r *http.Request)
	DeleteMemberByNameAndLastName(w http.ResponseWriter, r *http.Request)
	GetMembersWithRentedBooks(w http.ResponseWriter, r *http.Request)

	CreateRent(w http.ResponseWriter, r *http.Request)
	DeleteRentByTitleAndMemberName(w http.ResponseWriter, r *http.Request)
}

type bookController struct{}

func (b *bookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []methods.Book
	books, err := methods.GetAllBooks()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	if err := json.NewEncoder(w).Encode(&books); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (b *bookController) GetBooksByTitle(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	var books []methods.Book
	books, err := methods.GetBooksByTitle(title)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	if err := json.NewEncoder(w).Encode(&books); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (b *bookController) GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	author := mux.Vars(r)["author"]

	var books []methods.Book
	books, err := methods.GetBooksByAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	if err := json.NewEncoder(w).Encode(&books); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (b *bookController) GetAvailableBooks(w http.ResponseWriter, r *http.Request) {
	var books []methods.Book
	books, err := methods.GetAvailableBooks()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	if err := json.NewEncoder(w).Encode(&books); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (b *bookController) DeleteBookByTitle(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	err := methods.DeleteBookByTitle(title)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
}

func (b *bookController) InsertBook(w http.ResponseWriter, r *http.Request) {
	var book methods.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := methods.InsertBook(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
}

func (b *bookController) GetMembersByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	var members []methods.Club_Member
	members, err := methods.GetMembersByName(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application-json")
	if err := json.NewEncoder(w).Encode(&members); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (b *bookController) GetAllMembers(w http.ResponseWriter, r *http.Request) {
	var members []methods.Club_Member
	members, err := methods.GetAllMembers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	if err := json.NewEncoder(w).Encode(&members); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (b *bookController) GetMembersByLastName(w http.ResponseWriter, r *http.Request) {
	lastname := mux.Vars(r)["lastname"]
	var members []methods.Club_Member
	members, err := methods.GetMembersByLastName(lastname)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application-json")
	if err := json.NewEncoder(w).Encode(&members); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (b *bookController) InsertMember(w http.ResponseWriter, r *http.Request) {
	var club_member methods.Club_Member
	if err := json.NewDecoder(r.Body).Decode(&club_member); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := methods.InsertMember(club_member)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
}

func (b *bookController) DeleteMemberByNameAndLastName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	lastname := mux.Vars(r)["lastname"]

	err := methods.DeleteMember(name, lastname)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
}

func (b *bookController) GetMembersWithRentedBooks(w http.ResponseWriter, r *http.Request) {
	var members []methods.Club_Member
	members, err := methods.GetMembersWithRentedBooks()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	if err := json.NewEncoder(w).Encode(&members); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (b *bookController) CreateRent(w http.ResponseWriter, r *http.Request) {
	book := mux.Vars(r)["book_id"]
	member := mux.Vars(r)["member_id"]
	book_id, err := strconv.Atoi(book)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	member_id, err := strconv.Atoi(member)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = methods.CreateRent(book_id, member_id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
}

func (b *bookController) DeleteRentByTitleAndMemberName(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	member_name := mux.Vars(r)["member_name"]
	member_lastname := mux.Vars(r)["member_lastname"]

	err := methods.DeleteRentByTitleAndMemberName(title, member_name, member_lastname)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
}
