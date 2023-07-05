import { ItemType } from "../../types/item";
import Item from "./Item";

type Props = {
  name: string;
  items: ItemType[];
  isDragging: boolean;
  handleDragStart: (
    e: React.DragEvent<HTMLDivElement>,
    data: ItemType | null
  ) => void;
  handleDragEnd: () => void;
  handleDragOver: (e: React.DragEvent<HTMLDivElement>) => void;
  handleDrop: (e: React.DragEvent<HTMLDivElement>) => void;
};

const Card = ({
  name,
  items,
  isDragging,
  handleDragStart,
  handleDragEnd,
  handleDragOver,
  handleDrop,
}: Props) => {
  return (
    <div
      className={`${isDragging ? "card-dragged" : null} card`}
      onDragOver={handleDragOver}
      onDrop={handleDrop}
    >
      <p className="text-white">{name}</p>
      {items.map((item, index) => (
        <Item
          key={index}
          item={item}
          handleDragStart={handleDragStart}
          handleDragEnd={handleDragEnd}
        />
      ))}
    </div>
  );
};

export default Card;
