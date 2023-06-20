import { HttpClient } from "@angular/common/http"
import { Injectable } from "@angular/core"
import { Observable, tap, switchMap, EMPTY, take, map, BehaviorSubject } from "rxjs"
import { environment } from "src/environments/environment"
import { UserDataService } from "src/app/pages/login/log-user-data.service"
import { Router } from "@angular/router"

export interface LoginDTO {
  username: string;
  password: string;
}

@Injectable({
  providedIn: "root"
})
export class AuthService {
  loginStatus = new BehaviorSubject<boolean>(this.checkLoginStatus());
  role = new BehaviorSubject<string>(this.extractRole());
  userId: string = '';

  constructor(private m_UserDataService: UserDataService, private m_Http: HttpClient, private router: Router) {
    this.m_UserDataService.m_Token$.pipe(
      take(1),
      switchMap(token => {
        if (token) return this.getUserData();
        return EMPTY;
      })
    ).subscribe();
  }

  login(loginDTO: LoginDTO): Observable<any> {
    return this.m_Http.post(`${environment.hospitalApiUrl}/auth/login`, loginDTO).pipe(
      map((res: any) => {
        this.m_UserDataService.setToken = res['accessToken'];
        this.setRole();
        this.setUserId()
        this.router.navigate(['']);
        this.loginStatus.next(true);
      })/*,
      switchMap(_ => this.getUserData())*/
    );
  }

  logout(): void {
    this.m_UserDataService.setToken = null;
    this.m_UserDataService.setUserData = null;
    this.role.next('');
    this.userId = '';
    this.loginStatus.next(false);
    localStorage.clear();
  }

  getUserData(): Observable<any> {
    return this.m_Http.get(`${environment.hospitalApiUrl}/users/data`).pipe(
      tap((res: any) => {
        this.m_UserDataService.setUserData = res;
      })
    )
  }

  setRole() {
    let decodedJWT;
    let accessToken = localStorage.getItem('token');
    if (accessToken != null) {
        decodedJWT = JSON.parse(window.atob(accessToken.split('.')[1]));
    }
    this.role.next(decodedJWT.role);
  }

  private extractRole() {
    if (this.checkLoginStatus() === false) {
      return '';
    }
    let decodedJWT;
    let accessToken = localStorage.getItem('token');
    if (accessToken != null) {
        decodedJWT = JSON.parse(window.atob(accessToken.split('.')[1]));
    }
    return decodedJWT.role;
  }

  private setUserId() {
    if (this.checkLoginStatus() === false) {
      return;
    }
    let decodedJWT;
    let accessToken = localStorage.getItem('token');
    if (accessToken != null) {
        decodedJWT = JSON.parse(window.atob(accessToken.split('.')[1]));
    }
    this.userId = decodedJWT.user_id;
  }

  getRole() {
    return this.role.asObservable();
  }

  getUserId() {
    return this.userId;
  }

  checkLoginStatus(): boolean {
    var token = localStorage.getItem('token');

    if(token) {
      return true;
    }

    return false;

  }

  isLoggedIn() {
    return this.loginStatus.asObservable();
  }

}
