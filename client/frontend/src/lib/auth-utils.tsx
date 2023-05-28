import { useNavigate } from "react-router-dom";

export const SetAuthUser = (token: string) => {
  localStorage.setItem("token", token);
};

export const checkIfAuth = () => {
  return localStorage.getItem("token") ? true : false;
};

export const retrieveToken = () => {
  const token = localStorage.getItem("token");
  return token ? token : "";
};
