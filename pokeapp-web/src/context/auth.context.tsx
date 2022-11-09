import { loginService, registerService } from "@/service/api";
import { User, UserLogin, UserRegister } from "@/types/user.type";
import { createContext, ReactNode, useContext, useState } from "react";

interface IAuthContext {
  login: ({ username, password }: UserLogin) => void;
  register: ({ username, password }: UserRegister) => void;
  logout: () => void;
  isLogged: Boolean;
  token: string;
}

const userDefault: User = {
  user_id: "",
  username: "",
  password: "",
  role: "",
}

const AuthContext = createContext<IAuthContext>({
  login: () => null,
  logout: () => null,
  register: () => null,
  isLogged: false,
  token: "",
})

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [token, setToken] = useState<string>(() => window.localStorage.getItem("token") as string)

  const login = ({ username, password }: UserLogin) => {
    async function fetchLogin() {
      try {
        const data = await loginService({ username, password })
        setToken(data.token)
        window.localStorage.setItem("token", data.token)

      } catch (error) {
        console.log(error)
      }
    }
    fetchLogin()
  }

  const register = ({ username, role, password }: UserRegister) => {
    registerService({ username, role, password })
      .then(data => {
        setToken(data.token)
      })
      .catch(err => console.log(err))
  }

  const logout = () => {
    window.localStorage.removeItem("token")
    setToken("")
  }

  return <AuthContext.Provider
    value={{
      login,
      register,
      logout,
      isLogged: Boolean(token),
      token,
    }}
  >
    {children}
  </AuthContext.Provider>
}

export const useAuth = () => {
  const { login, register, logout, isLogged, token } = useContext(AuthContext)
  return {
    login,
    register,
    logout,
    isLogged,
    token,
  }
} 