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
  displayedColumns: string[] = ['id', 'name', 'location', 'benefits', 'photos', 'minGuest', 'maxGuest'];
  accommodations: AccommodationDTO[] = [];
  dataSource = new MatTableDataSource(this.accommodations);

  constructor(private accommodationService: AccommodationsService, private router: Router) {
  }

  ngOnInit(): void {
    this.accommodationService.getAccommodations().subscribe((response: AccommodationDTO[]) => {
      this.accommodations = response;
      console.log(response);
      console.log(this.accommodations);
      this.dataSource = new MatTableDataSource(response);
    });
  }


  applyFilter(event: Event) {

    const filterValue = (event.target as HTMLInputElement).value;

    this.dataSource.filterPredicate = (data, filter) => {
      const row = data as AccommodationDTO;
      return this.filterRow(row, filter);
    };

    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  applyFilter2(event: Event) {

    const filterValue = (event.target as HTMLInputElement).value;

    this.dataSource.filterPredicate = (data, filter) => {
      const row = data as AccommodationDTO;
      return this.filterRow2(row, filter);
    };

    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  applyFilter3(event: Event) {

    const filterValue = (event.target as HTMLInputElement).value;

    this.dataSource.filterPredicate = (data, filter) => {
      const row = data as AccommodationDTO;
      return this.filterRow3(row, filter);
    };

    this.dataSource.filter = filterValue.trim().toLowerCase();
  }




  private filterRow(row: AccommodationDTO, filterValue: string): boolean {
    // replace 'name' with the name of the column you want to filter
    return row.name.toLowerCase().includes(filterValue.toLowerCase());
  }

  private filterRow2(row: AccommodationDTO, filterValue: string): boolean {
    // replace 'name' with the name of the column you want to filter
    return row.location.toLowerCase().includes(filterValue.toLowerCase());
  }

  private filterRow3(row: AccommodationDTO, filterValue: string): boolean {
    // broj gostiju
    return row.minGuest<= Number(filterValue) && row.maxGuest>= Number(filterValue);
  }

  createAccommodation() {
    this.router.navigate(['/accommodations/create'])
  }

}
