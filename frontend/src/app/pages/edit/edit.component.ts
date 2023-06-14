import { Component, OnInit } from '@angular/core';
import { EditDTO, EditService } from './edit.service';
import { HttpClient } from '@angular/common/http';
import { UserDTO } from '../users/userDTO';

@Component({
  selector: 'app-edit',
  templateUrl: './edit.component.html',
  styleUrls: ['./edit.component.css']
})
export class EditComponent implements OnInit{

  constructor(private http: HttpClient, private editService: EditService){}
  userDTO!: UserDTO;
  ngOnInit(): void {
  }

  public onSubmit(editDTO: EditDTO){
    this.editService.edit(editDTO).subscribe()
    alert("User Edited")
    window.location.href = './users'
  }

}
