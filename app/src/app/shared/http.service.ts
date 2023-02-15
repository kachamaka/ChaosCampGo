import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, map, of } from 'rxjs';

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
    return this.http.post(this.domain + "register", user)
    .pipe(
      map(response => JSON.parse(JSON.stringify(response))),
      catchError(error => {
        console.log(error);
        return of(<Response>{success: false, message: "Server error"});
    }));
  }

  login(user: any) {
    return this.http.post(this.domain + "login", user)
    .pipe(
      map(response => JSON.parse(JSON.stringify(response))),
      catchError(error => {
        console.log(error);
        return of(<Response>{success: false, message: "Server error"});
    }));
  }

  getEvents() {
    return this.http.get(this.domain + "getEvents", this.headers())
    .pipe(
      map(response => JSON.parse(JSON.stringify(response))),
      catchError(error => {
        console.log(error);
        return of(<Response>{success: false, message: "Server error"});
    }));
  }

  addEvent(event: any) {
    return this.http.post(this.domain + "addEvent", event, this.headers())
    .pipe(
      map(response => JSON.parse(JSON.stringify(response))),
      catchError(error => {
        console.log(error);
        return of(<Response>{success: false, message: "Server error"});
    }));
  }

  removeEvent(event: any) {
    return this.http.post(this.domain + "deleteEvent", event, this.headers())
    .pipe(
      map(response => JSON.parse(JSON.stringify(response))),
      catchError(error => {
        console.log(error);
        return of(<Response>{success: false, message: "Server error"});
    }));
  }

  addReminder(reminder: any) {
    return this.http.post(this.domain + "addReminder", reminder, this.headers())
    .pipe(
      map(response => JSON.parse(JSON.stringify(response))),
      catchError(error => {
        console.log(error);
        return of(<Response>{success: false, message: "Server error"});
    }));
  }
}


export interface Response {
  success: boolean
  message: string
}