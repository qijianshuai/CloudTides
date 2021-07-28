import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { RESET_API_URL, RESET_PATH } from '@tide-config/path';
import { environment } from '@tide-environments/environment';
import { tap } from 'rxjs/operators';
import { DOCUMENT } from '@angular/common';
import { LoginComponent } from '../login/login.component';
import { LoginService } from '../login/login.service';

@Injectable()
export class ResetService {

constructor(
  private readonly http: HttpClient,
  @Inject(DOCUMENT) private readonly document: Document,
) { }

reset(
  password = '',
  newPassword = '',
) {
  return this.http.post<ResetResult>(environment.apiPrefix + RESET_API_URL,
    {  password, newPassword }).pipe(
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
    priority: string,
    password: string,
  }
}
