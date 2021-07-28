import { UserInfo } from './../login/login.service';
import { LoginService } from 'src/app/login/login.service';
import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { RESET_API_URL, RESET_PATH } from '@tide-config/path';
import { environment } from '@tide-environments/environment';
import { tap } from 'rxjs/operators';
import { DOCUMENT } from '@angular/common';
import { LoginComponent } from '../login/login.component';

@Injectable()
export class ResetService {

constructor(
  private readonly http: HttpClient,
  public readonly loginService: LoginService,
  @Inject(DOCUMENT) private readonly document: Document,
) { }

reset(
  username = '',
  password = '',
  newPassword = '',
) {
  return this.http.post<ResetResult>(environment.apiPrefix + RESET_API_URL,
    {  username, password, newPassword }).pipe(
    tap(val => {

    }),
  );
}

inResetPage() {
  return this.document.location.pathname === '/cloudtides' + RESET_PATH;
}
}

export interface ResetResult {
  userinfo: {
    username: string,
    password: string,
    pwReset: string,
  }
}
