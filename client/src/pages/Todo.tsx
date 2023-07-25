import { observer } from "@legendapp/state/react";
import { useEffect } from "react";
import TodoLayout from "../components/layout/Todo";
import Card from "../components/todo/Card";
import context from "../context";
import useAuth from "../hooks/useAuth";
import useDragAndDrop from "../hooks/useDragAndDrop";
import service from "../service";
import { ContainerType } from "../types/todo";

const Todo = observer(() => {
  const todoData = context.getContext("todo", "data");

  const {
    user: { username, id },
  } = useAuth();

  const {
    isDragging,
    handleDragStart,
    handleDragEnd,
    handleDragOver,
    handleDrop,
  } = useDragAndDrop(todoData);

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
              handleDrop={(draggedData) =>
                // index + 1 since database data starts from 1
                handleDrop(draggedData, index + 1, id)
              }
            />
          ))}
      </>
    </TodoLayout>
  );
});

export default Todo;
