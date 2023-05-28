import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";

export const useAllPluginsQuery = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery(
    ["plugins"],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(getServerUrl() + "/v1/plugins/all", {
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


