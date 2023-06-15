import { Component, OnInit } from '@angular/core';
import { AccommodationDTO } from './model/accommodationDTO';
import { AccommodationsService } from './accommodations.service';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';

@Component({
  selector: 'app-accommodations',
  templateUrl: './accommodations.component.html',
  styleUrls: ['./accommodations.component.css']
})
export class AccommodationsComponent implements OnInit {
  displayedColumns: string[] = ['id', 'name', 'location', 'benefits', 'photos', 'minGuest', 'maxGuest', 'defineDates'];
  accommodations: AccommodationDTO[] = [];

  constructor(private accommodationService: AccommodationsService, private router: Router) {
  }

  ngOnInit(): void {
    this.accommodationService.getAccommodations().subscribe((response: AccommodationDTO[]) => {
      this.accommodations = response;
      console.log(response);
      console.log(this.accommodations);
    });
  }

  createAccommodation() {
    this.router.navigate(['/accommodations/create'])
  }

  defineDates(id: string) {
    this.router.navigate(['/accommodations/define-dates', id])
  }

}
