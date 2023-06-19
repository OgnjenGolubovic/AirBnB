import { Component } from '@angular/core';
import { GuestReservationService, Reservation } from './guest-reservation.service';
import { MatTableDataSource } from '@angular/material/table';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Observable } from 'rxjs';
import { AuthService } from '../login/log-auth.service';

@Component({
  selector: 'app-guest-reservations',
  templateUrl: './guest-reservations.component.html',
  styleUrls: ['./guest-reservations.component.css']
})
export class GuestReservationsComponent {
    roleObs?: Observable<string>;
    role?: string;
  
    public reservations: Reservation[] = [];
    
    displayedColumns: string[] = ['accommodation', 'startDate', 'endDate', 'cancel'];
    dataSource = new MatTableDataSource(this.reservations);
  
    constructor(public dialog: MatDialog, private _reservationService: GuestReservationService, 
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
  }
