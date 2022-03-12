create database bookstore;

use bookstore;

create table book(
id int auto_increment primary key,
title varchar(50),
author varchar (50),
amount_available int
);

create table club_member(
id int auto_increment primary key,
name varchar(20),
last_name varchar(20));

create table rent(
book_id int references book(id),
club_member_id int references club_member(id),
primary key (book_id, club_member_id));
