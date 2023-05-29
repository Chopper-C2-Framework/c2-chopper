import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";

interface SetTasksResultsSeenRequest {}

interface SetTasksResultSeenResponse {}
export const useSetTasksSeen = () => {
  const navigate = useNavigate();
  const { toast } = useToast();

  return useMutation<
    SetTasksResultSeenResponse,
    any,
    SetTasksResultsSeenRequest,
    any
  >(
    ["tasks"],
    async (data: SetTasksResultsSeenRequest) => {
      return axios
        .patch(getServerUrl() + "/task/result/seen", data)
        .then((r) => r.data);
    },
    {
      onSuccess: (data) => {},
      onError: (error) => {},
    }
  );
};
