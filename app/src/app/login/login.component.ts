import { Component } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { HttpService } from '../shared/http.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  username: string = "";
  password: string = "";

  constructor(
    private httpService: HttpService, 
    private snackBar: MatSnackBar,
    private router: Router) {

  }
  
  onLogin() {
    let user = {
      "username": this.username,
      "password": this.password
    };
    this.httpService.login(user)
    .subscribe((res => {
      if(res["success"] == true) {
        this.httpService.auth(res["token"]);
        this.router.navigate(["/home"]);
        this.msg("Logged in!");
      }
      else {  
        console.log(res["message"]);
        this.msg(res["message"]);
      }
    }))
  }

  onRegister() {
    this.router.navigate(['register']);
  }
  
  msg(message: string) {
      this.snackBar.open(message, '', {
        horizontalPosition: 'end',
        verticalPosition: 'top',
        duration: 2000,
      });
  }
  
}

