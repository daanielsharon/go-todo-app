import { api } from "../api/api";

type Auth = {
  username: string;
};

const login = async ({ username }: Auth) => {
  const url = "/users/login";
  const data = {
    username,
  };

  const response = await api.post(url, data);
  return response;
};

const register = async ({ username }: Auth) => {
  const url = "/users/register";
  const data = {
    username,
  };

  const response = await api.post(url, data);
  return response;
};

export default {
  login,
  register,
};
