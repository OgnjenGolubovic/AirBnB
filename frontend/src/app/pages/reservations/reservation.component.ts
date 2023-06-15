import { Component, OnInit } from '@angular/core';
import { ReservationService, Reservation } from './services/reservation.service';
import { MatTableDataSource } from '@angular/material/table';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Observable } from 'rxjs';
import { AuthService } from '../login/log-auth.service';

@Component({
  selector: 'app-reservation',
  templateUrl: './reservation.component.html',
  styleUrls: ['./reservation.component.css']
})
export class ReservationComponent implements OnInit{

  roleObs?: Observable<string>;
  role?: string;

  public reservations: Reservation[] = [];
  
  displayedColumns: string[] = ['accommodation', 'startDate', 'endDate', 'cancel', 'reject', 'accept'];
  dataSource = new MatTableDataSource(this.reservations);

  constructor(public dialog: MatDialog, private _reservationService: ReservationService, 
    private snackBar: MatSnackBar, private authService: AuthService) {
  }

  ngOnInit(): void {
    this.getReservations();

    this.roleObs = this.authService.getRole();
    this.roleObs.subscribe((res: string) => {
      this.role = res;
    });
  }

  public getReservations() {
    this._reservationService.getReservations().subscribe(res => {
      this.reservations = res.reservation;
      this.dataSource = new MatTableDataSource(res.reservation);
    });
   }
  cancel(id : String) {
    this._reservationService.cancel(id).subscribe();
  }

  accept(id : string) {
    this._reservationService.accept(id).subscribe();
  }

  reject(id : string) {
    this._reservationService.reject(id).subscribe();
  }
}
