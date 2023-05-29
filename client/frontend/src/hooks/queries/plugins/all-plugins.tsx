import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";

interface AllPluginsResponse {
  names: string[];
  success: true;
}

export const useAllPluginsQuery = () => {
  // const isAuthenticated = checkIfAuth();
  return useQuery<string[]>(
    ["plugins"],
    () => {
      // if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get<AllPluginsResponse>(getServerUrl() + "/plugins/all", {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((res) => {
          if (!res.data.success)
            throw new Error("Error retrieving plugins" + res.data.success);
          return res.data.names;
        });
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
    }
  );
};
