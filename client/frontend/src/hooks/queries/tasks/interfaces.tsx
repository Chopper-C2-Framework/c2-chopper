import { Task, TaskResult } from "types";

export type TaskListResponse = {
  tasks: Task[];
  count: number;
}

export type TaskResultListResponse = {
  results: TaskResult[];
  count: number;
}
