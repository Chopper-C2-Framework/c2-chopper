import { useNavigate } from "react-router-dom";

export const useLogout = () => {
  const navigation = useNavigate();

  const logout = () => {
    localStorage.removeItem("token");
    setTimeout(() => navigation("/"), 1000);
  };

  return { logout };
};
