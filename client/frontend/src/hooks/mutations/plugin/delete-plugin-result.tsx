import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";

interface DeletePluginResultRequest {
  result_id: string;
}

interface DeletePluginResultResponse {}

export const useDeletePluginResult = () => {
  const { toast } = useToast();

  return useMutation<DeletePluginResultResponse, any, DeletePluginResultRequest, any>(
    ["result", "delete"],
    async (data: DeletePluginResultRequest) => {
      return axios
        .delete(`${getServerUrl()}/plugins/results/${data.result_id}`)
        .then((r) => r.data);
    },
    {
      onSuccess: () => {
        toast({
          title: "Plugin result deleted.",
          description: "We've successfuly deleted the plugin result.",
          variant: "success",
        });
      },
      onError: (error) => {
        toast({
          title: "Plugin result deletion failed.",
          description: "Unable to delete plugin result. Error: " + error,
          variant: "destructive",
        });
      },
    }
  );
};
