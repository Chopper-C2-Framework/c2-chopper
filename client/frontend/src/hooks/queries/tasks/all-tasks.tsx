import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";
import { Task } from "types";

interface AllTasksResponse {
    tasks: (Task&{taskId:string})[],
    count:number
}

export const useAllTasks = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery(
    ["tasks"],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get<AllTasksResponse>(getServerUrl() + "/task/all", {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((res) => res.data);
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
    }
  );
};
