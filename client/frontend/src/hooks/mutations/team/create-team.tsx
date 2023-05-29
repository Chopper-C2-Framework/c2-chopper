import { useToast } from "@components/ui/use-toast";
import { retrieveToken } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation, useQueryClient } from "react-query";
import { Team } from "types";
import * as z from "zod";

interface CreateTeamRequest extends Partial<Team>{
}

interface CreateTeamResponse {
  success: boolean;
  data: Team;
}
export const createTeamSchema = z.object({
  name: z.string(),
});

export const useCreateTeam = () => {
  const { toast } = useToast();
  const queryClient=useQueryClient()

  return useMutation<CreateTeamResponse, any, CreateTeamRequest, any>(
    ["teams"],
    async (data: CreateTeamRequest) => {
      return axios
        .post(getServerUrl() + "/management/team", data, {
          headers: {
            "Authorization":retrieveToken(),
          }
        })
        .then((r) => r.data);
    },
    {
      onSuccess: (_) => {
        queryClient.invalidateQueries("teams")
        
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
