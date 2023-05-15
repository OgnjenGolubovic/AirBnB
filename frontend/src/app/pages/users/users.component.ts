import { Component, OnInit } from '@angular/core';
import { UserDTO } from './userDTO';
import { UserService } from './users.service';
import { HttpClient } from '@angular/common/http';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { DeleteDTO } from './deleteDTO';
import { User } from '../login/log-user-data.service';


@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit{

  usersList: UserDTO[] = [];
  del: DeleteDTO = {} as DeleteDTO

  constructor(private http : HttpClient, private userService: UserService, private snackBar: MatSnackBar, private m_Router: Router) { }


  ngOnInit(): void {
    this.userService.getAll().subscribe((res: any)=> {
      res.users.forEach((user: UserDTO) => {
        this.usersList.push(user)
      })
    });
    console.log(this.usersList)
    console.log("SDFSDFSDFSDF")
  }

  public delete(id: String): void{
    console.log(id)
    this.del.id = id
    this.userService.delete(this.del).subscribe()
    alert("User Deleted!")
    window.location.href = "http://localhost:4200/users"
  }

  public edit(): void{
    window.location.href = "http://localhost:4200/edit"
  }
}
