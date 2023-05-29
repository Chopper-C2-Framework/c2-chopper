import { Skeleton } from "@components/ui/skeleton";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@components/ui/card";
import { TasksDisplay } from "@components/tasks-display";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@components/ui/tabs"

import { useAllActiveTasks } from "@hooks/queries/tasks/all-active-tasks";
import { useAllNewlyExecutedTasks } from "@hooks/queries/tasks/all-new-exec-tasks";
import { useAllTasks } from "@hooks/queries/tasks/all-tasks";

export default function TasksTabDisplay() {
  const activeTasks = useAllActiveTasks()
  const newlyExecTasks = useAllNewlyExecutedTasks()
  const allTasks = useAllTasks()

  const reloadActiveTasks = () => {
    activeTasks.refetch()
    allTasks.refetch()
  }

  const reloadNewlyExecTasks = () => {
    newlyExecTasks.refetch()
    allTasks.refetch()
  }

  const reloadAllTasks = () => {
    allTasks.refetch()
  }

  return (
    <Tabs defaultValue="active">
      <TabsList className="flex w-full lg:w-1/2">
        <TabsTrigger value="active" className="w-1/3">Active</TabsTrigger>
        <TabsTrigger value="done" className="w-1/3">Done</TabsTrigger>
        <TabsTrigger value="all" className="w-1/3">All</TabsTrigger>
      </TabsList>
      <TabsContent value="active">
        <Card>
          <CardHeader>
            <CardTitle>Active Tasks</CardTitle>
            <CardDescription className="pl-3">
              { 
                activeTasks.isLoading || activeTasks.data == null ? 
                  <Skeleton className="h-4 w-[250px]" /> : 
                  `You have ${activeTasks.data.count} active tasks.` 
              }
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-2">
          {
              (!activeTasks.isLoading && activeTasks.data != null) ?
                <TasksDisplay onRefresh={reloadActiveTasks} isLoading={false} tasks={activeTasks.data.tasks} /> :
                <TasksDisplay isLoading={true} tasks={[]} />
            }
          </CardContent>
        </Card>
      </TabsContent>
      <TabsContent value="done">
        <Card>
          <CardHeader>
            <CardTitle>All Tasks</CardTitle>
            <CardDescription className="pl-5">
              { 
                newlyExecTasks.isLoading || newlyExecTasks.data == null ? 
                  <Skeleton className="h-4 w-[250px]" /> : 
                  `You have ${newlyExecTasks.data.count} total of tasks.` 
              }
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-2">
            {
              (!newlyExecTasks.isLoading && newlyExecTasks.data != null) ?
                <TasksDisplay onRefresh={reloadNewlyExecTasks} isLoading={false} tasks={newlyExecTasks.data.tasks} /> :
                <TasksDisplay isLoading={true} tasks={[]} />
            }
          </CardContent>
        </Card>
      </TabsContent>
      <TabsContent value="all">
        <Card>
          <CardHeader>
            <CardTitle>All Tasks</CardTitle>
            <CardDescription className="pl-5">
              { allTasks.isLoading || allTasks.data == null ? <Skeleton className="h-4 w-[250px]" /> : `You have a total of ${allTasks.data.count} tasks` }
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-2">
            {
              (!allTasks.isLoading && allTasks.data != null) ?
                <TasksDisplay onRefresh={reloadAllTasks} isLoading={false} tasks={allTasks.data.tasks} /> :
                <TasksDisplay isLoading={true} tasks={[]} />
            }
          </CardContent>
        </Card>
      </TabsContent>
    </Tabs>
  )
}