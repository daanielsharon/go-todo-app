import { api } from "../../api/api";

const update = async (
  originId: number,
  originPriority: number,
  targetId: number,
  targetPriority: number
) => {
  const url = `todo/container/priority/${originId}`;
  const data = {
    originPriority,
    targetId,
    targetPriority,
  };

  const response = await api.patch(url, data);
  return response;
};

const create = async (userId: number, groupName: string, priority: number) => {
  const url = `/todo/container/`;
  const data = {
    userId,
    groupName,
    priority,
  };

  const response = await api.post(url, data);
  return response;
};

export default {
  create,
  update,
};
