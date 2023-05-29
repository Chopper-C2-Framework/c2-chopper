import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import { Task } from "@src/types";
import axios, { AxiosResponse } from "axios";
import { useQuery } from "react-query";

interface GetTaskResponse {
  task: Task;
}

export const useTask = (task_id: string) => {
  const isAuthenticated = checkIfAuth();
  return useQuery<Task>(
    ["task", task_id],
    async () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(`${getServerUrl()}/task/${task_id}`, {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((res: AxiosResponse<GetTaskResponse>) => res.data.task);
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
    }
  );
};
