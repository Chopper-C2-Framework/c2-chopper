import LoadingPage from "@components/loading-page";
import { checkIfAuth } from "@lib/auth-utils";
import { getServerUrl } from "@lib/get-server-url";
import React, { useEffect, useState } from "react";
import { useQuery } from "react-query";
import { useNavigate } from "react-router-dom";

interface ProtectedRouteProps {
  reverseProtect: boolean;
  children: React.ReactNode;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({
  reverseProtect,
  children,
}) => {
  const isAuthenticated = checkIfAuth();
  const [isHidden, setIsHidden] = useState<boolean>(true);

  const { data, isLoading, isError } = useQuery(["user", "me"], () => {
    if (!isAuthenticated) throw new Error("Unable to login ");
    return fetch(getServerUrl() + "/user/me").then((res) => res.json());
  });

  const navigate = useNavigate();

  useEffect(() => {
    if (((!isAuthenticated||isError) && reverseProtect) || data) {
      setIsHidden(false);
    } else {
      navigate(-1);
    }
  }, [isError, data]);

  if (isLoading && isHidden) {
    return <LoadingPage />;
  }

  return <>{children}</>;
};

export default ProtectedRoute;
