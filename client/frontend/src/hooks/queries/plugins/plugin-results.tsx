import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import { PluginResult } from "@src/types";
import axios, { AxiosResponse } from "axios";
import { useQuery } from "react-query";

interface IPluginResultsResponse {
  results: PluginResult[];
  count: number;
}

export const usePluginResults = (path: string) => {
  const isAuthenticated = checkIfAuth();

  return useQuery<IPluginResultsResponse>(
    ["plugin", path, "results"],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(getServerUrl() + "/plugins/results/" + path, {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((res: AxiosResponse<IPluginResultsResponse>) => res.data);
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
    }
  );
};
