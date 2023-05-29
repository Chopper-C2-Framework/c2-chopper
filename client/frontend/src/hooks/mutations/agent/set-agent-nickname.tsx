import { useToast } from "@components/ui/use-toast";
import { checkIfAuth } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";

interface ISetAgentNicknameRequest {
  agent_id: string;
  nickname: string;
}

export const useSetAgentNickname = () => {
  const isAuthenticated = checkIfAuth();
  const { toast } = useToast();

  return useMutation<any, any, ISetAgentNicknameRequest, any>(
    ["plugins"],
    async (data: ISetAgentNicknameRequest) => {
      if (!isAuthenticated) throw new Error("Unable to login ");
      
      return axios
        .patch(getServerUrl() + "/agent/" + data.agent_id, data)
        .then((res) => res.data);
    },
    {
      onSuccess: () => {
        toast({
          title: "Nickname edited.",
          description: "We've successfuly edited the agent nickname.",
          variant: "success",
        });
      },
      onError: (error) => {
        toast({
          title: "Agent editing failed.",
          description: "Unable to edit agent nickname. Error: " + error,
          variant: "destructive",
        });
      },
    }
  );
};
