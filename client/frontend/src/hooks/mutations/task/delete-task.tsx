import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";

interface DeleteTaskRequest {
  task_id: string;
}

interface DeleteTaskResponse {}

export const useDeleteTask = () => {
  const { toast } = useToast();

  return useMutation<DeleteTaskResponse, any, DeleteTaskRequest, any>(
    ["task", "delete"],
    async (data: DeleteTaskRequest) => {
      return axios
        .delete(`${getServerUrl()}/task/${data.task_id}`)
        .then((r) => r.data);
    },
    {
      onSuccess: () => {
        toast({
          title: "Task deleted.",
          description: "We've successfuly deleted the task.",
          variant: "success",
        });
      },
      onError: (error) => {
        toast({
          title: "Task deletion failed.",
          description: "Unable to delete task. Error: " + error,
          variant: "destructive",
        });
      },
    }
  );
};
