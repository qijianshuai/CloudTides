import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PRODUCT_NAME } from '@tide-config/const';

import { LoginService, UserInfo } from './login/login.service';
import { TranslateService } from '@ngx-translate/core';
import { I18nService } from '@tide-shared/service/i18n';
import { Observable, Subject } from 'rxjs';
import { RegisterService } from './register/register.service';

import { Location } from '@angular/common';

@Component({
  selector: 'tide-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  path = '';
  constructor(
    readonly loginService: LoginService,
    readonly registerService: RegisterService,
    private readonly router: Router,
    translate: TranslateService,
    public readonly i18nService: I18nService,
    private location: Location
  ) {
    // translate.addLangs(['en', 'zh-CN']);
    // // this language will be used as a fallback when a translation isn't found in the current language
    // translate.setDefaultLang('zh-CN');

    this.router.events.subscribe((val) => {
      this.path = this.location.path();
    });
  }
  
  // useLanguage(language: string): void {
  //   this.translate.use(language);
  // }

  readonly vo = {
    title: PRODUCT_NAME,
  };

  subject = new Subject();

  signOut() {
    this.loginService.logout();
  }

  ngOnInit() {

  }
}
