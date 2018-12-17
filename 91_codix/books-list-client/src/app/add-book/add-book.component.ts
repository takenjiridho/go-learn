import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { BookService } from '../services/book.service';
import {Book} from "../model/book.model";
import { first } from 'rxjs/operators';
import { Router }  from '@angular/router';

@Component({
  selector: 'app-add-book',
  templateUrl: './add-book.component.html',
  styleUrls: ['./add-book.component.css']
})
export class AddBookComponent implements OnInit {
  addForm: FormGroup

  constructor(private formBuilder: FormBuilder, private router: Router,
  private bookService: BookService) { }

  ngOnInit() {
    this.addForm = this.formBuilder.group({
      id: [],
      title: ['', Validators.required],
      author: ['', Validators.required],
      year: ['', Validators.required]
    })
  }

  cancel() {
    this.router.navigate(['list-books']);
  }

  onSubmit() {
    if (this.addForm.value.id) {
      delete this.addForm.value.id
    }

    if (this.addForm.status === 'INVALID') {
      console.log("invalid form.")
      return;
    }

    this.bookService.addBook(this.addForm.value).subscribe(data => {
      this.router.navigate(['list-books'])
    })
  }

}
