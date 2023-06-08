import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../pages/login/log-auth.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  isLoggedIn: boolean = false;
  role: string = '';
  
  constructor(private authService: AuthService, private router: Router) {

  }

  ngOnInit(): void {
    this.router.events.subscribe(event => {
      if (event.constructor.name === "NavigationEnd") {
       this.isLoggedIn = this.authService.isLoggedIn;
       this.role = this.authService.role;
      }
    })
  }

  logout() {
    this.authService.logout();
    this.isLoggedIn = false;
    this.role = '';
    this.router.navigate(['']);
  }

}
