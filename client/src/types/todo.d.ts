export type ContainerType = {
  id: number;
  group_name: string;
  item: ItemType[];
  priority: number;
};

export type ItemType = {
  id: number;
  name: string;
};
