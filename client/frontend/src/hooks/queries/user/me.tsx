import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";

export const useMeQuery = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery(
    ["user", "me"],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get(getServerUrl() + "/me", {
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
