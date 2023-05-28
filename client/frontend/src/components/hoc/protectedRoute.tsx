import LoadingPage from "@components/loading-page";
import { useMeQuery } from "@hooks/queries/user/me";
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

interface ProtectedRouteProps {
  reverseProtect: boolean;
  children: React.ReactNode;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({
  reverseProtect,
  children,
}) => {
  const [isHidden, setIsHidden] = useState<boolean>(true);
  const { data, isLoading, isError } = useMeQuery();

  const navigate = useNavigate();

  useEffect(() => {
    if ((isError || !data) && !reverseProtect) {
      navigate("/login");
    } else if ((isError || (!data && !isLoading)) && reverseProtect) {
      setIsHidden(false);
    } else if (data && reverseProtect) {
      navigate("/app/dashboard");
    } else if (data && !reverseProtect) {
      setIsHidden(false);
    }
  }, [isError, data]);

  if (isHidden) {
    return <LoadingPage />;
  }

  return <>{children}</>;
};

export default ProtectedRoute;
