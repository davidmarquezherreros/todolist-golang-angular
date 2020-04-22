import { Component, OnInit } from '@angular/core';
import { TaskService } from '../task.service';
import { Task } from '../task.model';

@Component({
  selector: 'app-task-list',
  templateUrl: './task-list.component.html',
  styleUrls: ['./task-list.component.css']
})


export class TaskListComponent implements OnInit {

  private buttonDisplayTexts = [{"text": "Show cards"}, {"text": "Show Table"}];

  public cardRender: boolean = true;
  public renderButtonDisplay: string;

  public tasks : Task;

  constructor(private taskService: TaskService) { }

  // Initializes the component
  ngOnInit(): void {
    this.loadAllTasks();
    this.renderButtonDisplay = this.buttonDisplayTexts[+this.cardRender].text;
  }

  // Function that retrieves all tasks
  loadAllTasks(): void {
    this.taskService.getAllTasks().subscribe(
      (data) => {
        this.tasks = data;
      }, 
      (error) => console.log(error)
    )
  }

  /// Functions that toggles between card and table view
  changeRenderClicked(): void {
    this.cardRender = !this.cardRender;
    this.renderButtonDisplay = this.buttonDisplayTexts[+this.cardRender].text;
  }

}
