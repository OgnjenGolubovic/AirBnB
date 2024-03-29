import { HttpHeaders, HttpClient, HttpParams } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";

export interface Reservation {
    id: string;
    startDate : string;
    endDate : string;
    guestNumber : string;
    status : string;
    userId: string;
}

@Injectable({
  providedIn: 'root'
})
export class GuestReservationService {
  
     apiHost: string = 'http://localhost:8000/';
      headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });
  
      constructor(private http: HttpClient) {}
  
      getReservations() : Observable<any> {
          return this.http.get(this.apiHost + 'reservation/getAll');
      }
  
      cancel(id : String) : Observable<any> {
          return this.http.post(this.apiHost + 'user/reservation-cancel/' + id, null);
      }
  }
