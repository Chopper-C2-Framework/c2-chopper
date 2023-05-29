import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@components/ui/dialog"
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@components/ui/select"
import { Button } from "@components/ui/button";
import { Input } from "@components/ui/input";
import { Label } from "@components/ui/label";
import { useState } from "react";
import { Task, TaskType } from "@src/types.ts";
import { useCreateTask } from "@hooks/mutations/task/create-task";
import { useEditTask } from "@hooks/mutations/task/edit-task";

interface ICreateTaskDialog{
  taskEdit?: Task;
  onAction: () => void;
}

export default function CreateTaskDialog({taskEdit, onAction}: ICreateTaskDialog) {
  const [arg, setArg] = useState(taskEdit?.args[0] ?? "")
  const [type, setType] = useState<TaskType>(taskEdit?.type ?? TaskType.UNKNOWN)
  const [name, setName] = useState(taskEdit?.name ?? "")
  const [agentId, setAgentId] = useState(taskEdit?.agentId ?? "")
  const [dialogOpen, setDialogOpen] = useState(false)

  const createTaskHook = useCreateTask()
  const editTaskHook = useEditTask()

  const createNewTask = async () => {
    await createTaskHook.mutateAsync({
      task: {
        agentId,
        name,
        type,
        args: [arg],
      }
    })
    setArg("")
    setType(TaskType.UNKNOWN)
    setName("")
    setAgentId("")
    setDialogOpen(false)
    onAction()
  }

  const editTask = async () => {
    await editTaskHook.mutateAsync({
      task: {
        taskId: taskEdit!.taskId,
        agentId,
        name,
        type,
        args: [arg],
      }
    })
    setArg("")
    setType(TaskType.UNKNOWN)
    setName("")
    setAgentId("")
    setDialogOpen(false)
    onAction()
  }

  return (
    <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
      <DialogTrigger asChild>
        <Button size="sm">
          {
            taskEdit == null ? "New task" : "Edit task"
          }
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>{taskEdit == null ? "Create a new" : "Edit a"} task</DialogTitle>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="name" className="text-right">
              Name
            </Label>
            <Input value={name} onChange={(e) => setName(e.target.value)} id="name" placeholder="Task name" className="col-span-3" />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="agentId" className="text-right">
              Agent Id
            </Label>
            <Input value={agentId} onChange={(e) => setAgentId(e.target.value)} id="agentId" placeholder="Agent Id" className="col-span-3" />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label className="text-right">
              Task Type
            </Label>
            <Select value={`${type}`} onValueChange={(value) => setType(value as TaskType) }>
              <SelectTrigger className="w-[180px]">
                <SelectValue placeholder="Select a task type" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectLabel>Type</SelectLabel>
                  <SelectItem value={TaskType.PING}>Ping</SelectItem>
                  <SelectItem value={TaskType.SHELL}>Shell</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="arg" className="text-right">
              Argument
            </Label>
            <Input id="arg" placeholder="Task argument" value={arg} onChange={(e) => setArg(e.target.value)} className="col-span-3" />
          </div>
        </div>
        <DialogFooter>
          {
            taskEdit == null && <Button type="submit" onClick={createNewTask}>Create</Button>
          }
          {
            taskEdit != null && <Button type="submit" onClick={editTask}>Edit</Button>
          }
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}