import { checkIfAuth, retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";
import { TaskResult } from "types";

interface LatestTasksResultsRequest {
  limit: number;
  page: number;
  unseen: boolean;
}

interface LatestTasksResultsResponse {
    count: number;
    results:TaskResult
}

export const useLatestTasksResults = (
  latestTasksResultsRequest: LatestTasksResultsRequest
) => {
  const isAuthenticated = checkIfAuth();
  return useQuery(
    ["tasks","results"],
    () => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      return axios
        .get<LatestTasksResultsResponse>(
          getServerUrl() +
            `/task/result/latest?limit=${latestTasksResultsRequest.limit}&page=${latestTasksResultsRequest.page}&unseen=${latestTasksResultsRequest.unseen}`,
          {
            headers: {
              Authorization: retrieveToken(),
            },
          }
        )
        .then((res) => res.data);
    },
    {
      retry: false,
      refetchOnWindowFocus: false,
    }
  );
};
