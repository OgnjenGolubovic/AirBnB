import { Component, OnInit } from '@angular/core';
import { Accommodation } from './model/accommodation.model';
import { AccommodationService } from './services/accommodation.service';

@Component({
  selector: 'app-accommodation',
  templateUrl: './accommodation.component.html',
  styleUrls: ['./accommodation.component.css']
})
export class AccommodationComponent implements OnInit{

  public accommodations: Accommodation[] = [];

  constructor(private _accommodationService: AccommodationService) {
  }

  ngOnInit(): void {
    this.getAccommodations();
  }

  public getAccommodations() {
    this._accommodationService.getAccommodations().subscribe(res => {
      this.accommodations = res.accommodations;
      console.log(this.accommodations);
    });
   }

}
