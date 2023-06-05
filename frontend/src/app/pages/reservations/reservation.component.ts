import { Component, OnInit } from '@angular/core';
import { ReservationService, Reservation } from './services/reservation.service';
import { MatTableDataSource } from '@angular/material/table';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-reservation',
  templateUrl: './reservation.component.html',
  styleUrls: ['./reservation.component.css']
})
export class ReservationComponent implements OnInit{

  public reservations: Reservation[] = [];
  
  displayedColumns: string[] = ['accommodation', 'startDate', 'endDate', 'cancel'];
  dataSource = new MatTableDataSource(this.reservations);

  constructor(public dialog: MatDialog, private _reservationService: ReservationService, private snackBar: MatSnackBar) {
  }

  ngOnInit(): void {
    this.getReservations();
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
}
