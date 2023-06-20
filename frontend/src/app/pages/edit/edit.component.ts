import { Component, OnInit } from '@angular/core';
import { EditDTO, EditService } from './edit.service';
import { HttpClient } from '@angular/common/http';
import { UserDTO } from '../users/userDTO';
import { UserService } from '../users/users.service';
import { AuthService } from '../login/log-auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-edit',
  templateUrl: './edit.component.html',
  styleUrls: ['./edit.component.css']
})
export class EditComponent implements OnInit{
  editDTO!: EditDTO;
  password: string = '';
  userId: string = '';

  constructor(private http: HttpClient, private editService: EditService, 
    private authService: AuthService, private userService: UserService,
    private router: Router){}

  ngOnInit(): void {
    this.userId = this.authService.getUserId();
    this.userService.getById(this.userId).subscribe((res: EditDTO) => {
      this.editDTO = res;
    });
  }

  public onSubmit(){
    this.editDTO.id = this.userId;
    this.authService.logout();
    this.editService.edit(this.editDTO).subscribe()
    alert("User Edited")
    this.router.navigate(['/login']);
  }

}
