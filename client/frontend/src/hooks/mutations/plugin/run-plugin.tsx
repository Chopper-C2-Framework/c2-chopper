import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
import { Plugin } from "types";
import * as z from "zod";

interface RunPluginRequest {
  FileName: string;
  Args: {
    [key: string]:
      | {
          type: "string_value";
          string_value: string;
        }
      | {
          type: "number_value";
          string_value: number;
        };
  };
}

//  "Args": { "arg0": { "type": "string_value", "string_value": "192.168.1.1" }, "arg1": { "type": "string_value", "string_value": "1-100" }, "arg2": { "type": "map_value", "map_value": { "items": [ { "key": "key1", "value": { "type": "string_value", "string_value": "Hello world" } } ] } } }

interface RunPluginsResponse {
  success: boolean;
  data: Plugin;
}
export const loadPluginSchema = z.object({});

export const useRunPluginMutation = () => {
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
