import { api } from "../api/api";

type Auth = {
  username: string;
};

export const login = async ({ username }: Auth) => {
  const url = "/users/login";
  const data = {
    username,
  };

  const response = await api.post(url, data);
  return response;
};

export const register = async ({ username }: Auth) => {
  const url = "/users/register";
  const data = {
    username,
  };

  const response = await api.post(url, data);
  return response;
};
