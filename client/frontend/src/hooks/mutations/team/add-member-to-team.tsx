import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { Team } from "types";
import * as z from "zod";

interface AddMemberToTeamRequest {
  user_id: string;
}

interface AddMemberToTeamResponse {
  success: boolean;
  data: Team;
}
export const addMemberToTeamSchema = z.object({
  member_id: z.string(),
});

export const useAddMemberToTeam = (team_id: string) => {
  const { toast } = useToast();

  return useMutation<AddMemberToTeamResponse, any, AddMemberToTeamRequest, any>(
    ["teams"],
    async (data: AddMemberToTeamRequest) => {
      return axios
        .post(
          getServerUrl() + "/management/team/members/" + team_id,
          data.user_id
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
