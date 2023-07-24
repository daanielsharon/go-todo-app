import context from "../../context";
import service from "../../service";
import { ContainerType, ItemType } from "../../types/todo";

type Props = {
  item: ItemType | null;
  handleDragStart: (
    e: React.DragEvent<HTMLDivElement>,
    data: ItemType | null
  ) => void;
  handleDragEnd: () => void;
};

const Item = ({ item, handleDragStart, handleDragEnd }: Props) => {
  const handleClick = async (id: number) => {
    // api logic
    const res = await service.todo.remove(id);
    if (res) {
      const data: ContainerType[] = context.getContext("todo", "data");
      const newData = [...data];
      newData.forEach((item, idx) => {
        item.item.forEach((element, itemIdx) => {
          if (element.id === id) {
            newData[idx].item.splice(itemIdx, 1);
          }
        });
      });

      context.setContext("todo", "data", newData);
    }
  };
  return (
    <>
      <div className="relative">
        <button
          className="absolute right-4 top-0"
          onClick={() => item?.id && handleClick(item.id)}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="10"
            height="10"
            viewBox="0 0 24 24"
          >
            <path d="M24 20.188l-8.315-8.209 8.2-8.282-3.697-3.697-8.212 8.318-8.31-8.203-3.666 3.666 8.321 8.24-8.206 8.313 3.666 3.666 8.237-8.318 8.285 8.203z" />
          </svg>
        </button>
        <div
          className="item-container"
          draggable
          onDragStart={(e) => handleDragStart(e, item)}
          onDragEnd={handleDragEnd}
        >
          {item?.name}
        </div>
      </div>
    </>
  );
};

export default Item;
