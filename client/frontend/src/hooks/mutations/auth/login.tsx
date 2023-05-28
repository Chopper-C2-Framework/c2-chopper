import { useMutation } from "react-query";
import axios from "axios";
import * as z from "zod"
import { getServerUrl } from "@lib/get-server-url";

interface LoginRequest {
  username: string;
  password: string;
}

 
export const loginSchema = z.object({
  username: z.string().min(2).max(50),
  password : z.string().min(8).max(50)
})

export const useLoginMutation = () =>
  useMutation(
    async (data: LoginRequest) => {
      console.log(getServerUrl()+"/v1/login")
      return axios.post(getServerUrl()+"/v1/login", data);
    },
    {
      onSuccess: (data) => {},

      onError: (error) => {},
    }
  );
