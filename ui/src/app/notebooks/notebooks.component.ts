import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router'
// import '@cds/core/icon/register.js';
import { NotebooksService } from './notebooks.service'
// import { ClarityIcons, plusCircleIcon, refreshIcon, playIcon, stopIcon, powerIcon, trashIcon, listIcon, pencilIcon, searchIcon, filter2Icon, undoIcon, angleIcon} from '@cds/core/icon'
// ClarityIcons.addIcons(plusCircleIcon)
// ClarityIcons.addIcons(refreshIcon)
// ClarityIcons.addIcons(playIcon)
// ClarityIcons.addIcons(stopIcon)
// ClarityIcons.addIcons(powerIcon)
// ClarityIcons.addIcons(trashIcon)
// ClarityIcons.addIcons(listIcon)
// ClarityIcons.addIcons(pencilIcon)
// ClarityIcons.addIcons(searchIcon)
// ClarityIcons.addIcons(filter2Icon)
// ClarityIcons.addIcons(undoIcon)
// ClarityIcons.addIcons(angleIcon)
@Component({
  selector: 'tide-notebooks',
  templateUrl: './notebooks.component.html',
  styleUrls: ['./notebooks.component.scss']
})
export class NotebooksComponent implements OnInit {

  constructor(public router: Router, private noteBooks: NotebooksService) {
    this.noteBook = this.noteBooks
  }
  ngOnInit(): void {
  }
  noteBook: NotebooksService
  cancel () {
    this.noteBook.createInstance = false
  }
  create () {
    this.noteBook.createInstance = false
  }
  advancedOptions () {
    this.noteBook.createInstance = false
    this.router.navigate(['/cloudtides/notebooks/newInstance'])
  }
  createInstanceHandle () {
    if (!this.noteBook.createInstance) {
      this.noteBook.createInstance = true
    }
  }
  flag = false
  aa () {
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

  requestFullScreen(element) {
    // 判断各种浏览器，找到正确的方法
    let requestMethod = element.requestFullScreen || //W3C
    element.webkitRequestFullScreen || //Chrome等
    element.mozRequestFullScreen || //FireFox
    element.msRequestFullScreen; //IE11
    if (requestMethod) {
     requestMethod.call(element);
    }
    // else if (typeof window.ActiveXObject !== "undefined") {//for Internet Explorer
    //  var wscript = new ActiveXObject("WScript.Shell");
    //  if (wscript !== null) {
    //   wscript.SendKeys("{F11}");
    //  }
    // }
   }
}
