import { Task } from "types";

export type TaskListResponse = {
  tasks: Task[];
  count: number;
}