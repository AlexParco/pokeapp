import { loginService, registerService } from "@/service/api";
import { State } from "@/types/state.type";
import { User, UserLogin, UserRegister } from "@/types/user.type";
import { createContext, ReactNode, useCallback, useContext, useState } from "react";

interface IAuthContext {
  login: ({ username, password }: UserLogin) => void;
  register: ({ username, password }: UserRegister) => void;
  logout: () => void;
  isLogged: Boolean;
  token: string;
  state: State;
}

const userDefault: User = {
  user_id: "",
  username: "",
  password: "",
  role: "",
}

const stateDefault: State = { error: false, loading: false, succes: false }

const AuthContext = createContext<IAuthContext>({
  login: () => null,
  logout: () => null,
  register: () => null,
  isLogged: false,
  token: "",
  state: stateDefault,
})

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [token, setToken] = useState<string>(() => window.localStorage.getItem("token") as string)
  const [state, setState] = useState<State>(stateDefault)

  const login = ({ username, password }: UserLogin) => {
    setState({ error: false, loading: true, succes: false })
    loginService({ username, password })
      .then(data => {
        setToken(data.token)
        window.localStorage.setItem("token", data.token)

        setState({ error: false, loading: false, succes: true })
      })
      .catch(err => {
        setState({ error: true, loading: false, succes: false })
      })
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
      state,
    }}
  >
    {children}
  </AuthContext.Provider>
}

export const useAuth = () => {
  const { login, register, logout, isLogged, token, state } = useContext(AuthContext)
  return {
    login,
    register,
    logout,
    isLogged,
    token,
    state
  }
} 