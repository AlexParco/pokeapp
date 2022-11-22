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

export type UserRegister = Omit<User, "user_id">

export interface UserWithToken {
  user: User;
  token: string;
}
