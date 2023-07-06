import { Navigate } from "react-router";
import useAuth from "../../hooks/useAuth";

const ProtectedRoute = ({ children }: { children: JSX.Element }) => {
  const { username } = useAuth();

  if (!username) {
    return <Navigate to="/login" />;
  }

  return children;
};

export default ProtectedRoute;
