import { Component, OnInit } from '@angular/core';
import { UserDTO } from './userDTO';
import { UserService } from './users.service';
import { HttpClient } from '@angular/common/http';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { DeleteDTO } from './deleteDTO';
import { User } from '../login/log-user-data.service';
import { EditDTO, EditService } from '../edit/edit.service';


@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit{

  del: DeleteDTO = {} as DeleteDTO
  editDTO: EditDTO = {
    id: '',
    username: '',
    password: '',
    name: '',
    surname: '',
    email: '',
    address: ''
  }
  userDTO: UserDTO = {
    id: '',
    name: '',
    surname: '',
    address: '',
    username: '',
    email: ''
  }
  constructor(private http : HttpClient,private editService:EditService, private userService: UserService, private snackBar: MatSnackBar, private m_Router: Router) { }

  public flag = false
  ngOnInit(): void {
    this.userService.getAll().subscribe((res: UserDTO)=> {
      this.editDTO.username = res.username
      this.editDTO.address= res.address
      this.editDTO.email = res.email
      this.editDTO.name = res.name
      this.editDTO.surname = res.username
      this.userDTO = res
    });
    console.log(this.editDTO)
    console.log("SDFSDFSDFSDF")
  }

  public delete(id: String): void{
    console.log(id)
    this.del.id = id
    this.userService.delete(this.del).subscribe()
    alert("User Deleted!")
    window.location.href = "http://localhost:4200/users"
  }

  public edit(user: UserDTO): void{
    console.log(user)
    this.flag = true
  }

  public onSubmit(editDTO: EditDTO){
    this.editService.edit(editDTO).subscribe()
    alert("User Edited")
    window.location.href = './users'
  }


}
