import { useToast } from "@components/ui/use-toast";
import { retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { Team } from "types";
import * as z from "zod";

interface JoinToTeamRequest {
  user_id: string;
}

interface JoinToTeamResponse {
  success: boolean;
  data: Team;
}
export const addMemberToTeamSchema = z.object({
  member_id: z.string(),
});

export const useAddMemberToTeam = (team_id: string) => {
  const { toast } = useToast();

  return useMutation<JoinToTeamResponse, any, JoinToTeamRequest, any>(
    ["teams", team_id],
    async (data: JoinToTeamRequest) => {
      return axios
        .post(
          getServerUrl() + "/management/team/join/" + team_id,
          data.user_id,
          {
            headers: {
              Authorization: retrieveToken(),
            },
          }
        )
        .then((r) => r.data);
    },
    {
      onSuccess: (_) => {
        toast({
          title: "Member was added successfully",
        });
      },
      onError: (error) => {
        toast({
          title: "Error adding new member to team",
          description: error.message,
        });
      },
      onMutate: () => {
        toast({
          title: "Member is being added",
          description: "Please wait",
        });
      },
    }
  );
};
