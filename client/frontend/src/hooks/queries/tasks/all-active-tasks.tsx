import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";
import { Task } from "types";

interface AllActiveTasksResponse{
  tasks: Task[],
  count:number
}

export const useAllActiveTasks = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery(
    ["tasks","unexectued"],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get<AllActiveTasksResponse>(getServerUrl() + "/task/unexecuted/all", {
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
