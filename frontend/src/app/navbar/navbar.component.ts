import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthService } from '../pages/login/log-auth.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  loginStatus? : Observable<boolean>;
  isLoggedIn? : boolean;
  roleObs?: Observable<string>;
  role?: string;
  
  constructor(private authService: AuthService, private router: Router) {
  }

  ngOnInit(): void {
    this.loginStatus = this.authService.isLoggedIn();
    this.loginStatus.subscribe((res: boolean) => {
      this.isLoggedIn = res;
      
    });

    this.roleObs = this.authService.getRole();
    this.roleObs.subscribe((res: string) => {
      this.role = res;
    });

    console.log(this.role);
  }

  logout() {
    this.authService.logout();
    this.isLoggedIn = false;
    this.role = '';
    this.router.navigate(['']);
  }

}
