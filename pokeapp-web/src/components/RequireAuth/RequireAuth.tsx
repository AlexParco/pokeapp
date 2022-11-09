import { useAuth } from "@/context/auth.context";
import { ReactNode } from "react";
import { Navigate } from "react-router-dom";

const RequireAuth = ({ children }: { children: ReactNode }) => {
  const { isLogged } = useAuth()

  if (!isLogged) {
    return (
      < Navigate to="/login" replace />
    )
  }

  return (
    <>
      {children}
    </>
  )
}

export default RequireAuth