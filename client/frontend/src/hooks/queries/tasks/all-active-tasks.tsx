import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios, { AxiosResponse } from "axios";
import { useQuery } from "react-query";
import { TaskListResponse } from "./interfaces.tsx";

export const useAllActiveTasks = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery<TaskListResponse>(
    ["tasks", "active"],
    async () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(getServerUrl() + "/task/unexecuted/all", {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((res: AxiosResponse<TaskListResponse>) => res.data);
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
      cacheTime: 0
    }
  );
};
