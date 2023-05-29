import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";

export const useAgentTasks = (agent_id: string) => {
  const isAuthenticated = checkIfAuth();
  return useQuery(
    ["tasks", "agent", agent_id],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(getServerUrl() + "/task/agent/" + agent_id, {
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
