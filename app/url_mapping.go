package app

import (
	"net/http"

	"github.com/octavio-luna/bookstore/controllers"
)

func mapUrls() {

	books_router := router.PathPrefix("/books").Subrouter()
	books_router.Path("").Methods(http.MethodGet).HandlerFunc(controllers.BookController.GetAllBooks)
	books_router.Path("").Methods(http.MethodPut).HandlerFunc(controllers.BookController.InsertBook)
	books_router.Path("/title/{title}").Methods(http.MethodGet).HandlerFunc(controllers.BookController.GetBooksByTitle)
	books_router.Path("/title/{title}").Methods(http.MethodDelete).HandlerFunc(controllers.BookController.DeleteBookByTitle)
	books_router.Path("/author/{author}").Methods(http.MethodGet).HandlerFunc(controllers.BookController.GetBooksByAuthor)
	books_router.Path("/available").Methods(http.MethodGet).HandlerFunc(controllers.BookController.GetAvailableBooks)

	members_router := router.PathPrefix("/members").Subrouter()
	members_router.Path("").Methods(http.MethodGet).HandlerFunc(controllers.BookController.GetAllMembers)
	members_router.Path("/withbooks").Methods(http.MethodGet).HandlerFunc(controllers.BookController.GetMembersWithRentedBooks)
	members_router.Path("").Methods(http.MethodPut).HandlerFunc(controllers.BookController.InsertMember)
	members_router.Path("/{name}-{lastname}").Methods(http.MethodDelete).HandlerFunc(controllers.BookController.DeleteMemberByNameAndLastName)
	members_router.Path("/name/{name}").Methods(http.MethodGet).HandlerFunc(controllers.BookController.GetMembersByName)
	members_router.Path("/lastname/{lastname}").Methods(http.MethodGet).HandlerFunc(controllers.BookController.GetMembersByLastName)
	//members_router.Path("/rent/{book_id}-{member-id}").Methods(http.MethodPut).HandlerFunc(controllers.BookController.CreateRent)
	members_router.Path("/rent/{title}-{member_name}-{member_lastname}").Methods(http.MethodDelete).HandlerFunc(controllers.BookController.DeleteRentByTitleAndMemberName)

}
