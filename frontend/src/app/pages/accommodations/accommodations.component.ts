import { Component, OnInit } from '@angular/core';
import { AccommodationDTO } from './model/accommodationDTO';
import { AccommodationsService } from './accommodations.service';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthService } from '../login/log-auth.service';
import { ReservationService } from '../reservations/services/reservation.service';
import { AccommodationsRequest } from 'src/app/navbar/navbar.component';

@Component({
  selector: 'app-accommodations',
  templateUrl: './accommodations.component.html',
  styleUrls: ['./accommodations.component.css']
})
export class AccommodationsComponent implements OnInit {
  displayedColumns: string[] = ['photo', 'id', 'name', 'location', 'benefits', 'minGuest', 
  'maxGuest', 'price', 'defineDates', 'changePrice'];
  displayedColumnsNonHost: string[] = ['photo', 'name', 'location', 'benefits', 'minGuest', 
  'maxGuest', 'price'];
  accommodations: AccommodationDTO[] = [];
  dataSource = new MatTableDataSource(this.accommodations);
  roleObs?: Observable<string>;
  role: string = '';
  hasActive: boolean = false;

  constructor(private accommodationService: AccommodationsService, private router: Router,
      private authService: AuthService, private reservationService: ReservationService) {
  }

  ngOnInit(): void {
    if (this.authService.isLoggedIn()) {
      this.roleObs = this.authService.getRole();
      this.roleObs.subscribe((res: string) => {
        this.role = res;
      });
    }

    if (this.role === 'Host') {
      this.accommodationService.getAccommodationsByHost(this.authService.getUserId()).subscribe((response: AccommodationDTO[]) => {
        this.accommodations = response;
        console.log(response);
        console.log(this.accommodations);
        this.dataSource = new MatTableDataSource(response);
      });
    } else {
      this.accommodationService.getAccommodations().subscribe((response: AccommodationDTO[]) => {
        this.accommodations = response;
        console.log(response);
        console.log(this.accommodations);
        this.dataSource = new MatTableDataSource(response);
      });
    }
  }


  applyFilter(event: Event) {

    const filterValue = (event.target as HTMLInputElement).value;

    this.dataSource.filterPredicate = (data, filter) => {
      const row = data as AccommodationDTO;
      return this.filterRow(row, filter);
    };

    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  applyFilter2(event: Event) {

    const filterValue = (event.target as HTMLInputElement).value;

    this.dataSource.filterPredicate = (data, filter) => {
      const row = data as AccommodationDTO;
      return this.filterRow2(row, filter);
    };

    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  applyFilter3(event: Event) {

    const filterValue = (event.target as HTMLInputElement).value;

    this.dataSource.filterPredicate = (data, filter) => {
      const row = data as AccommodationDTO;
      return this.filterRow3(row, filter);
    };

    this.dataSource.filter = filterValue.trim().toLowerCase();
  }




  private filterRow(row: AccommodationDTO, filterValue: string): boolean {
    // replace 'name' with the name of the column you want to filter
    return row.name.toLowerCase().includes(filterValue.toLowerCase());
  }

  private filterRow2(row: AccommodationDTO, filterValue: string): boolean {
    // replace 'name' with the name of the column you want to filter
    return row.location.toLowerCase().includes(filterValue.toLowerCase());
  }

  private filterRow3(row: AccommodationDTO, filterValue: string): boolean {
    // broj gostiju
    return row.minGuest<= Number(filterValue) && row.maxGuest>= Number(filterValue);
  }

  createAccommodation() {
    this.router.navigate(['/accommodations/create'])
  }

  defineDates(id: string) {
    this.router.navigate(['/accommodations/define-dates', id])
    // this.reservationService.hasActiveReservationsForAccommodations(this.prepareRequest(id)).subscribe((res: boolean) => {
    //   this.hasActive = Object.values(res)[0];
    //   if (this.hasActive) {
    //     alert('Can\'t edit accommodation because it has active reservations')
    //     return;
    //   } else {
    //     this.router.navigate(['/accommodations/define-dates', id])
    //   }
    // });
  }

  changePrice(id: string) {
    this.reservationService.hasActiveReservationsForAccommodations(this.prepareRequest(id)).subscribe((res: boolean) => {
      this.hasActive = Object.values(res)[0];
      if (this.hasActive) {
        alert('Can\'t edit accommodation because it has active reservations')
        return;
      } else {
        this.router.navigate(['/accommodations/change-price', id])
      }
    });
  }

  prepareRequest(id: string): AccommodationsRequest {
    let accommodation: AccommodationDTO = {} as AccommodationDTO;
    for (let i = 0; i < this.accommodations.length; ++i) {
      if (id === this.accommodations[i].id) {
        accommodation = this.accommodations[i];
      }
    }

    let tempAccommodations: AccommodationDTO[] = [];
    tempAccommodations.push(accommodation);

    let accommodationsRequest: AccommodationsRequest = {} as AccommodationsRequest;
    accommodationsRequest.accommodations = tempAccommodations;

    return accommodationsRequest;
  }

}
