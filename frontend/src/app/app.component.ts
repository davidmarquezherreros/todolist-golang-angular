import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'ToDo List App';

  links = [
    {
      name: "Home",
      url: "/home"
    },
    {
      name: "Create a new task",
      url: "/new"
    }
  ]

}
