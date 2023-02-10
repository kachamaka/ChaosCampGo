import { Component, OnInit } from '@angular/core';
import { HttpService } from '../shared/http.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit{

  constructor(private httpService: HttpService) {
    let user = {
      "username":"testuser1",
      "password":"testpass",
      "email": "test@test.test"
  }
    this.httpService.register(user);
  }

  ngOnInit(): void {

  }

}
