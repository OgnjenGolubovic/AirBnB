import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AppRoutingModule } from 'src/app/app-routing.module';
import { MatTableModule } from '@angular/material/table';
import { MatCardModule } from '@angular/material/card';
import { FormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { ReactiveFormsModule } from '@angular/forms';
import { LoginComponent } from './login/login.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { MatDialogModule } from '@angular/material/dialog';
import { AccommodationsComponent } from './accommodations/accommodations.component';
import { AccommodationCreateComponent } from './accommodations/create/accommodation-create.component';
import { AccommodationComponent } from './accommodation/accommodation.component';
import { ReservationComponent } from './reservations/reservation.component';
import {MatMenuModule} from '@angular/material/menu';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatRadioModule} from '@angular/material/radio';
import { DefineDatesComponent } from './accommodations/define-dates/define-dates.component';

@NgModule({
  declarations: [
    LoginComponent,
    AccommodationsComponent,
    AccommodationComponent,
    AccommodationCreateComponent,
    ReservationComponent,
    DefineDatesComponent
  ],
  imports: [
    FormsModule,
    MatDialogModule,
    CommonModule,
    MatCardModule,
    MatInputModule,
    AppRoutingModule,
    FormsModule,
    MatSnackBarModule,
    MatFormFieldModule,
    MatTableModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatIconModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatMenuModule,
    MatCheckboxModule,
    MatRadioModule
  ]
})
export class PagesModule { }
