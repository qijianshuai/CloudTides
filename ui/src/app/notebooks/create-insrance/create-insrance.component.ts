import { Component, OnInit } from '@angular/core';
import { NotebooksService } from '../notebooks.service'
import { Router } from '@angular/router'
interface InstanceModel {
  name: string
}
@Component({
  selector: 'tide-create-insrance',
  templateUrl: './create-insrance.component.html',
  styleUrls: ['./create-insrance.component.scss']
})
export class CreateInsranceComponent implements OnInit {
  noteBook: NotebooksService
  constructor(private noteBooks: NotebooksService, private router: Router) {
    this.noteBook = this.noteBooks
  }

  ngOnInit(): void {
  }
  currentButton = true
  createInstanceFlag = false
  instanceForm:InstanceModel = {
    name:""
  }
  backCreateInstance () {
    this.router.navigate(['/cloudtides/notebooks/list'])
    this.noteBook.createInstance = true
  }
  openInstanceModal () {
    this.createInstanceFlag = true
  }
  cancel () {
    this.createInstanceFlag = false
  }
  create () {
    this.createInstanceFlag = false
  }
  toggleCurrentButton () {
    this.currentButton = !this.currentButton
  }
}
