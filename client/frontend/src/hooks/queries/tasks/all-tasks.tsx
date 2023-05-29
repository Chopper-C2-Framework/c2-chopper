import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";

export const useAllTasks = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery(
    ["tasks", "all"],
    async () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(getServerUrl() + "/task/all", {
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
