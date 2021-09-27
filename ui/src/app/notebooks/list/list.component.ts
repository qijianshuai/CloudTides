import { Component, OnInit } from '@angular/core';
// import '@cds/core/icon/register.js';
// import { ClarityIcons, searchIcon, filter2Icon} from '@cds/core/icon'
// ClarityIcons.addIcons(searchIcon)
// ClarityIcons.addIcons(filter2Icon)
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
  filterSearchValue: string = ''
}
