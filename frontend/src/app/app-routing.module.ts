import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { AccommodationsComponent } from './pages/accommodations/accommodations.component';
import { AccommodationCreateComponent } from './pages/accommodations/create/accommodation-create.component';
import { AccommodationComponent } from './pages/accommodation/accommodation.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent, pathMatch: 'full' },
  { path: 'accommodation', component: AccommodationComponent},
  { path: 'accommodations', component: AccommodationsComponent},
  { path: 'accommodations/create', component: AccommodationCreateComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
