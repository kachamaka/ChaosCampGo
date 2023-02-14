import { Component, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSelect } from '@angular/material/select';
import { MatSnackBar } from '@angular/material/snack-bar';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';
import { HttpService } from '../shared/http.service';

@Component({
  selector: 'app-schedule',
  templateUrl: './schedule.component.html',
  styleUrls: ['./schedule.component.css']
})
export class ScheduleComponent {
  
  days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"];
  currentDay = new Date().getUTCDay();
  data: Array<Array<Event>> = [[], [], [], [], [], [], []];
  events: Array<Event> = [];
  reminderOptions = ["15m", "30m", "1h", "1h30m", "2h", "2h30m", "3h"];

  displayedColumns: string[] = ["subject", "start", "end"];
  dataSource: MatTableDataSource<Event> = new MatTableDataSource<Event>;

  @ViewChild('matSelect') day!: MatSelect;
  eventDay!: number;
  eventSubject!: string;
  eventStart!: string;
  eventEnd!: string;
  
  reminderEvent!: Event;
  reminderAhead!: number;

  selectedEvent: Event = { day: 1, subject: "", start: "", end: ""};

  constructor(private httpService: HttpService, private snackBar: MatSnackBar, private router: Router) {
    this.previous();
    console.log(this.currentDay, this.days[this.currentDay], "DAY!");
    this.getEvents();
  }

  nullify() {
    this.data = [[], [], [], [], [], [], []];
    this.events = [];
  }

  getEvents() {
    this.httpService.getEvents().subscribe(res => {
      console.log(res);
      if(res["success"] == true) {
        this.parseEvents(res["events"]);
        this.updateSource();
      }
      else {
        this.validateToken(res);
        this.msg(res["message"]);
      }
    })
  }

  parseEvents(events: any) {
    this.nullify();
    events.forEach((event:Event) => {
      this.data[event.day].push(event);
      this.events.push(event);
    });
  }

  updateSource() {
    this.dataSource = new MatTableDataSource(this.data[this.currentDay]);
  }
  
  next() {
    this.currentDay = (this.currentDay + 1) % 7;
    this.updateSource();
  }

  previous() {
    this.currentDay = (this.currentDay == 0) ? 6 : ((this.currentDay - 1) % 7);
    this.updateSource();
  }

  addEvent() {
    if(this.eventSubject == undefined || this.eventStart == undefined || this.eventEnd == undefined || this.eventDay == undefined) {
      this.msg("please fill every field");
      return;
    }
    let event: Event = {
      subject: this.eventSubject,
      start: this.eventStart,
      end: this.eventEnd,
      day: this.eventDay
    };
    console.log(event);
    this.httpService.addEvent(event).subscribe(res => {
      if(res["success"]) {
        this.getEvents();
        this.msg(res["message"]);
      }
      else {
        this.msg(res["message"]);
      }
    });
  }

  removeEvent() {
    if(this.selectedEvent) {
      this.httpService.removeEvent(this.selectedEvent).subscribe(res => {
        if(res["success"]) {
          this.selectedEvent = <Event>{};
          this.getEvents();
          this.msg(res["message"]);
        } else {
          this.msg(res["message"]);
        }
      });
    }
    else {
      this.msg("no event selected");
    }
  }

  addReminder() {
    if(this.reminderEvent == undefined || this.reminderAhead == undefined) {
      this.msg("please fill every field");
      return;
    }
    let eventTime = this.getTime(this.reminderEvent.start, this.reminderEvent.day);
    let timeAhead = this.parseTimeAhead();
    let reminder: Reminder = {
      subject: this.reminderEvent.subject,
      eventStart: eventTime,
      timeAhead: timeAhead
    };
    console.log(reminder);
    this.httpService.addReminder(reminder).subscribe(res => {
      if(res["success"]) {
        // this.getEvents();
        this.msg(res["message"]);
      }
      else {
        this.msg(res["message"]);
      }
    });
  }
  parseTimeAhead() {
    console.log(this.reminderAhead);
    if(this.reminderAhead == 0) {
      return 15 * 60;
    }
    return (this.reminderAhead) * 30 * 60;
  }

  selectEvent(event: Event) {
    this.selectedEvent = event;
    console.log(event);
    // let startTime = this.getTime(event.start, event.day);
    // let timeAhead = 3600;
    // console.log(startTime - timeAhead);
    
  }

  getTime(start: string, day: number): number {
    let startTime = start.split(":");
    let startHour = parseInt(startTime[0]);
    let startMinutes = parseInt(startTime[1]);

    let d = new Date();
    let dateDay = d.getDay();
    dateDay = (dateDay == 0) ? 6 : ((dateDay - 1) % 7);
    if(day < dateDay) day += 7;
    let addDays = day - dateDay;
    d.setHours(startHour);
    d.setMinutes(startMinutes);
    d.setSeconds(0);
    d.setDate(d.getDate() + addDays);
    // console.log(d, startMinutes);
    let time = Math.floor(d.getTime()/1000);
    // console.log(d, time);

    
    return time;
  }
  
  msg(message: string) {
      this.snackBar.open(message, '', {
        horizontalPosition: 'end',
        verticalPosition: 'top',
        duration: 2000,
      });
  }

  validateToken(res: any) {
    if(res["status"] == 401 || res["status"] == 3) {
      this.httpService.logout();
      this.router.navigate(['/login']);
    } else {
      this.httpService.isLoggedIn = true
    }
  }

}

export interface Event {
  subject: string;
  day: number;
  start: string;
  end: string;
}
export interface Reminder {
  subject: string;
  eventStart: number;
  timeAhead: number;
}

