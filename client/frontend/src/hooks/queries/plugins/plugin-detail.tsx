import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios, { AxiosResponse } from "axios";
import { useQuery } from "react-query";
import { Plugin } from "types";

interface PluginDetailsResponse {
  data: Plugin;
}

export const usePluginsDetails = (plugin_id: string, enabled?: boolean) => {
  const isAuthenticated = checkIfAuth();
  return useQuery<Plugin>(
    ["plugins", plugin_id],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get<{ data: Plugin }>(
          getServerUrl() + "/plugins/details/" + plugin_id,
          {
            headers: {
              Authorization: retrieveToken(),
            },
          }
        )
        .then((res: AxiosResponse<PluginDetailsResponse>) => res.data.data);
    },
    {
      retry: false,
      enabled,
      refetchOnWindowFocus: false,
    }
  );
};
