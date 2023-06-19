import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";

export interface EditDTO{
  id: string;
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

export class EditService{

  constructor(private http: HttpClient){}

  edit(editDTO: EditDTO): Observable<String>{
    return this.http.post<String>('http://localhost:8000/user/edit',editDTO)
  }





}



