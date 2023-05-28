import { getServerUrl } from "@lib/get-server-url";
import axios from "axios";
import { useQuery } from "react-query";
import { useLogout } from "./useLogout";
import { useToast } from "@components/ui/use-toast";

const UseMe = () => {
  const { logout } = useLogout();
  const { toast } = useToast();
  return useQuery(["user", "me"], () => axios.get(getServerUrl() + "/me"), {
    onError(err) {
      toast({
        title: "Error",
        description: JSON.stringify(err),
        variant: "destructive",
      });
      logout();
    },
  });
};

export default UseMe;
