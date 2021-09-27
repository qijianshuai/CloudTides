import { Injectable } from '@angular/core';
interface Form {
  name: string
}
@Injectable({
  providedIn: 'root'
})
export class NotebooksService {

  constructor() { }
  form: Form = {
    name: 'Abc'
  }
  createInstance = false
}
