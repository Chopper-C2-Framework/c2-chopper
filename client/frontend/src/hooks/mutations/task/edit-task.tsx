import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import { Task } from "@src/types";
import axios from "axios";
import { useMutation } from "react-query";

interface EditTaskRequest {
  task: Omit<Task, "creatorId">;
}

interface EditTaskResponse {}

export const useEditTask = () => {
  const { toast } = useToast();

  return useMutation<EditTaskResponse, any, EditTaskRequest, any>(
    ["task", "edit"],
    async (data: EditTaskRequest) => {
      return axios
        .patch(`${getServerUrl()}/task/${data.task.taskId}`, data.task)
        .then((r) => r.data);
    },
    {
      onSuccess: () => {
        toast({
          title: "Task edited.",
          description: "We've successfuly edited the task.",
          variant: "success",
        });
      },
      onError: (error) => {
        toast({
          title: "Task editing failed.",
          description: "Unable to edit task. Error: " + error,
          variant: "destructive",
        });
      },
    }
  );
};
