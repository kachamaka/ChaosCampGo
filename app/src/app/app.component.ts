import { Component, ViewChild } from '@angular/core';
import { MatSidenav } from '@angular/material/sidenav';
import { Router } from '@angular/router';
import { HttpService } from './shared/http.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  @ViewChild('sidenav') sidenav!: MatSidenav;

  reason = '';

  constructor(public httpService: HttpService, private router: Router) {

  }

  close(reason: string) {
    this.reason = reason;
    if(this.reason == "logout") {
      this.httpService.logout();
      this.router.navigate(['/login']);
    }
    this.sidenav.close();
  }
  // shouldRun = [/(^|\.)plnkr\.co$/, /(^|\.)stackblitz\.io$/].some(h => h.test(window.location.host));
}
