import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { MatCalendarCellCssClasses } from '@angular/material/datepicker';
import { ActivatedRoute } from '@angular/router';
import { AccommodationsService } from '../accommodations.service';
import { AccommodationDTO, DateRequest, DatesRange } from '../model/accommodationDTO';

@Component({
  selector: 'app-define-dates',
  templateUrl: './define-dates.component.html',
  styleUrls: ['./define-dates.component.css']
})
export class DefineDatesComponent implements OnInit {
  range = new FormGroup({
    start: new FormControl<Date | null>(null),
    end: new FormControl<Date | null>(null),
  });

  id: string = '';
  accommodation: AccommodationDTO = {} as AccommodationDTO;
  freeDates: Date[] = [];
  stringDates: string[] = [];
  dates: DatesRange[] = [];
  newRange: DatesRange = {} as DatesRange;

  constructor(private accommodationService: AccommodationsService, private route: ActivatedRoute) {}

  ngOnInit(): void {
    const accId = this.route.snapshot.paramMap.get('id');
    if (accId) {
      this.id = accId;
    }

    this.accommodationService.getAccommodation(this.id).subscribe((res: AccommodationDTO) => {
      this.accommodation = res;
      this.accommodationService.getFreeDates(this.id).subscribe((res: DatesRange[]) => {

        for (let i = 0; i < Object.values(Object.values(res)[0]).length; ++i) {
          let range: DatesRange = Object.values(Object.values(res)[0])[i];
          this.dates.push(range);
        }

        for (let i = 0; i < this.dates.length; ++i) {
          let start = this.changeToCorrectFormat(this.dates[i].startDate);
          let end = this.changeToCorrectFormat(this.dates[i].endDate);

          let stringDates = this.getDatesBetween(new Date(start), new Date(end));

          for (let j = 0; j < stringDates.length; ++j) {
            this.freeDates.push(stringDates[j]);
          }
        }

        console.log(this.freeDates);
      });
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

  dateClass() {
    return (date: Date): MatCalendarCellCssClasses => {
      const highlightDate = this.freeDates.map(strDate => new Date(strDate))
      .some(d => d.getDate() === date.getDate() && d.getMonth() === date.getMonth() && d.getFullYear() === date.getFullYear());
    
      return highlightDate ? 'highlight' : '';
    };
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

  addNewRange() {
    console.log(this.changeToBackendFormat(this.newRange.startDate));
    console.log(this.changeToBackendFormat(this.newRange.endDate));
    let dateRequest: DateRequest = {} as DateRequest;

    dateRequest.startDate = this.changeToBackendFormat(this.newRange.startDate);
    dateRequest.endDate = this.changeToBackendFormat(this.newRange.endDate);
    dateRequest.id = this.accommodation.id;

    this.accommodationService.addFreeDates(dateRequest).subscribe();
  }

}
