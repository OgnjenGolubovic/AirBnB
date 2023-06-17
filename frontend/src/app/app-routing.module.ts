import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { RegistrationComponent } from './pages/registration/registration.component';
import { UsersComponent } from './pages/users/users.component';
import { EditComponent } from './pages/edit/edit.component';
import { AccommodationsComponent } from './pages/accommodations/accommodations.component';
import { AccommodationCreateComponent } from './pages/accommodations/create/accommodation-create.component';
import { ReservationComponent } from './pages/reservations/reservation.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent, pathMatch: 'full' },
<<<<<<< HEAD
  { path: 'register', component: RegistrationComponent, pathMatch:'full'},
  {path: 'users', component: UsersComponent, pathMatch:'full'},
  {path: 'edit', component: EditComponent, pathMatch:'full'},
  { path: 'accommodation', component: AccommodationComponent},
=======
  { path: 'reservations', component: ReservationComponent, pathMatch: 'full' },
>>>>>>> b967a30dbdf9054c63c82d19bc9b41c1a7d7f95b
  { path: 'accommodations', component: AccommodationsComponent},
  { path: 'accommodations/create', component: AccommodationCreateComponent }
]
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
