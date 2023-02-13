import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HttpService {
  
  domain = 'http://localhost:8888/';
  isLoggedIn = false;
  authToken = "";
  options!: HttpHeaders;

  constructor(private http: HttpClient) {
    this.loadToken();
    if(this.authToken != "") {
      this.isLoggedIn = true;
    }
  }

  loadToken() {
    let token = localStorage.getItem("token");
    this.authToken = (token == null ? "" : token);
  }  

  auth(token: string) {
    localStorage.setItem("token", token);
    this.isLoggedIn = true;
  }

  logout() {
    this.authToken = "";
    localStorage.removeItem("token");
    this.isLoggedIn = false;
  }

  headers() {
    this.loadToken();
    this.options = new HttpHeaders({
      "Content-Type": "application/json",
      "Authorization": this.authToken
    });

    return {headers: this.options};
  }

  register(user: any) {
    // return this.http.post(this.domain + '/register', user).subscribe
    return this.http.post(this.domain + "register", user)
    .pipe(map(response => JSON.parse(JSON.stringify(response))));
  }

  login(user: any) {
    return this.http.post(this.domain + "login", user)
    .pipe(map(response => JSON.parse(JSON.stringify(response))));
  }

  getEvents() {
    return this.http.get(this.domain + "getEvents", this.headers())
    .pipe(map(response => JSON.parse(JSON.stringify(response))));
  }

  addEvent(event: any) {
    return this.http.post(this.domain + "addEvent", event, this.headers())
    .pipe(map(response => JSON.parse(JSON.stringify(response))));
  }

  removeEvent(event: any) {
    return this.http.post(this.domain + "deleteEvent", event, this.headers())
    .pipe(map(response => JSON.parse(JSON.stringify(response))));
  }

  addReminder(reminder: any) {
    return this.http.post(this.domain + "addReminder", reminder, this.headers())
    .pipe(map(response => JSON.parse(JSON.stringify(response))));
  }
}
