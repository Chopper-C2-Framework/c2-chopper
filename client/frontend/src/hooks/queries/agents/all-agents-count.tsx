import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";
import { Agent } from "types";


interface AllAgentsResponse {
  data: Agent[],
  count:number
}

export const useAllAgentsQuery = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery<AllAgentsResponse>(
    ["agents"],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get<AllAgentsResponse>(getServerUrl() + "/agent/all", {
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
