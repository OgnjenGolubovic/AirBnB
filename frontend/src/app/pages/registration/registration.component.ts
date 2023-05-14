import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RegistrationDTO, RegistrationService } from './registration.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {

  constructor(private http : HttpClient, private registerService:RegistrationService, private snackBar: MatSnackBar, private m_Router: Router) { }

  ngOnInit(): void {
  }

  onSubmit(obj: RegistrationDTO){
    var sth = this.registerService.onSubmit(obj).subscribe(response =>{
      this.snackBar.open('User Registered!','Ok', {
        duration: 3000

        })
      this.m_Router.navigate(['/login']);
    }, error=>{
      this.snackBar.open(error.error.message,'Ok', {
        duration: 3000

        })
    })
  }
}
