import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { AccommodationsComponent } from './pages/accommodations/accommodations.component';
import { AccommodationCreateComponent } from './pages/accommodations/create/accommodation-create.component';
import { ReservationComponent } from './pages/reservations/reservation.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent, pathMatch: 'full' },
  { path: 'reservation', component: ReservationComponent, pathMatch: 'full' },
  { path: 'accommodations', component: AccommodationsComponent},
  { path: 'accommodations/create', component: AccommodationCreateComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
