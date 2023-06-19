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
import { RegistrationComponent } from './registration/registration.component';
import { UsersComponent } from './users/users.component';
import { EditComponent } from './edit/edit.component';
import { AccommodationsComponent } from './accommodations/accommodations.component';
import { AccommodationCreateComponent } from './accommodations/create/accommodation-create.component';
import { ReservationComponent } from './reservations/reservation.component';
import { GuestReservationComponent } from './guest-reservation/guest-reservation.component';
import { MatSelectModule } from '@angular/material/select';
import {MatMenuModule} from '@angular/material/menu';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatRadioModule} from '@angular/material/radio';
import { DefineDatesComponent } from './accommodations/define-dates/define-dates.component';
import {MatListModule} from '@angular/material/list';
import { ChangePriceComponent } from './accommodations/change-price/change-price.component';

@NgModule({
  declarations: [
    LoginComponent,
    RegistrationComponent,
    UsersComponent,
    EditComponent,
    AccommodationsComponent,
    AccommodationCreateComponent,
    ReservationComponent,
    GuestReservationComponent,
    DefineDatesComponent,
    ChangePriceComponent
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
    MatSelectModule,
    MatMenuModule,
    MatCheckboxModule,
    MatRadioModule,
    MatListModule
  ]
})
export class PagesModule { }
