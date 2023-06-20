import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable, switchMap } from "rxjs";
import { environment } from "src/environments/environment"


export interface RegistrationDTO{
  username: string;
	password: string;
	name: string ;
  surname:  string;
  email: string;
  address: string;
}

@Injectable({
  providedIn: "root"
})
export class RegistrationService {
  constructor( private http: HttpClient){

  }

  onSubmit( obj: RegistrationDTO): Observable<any>{
    console.log(obj)
    return this.http.post(`${environment.hospitalApiUrl}/user/register`, obj)
  }

}



