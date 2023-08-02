import { api } from "../../api/api";

const update = async (
  id: number,
  name: string,
  groupId: number,
  userId: number
) => {
  const url = "todo/";
  const data = {
    id,
    name,
    groupId,
    userId,
  };

  const response = await api.patch(url, data);
  return response;
};

export default {
  update,
};
