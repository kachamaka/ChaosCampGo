import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class HttpService {
  
  domain = 'http://localhost:8888/';
  isLoggedIn = false;
  authToken: any;
  options!: HttpHeaders;

  constructor(private http: HttpClient) { }

  loadToken() {
    this.authToken = localStorage.getItem('token');
  }  

  createAuthenticationHeaders() {
    this.loadToken();
    this.options = new HttpHeaders({
      'Content-Type': 'application/json',
      'Authorization': this.authToken
    });
  }

  register(user: any) {
    // return this.http.post(this.domain + '/register', user).subscribe
    return this.http.post(this.domain + 'register', user).subscribe((res => {
      console.log(res);
      return res;
    }));
  }
}
