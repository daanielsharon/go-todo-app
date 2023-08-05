import { api } from "../../api/api";

const update = async (
  originId: number,
  originPriority: number,
  targetId: number,
  targetPriority: number
) => {
  const url = `todo/priority/${originId}`;
  const data = {
    originPriority,
    targetId,
    targetPriority,
  };

  const response = await api.patch(url, data);
  return response;
};

export default {
  update,
};
