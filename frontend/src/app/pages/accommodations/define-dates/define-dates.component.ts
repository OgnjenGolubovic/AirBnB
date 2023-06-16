import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { MatCalendarCellCssClasses } from '@angular/material/datepicker';
import { ActivatedRoute } from '@angular/router';
import { AccommodationsService } from '../accommodations.service';
import { AccommodationDTO, DateRequest, DatesRange } from '../model/accommodationDTO';
import { Location } from '@angular/common';

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
  removeRange: string = '';

  constructor(private accommodationService: AccommodationsService, 
    private route: ActivatedRoute, private location: Location) {}

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
        // console.log(this.dates[0])

        for (let i = 0; i < this.dates.length; ++i) {
          let start = this.changeToCorrectFormat(this.dates[i].startDate);
          let end = this.changeToCorrectFormat(this.dates[i].endDate);

          let stringDates = this.getDatesBetween(new Date(start), new Date(end));

          for (let j = 0; j < stringDates.length; ++j) {
            this.freeDates.push(stringDates[j]);
          }
        }

        // console.log(this.freeDates);
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
    dateRequest.id = this.id;

    this.accommodationService.addFreeDates(dateRequest).subscribe();
    this.location.back();
  }

  removeDates() {
    // console.log(this.removeRange[0])
    let parts = this.removeRange[0].split('-')

    console.log(parts[0]);
    console.log(parts[1]);
    let dateRequest: DateRequest = {} as DateRequest;

    dateRequest.startDate = parts[0];
    dateRequest.endDate = parts[1];
    dateRequest.id = this.id;

    console.log(dateRequest);

    // let removeDates = this.getDatesBetween(new Date(this.newRange.startDate), new Date(this.newRange.endDate));

    // // const test = this.changeToBackendFormat(this.freeDates[0].toString())
    // console.log('remove:' + removeDates);
    // console.log(this.freeDates);

    // for (let i = 0; i < this.freeDates.length; ++i) {
    //   let currentDate = this.changeToBackendFormat(this.freeDates[i].toString());

    //   for (let j = 0; j < removeDates.length; ++j) {
    //     let dateToRemove = this.changeToBackendFormat(removeDates[j].toString());

    //     // console.log('current:' + currentDate);
    //     // console.log('remove:' + dateToRemove);
    //     if (currentDate === dateToRemove) {
    //       // console.log('deleted');
    //       this.freeDates.splice(i, 1);
    //       --i;
    //       continue;
    //     }
    //   }
      
    // }

    // console.log(this.freeDates);

    this.accommodationService.removeFreeDates(dateRequest).subscribe();
    this.location.back();
  }

}
