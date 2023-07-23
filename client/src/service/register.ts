import { api } from "../api/api";

type Register = {
  username: string;
};

export const register = async ({ username }: Register) => {
  const url = "/users/register";
  const data = {
    username,
  };

  const response = await api.post(url, data);
  return response;
};
