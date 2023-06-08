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
  
  constructor(private authService: AuthService, private router: Router) {

  }

  ngOnInit(): void {
    this.router.events.subscribe(event => {
      if (event.constructor.name === "NavigationEnd") {
       this.isLoggedIn = this.authService.isLoggedIn;
      }
    })
  }

  logout() {
    this.authService.logout();
    this.isLoggedIn = false;
    this.router.navigate(['']);
  }

}
