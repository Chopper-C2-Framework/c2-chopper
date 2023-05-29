import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import { Cred } from "@src/types";
import axios, { AxiosResponse } from "axios";
import { useQuery } from "react-query";

interface AllCredsRequest {}

interface AllCredsResponse {
  creds: Cred[];
  success: boolean;
}

export const useAllCreds = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery<AllCredsResponse>(
    ["findings", "creds"],
    async () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(getServerUrl() + "/tracking/creds", {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((res: AxiosResponse<AllCredsResponse>) => res.data);
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
      cacheTime: 0,
    }
  );
};
