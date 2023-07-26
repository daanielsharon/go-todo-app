import { observer } from "@legendapp/state/react";
import { useEffect } from "react";
import TodoLayout from "../components/layout/Todo";
import Card from "../components/todo/Card";
import context from "../context";
import useAuth from "../hooks/useAuth";
import useDragAndDrop from "../hooks/useDragAndDrop";
import service from "../service";
import { ContainerType } from "../types/todo";
import { useNavigate } from "react-router";

const Todo = observer(() => {
  const todoData = context.getContext("todo", "data");

  const {
    user: { username, id },
  } = useAuth();

  const navigateTo = useNavigate();

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
        <div className="flex flex-row items-center gap-5">
          <h1 className="text-5xl font-bold leading-normal mt-2 mb-2 text-gray-900">
            {username}'s todo list
          </h1>
          <button
            className="text-white bg-gray-700 hover:bg-gray-800 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus:ring-gray-800"
            onClick={() => {
              localStorage.removeItem("todo");
              service.auth.logout();
              navigateTo("/");
            }}
          >
            Logout
          </button>
        </div>
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
              handleDrop={(draggedData) => handleDrop(draggedData, item.id, id)}
            />
          ))}
      </>
    </TodoLayout>
  );
});

export default Todo;
