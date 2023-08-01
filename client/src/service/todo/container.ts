import { api } from "../../api/api";

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

export default {
  update,
};
