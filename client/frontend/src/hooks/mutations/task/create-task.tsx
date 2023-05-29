import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { Task } from "types";

interface CreateTaskRequest {
  task:Omit<Omit<Task, 'creatorId'>, 'taskId'>;
}

interface CreateTaskResponse {}
export const useCreateTask = () => {
  const { toast } = useToast();

  return useMutation<CreateTaskResponse, any, CreateTaskRequest, any>(
    ["task", "create"],
    async (data: CreateTaskRequest) => {
      return axios

        .post(`${getServerUrl()}/task/${data.task.agentId}`, data.task)
        .then((r) => r.data);
    },
    {
      onSuccess: () => {
        toast({
          title: "Task created.",
          description: "We've successfuly created the task.",
          variant: "success",
        });
      },
      onError: (error) => {
        toast({
          title: "Task creation failed.",
          description: "Unable to create task. Error: " + error,
          variant: "destructive",
        });
      },
    }
  );
};
