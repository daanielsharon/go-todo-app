import { ItemType } from "../../types/item";
import Item from "./Item";

const Card = ({ name, items }: { name: string; items: ItemType[] }) => {
  return (
    <div className={`card`}>
      <p className="text-white">{name}</p>
      {items.map((item) => (
        <Item item={item} />
      ))}
    </div>
  );
};

export default Card;
