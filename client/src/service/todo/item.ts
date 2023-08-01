import { api } from "../../api/api";

const get = async (username: string) => {
  const url = `todo/${username}`;
  const response = await api.get(url);
  return response;
};

const create = async (user_id: number, group_id: number, name: string) => {
  const url = "todo/";
  const data = {
    user_id,
    group_id,
    name,
  };

  const response = await api.post(url, data);
  return response;
};

const update = async (
  id: number,
  name: string,
  group_id: number,
  user_id: number
) => {
  const url = "todo/";
  const data = {
    id,
    name,
    group_id,
    user_id,
  };

  const response = await api.put(url, data);
  return response;
};

const remove = async (todoId: number) => {
  const url = `todo/${todoId}`;
  const response = await api.delete(url);

  return response;
};

export default {
  get,
  create,
  update,
  remove,
};
