import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AccommodationsService } from '../accommodations.service';
import { AccommodationDTO } from '../model/accommodationDTO';
import { Location } from '@angular/common';

export interface UpdatePriceRequest {
  id: string;
  price: number;
  isPerGuest: boolean;
  hasWeekend: boolean;
  hasSummer: boolean;
}

@Component({
  selector: 'app-change-price',
  templateUrl: './change-price.component.html',
  styleUrls: ['./change-price.component.css']
})
export class ChangePriceComponent implements OnInit {

  id: string = '';
  accommodation: AccommodationDTO = {} as AccommodationDTO;

  constructor(private accommodationService: AccommodationsService, 
    private route: ActivatedRoute, private location: Location) {}

  ngOnInit(): void {
    const accId = this.route.snapshot.paramMap.get('id');
    if (accId) {
      this.id = accId;
    }

    this.accommodationService.getAccommodation(this.id).subscribe((res: AccommodationDTO) => {
      this.accommodation = Object.values(res)[0];
      console.log(res);
    });

  }

  done() {
    let updatePrice: UpdatePriceRequest = {} as UpdatePriceRequest;

    updatePrice.id = this.id;
    updatePrice.price = this.accommodation.price;
    updatePrice.isPerGuest = this.accommodation.isPerGuest;
    updatePrice.hasWeekend = this.accommodation.hasWeekend;
    updatePrice.hasSummer = this.accommodation.hasSummer;

    this.accommodationService.updatePrice(updatePrice).subscribe();
    this.location.back()
  }

}
