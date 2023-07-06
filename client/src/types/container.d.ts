import { ItemType } from "./item";

export type ContainerType = {
  id: number;
  name: string;
  item: ItemType[];
  priority: number;
};
