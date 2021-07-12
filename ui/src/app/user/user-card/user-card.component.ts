import { Component, EventEmitter, OnInit, Input, Output } from '@angular/core';
import { Item, UserService } from '../user.service';
import { TranslateService } from '@ngx-translate/core';
import { NOTIFICATION_EXIST_TIME } from '@tide-config/const';
import { LoginService } from '../../login/login.service';

@Component({
  selector: 'tide-user-card',
  templateUrl: './user-card.component.html',
  styleUrls: ['./user-card.component.scss'],
})
export class UserCardComponent implements OnInit {

  constructor(
    public readonly translate: TranslateService,
    public readonly userService: UserService,
    public readonly loginService: LoginService,
  ) { }

  @Input() userid = 1;
  @Input() opened = false;
  @Output() save = new EventEmitter();
  @Output() cancel = new EventEmitter();
  
  readonly vo = {
    alertType: '',
    alertText: '',
  };

  ngOnInit() {

  }


  async resetAlert(time?: number) {
    window.setTimeout(() => {
      this.vo.alertText = '';
    }, time || NOTIFICATION_EXIST_TIME);
  }

}
