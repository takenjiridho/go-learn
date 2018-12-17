import { Component, OnInit } from '@angular/core';
import {BookService} from "../services/book.service";
import {Router} from "@angular/router";
import {Book} from "../model/book.model";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {first} from "rxjs/operators";

@Component({
  selector: 'app-edit-book',
  templateUrl: './edit-book.component.html',
  styleUrls: ['./edit-book.component.css']
})
export class EditBookComponent implements OnInit {

  book: Book
  editForm: FormGroup

  constructor(private formBuilder: FormBuilder,private router: Router,
    private bookService: BookService) { }

  ngOnInit() {
    let bookId = localStorage.getItem('bookId')

    if (!bookId) {
      this.router.navigate(['list-books']);
      return;
    }

    this.editForm = this.formBuilder.group({
      id: [],
      title: ['', Validators.required],
      author: ['', Validators.required],
      year: ['', Validators.required]
    });

    this.bookService.getBookById(+bookId).subscribe(data => {
      for (let key in data) {
        const newKey = key.toLowerCase();
        data[newKey] = data[key]
        delete data[key];
      }

      this.editForm.setValue(data);
    })
  }

  cancel() {
    this.router.navigate(['list-books']);
  }

  onSubmit() {
    this.bookService.updateBook(this.editForm.value).pipe(first())
      .subscribe(
        data => {
          this.router.navigate(['list-books'])
        },
        error => {
          console.log(error)
        });
  }
}
