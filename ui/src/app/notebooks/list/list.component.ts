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
  users = [
    {
      name: 'Jupyter',
      logo: 'assets/img/jupyter.svg'
    }
  ]
  selected= 'selected'
  filterSearchValue: string = ''
  toJupyter () {
    let form = document.createElement('form')
    form.action="http://120.133.15.12:8888/lab"
    form.method='get'
    form.target = '_blank'
    const body = document.documentElement
    body.appendChild(form)
    const input = document.createElement('input')
    input.type = 'hidden'
    input.name = 'token'
    input.value = '1dc53b34f46aff0f91f8c65ec96f55eb3057d3770e2253b8'
    form.appendChild(input)
    setTimeout(() => {
      form.submit()
      form = null
    }, 1000)
  }

}
