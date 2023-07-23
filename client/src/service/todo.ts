import { api } from "../api/api";

export const fetchTodo = async (username: string) => {
  const url = `todo/${username}`;
  const response = await api.get(url);
  return response;
};
