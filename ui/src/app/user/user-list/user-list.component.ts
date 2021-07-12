import { Component, OnDestroy, OnInit } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Item, UserService } from '../user.service';
import { TranslateService } from '@ngx-translate/core';
import { NOTIFICATION_EXIST_TIME, RESOURCE_USAGE_REFRESH_PERIOD } from '@tide-config/const';
import { LoginService } from 'src/app/login/login.service';

@Component({
  selector: 'tide-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss'],
})
export class UserListComponent implements OnInit, OnDestroy {

  constructor(
    public userService: UserService,
    public readonly translate: TranslateService,
    public readonly loginService: LoginService,
  ) {

  }

  readonly vo = {
    alertType: '',
    alertText: '',
  };

  async delete(id: string) {
    await this.userService.removeItem(id).then(() => {
      this.vo.alertText = `Successfully delete User with id ${id}`;
      this.vo.alertType = 'success';
    }, (error) => {
      this.vo.alertType = 'danger';
      this.vo.alertText = error;
    }).then(() => {
      this.resetAlert();
    });
    this.refreshList();
  }

  async resetAlert(time?: number) {
    window.setTimeout(() => {
      this.vo.alertText = '';
    }, time || NOTIFICATION_EXIST_TIME);
  }

  
  list$: Observable<Item[]> = of([]);
  opened = false;
  refreshInterval: number;
  // selected: Observable<Item[]> = of([])
  orgList: Object = {};


  async save() {
    await this.refreshList();
  }

  cancel() {
    this.opened = false;
  }

  async ngOnInit() {
    await this.refreshList();
  }

  async refreshList() {
    this.list$ = of(await this.userService.getUserList());
    this.refreshInterval = window.setInterval(async () => {
      this.list$ = of(await this.userService.getUserList());
    }, RESOURCE_USAGE_REFRESH_PERIOD);
    this.orgList = Object(await this.userService.getOrgList())
  }

  async displayDetail(id: number){
  }

  ngOnDestroy(): void {
    window.clearInterval(this.refreshInterval);
  }

}
