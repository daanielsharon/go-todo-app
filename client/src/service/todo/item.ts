import { api } from "../../api/api";

const get = async (username: string) => {
  const url = `todo/${username}`;
  const response = await api.get(url);
  return response;
};

const create = async (userId: number, groupId: number, name: string) => {
  const url = "todo/";
  const data = {
    userId,
    groupId,
    name,
  };

  const response = await api.post(url, data);
  return response;
};

const update = async (
  id: number,
  name: string,
  groupId: number,
  userId: number
) => {
  const url = `todo/${id}`;
  const data = {
    name,
    groupId,
    userId,
  };

  const response = await api.patch(url, data);
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
