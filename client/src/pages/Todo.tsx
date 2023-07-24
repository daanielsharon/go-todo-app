import { useEffect } from "react";
import TodoLayout from "../components/layout/Todo";
import Card from "../components/todo/Card";
import context from "../context";
import useAuth from "../hooks/useAuth";
import useDragAndDrop from "../hooks/useDragAndDrop";
import service from "../service";
import { ContainerType } from "../types/todo";

const Todo = () => {
  const {
    user: { username },
  } = useAuth();

  const {
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
      }
    };
    fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const todoData = context.getContextUpdate("todo", "data");
  console.log("todoData", todoData);

  return (
    <TodoLayout>
      <>
        <h1 className="text-5xl font-bold leading-normal mt-2 mb-2 text-gray-900">
          {username}'s todo list
        </h1>
      </>
      <>
        {todoData &&
          todoData.map((item: ContainerType, index: number) => (
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
