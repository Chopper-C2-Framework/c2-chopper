import { Task, TaskResult } from "@src/types"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@components/ui/card"
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@components/ui/table"
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "./ui/dialog";
import { useSetTasksSeen } from "@hooks/mutations/task/set-task-result-seen";
import { useAgentInfo } from "@hooks/queries/agents/one-agent";
import { Skeleton } from "./ui/skeleton";

interface ITaskResultsDisplay{
  results: TaskResult[];
  task: Task;
}

export default function TaskResultsDisplay({task, results}: ITaskResultsDisplay){
  const setTaskSeen = useSetTasksSeen()
  const agent = useAgentInfo(task.agentId)
  return (
    <div className="container gap-5 px-8 py-5 flex flex-col align-center">
      <Card >
        <CardHeader>
          <CardTitle className="flex items-center">
            <div>
              {task.name}
              <CardDescription className="pl-2 font-normal">Type: {task.type.toLowerCase()}</CardDescription>
            </div>
          </CardTitle>
        </CardHeader>
        <CardContent className="flex flex-row pr-20">
          <div className="w-1/2 flex flex-col gap-2">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="agentId" className="text-right">
                Agent
              </Label>
              {
                (agent.data == null || agent.isLoading) && <Skeleton className="h-4 w-[70px]" />
              }
              {
                (agent.data != null && !agent.isLoading) && <Input id="agentId" value={`${agent.data.hostname} | ${agent.data.nickname} (${agent.data.username} - ${agent.data.userId})`} readOnly className="col-span-3"/>
              }
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="args" className="text-right">
                Argument
              </Label>
              <Input id="args" value={task.args[0]} readOnly className="col-span-3"/>
            </div>
          </div>
          <div className="w-1/2 flex flex-col gap-2">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="creatorId" className="text-right">
                Created By
              </Label>
              <Input id="creatorId" value={task.creatorId} readOnly className="col-span-3"/>
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="createdAt" className="text-right">
                Created At
              </Label>
              <Input id="createdAt" value={task.createdAt} readOnly className="col-span-3"/>
            </div>
          </div>
        </CardContent>
        <CardFooter className="">
          Hello
        </CardFooter>
      </Card>

      <Table>
        <TableCaption>
          {
            results.length == 0 && "No executions for this task."
          }
        </TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[200px]">Executed At</TableHead>
            <TableHead className="w-[100px]">Status</TableHead>
            <TableHead>Output</TableHead>
            <TableHead className="w-[100px] text-center">Manage</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {results.map((result, idx) => (
            <TableRow key={idx}>
              <TableCell className="font-medium">{result.createdAt}</TableCell>
              <TableCell>{result.status}</TableCell>
              <Dialog onOpenChange={async (open) => {
                if(open && !result.seen) {
                  await setTaskSeen.mutateAsync({
                    result_ids: [result.id]
                  })
                  result.seen = true;
                }
              }}>
                <DialogTrigger asChild>
                  <TableCell className="line-clamp-1 cursor-pointer">{result.output}</TableCell>
                </DialogTrigger>
                <DialogContent className="sm:max-w-[425px]">
                  <DialogHeader>
                    <DialogTitle>Output</DialogTitle>
                  </DialogHeader>
                  <p className="text-sm font-medium leading-none">
                    {result.output}
                  </p>
                </DialogContent>
              </Dialog>
              <TableCell className="w-[100px]">
                <Button>Delete</Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  )
}