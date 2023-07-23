import { api } from "../api/api";

type Login = {
  username: string;
};

export const login = async ({ username }: Login) => {
  const url = "/users/login";
  const data = {
    username,
  };

  const response = await api.post(url, data);
  return response;
};
