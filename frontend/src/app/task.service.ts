import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Task } from './task.model';

@Injectable({
  providedIn: 'root'
})
export class TaskService {

  /// Endpoint
  private readonly TASK_ENDPOINT : string = "http://localhost:9090"

  constructor(protected httpClient: HttpClient) { }

  ///
  /// Returns all available tasks
  ///
  getAllTasks() : Observable<Task>{
    return this.httpClient.get<Task>(`${this.TASK_ENDPOINT}/tasks`)
  }

}
