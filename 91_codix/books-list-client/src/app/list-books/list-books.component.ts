import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {BookService} from "../services/book.service";
import {Book} from "../model/book.model";


@Component({
  selector: 'app-list-books',
  templateUrl: './list-books.component.html',
  styleUrls: ['./list-books.component.css']
})
export class ListBooksComponent implements OnInit {

  books: Book[]
  constructor(private router: Router, private bookService: BookService) { }

  ngOnInit() {
    this.bookService.getBooks().subscribe(data => {
      if (data && Array.isArray(data)) {
        this.books = data.map((item) => {
          for (let key in item) {
              const newKey = key.toLowerCase();
              item[newKey] = item[key]
              delete item[key];
          }
          return item;
        });
      }
    })
  }

  editBook(book: Book): void {
    if (book.id) {
      localStorage.removeItem("bookId");
      localStorage.setItem("bookId", book.id.toString())
      this.router.navigate(['edit-book'])
    }
  }

  deleteBook(book: Book): void {
    this.bookService.deleteBook(book.id).subscribe(data => {
      this.books = this.books.filter(b => b !== book);
    })
  }

  addBook(): void {
    this.router.navigate(['add-book']);
  }

}
