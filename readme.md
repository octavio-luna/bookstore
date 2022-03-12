API built for the "Biblioteca Aristobulo Del Valle" for local management.
Endpoints: 
    /books (method get): returns an array containing all the books in the database
    
    /books (method put): takes a book object in the request body and inserts it into the database
    
    /books/title/{title} (method get): returns an array of book object with mismatching title
    
    /books/title/{title} (method delete): deletes from the database the book with mismatching Title
    
    /books/author/{author} (method get): returns an array of book objects with the same author name

    /books/available (method get): returns an array of book book object which has at least one available unit

    /members (method get): returns an array compound of all the members

    /members (method put): takes a member object in the request body and inserts it into the database

    /members/withbooks (method get): returns an array of all the members with borrowed books 

    /members/{name}-{lastname} (method delete): drops from the database the member with the given name and last name (if more than one with same name and last name exists, should be done with member_id)

    /members/name/{name} (method get): returns an array of all the members with the given name

    /members/lastname/{lastname} (method get): returns an array of all the members with the given last name

    /members/rent/{title}-{member_name}-{member_lastname} (method delete): deletes from the database the rent row for given user and title


Steps for running:

1. Change the database credentials in /datasources/database.go

2. In the database console run the settings.sql commands

3. (OPTIONAL) for testing purposes, copy the following sql statements into the database console:

    insert into book (title, author, amount_available)
    values 
    ("Atlas shruggled", "Ayn Rand", 1),
    ("1984", "George Orwell", 2),
    ("Moby Dick", "Herman Melville", 1),
    ("Lolita", "Vladimir Nabokov", 1);

    insert into club_member (name, last_name)
    values 
    ("Juan", "perez"),
    ("Emilio", "Rodriguez"),
    ("Irma", "Grill"),
    ("Aaron", "Smith"),
    ("Irina", "Orwell");

    insert into rent (book_id, club_member_id) values (1, 1),
    (2,1),
    (3,2);

4. Run and build
