import { api } from "../api/api";

const login = async ({ username, password }: AuthType) => {
  const url = "/users/login";
  const data = {
    username,
    password,
  };

  const response = await api.post(url, data);
  return response;
};

const register = async ({ username, password }: AuthType) => {
  const url = "/users/register";
  const data = {
    username,
    password,
  };

  const response = await api.post(url, data);
  return response;
};

const logout = async () => {
  const url = "/users/logout";
  const response = await api.post(url, null);
  return response;
};

export default {
  login,
  register,
  logout,
};
