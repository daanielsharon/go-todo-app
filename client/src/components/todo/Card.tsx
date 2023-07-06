import { useState } from "react";
import { ItemType } from "../../types/item";
import AddTodoModal from "../modal/AddTodoModal";
import Item from "./Item";

type Props = {
  index: number;
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
  index,
  items,
  isDragging,
  handleDragStart,
  handleDragEnd,
  handleDragOver,
  handleDrop,
}: Props) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  console.info(isOpen);

  return (
    <>
      <AddTodoModal open={isOpen} handleClose={() => setIsOpen(false)} />
      <div
        className={`${isDragging ? "card-dragged" : null} card`}
        onDragOver={handleDragOver}
        onDrop={handleDrop}
      >
        <div className="card-container">
          <p className="text-white card-container-title">{name}</p>
          {index === 0 && (
            <button
              className="text-white card-container-button"
              onClick={() => setIsOpen(true)}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="20"
                height="20"
                viewBox="0 0 24 24"
              >
                <path
                  fill="white"
                  d="M19 13h-6v6c0 .6-.4 1-1 1s-1-.4-1-1v-6H5c-.6 0-1-.4-1-1s.4-1 1-1h6V5c0-.6.4-1 1-1s1 .4 1 1v6h6c.6 0 1 .4 1 1s-.4 1-1 1z"
                />
              </svg>
            </button>
          )}
        </div>
        {items.map((item, index) => (
          <Item
            key={index}
            item={item}
            handleDragStart={handleDragStart}
            handleDragEnd={handleDragEnd}
          />
        ))}
      </div>
    </>
  );
};

export default Card;
