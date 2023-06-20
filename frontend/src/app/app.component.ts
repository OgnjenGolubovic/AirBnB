import { Component } from '@angular/core';
import { AccommodationsService } from './pages/accommodations/accommodations.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [AccommodationsService]
})
export class AppComponent {
  title = 'AirBnB';
}
