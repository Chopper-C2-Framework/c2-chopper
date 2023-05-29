import React from "react";
import CreateTaskDialog from "@components/create-task-dialog";
import TasksTabDisplay from "@components/tasks-tab-display";

interface TasksProps {}

export const Tasks: React.FC<TasksProps> = ({}) => {
  return (
    <div className="px-10">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <div className="flex items-center justify-between space-y-2">
          <h2 className="text-3xl font-bold tracking-tight">Tasks</h2>
          <CreateTaskDialog onAction={()=>window.location.reload()} />
        </div>
      </div>
      <div className="px-16">
        <TasksTabDisplay />
      </div>
    </div>
  )
};

export default Tasks;
