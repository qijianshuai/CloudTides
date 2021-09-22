import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'tide-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }
  users = []
  selected= 'selected'
}
