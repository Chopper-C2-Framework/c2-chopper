import { useToast } from "@components/ui/use-toast";
import { retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { Team } from "types";

interface DeleteTeamRequest {}

interface DeleteTeamResponse {
  success: boolean;
  data: Team;
}

export const useDeleteTeamMutation = (team_id: string) => {
  const { toast } = useToast();

  return useMutation<DeleteTeamResponse, any, DeleteTeamRequest, any>(
    ["teams"],
    async (_: DeleteTeamRequest) => {
      return axios
        .delete(getServerUrl() + "/management/team/" + team_id, {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((r) => r.data);
    },
    {
      onSuccess: (_) => {
        toast({
          title: "Team was deleted successfully",
        });
      },
      onError: (error) => {
        toast({
          title: "Error deleting the team",
          description: error.message,
        });
      },
      onMutate: () => {
        toast({
          title: "Team is being deleted",
          description: "Please wait",
        });
      },
    }
  );
};
