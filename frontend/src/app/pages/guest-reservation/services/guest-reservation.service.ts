import { HttpHeaders, HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable, tap } from "rxjs";
import { Reservation } from "../interfaces/guest-reservation.interface";
import { MatSnackBar } from "@angular/material/snack-bar";


@Injectable({
    providedIn:'root'
})
export class GuestReservationService {
   
    apiHost: string = 'http://localhost:8000/';
    headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });
  
    constructor(
      private http: HttpClient,
      private snackbar: MatSnackBar
    ) { }
  
    getAccommodations(): Observable<any> {
      return this.http.get<any>(this.apiHost + 'accommodation', {headers: this.headers})
    }

    guestReservation(reservationRequest: any): Observable<Reservation> {
        return this.http.post<Reservation>(this.apiHost + 'reservation/reserve', reservationRequest, {headers: this.headers}).pipe(
            tap((res: Reservation) => this.snackbar.open(`Reservation request created successfully`, 'Close', {
              duration: 2000, horizontalPosition: 'right', verticalPosition: 'top'
            }))
          );
    }
    
      

}