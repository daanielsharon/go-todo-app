import { useEffect, useState } from "react";
import TodoLayout from "../components/layout/Todo";
import Card from "../components/todo/Card";
import useAuth from "../hooks/useAuth";
import useDragAndDrop from "../hooks/useDragAndDrop";
import service from "../service";
import context from "../context";
import { todoState } from "../context/todo";
import { ContainerType } from "../types/todo";

const Todo = () => {
  const [data, setData] = useState<ContainerType[]>([]);
  const {
    user: { username },
  } = useAuth();

  const {
    // data,
    isDragging,
    handleDragStart,
    handleDragEnd,
    handleDragOver,
    handleDrop,
  } = useDragAndDrop();

  useEffect(() => {
    const fetchData = async () => {
      const response = await service.todo.get(username);
      if (Array.isArray(response.data)) {
        context.setContext("todo", "data", response.data);
        setData(response.data);
      }
    };
    fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // const todoData = context.getContext("todo", "data");
  const todoData = todoState.data.get();
  console.log("todoData", todoData);

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
            index={index}
            items={item.item}
            groupId={item.id}
            name={item.group_name}
            isDragging={isDragging}
            handleDragStart={handleDragStart}
            handleDragEnd={handleDragEnd}
            handleDragOver={handleDragOver}
            handleDrop={(e) => handleDrop(e, index)}
          />
        ))}
      </>
    </TodoLayout>
  );
};

export default Todo;
