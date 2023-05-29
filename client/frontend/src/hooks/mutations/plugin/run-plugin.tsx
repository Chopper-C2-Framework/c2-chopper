import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
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

export const useRunPluginMutation = () => {
  const navigate = useNavigate();
  const { toast } = useToast();

  return useMutation<LoadPluginResponse, any, LoadPluginRequest, any>(
    ["plugins"],
    async (data: LoadPluginRequest) => {
      return axios
        .post(getServerUrl() + "/plugins/run", data)
        .then((r) => r.data);
    },
    {
      onSuccess: (data) => {},
      onError: (error) => {},
    }
  );
};
