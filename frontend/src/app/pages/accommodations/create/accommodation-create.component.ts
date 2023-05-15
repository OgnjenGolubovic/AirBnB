import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { AccommodationsService } from '../accommodations.service';
import { AccommodationBody } from '../model/accommodation-body';
import { AccommodationDTO } from '../model/accommodationDTO';

@Component({
  selector: 'app-accommodation-create',
  templateUrl: './accommodation-create.component.html',
  styleUrls: ['./accommodation-create.component.css']
})
export class AccommodationCreateComponent implements OnInit {
  accommodation: AccommodationDTO = {} as AccommodationDTO;
  ab: AccommodationBody = {} as AccommodationBody;

  constructor(private location: Location, private accommodationService: AccommodationsService) {
  }

  ngOnInit(): void {

  }

  public goBack() {
    this.location.back()
  }

  public create() {
    this.ab.accommodation = this.accommodation;

    this.accommodationService.create(this.ab).subscribe((response) => {
      console.log(response);
      this.location.back();
    });
  }

}
