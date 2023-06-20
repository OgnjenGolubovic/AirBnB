import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { AccommodationsService } from '../pages/accommodations/accommodations.service';
import { AccommodationDTO } from '../pages/accommodations/model/accommodationDTO';
import { AuthService } from '../pages/login/log-auth.service';
import { Reservation, ReservationService } from '../pages/reservations/services/reservation.service';
import { DeleteDTO } from '../pages/users/deleteDTO';
import { UserService } from '../pages/users/users.service';

export interface AccommodationsRequest {
  accommodations: AccommodationDTO[];
}

export interface ReservationsResponse {
  reservation: Reservation[];
}

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
  
  constructor(private authService: AuthService, private router: Router, 
    private userService: UserService, private accomService: AccommodationsService,
    private resService: ReservationService) {
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

    if (this.role === 'Guest') {
      this.userService.hasActiveReservations(userId).subscribe((res: ReservationsResponse) => {
        console.log('guest');
        let reservations = Object.values(res)[0]
        if (reservations.length != 0) {
          alert('Can\'t delete profile because you have active reservations');
          return;
        }
  
        this.del.id = userId
        this.userService.delete(this.del).subscribe()
        this.authService.logout();
        this.router.navigate(['login']);
      });
    } else if (this.role === 'Host') {
      console.log('host');
      this.accomService.getAccommodationsByHost(userId).subscribe((res: AccommodationDTO[]) => {
        let accommodations: AccommodationDTO[] = [];
        for (let i = 0; i < Object.values(res).length; ++i) {
          accommodations.push(Object.values(res)[i])
        }

        if (accommodations.length != 0) {
          let accommodationsRequest: AccommodationsRequest = {} as AccommodationsRequest;
          accommodationsRequest.accommodations = accommodations;
          this.resService.hasActiveReservationsForAccommodations(accommodationsRequest).subscribe((res: boolean) => {
            let hasActive = Object.values(res)[0];
            if (hasActive) {
              alert('Can\'t delete profile because you have active reservations for one of the accommodations');
              return;
            } else {
              this.del.id = userId

              this.accomService.deleteAccommodationsForHost(userId).subscribe();
              this.userService.delete(this.del).subscribe()
              this.authService.logout();
              this.router.navigate(['login']);
            }
          });
        } else {
          this.del.id = userId
          this.userService.delete(this.del).subscribe()
          this.authService.logout();
          this.router.navigate(['login']);
        }
      });
    }


  }



}
