import TodoLayout from "../components/layout/Todo";
import Card from "../components/todo/Card";
import fake from "../data/card.json";
import useAuth from "../hooks/useAuth";
import useDragAndDrop from "../hooks/useDragAndDrop";

const Todo = () => {
  const { username } = useAuth();
  const {
    data,
    isDragging,
    handleDragStart,
    handleDragEnd,
    handleDragOver,
    handleDrop,
  } = useDragAndDrop(fake.data);

  // const [data] = useState(fake);

  return (
    <TodoLayout>
      <>
        <h1 className="text-5xl font-bold leading-normal mt-2 mb-2 text-gray-900">
          {username}'s todo list
        </h1>
      </>
      <>
        {data.map((item, index) => (
          <Card
            key={index}
            name={item.name}
            items={item.item}
            isDragging={isDragging}
            handleDragStart={handleDragStart}
            handleDragEnd={handleDragEnd}
            handleDragOver={handleDragOver}
            handleDrop={(e) => handleDrop(e, item)}
          />
        ))}
      </>
    </TodoLayout>
  );
};

export default Todo;
