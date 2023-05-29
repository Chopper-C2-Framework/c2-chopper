import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
import { Plugin } from "types";
import * as z from "zod";

interface RunPluginRequest {
  file_name: string;
  items: { [key: string]: "number" | "string" };
}

interface RunPluginsResponse {
  success: boolean;
  data: Plugin;
}
export const loadPluginSchema = z.object({});

export const useRunPluginMutation = () => {
  const navigate = useNavigate();
  const { toast } = useToast();

  return useMutation<RunPluginsResponse, any, RunPluginRequest, any>(
    ["plugins"],
    async (data: RunPluginRequest) => {
      return axios
        .post(getServerUrl() + "/plugins/run", data)
        .then((r) => r.data);
    },
    {
      onSuccess: (data) => {
        toast({
          title: "Plugin has ran",
          description: "You can view the results",
        });
      },
      onError: (error) => {
        toast({
          title: "Error",
          description: error.message,
        });
      },
      onMutate: () => {
        toast({
          title: "Plugin is running",
          description: "Please wait",
        });
      },
    }
  );
};
