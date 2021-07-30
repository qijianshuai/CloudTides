import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PRODUCT_NAME } from '@tide-config/const';

import { LoginService, UserInfo } from './login/login.service';
import { TranslateService } from '@ngx-translate/core';
import { I18nService } from '@tide-shared/service/i18n';
import { Observable, Subject } from 'rxjs';
import { RegisterService } from './register/register.service';
import { ResetService } from './reset/reset.service';

import { Location } from '@angular/common';
import { NavigationEnd } from '@angular/router';

@Component({
  selector: 'tide-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  path = '';
  constructor(
    readonly loginService: LoginService,
    readonly resetService: ResetService,
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
      // if (this.path == '/login') {
      //   window.location.reload();
      // }
    });

  }
  
  // useLanguage(language: string): void {
  //   this.translate.use(language);
  // }

  // redirectTo(uri:string){
  //   this.router.navigateByUrl('/', {skipLocationChange: true}).then(()=>
  //   this.router.navigate([uri]));
  //   }

  redirectToLogin() {
    this.router.navigate(['/login'])
    .then(() => {
      window.location.reload()
    })
  }

  routeToVcpp() {
    this.router.navigate(['/vcpp'])
  }

  reloadCurrentPage() {
    window.location.reload();
    }
  readonly vo = {
    title: PRODUCT_NAME,
  };

  subject = new Subject();

  signOut() {
    this.loginService.logout();
  }

  cloudtides_logout() {
    this.loginService.cloudtides_logout()
  }

  cloudtides_reset_code() {
    this.resetService.cloud_reset_code(this.loginService.session.username);
    console.log("here!");
    this.router.navigate(['/cloudtides/reset']);
    
  }

  ngOnInit() {

  }
}
