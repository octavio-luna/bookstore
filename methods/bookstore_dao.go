package methods

import (
	"context"
	"errors"
	"fmt"

	"github.com/octavio-luna/bookstore/datasources"
)

const (
	queryInsertBook        = `Insert into book (title, author, amount_available) values ("%s", "%s", %d);`
	queryDeleteBook        = "Delete from book where title like '%s';"
	queryGetBookByTitle    = "Select * from book where title like '%s';"
	queryGetBookByAuthor   = "Select * from book where author like '%s'"
	queryGetAvailableBooks = "Select * from book where amount_available > 0"

	queryGetMemberByName     = "Select * from club_member where name like '%s';"
	queryGetMemberByLastName = "Select * from club_member where name like '%s'"
	queryInsertMember        = `Insert into club_member (name, last_name) values ("%s", "%s");`
	queryDeleteMember        = "Delete from club_member where name like '%s' and last_name like '%s'"
	queryGetMemberWithBooks  = "Select * from club_member where id in (select club_member_id from rent);"
	queryGetAllMembers       = "Select * from club_member"

	queryCreateRent         = "Insert into rent (book_id, club_member_id) values (%d, %d);"
	queryUpdateOneMinusBook = "Update book set amount_available = (amount_available - 1) where id = %d;"
	queryDeleteRent         = `Delete from rent where book_id in (select id from book where title like "%s") and club_member_id in (select club_member_id from club_member where name like "%s" and last_name like "%s");`
	queryUpdateOneMoreBook  = `Update book set amount_available = (amount_available + 1) where title like "%s";`
)

func GetAllBooks() ([]Book, error) {
	ctx := context.Background()

	psql := "Select * from book"

	var b []Book
	result, err := datasources.DB.QueryContext(
		ctx,
		psql)
	if err != nil {
		return b, err
	}
	for result.Next() {
		var book Book
		err = result.Scan(&book.ID, &book.Title, &book.Author, &book.Amount_available)
		if err != nil {
			return b, err
		}
		b = append(b, book)
	}
	return b, nil
}

func InsertBook(newBook Book) error {
	ctx := context.Background()

	var err error
	if datasources.DB == nil {
		err = errors.New("Db is null")
		return err
	}

	err = datasources.DB.PingContext(ctx)
	if err != nil {
		return err
	}

	tsql := fmt.Sprintf(queryInsertBook, newBook.Title, newBook.Author, newBook.Amount_available)

	_, err = datasources.DB.ExecContext(
		ctx,
		tsql)
	if err != nil {
		return err
	}
	return nil
}

func GetAvailableBooks() ([]Book, error) {
	ctx := context.Background()
	result, err := datasources.DB.QueryContext(
		ctx,
		queryGetAvailableBooks,
	)
	if err != nil {
		return nil, err
	}
	var b []Book
	for result.Next() {
		var s Book
		result.Scan(&s.ID, &s.Title, &s.Author, &s.Amount_available)
		b = append(b, s)
	}
	if err != nil {
		return nil, err
	}
	return b, nil
}

func DeleteBookByTitle(title string) error {
	ctx := context.Background()

	psql := fmt.Sprintf(queryDeleteBook, title)
	_, err := datasources.DB.QueryContext(
		ctx,
		psql)
	if err != nil {
		return err
	}
	return nil
}

func GetBooksByTitle(title string) ([]Book, error) {
	ctx := context.Background()
	psql := fmt.Sprintf(queryGetBookByTitle, title)

	var b []Book
	result, err := datasources.DB.QueryContext(
		ctx,
		psql)
	if err != nil {
		return b, err
	}
	for result.Next() {
		var s Book
		err = result.Scan(&s.ID, &s.Title, &s.Author, &s.Amount_available)
		if err != nil {
			return b, err
		}
		b = append(b, s)
	}
	return b, nil

}

func GetBooksByAuthor(title string) ([]Book, error) {
	ctx := context.Background()
	psql := fmt.Sprintf(queryGetBookByAuthor, title)

	var b []Book
	result, err := datasources.DB.QueryContext(
		ctx,
		psql)
	if err != nil {
		return b, err
	}
	for result.Next() {
		var s Book
		err = result.Scan(&s.ID, &s.Title, &s.Author, &s.Amount_available)
		if err != nil {
			return b, err
		}
		b = append(b, s)
	}
	return b, nil

}

func GetMembersByName(name string) ([]Club_Member, error) {
	ctx := context.Background()
	psql := fmt.Sprintf(queryGetMemberByName, name)

	var m []Club_Member
	result, err := datasources.DB.QueryContext(
		ctx,
		psql)
	if err != nil {
		return m, err
	}
	for result.Next() {
		var s Club_Member
		err = result.Scan(&s.ID, &s.Name, &s.LastName)
		if err != nil {
			return m, err
		}
		m = append(m, s)
	}
	return m, nil

}

func GetMembersByLastName(name string) ([]Club_Member, error) {
	ctx := context.Background()
	psql := fmt.Sprintf(queryGetMemberByName, name)

	var m []Club_Member
	result, err := datasources.DB.QueryContext(
		ctx,
		psql)
	if err != nil {
		return m, err
	}
	for result.Next() {
		var s Club_Member
		err = result.Scan(&s.ID, &s.Name, &s.LastName)
		if err != nil {
			return m, err
		}
		m = append(m, s)
	}
	return m, nil

}

func InsertMember(newMember Club_Member) error {
	ctx := context.Background()

	psql := fmt.Sprintf(queryInsertMember, newMember.Name, newMember.LastName)

	_, err := datasources.DB.QueryContext(
		ctx,
		psql)
	return err
}

func DeleteMember(name string, last_name string) error {
	ctx := context.Background()

	psql := fmt.Sprintf(queryDeleteMember, name, last_name)

	_, err := datasources.DB.QueryContext(
		ctx,
		psql)
	return err
}

func GetMembersWithRentedBooks() ([]Club_Member, error) {
	ctx := context.Background()

	// psql := fmt.Sprintf(queryGetMemberWithBooks)

	var m []Club_Member
	result, err := datasources.DB.QueryContext(
		ctx,
		queryGetMemberWithBooks)
	if err != nil {
		return m, err
	}
	for result.Next() {
		var s Club_Member
		err = result.Scan(&s.ID, &s.Name, &s.LastName)
		if err != nil {
			return m, err
		}
		m = append(m, s)
	}
	return m, nil
}

func CreateRent(b_id int, m_id int) error {
	ctx := context.Background()

	psql := fmt.Sprintf(queryCreateRent, b_id, m_id)

	_, err := datasources.DB.QueryContext(
		ctx,
		psql)
	_, err = datasources.DB.QueryContext(
		ctx,
		fmt.Sprintf(queryUpdateOneMinusBook, b_id))
	return err
}

func DeleteRentByTitleAndMemberName(title string, memberName string, lastName string) error {
	ctx := context.Background()

	psql := fmt.Sprintf(queryDeleteRent, title, memberName, lastName)

	_, err := datasources.DB.QueryContext(
		ctx,
		psql)
	_, err = datasources.DB.QueryContext(
		ctx,
		fmt.Sprintf(queryUpdateOneMoreBook, title))
	return err
}

func GetAllMembers() ([]Club_Member, error) {
	ctx := context.Background()

	var m []Club_Member
	result, err := datasources.DB.QueryContext(
		ctx,
		queryGetAllMembers)
	if err != nil {
		return m, err
	}
	for result.Next() {
		var s Club_Member
		err = result.Scan(&s.ID, &s.Name, &s.LastName)
		if err != nil {
			return m, err
		}
		m = append(m, s)
	}
	return m, nil
}
