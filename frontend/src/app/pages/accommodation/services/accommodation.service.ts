import { HttpHeaders, HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { Accommodation } from "../model/accommodation.model";

@Injectable()
export class AccommodationService {

    apiHost: string = 'http://localhost:8000/';
    headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });

    constructor(private http: HttpClient) {}

    getAccommodations() : Observable<Accommodation[]> {
        return this.http.get<Accommodation[]>(this.apiHost + 'accommodationAll', {headers: this.headers});
    }
}