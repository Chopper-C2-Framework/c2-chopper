import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation, useQueryClient } from "react-query";
import { useNavigate } from "react-router-dom";
import { Plugin } from "types";
import * as z from "zod";

interface LoadPluginRequest {
  file_name: string;
}

interface LoadPluginResponse {
  success: boolean;
  data: Plugin;
}
export const loadPluginSchema = z.object({});

export const useLoadPluginMutation = () => {
  const navigate = useNavigate();
  const { toast } = useToast();
  const client= useQueryClient()

  return useMutation<LoadPluginResponse, any, LoadPluginRequest, any>(
    ["plugins", "loaded"],
    async (data: LoadPluginRequest) => {
      return axios
        .post<LoadPluginResponse>(getServerUrl() + "/plugins/load", data)
        .then((r) => r.data);
    },
    {
      onSuccess: (data) => {
        client.invalidateQueries("plugins")
        

      },
      
      onError: (error) => {},
    }
  );
};
