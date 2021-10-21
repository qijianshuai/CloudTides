import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { environment } from '@tide-environments/environment';
import { Router } from '@angular/router'
import { tap } from 'rxjs/operators';
@Injectable({
  providedIn: 'root'
})
export class NotebooksService {

  constructor(private readonly http: HttpClient,private readonly router: Router) { }
  getJupyterToken () {
    const CloudTidesUserToken = localStorage.getItem('CloudTidesUserToken')
    if (!CloudTidesUserToken) {
      this.router.navigateByUrl('/login')
      return false
    }
    return this.http.post(environment.apiPrefix + '/api', {CloudTidesUserToken}).pipe(
      tap(data => {})
    )
  }
  createNewNotebook (form) {
    return this.http.post(environment.apiPrefix + '/api', {form}).pipe(
      tap(data => {})
    )
  }
  getDefaultInfo () {
    return this.http.get(environment.apiPrefix + '/api').pipe(
      tap(data => {})
    )
  }
  getInstanceTempldate () {
    return this.http.get(environment.apiPrefix + '/api').pipe(
      tap(data => {})
    )
  }
}
