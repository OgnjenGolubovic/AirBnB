import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { AccommodationsComponent } from './pages/accommodations/accommodations.component';
import { AccommodationCreateComponent } from './pages/accommodations/create/accommodation-create.component';
import { ReservationComponent } from './pages/reservations/reservation.component';
import { GuestReservationComponent } from './pages/guest-reservation/guest-reservation.component';
import { AuthGuard } from './pages/login/log-auth.guard';
import { DefineDatesComponent } from './pages/accommodations/define-dates/define-dates.component';
import { ChangePriceComponent } from './pages/accommodations/change-price/change-price.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent, pathMatch: 'full' },
  { path: 'guest-reservation', component: GuestReservationComponent, pathMatch: 'full' },
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
