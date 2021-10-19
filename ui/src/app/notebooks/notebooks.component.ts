import { Component, OnInit, HostListener } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router'
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
    // const iframe = document.getElementById('inner') as HTMLIFrameElement
    // iframe.onload = () => {
    //   iframe.contentWindow.postMessage('数据', '')
    // }
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
