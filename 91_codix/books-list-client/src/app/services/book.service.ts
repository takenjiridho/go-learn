import { HttpClient, HttpClientModule, HttpHeaders} from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Book } from "../model/book.model"

@Injectable({
  providedIn: 'root'
})
export class BookService {

  baseUrl: string = "http://localhost:8000/books"

  constructor(private http: HttpClient) { }

  getBooks() {
    return this.http.get<Book[]>(this.baseUrl)
  }

  getBookById(id: number) {
    return this.http.get<Book>(this.baseUrl + "/" + id);
  }

  addBook(book: Book) {
    return this.http.post(this.baseUrl, book);
  }

  updateBook(book: Book) {
    return this.http.put(this.baseUrl, book);
  }

  deleteBook(id: number) {
    return this.http.delete(this.baseUrl + "/" + id);
  }

}
