import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
import { Task } from "types";

interface CreateTaskRequest {
    task:Task

}

interface CreateTaskResponse {
}
export const useCreateTask = (agent_id: string) => {
  const navigate = useNavigate();
  const { toast } = useToast();

  return useMutation<CreateTaskResponse, any, CreateTaskRequest, any>(
    ["tasks"],
    async (data: CreateTaskRequest) => {
      return axios
        .post(getServerUrl() + "/v1/task" + agent_id, data)
        .then((r) => r.data);
    },
      {
          onSuccess: (data) => {
        
          },
          onError: (error) => {
          }
      }
  );
};
