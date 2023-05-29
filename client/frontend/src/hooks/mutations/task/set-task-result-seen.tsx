import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";

interface SetTasksResultsSeenRequest {
  result_ids: string[];
}

interface SetTasksResultSeenResponse {}

export const useSetTasksSeen = () => {

  return useMutation<
    SetTasksResultSeenResponse,
    any,
    SetTasksResultsSeenRequest,
    any
  >(
    ["task", "seen"],
    async (data: SetTasksResultsSeenRequest) => {
      return axios
        .patch(getServerUrl() + "/task/result/seen", data)
        .then((r) => r.data);
    },
  );
};
