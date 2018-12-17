import { RouterModule, Routes } from '@angular/router';

import {AddBookComponent} from "./add-book/add-book.component";
import {EditBookComponent} from "./edit-book/edit-book.component";
import {ListBooksComponent} from "./list-books/list-books.component";

const routes: Routes = [
  { path: 'add-book', component: AddBookComponent },
  { path: 'edit-book', component: EditBookComponent },
  { path: 'list-books', component: ListBooksComponent },
  { path: '', component: ListBooksComponent }
];
export const routing = RouterModule.forRoot(routes);
