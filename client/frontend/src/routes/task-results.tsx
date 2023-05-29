import React from "react";
import { useParams } from "react-router-dom";
import { useTaskResults } from "@hooks/queries/tasks/one-task-result";
import TaskResultsDisplay from "@components/task-results-display";
import { useTask } from "@hooks/queries/tasks/one-task";

interface TaskResultsProps {}

export const TaskResults: React.FC<TaskResultsProps> = ({}) => {
  const { taskId } = useParams();
  if (taskId == undefined) {
    window.location.href = "/app/tasks";
    return <></>;
  }
  const taskResults = useTaskResults(taskId);
  const task = useTask(taskId);

  if (taskResults.isError || task.isError) {
    window.location.href = "/app/tasks";
    return <></>;
  }

  return (
    <div className="px-10">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <div className="flex items-center justify-between space-y-2">
          <h2 className="text-3xl font-bold tracking-tight">Tasks</h2>
        </div>
      </div>
      <div className="px-16">
        {taskResults.data != null && !taskResults.isLoading && task.data && (
          <TaskResultsDisplay
            task={task.data}
            results={taskResults.data.results}
          />
        )}
      </div>
    </div>
  );
};

export default TaskResults;
