import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import { Cred, Host } from "@src/types";
import axios, { AxiosResponse } from "axios";
import { useQuery } from "react-query";

interface AllHostsRequest {}

interface AllHostsResponse {
  hosts: Host[];
  success: boolean;
}

export const useAllHosts = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery<AllHostsResponse>(
    ["findings", "hosts"],
    async () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(getServerUrl() + "/tracking/host", {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((res: AxiosResponse<AllHostsResponse>) => res.data);
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
      cacheTime: 0,
    }
  );
};
