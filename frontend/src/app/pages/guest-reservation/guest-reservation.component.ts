import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { GuestReservationService } from './services/guest-reservation.service';

@Component({
  selector: 'app-guest-reservation',
  templateUrl: './guest-reservation.component.html',
  styleUrls: ['./guest-reservation.component.css']
})
export class GuestReservationComponent implements OnInit {
  

  reservationForm = new FormGroup({
    accommodationId: new FormControl(null, [Validators.required]),
    startDate: new FormControl(null, [Validators.required]),
    endDate: new FormControl(null, [Validators.required]),
    guestNumber: new FormControl(null, [Validators.required])
  })

  accommodations: any = [];

  constructor( 
    private guestReservationService: GuestReservationService
    ) { }

  ngOnInit(): void {
    this.guestReservationService.getAccommodations()
      .subscribe(res => this.accommodations = res);
  }

  guestReservation(){
    if (!this.reservationForm.valid) {
      return;
    }
    this.guestReservationService.guestReservation(this.reservationForm.value).subscribe();
  }



}
