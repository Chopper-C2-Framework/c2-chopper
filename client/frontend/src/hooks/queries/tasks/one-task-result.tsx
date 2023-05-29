import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios, { AxiosResponse } from "axios";
import { useQuery } from "react-query";
import { TaskResultListResponse } from "./interfaces";

export const useTaskResults = (task_id: string) => {
  const isAuthenticated = checkIfAuth();
  return useQuery<TaskResultListResponse>(
    ["task", task_id, "results"],
    async () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(`${getServerUrl()}/task/result/${task_id}`, {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((res: AxiosResponse<TaskResultListResponse>) => res.data);
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
      cacheTime: 0,
    }
  );
};
