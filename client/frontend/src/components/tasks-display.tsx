import { Task } from "types";
import TaskCardDisplay from "./task-card-display";


interface TasksDisplayProps {
  tasks: Task[];
  isLoading: boolean;
  onRefresh?: () => void;
}

export function TasksDisplay({tasks, isLoading, onRefresh}: TasksDisplayProps) {
  return (
    <div className="container gap-5 px-8 py-5 flex flex-wrap align-center">
      {
        isLoading && (
          [1,2,3].map(() => {
            return (
              <TaskCardDisplay loading={true} />
            )
          })
        )
      }
      {
        !isLoading && tasks.map((task) => {
          return (
            <TaskCardDisplay onRefresh={onRefresh} task={task} />
          )
        })
      }
    </div>
  );
}
