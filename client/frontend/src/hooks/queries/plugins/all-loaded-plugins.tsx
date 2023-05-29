import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";

interface AllLoadedPluginsResponse {
  success: boolean;
  names: string[];
}

export const useAllLoadedPlugins = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery<string[]>(
    ["plugins", "loaded"],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get<AllLoadedPluginsResponse>(getServerUrl() + "/plugins/loaded", {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((res) => res.data.names);
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
    }
  );
};
