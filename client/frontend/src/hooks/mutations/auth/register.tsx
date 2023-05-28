import { useMutation } from "react-query";
import * as z from "zod"
import axios from "axios";
import { getServerUrl } from "@lib/get-server-url";

interface RegisterRequest {
  username: string;
  password: string;
}

export const registerSchema = z.object({
  username: z.string().min(2).max(50),
  password : z.string().min(8).max(50)
})

export const useRegisterMutation = () =>
  useMutation(
    async (data: RegisterRequest) => {
      return axios.post(getServerUrl()+"/v1/register", data);
    },
    {
      onSuccess: (data) => {},

      onError: (error) => {},
    }
  );


