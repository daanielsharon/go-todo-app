import { useState } from "react";
import AddTodoModal from "../modal/AddTodoModal";
import Item from "./Item";
import { ContainerDrag, ContainerType, ItemType } from "../../types/todo";

type Props = {
  index: number;
  name: string;
  item: ContainerType;
  items: ItemType[];
  groupId: number;
  isDragging: boolean;
  isContainerDragging: ContainerDrag;
  handleItemDrag: (
    e: React.DragEvent<HTMLDivElement>,
    data: ItemType | null
  ) => void;
  handleDragEnd: () => void;
  handleContainerDrag: (
    e: React.DragEvent<HTMLDivElement>,
    data: ContainerType
  ) => void;
  handleDragOver: (e: React.DragEvent<HTMLDivElement>) => void;
  handleItemDrop: (e: React.DragEvent<HTMLDivElement>) => void;
};

const Card = ({
  name,
  index,
  item,
  items,
  groupId,
  isDragging,
  isContainerDragging,
  handleItemDrag,
  handleDragEnd,
  handleDragOver,
  handleItemDrop,
  handleContainerDrag,
}: Props) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  return (
    <>
      <AddTodoModal
        open={isOpen}
        handleClose={() => setIsOpen(false)}
        groupId={groupId}
      />
      <div
        className={`${isDragging ? "card-dragged" : null} card cursor-pointer ${
          isContainerDragging.status
            ? isContainerDragging.containerIndex != index
              ? "card-dragged"
              : null
            : null
        }`}
        draggable={isContainerDragging.status}
        // onDragStart={(e) => {
        //   handleContainerDrag(e, item);
        // }}
        onDragOver={handleDragOver}
        onDrop={handleItemDrop}
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
            handleItemDrag={handleItemDrag}
            handleDragEnd={handleDragEnd}
            isContainerDragging={isContainerDragging}
          />
        ))}
      </div>
    </>
  );
};

export default Card;
