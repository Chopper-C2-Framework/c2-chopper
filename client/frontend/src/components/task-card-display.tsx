import { Task } from "@src/types";
import { 
  Card, 
  CardDescription, 
  CardFooter, 
  CardHeader, 
  CardTitle 
} from "./ui/card";
import { Badge } from "./ui/badge";
import { Settings2 } from "lucide-react"
import { Button } from "@components/ui/button"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@components/ui/popover"
import { Skeleton } from "./ui/skeleton";
import CreateTaskDialog from "./create-task-dialog";
import { useDeleteTask } from "@hooks/mutations/task/delete-task";


interface ITaskCardDisplay{
  loading?: boolean;
  onRefresh?: () => void;
  task?: Task;
}

export default function TaskCardDisplay({loading, task, onRefresh}: ITaskCardDisplay){
  const deleteTask = useDeleteTask()

  const onDeleteTaskClick = async (task: Task) => {
    if(deleteTask.isLoading) return;
    console.log(task)
    await deleteTask.mutateAsync({task_id: task.taskId})
    if(onRefresh != null) onRefresh()
  }

  const navigateToResultsPage = (task: Task) => {
    window.location.href = `/app/tasks/${task.taskId}/results`
  }

  if(loading || task == null) {
    return (
      <Card className="min-w-[300px]">
        <CardHeader>
          <CardTitle className="flex justify-between items-center">
            <div>
              <Skeleton className="h-4 w-[150px]" />
              <CardDescription className="pl-2 pt-1 font-normal">
                <Skeleton className="h-4 w-[70px]" />
              </CardDescription>
            </div>
          </CardTitle>
        </CardHeader>
        <CardFooter className="">
          <Skeleton className="h-4 w-[70px]" />
        </CardFooter>
      </Card>
    )
  }
  return (
    <Card className="min-w-[300px] w-[31.5%]">
      <CardHeader>
        <CardTitle className="flex justify-between items-center">
          <div>
            {task.name}
            <CardDescription className="pl-2 font-normal">{task.agentId}</CardDescription>
          </div>
          <Popover>
            <PopoverTrigger asChild>
              <Button variant="outline" className="w-10 rounded-full p-0">
                <Settings2 className="h-4 w-4" />
                <span className="sr-only">Open popover</span>
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-80 flex flex-col gap-3">
              <Button onClick={() => onDeleteTaskClick(task)}>Delete</Button>
              <CreateTaskDialog onAction={()=> onRefresh && onRefresh()} taskEdit={task} />
              <Button onClick={() => navigateToResultsPage(task)}>View results</Button>
            </PopoverContent>
          </Popover>
        </CardTitle>
      </CardHeader>
      <CardFooter className="">
        <Badge>{task.type == "PING" ? "Ping" : (task.type == "SHELL" ? "Shell" : "Unknown")}</Badge>
      </CardFooter>
    </Card>
  )
}