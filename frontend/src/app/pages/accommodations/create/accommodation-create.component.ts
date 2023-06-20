import { Location } from '@angular/common';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { MatCalendarCellClassFunction } from '@angular/material/datepicker';
import { AuthService } from '../../login/log-auth.service';
import { AccommodationsService } from '../accommodations.service';
import { AccommodationBody } from '../model/accommodation-body';
import { AccommodationDTO } from '../model/accommodationDTO';

@Component({
  selector: 'app-accommodation-create',
  templateUrl: './accommodation-create.component.html',
  styleUrls: ['./accommodation-create.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class AccommodationCreateComponent implements OnInit {
  accommodation: AccommodationDTO = {} as AccommodationDTO;
  ab: AccommodationBody = {} as AccommodationBody;
  selectedDates: string[] = []

  constructor(private location: Location, private accommodationService: AccommodationsService,
    private authService: AuthService) {
  }

  ngOnInit(): void {
    this.accommodation.isPerGuest = false;
    this.accommodation.automaticApproval = false;
    this.accommodation.hasWeekend = false;
    this.accommodation.hasSummer = false;
  }

  public goBack() {
    this.location.back()
  }

  public create() {
    this.ab.accommodation = this.accommodation;
    this.ab.accommodation.dates = [];
    this.ab.accommodation.hostId = this.authService.getUserId();

    console.log(this.ab.accommodation);

    this.accommodationService.create(this.ab).subscribe((response) => {
      console.log(response);
      this.location.back();
    });
  }

  dateClass: MatCalendarCellClassFunction<Date> = (cellDate, view) => {
    if (view == 'month') {
      let dateToFind = this.getDateOnly(cellDate)
      let i = this.selectedDates.indexOf(dateToFind)
      if (i >= 0) {
        return 'selected'
      }
    }
    return ''
  }

  
  daySelected(date: Date | null,calendar: any) {
    console.log(date);
    if (date) {
      let dateSelected = this.getDateOnly(date)
      let i = this.selectedDates.indexOf(dateSelected)
      if (i >= 0) {
        this.selectedDates.splice(i,1)
      } else {
        this.selectedDates.push(dateSelected)
      }
      calendar.updateTodaysDate();
    }
  }

  getDateOnly(date: Date):string {
    return date.toISOString().split('T')[0];
  }

}
