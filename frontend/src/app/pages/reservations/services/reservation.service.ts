import { HttpHeaders, HttpClient, HttpParams } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { AccommodationsRequest } from "src/app/navbar/navbar.component";
import { environment } from "src/environments/environment";
import { AccommodationDTO } from "../../accommodations/model/accommodationDTO";

export interface Reservation {
    id: string;
    startDate : string;
    endDate : string;
    guestNumber : string;
    status : string;
    userId: string;
}

@Injectable({
    providedIn:'root'
})
export class ReservationService {

    apiHost: string = 'http://localhost:8000/';
    headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });

    constructor(private http: HttpClient) {}

    getReservations() : Observable<any> {
        return this.http.get(this.apiHost + 'reservation/getAllPending');
    }

    hasActiveReservationsForAccommodations(accommodations: AccommodationsRequest) : Observable<boolean> {
        return this.http.post<boolean>(this.apiHost + `reservation/accommodation`, accommodations);
    }

    cancel(id : String) : Observable<any> {
        return this.http.post(this.apiHost + 'user/reservation-cancel/' + id, null);
    }
    accept(id : string) : Observable<any> {
        const params = new HttpParams()
                        .set('id', id);

        return this.http.get(this.apiHost + 'reservation/approve', {params});
    }
    reject(id : string) : Observable<any> {
        const params = new HttpParams()
                        .set('id', id);

        return this.http.get(this.apiHost + 'reservation/reject', {params});
    }
}