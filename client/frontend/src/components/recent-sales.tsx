import { useAllTasks } from "@hooks/queries/tasks/all-tasks";

export function TasksTable() {
  const { data: allTasksData, isLoading: allTasksLoading } = useAllTasks();
  return (
    <div className="space-y-4 px-[50px]">
      <div className="flex items-center w-full space-x-[200px] h-10 bg-gray-800/40 rounded-md px-12 justify-between">
        <p className="font-bold text-lg text-primary">Task ID</p>
        <p className="font-bold text-lg text-primary">Task's name</p>
        <p className="font-bold text-lg text-primary">Task' type</p>
        <p className="font-bold text-lg text-primary">Task's arguments</p>
      </div>
      {!allTasksLoading && allTasksData !== undefined && allTasksData.tasks ? (
        allTasksData?.tasks.map((task) => {
          return (
            <div
              key={task.taskId}
              className="flex items-center justify-between w-full h-10 bg-gray-800/40 rounded-md space-x-[400px] px-12 space-x-[200px] hover:opacity-80 cursor-pointer "
            >
              <p>{task.taskId.slice(0, 8) + "..."}</p>
              <p>{task.name}</p>
              <p>{task.type}</p>
              <p>
                {task.args.map((arg) => (
                  <span>{arg}</span>
                ))}
              </p>
            </div>
          );
        })
      ) : (
        <p>No data now</p>
      )}
    </div>
  );
}
