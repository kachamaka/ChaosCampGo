import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { HttpService } from '../shared/http.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {

  username: string = "";
  password: string = "";

  constructor(
    private httpService: HttpService, 
    private snackBar: MatSnackBar,
    private router: Router) {

  }
  
  onLogin() {
    this.router.navigate(['login']);
  }

  onRegister() {
    let user = {
      "username": this.username,
      "password": this.password
    };
    this.httpService.register(user)
    .subscribe((res => {
      if(res["success"] == true) {
        this.httpService.auth(res["token"])
        this.router.navigate(["/home"]);
        this.msg("Registered!");
      }
      else {  
        console.log(res["message"]);
        this.msg(res["message"]);
      }
    }))
  }
  
  msg(message: string) {
      this.snackBar.open(message, '', {
        horizontalPosition: 'end',
        verticalPosition: 'top',
        duration: 2000,
      });
  }
}
