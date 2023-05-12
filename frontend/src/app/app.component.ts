import { Component } from '@angular/core';
import { AccommodationService } from './pages/accommodation/services/accommodation.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [AccommodationService]
})
export class AppComponent {
  title = 'AirBnB';
}
