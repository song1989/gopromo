package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1, Book2 Books

	/** book 1描述*/
	Book1.title = "go语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "go 语言教程"
	Book1.book_id = 1
	printBook(Book1)

	Book2.title = "php"
	Book2.author = "www.runoob.com"
	Book2.subject = "php 语言教程"
	Book2.book_id = 2
	printBook(Book2)
}

func printBook(book Books) {
	fmt.Printf("book title: %s\n", book.title)
	fmt.Printf("book author: %s\n", book.author)
	fmt.Printf("book subject: %s\n", book.subject)
	fmt.Printf("book book_id: %d\n", book.book_id)
}
