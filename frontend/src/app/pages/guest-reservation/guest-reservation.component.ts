import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { AccommodationsService } from '../accommodations/accommodations.service';
import { AccommodationDTO } from '../accommodations/model/accommodationDTO';
import { Reservation } from './interfaces/guest-reservation.interface';
import { GuestReservationService } from './services/guest-reservation.service';

@Component({
  selector: 'app-guest-reservation',
  templateUrl: './guest-reservation.component.html',
  styleUrls: ['./guest-reservation.component.css']
})
export class GuestReservationComponent implements OnInit {

  price: number = 0;
  aId: string = '';
  reservationForm = new FormGroup({
    accommodationId: new FormControl(null, [Validators.required]),
    startDate: new FormControl(),
    endDate: new FormControl(),
    guestNumber: new FormControl(null, [Validators.required])
  })
  accommodations: any = [];

  constructor( 
    private guestReservationService: GuestReservationService,
    private accommodationService: AccommodationsService
    ) { }

  ngOnInit(): void {
    this.guestReservationService.getAccommodations()
      .subscribe(res => this.accommodations = res.accommodations);
  }

  guestReservation(){
    if (!this.reservationForm.valid) {
      return;
    }
    this.guestReservationService.guestReservation(this.reservationForm.value).subscribe();
  }

  calculatePrice() {
    this.price = 0;
    console.log(this.reservationForm.value.accommodationId);
    let accommodationId = this.reservationForm.value.accommodationId+'';
    if (this.reservationForm.value.accommodationId) {
      accommodationId = this.reservationForm.value.accommodationId;
    }

    let guestNumber = this.reservationForm.value.guestNumber+'';
    let start = this.changeToBackendFormat(this.reservationForm.value.startDate+'');
    let end = this.changeToBackendFormat(this.reservationForm.value.endDate+'');
    // new Date(start), new Date(end)
    start = this.changeToCorrectFormat(start);
    end = this.changeToCorrectFormat(end);

    let dates: Date[] = this.getDatesBetween(new Date(start), new Date(end));
    console.log(dates);


    let accommodation: AccommodationDTO = {} as AccommodationDTO;
    this.accommodationService.getAccommodation(accommodationId).subscribe((res: AccommodationDTO) => {
      accommodation = Object.values(res)[0];
      console.log(accommodation);

      for (let i = 0; i < dates.length; ++i) {
        this.price = this.price + (+accommodation.price);
        if (accommodation.hasWeekend && (dates[i].getDay() === 0 || dates[i].getDay() === 6)) {
          this.price = this.price + (this.price * 20 / 100);
        }

        if (accommodation.hasSummer && this.checkIfInSummer(dates[i])) {
          this.price = this.price + (this.price * 5 / 100);
        }

      }
      if(accommodation.isPerGuest) {
        this.price = this.price * (+guestNumber);
      }
    });
  }

  private getDatesBetween(startDate: Date, endDate: Date) {
    const currentDate = new Date(startDate.getTime());
    const dates = [];
    while (currentDate <= endDate) {
      dates.push(new Date(currentDate));
      currentDate.setDate(currentDate.getDate() + 1);
    }
    return dates;
  }

  private changeToCorrectFormat(dateString: string): string {
    let parts: string[] = dateString.split('/');
    return parts[1] + '/' + parts[0] + '/' + parts[2];
  }

  private changeToBackendFormat(dateString: string): string {
    let date = new Date(dateString);
    let temp = ((date.getMonth() > 8) ? (date.getMonth() + 1) : ('0' + (date.getMonth() + 1))) + '/' + ((date.getDate() > 9) ? date.getDate() : ('0' + date.getDate())) + '/' + date.getFullYear();

    return this.changeToCorrectFormat(temp);
  }

  private checkIfInSummer(date: Date): boolean {
    var dateFrom = "06/21/2023";
    var dateTo = "09/23/2023";

    var currentDate = ((date.getMonth() > 8) ? (date.getMonth() + 1) : ('0' + (date.getMonth() + 1))) + '/' + ((date.getDate() > 9) ? date.getDate() : ('0' + date.getDate())) + '/' + date.getFullYear();
    var dateCheck = this.changeToCorrectFormat(this.changeToBackendFormat(currentDate));

    var from = Date.parse(dateFrom);
    var to   = Date.parse(dateTo);
    var check = Date.parse(dateCheck);

    if((check <= to && check >= from))
      return true;   

    return false;
  }


}
