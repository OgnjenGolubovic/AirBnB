import { Injectable } from "@angular/core";
import { ActivatedRoute, ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from "@angular/router";
import { Observable, map, take, switchMap } from "rxjs";
import { UserDataService } from "./log-user-data.service";

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private m_UserDataService: UserDataService, private m_Router: Router, private _activatedRoute: ActivatedRoute) { }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    // if (!localStorage.getItem('token')) {
    //   this.m_Router.navigate(['login']);
    //   return false;
    // }
    
    return this.m_UserDataService.m_Token$.pipe(take(1), switchMap(token => {
      return this.m_UserDataService.m_UserData$.pipe(map(user_data => {
        return !!token ? this.checkRole(user_data?.role, route) : this.m_Router.createUrlTree(['/login']);
      }));
    }));
  }
  checkRole(role: number | undefined, route : ActivatedRouteSnapshot): boolean{

    let jwt = localStorage.getItem('token');
    let decodedJWT;
    if (jwt != null) {
      decodedJWT = JSON.parse(window.atob(jwt.split('.')[1]));
    }
    console.log(route.url.toString());

    if(decodedJWT.role === 'RegisteredUser'){
      if(route.url.toString() === 'bought-tickets'){
        return true;
      }
    }else if(decodedJWT.role === 'Admin'){
      if(route.url.toString() === 'flights' || route.url.toString() === 'flights,create'){
        return true;
      }
    }
    return false;
  }
  getRole(role: number) : string{
    if(role==0)return 'admin';
    if(role==1)return 'registered-user';
    return '';
  }
}
