import { Component, OnInit } from '@angular/core';
import { Accommodation } from './model/accommodation.model';
import { AccommodationService } from './services/accommodation.service';
import { MatTableDataSource } from '@angular/material/table';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-accommodation',
  templateUrl: './accommodation.component.html',
  styleUrls: ['./accommodation.component.css']
})
export class AccommodationComponent implements OnInit{

  public accommodations: Accommodation[] = [];
  
  displayedColumns: string[] = ['id', 'name', 'location', 'benefits', 'photos', 'minGuest', 'maxGuest'];
  dataSource = new MatTableDataSource(this.accommodations);

  constructor(public dialog: MatDialog, private _accommodationService: AccommodationService, private snackBar: MatSnackBar) {
  }

  ngOnInit(): void {
    this.getAccommodations();
  }

  public getAccommodations() {
    this._accommodationService.getAccommodations().subscribe(res => {
      this.accommodations = res.accommodations;
      console.log(this.accommodations);
      this.dataSource = new MatTableDataSource(res.accommodations);
    });
   }

   applyFilter(event: Event) {

    const filterValue = (event.target as HTMLInputElement).value;

    this.dataSource.filterPredicate = (data, filter) => {
      const row = data as Accommodation;
      return this.filterRow(row, filter);
    };

    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  applyFilter2(event: Event) {

    const filterValue = (event.target as HTMLInputElement).value;

    this.dataSource.filterPredicate = (data, filter) => {
      const row = data as Accommodation;
      return this.filterRow2(row, filter);
    };

    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  applyFilter3(event: Event) {

    const filterValue = (event.target as HTMLInputElement).value;

    this.dataSource.filterPredicate = (data, filter) => {
      const row = data as Accommodation;
      return this.filterRow3(row, filter);
    };

    this.dataSource.filter = filterValue.trim().toLowerCase();
  }




  private filterRow(row: Accommodation, filterValue: string): boolean {
    // replace 'name' with the name of the column you want to filter
    return row.name.toLowerCase().includes(filterValue.toLowerCase());
  }

  private filterRow2(row: Accommodation, filterValue: string): boolean {
    // replace 'name' with the name of the column you want to filter
    return row.location.toLowerCase().includes(filterValue.toLowerCase());
  }

  private filterRow3(row: Accommodation, filterValue: string): boolean {
    // broj gostiju
    return row.minGuest<= Number(filterValue) && row.maxGuest>= Number(filterValue);
  }

}
