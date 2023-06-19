import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { UpdatePriceRequest } from './change-price/change-price.component';
import { AccommodationBody } from './model/accommodation-body';
import { AccommodationDTO, DateRequest, DatesRange } from './model/accommodationDTO';

@Injectable({
  providedIn: 'root'
})
export class AccommodationsService {

  apiHost: string = 'http://localhost:8000/';
  headers: HttpHeaders = new HttpHeaders({ 'Accept': 'application/json', 'Content-Type': 'application/json' });

  constructor(private http: HttpClient) { }

  getAccommodations() : Observable<AccommodationDTO[]> {
    return this.http.get<AccommodationDTO[]>(this.apiHost + 'accommodation').pipe(
      map((data: any) => data.accommodations)
    );
  }

  getAccommodationsByHost(id: string) : Observable<AccommodationDTO[]> {
    return this.http.get<AccommodationDTO[]>(this.apiHost + `accommodation/getByHost?id=${id}`).pipe(
      map((data: any) => data.accommodations)
    );
  }

  getAccommodation(id: string) : Observable<AccommodationDTO> {
    return this.http.get<AccommodationDTO>(this.apiHost + `accommodation/${id}`);
  }

  getFreeDates(id: string) : Observable<DatesRange[]> {
    return this.http.get<DatesRange[]>(this.apiHost + `accommodation/get?id=` + id);
  }

  create(accommodation: AccommodationBody) {
    // const body=JSON.stringify(accommodation);
    // console.log(body);
    return this.http.post(this.apiHost + 'accommodation', accommodation);
  }

  addFreeDates(dateRequest: DateRequest) {
    return this.http.put(this.apiHost + 'accommodation/dates', dateRequest);
  }

  removeFreeDates(dateRequest: DateRequest) {
    return this.http.put(this.apiHost + 'accommodation/dates/remove', dateRequest);
  }

  updatePrice(UpdatePriceRequest: UpdatePriceRequest) {
    return this.http.put(this.apiHost + 'accommodation/price/update', UpdatePriceRequest);
  }

}
