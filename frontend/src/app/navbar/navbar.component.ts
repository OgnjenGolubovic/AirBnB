import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthService } from '../pages/login/log-auth.service';
import { DeleteDTO } from '../pages/users/deleteDTO';
import { UserService } from '../pages/users/users.service';

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
  del: DeleteDTO = {} as DeleteDTO;
  
  constructor(private authService: AuthService, private router: Router, private userService: UserService) {
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

  delete() {
    let userId = this.authService.getUserId();
    this.del.id = userId
    this.userService.delete(this.del).subscribe()
    this.authService.logout();
    this.router.navigate(['login']);
  }

}
