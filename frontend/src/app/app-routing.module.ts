import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { RegistrationComponent } from './pages/registration/registration.component';
import { UsersComponent } from './pages/users/users.component';
import { EditComponent } from './pages/edit/edit.component';
import { AccommodationsComponent } from './pages/accommodations/accommodations.component';
import { AccommodationCreateComponent } from './pages/accommodations/create/accommodation-create.component';
import { ReservationComponent } from './pages/reservations/reservation.component';
import { GuestReservationComponent } from './pages/guest-reservation/guest-reservation.component';
import { AuthGuard } from './pages/login/log-auth.guard';
import { DefineDatesComponent } from './pages/accommodations/define-dates/define-dates.component';
import { ChangePriceComponent } from './pages/accommodations/change-price/change-price.component';
import { GuestReservationsComponent } from './pages/guest-reservations/guest-reservations.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent, pathMatch: 'full' },
  { path: 'register', component: RegistrationComponent, pathMatch:'full'},
  {path: 'users', component: UsersComponent, pathMatch:'full'},
  {path: 'edit', component: EditComponent, pathMatch:'full'},
  { path: 'accommodation', component: AccommodationsComponent},
  { path: 'reservations', component: ReservationComponent, pathMatch: 'full' },
  { path: 'accommodations', component: AccommodationsComponent},
  { path: 'accommodations/create', component: AccommodationCreateComponent },
  { path: 'guest-reservation', component: GuestReservationComponent, pathMatch: 'full' },
  { path: 'guest-reservations', component: GuestReservationsComponent, pathMatch: 'full' },
  { path: 'reservations', component: ReservationComponent, pathMatch: 'full'},
  { path: 'accommodations', component: AccommodationsComponent},
  { path: 'accommodations/create', component: AccommodationCreateComponent},
  { path: 'accommodations/define-dates/:id', component: DefineDatesComponent},
  { path: 'accommodations/change-price/:id', component: ChangePriceComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
