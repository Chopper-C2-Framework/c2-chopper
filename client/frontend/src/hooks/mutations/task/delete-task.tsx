import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";

interface DeleteTaskRequest {}

interface DeleteTaskResponse {}

export const useDeleteTask = (task_id: string) => {
  const navigate = useNavigate();
  const { toast } = useToast();

  return useMutation<DeleteTaskResponse, any, DeleteTaskRequest, any>(
    ["tasks"],
    async (data: DeleteTaskRequest) => {
      return axios
        .delete(getServerUrl() + "/v1/task" + task_id)
        .then((r) => r.data);
    },
    {
      onSuccess: (data) => {},
      onError: (error) => {},
    }
  );
};
