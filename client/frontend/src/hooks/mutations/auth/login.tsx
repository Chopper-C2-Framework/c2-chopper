import { useToast } from "@components/ui/use-toast";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
import * as z from "zod";

interface LoginRequest {
  username: string;
  password: string;
}

interface LoginResponse {
  token: string;
  success: boolean;
}
export const loginSchema = z.object({
  username: z.string().min(2).max(50),
  password: z.string().min(8).max(50),
});

export const useLoginMutation = () => {
  const navigate = useNavigate();
  const { toast } = useToast();

  return useMutation<LoginResponse, any, LoginRequest, any>(
    ["user"],
    async (data: LoginRequest) => {
      return axios.post(getServerUrl() + "/v1/login", data);
    },
    {
      onSuccess: (data) => {
        if (data.token) {
          toast({
            title: "Successfully logged in.",
            description: "We are taking you to your dashboard.",
            variant: "success",
          });
          setTimeout(() => navigate("/app/dashboard"), 1000);
        }
      },
      onError: (error) => {
        toast({
          title: "Account access failed.",
          description:
            "Please verify your credentials, if you have forgotten them please contact your admin." +
            error,
          variant: "destructive",
        });
      },
    }
  );
};
