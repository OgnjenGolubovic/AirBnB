import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { MatCalendarCellCssClasses } from '@angular/material/datepicker';
import { ActivatedRoute } from '@angular/router';
import { AccommodationsService } from '../accommodations.service';
import { AccommodationDTO } from '../model/accommodationDTO';

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
  freeDates: string[] = [];

  constructor(private accommodationService: AccommodationsService, private route: ActivatedRoute) {}

  ngOnInit(): void {
    const accId = this.route.snapshot.paramMap.get('id');
    if (accId) {
      this.id = accId;
    }

    this.accommodationService.getAccommodation(this.id).subscribe((res: AccommodationDTO) => {
      this.accommodation = res;
      // console.log(this.accommodation);
      this.accommodationService.getFreeDates(this.id).subscribe((res: string[]) => {
        this.accommodation.dates = Object.values(res).map(o => Object.values(o)[0]);
        this.freeDates = this.accommodation.dates

        this.freeDates.map(strDate => {
          this.freeDates.push(Object.values(strDate)[0]);
          this.freeDates.push(Object.values(strDate)[1]);
        });

        this.freeDates.shift();
        console.log(this.freeDates);
;
        // let test: Date[] = this.freeDates.map(strDate => new Date(strDate));

        // console.log(test);
      });
    });
  }

  dateClass() {
    return (date: Date): MatCalendarCellCssClasses => {
      const highlightDate = this.freeDates.map(strDate => new Date(strDate))
      .some(d => d.getDate() === date.getDate() && d.getMonth() === date.getMonth() && d.getFullYear() === date.getFullYear());
    
      return highlightDate ? 'highlight' : '';
    };
  }

}
