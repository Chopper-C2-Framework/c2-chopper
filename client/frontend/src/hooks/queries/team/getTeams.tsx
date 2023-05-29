import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import { Team } from "@src/types";
import axios from "axios";
import { useQuery } from "react-query";

interface GetTeamsResponse {
  success: boolean;
  teams: Team[];
}

export const useGetTeams = () => {
  const isAuthenticated = checkIfAuth();
  return useQuery<GetTeamsResponse>(
    ["teams"],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get<GetTeamsResponse>(getServerUrl() + "/management/team", {
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
