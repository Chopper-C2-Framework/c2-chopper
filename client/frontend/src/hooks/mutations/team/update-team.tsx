import { useToast } from "@components/ui/use-toast";
import { retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { Team } from "types";
import * as z from "zod";

interface UpdateTeamRequest {
  data: Partial<Team>;
}

interface UpdateTeamResponse {
  success: boolean;
  data: Team;
}
export const updateTeamSchema = z.object({
  name: z.string(),
});

export const useUpdateTeamMutation = (team_id: string) => {
  const { toast } = useToast();

  return useMutation<UpdateTeamResponse, any, UpdateTeamRequest, any>(
    ["teams", team_id],
    async (data: UpdateTeamRequest) => {
      return axios
        .patch(getServerUrl() + "/management/team/" + team_id, data, {
          headers: {
            Authorization: retrieveToken(),
          },
        })
        .then((r) => r.data);
    },
    {
      onSuccess: (_) => {
        toast({
          title: "Team was created successfully",
        });
      },
      onError: (error) => {
        toast({
          title: "Error creating the team",
          description: error.message,
        });
      },
      onMutate: () => {
        toast({
          title: "Team is being created",
          description: "Please wait",
        });
      },
    }
  );
};
