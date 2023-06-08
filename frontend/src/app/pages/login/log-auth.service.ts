import { HttpClient } from "@angular/common/http"
import { Injectable } from "@angular/core"
import { Observable, tap, switchMap, EMPTY, take, map } from "rxjs"
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
  isLoggedIn: boolean = false;


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
        this.router.navigate(['']);
        this.isLoggedIn = true;
      })/*,
      switchMap(_ => this.getUserData())*/
    );
  }

  logout(): void {
    this.m_UserDataService.setToken = null;
    this.m_UserDataService.setUserData = null;
    this.isLoggedIn = false;
    localStorage.clear();
  }

  getUserData(): Observable<any> {
    return this.m_Http.get(`${environment.hospitalApiUrl}/users/data`).pipe(
      tap((res: any) => {
        this.m_UserDataService.setUserData = res;
      })
    )
  }

}
