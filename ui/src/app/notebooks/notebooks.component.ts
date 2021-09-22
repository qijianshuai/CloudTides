import { Component, OnInit } from '@angular/core';
import '@cds/core/icon/register.js';
import { ClarityIcons, plusCircleIcon, refreshIcon, playIcon, stopIcon, powerIcon, trashIcon, listIcon } from '@cds/core/icon'
ClarityIcons.addIcons(plusCircleIcon)
ClarityIcons.addIcons(refreshIcon)
ClarityIcons.addIcons(playIcon)
ClarityIcons.addIcons(stopIcon)
ClarityIcons.addIcons(powerIcon)
ClarityIcons.addIcons(trashIcon)
ClarityIcons.addIcons(listIcon)
@Component({
  selector: 'tide-notebooks',
  templateUrl: './notebooks.component.html',
  styleUrls: ['./notebooks.component.scss']
})
export class NotebooksComponent implements OnInit {

  constructor() { }
  ngOnInit(): void {
  }

}
