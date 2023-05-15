import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import { UserDTO } from "./userDTO";
import { DeleteDTO } from "./deleteDTO";

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
export class UserService {
  constructor( private http: HttpClient){}

  getAll(): Observable<UserDTO[]> {
    return this.http.get<UserDTO[]>('http://localhost:8000/user/getAll');
  }

  delete( id: DeleteDTO): Observable<any>{
    console.log(id)
    return this.http.post<Observable<any>>('http://localhost:8000/user', id);
  }

}
