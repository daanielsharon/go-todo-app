import { ItemType } from "../../types/item";

const Item = ({ item }: { item: ItemType | null }) => {
  return (
    <div className="item-container" draggable>
      {item?.name}
    </div>
  );
};

export default Item;
