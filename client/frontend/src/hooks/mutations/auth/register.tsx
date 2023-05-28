import { useToast } from "@components/ui/use-toast";
import { SetAuthUser } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
import * as z from "zod";

interface RegisterRequest {
  username: string;
  password: string;
}

interface RegisterResponse {
  success: boolean;
  token: string;
}

export const registerSchema = z.object({
  username: z.string().min(2).max(50),
  password: z.string().min(8).max(50),
});

export const useRegisterMutation = () => {
  const navigate = useNavigate();
  const { toast } = useToast();

  return useMutation<RegisterResponse, any, RegisterRequest, any>(
    ["user"],
    async (data: RegisterRequest) => {
      return axios.post(getServerUrl() + "/register", data).then(r=>r.data);
    },
    {
      onSuccess: (data) => {
        console.log(data)
        if (data.token) {
          toast({
            title: "Account created.",
            description: "We've created your account for you.",
            variant: "success",
          });
          SetAuthUser(data.token)
          setTimeout(() => navigate("/app/dashboard"), 1000);
        }
      },
      onError: (error) => {
        toast({
          title: "Account creation failed.",
          description: "We've created your account for you." + error,
          variant: "destructive",
        });
      },
    }
  );
};
