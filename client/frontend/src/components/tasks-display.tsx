import { Task } from "types";
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
import { useDeleteTask } from "@hooks/mutations/task/delete-task";

interface TasksDisplayProps {
  tasks: Task[];
  isLoading: boolean;
  onRefresh?: () => void;
}

export function TasksDisplay({tasks, isLoading, onRefresh}: TasksDisplayProps) {
  const deleteTask = useDeleteTask()

  const onDeleteTaskClick = async (task: Task) => {
    if(deleteTask.isLoading) return;
    console.log(task)
    await deleteTask.mutateAsync({task_id: task.taskId})
    if(onRefresh != null) onRefresh()
  }

  return (
    <div className="container gap-5 px-8 py-5 flex flex-wrap align-center">
      {
        isLoading && (
          [1,2,3].map(() => {
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
          })
        )
      }
      {
        !isLoading && tasks.map((task) => {
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
                    <PopoverContent className="w-80">
                      <Button onClick={() => onDeleteTaskClick(task)}>Delete</Button>
                    </PopoverContent>
                  </Popover>
                </CardTitle>
                
              </CardHeader>
              {/* <CardContent className="grid gap-4">
                <p className="text-sm font-medium leading-none">{task.}</p>
              </CardContent> */}
              <CardFooter className="">
                <Badge>{task.type == 1 ? "Shell" : "Ping"}</Badge>
              </CardFooter>
            </Card>
          )
        })
      }
    </div>
  );
}
