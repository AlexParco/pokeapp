export interface User {
  user_id: string;
  username: string;
  password: string;
  role: string;

}

export interface UserLogin {
  username: string;
  password: string;
}

export interface UserRegister {
  username: string;
  role: string;
  password: string;
}

export interface UserWithToken {
  user: User;
  token: string;
}
