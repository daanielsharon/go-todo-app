import { ItemType } from "../../types/item";

type Props = {
  item: ItemType | null;
  handleDragStart: (
    e: React.DragEvent<HTMLDivElement>,
    data: ItemType | null
  ) => void;
  handleDragEnd: () => void;
};

const Item = ({ item, handleDragStart, handleDragEnd }: Props) => {
  return (
    <div
      className="item-container"
      draggable
      onDragStart={(e) => handleDragStart(e, item)}
      onDragEnd={handleDragEnd}
    >
      {item?.name}
    </div>
  );
};

export default Item;
