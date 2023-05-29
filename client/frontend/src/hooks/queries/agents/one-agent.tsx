import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import { Agent } from "@src/types";
import axios, { AxiosResponse } from "axios";
import { useQuery } from "react-query";

interface IGetAgentInfoResponse{
  agent: Agent;
}

export const useAgentInfo = (agent_id: string) => {
  const isAuthenticated = checkIfAuth();
  return useQuery<Agent>(
    ["plugins"],
    async () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(getServerUrl() + "/agent/" + agent_id, {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((res: AxiosResponse<IGetAgentInfoResponse>) => res.data.agent);
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
    }
  );
};
