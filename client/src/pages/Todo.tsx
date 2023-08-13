import { observer } from "@legendapp/state/react";
import { useEffect, useState } from "react";
import TodoLayout from "../components/layout/Todo";
import Card from "../components/todo/Card";
import context from "../context";
import useAuth from "../hooks/useAuth";
import useDragAndDrop from "../hooks/useDragAndDrop";
import service from "../service";
import { ContainerType } from "../types/todo";
import { useNavigate } from "react-router";
import AddTodoContainer from "../components/modal/AddTodoContainer";

const Todo = observer(() => {
  const todoData = context.getContextUpdate("todo", "data");
  const [open, setIsOpen] = useState(false);

  const {
    user: { username, id },
  } = useAuth();

  const navigateTo = useNavigate();

  const {
    isDragging,
    isContainerDragging,
    handleContainerStartDragging,
    handleItemDrag,
    handleDragEnd,
    handleDragOver,
    handleItemDrop,
    handleContainerDrag,
    handleContainerDrop,
    handleContainerDragEnd,
  } = useDragAndDrop(todoData);

  useEffect(() => {
    const fetchData = async () => {
      const response = await service.todo.item.get(username);
      if (Array.isArray(response.data)) {
        context.setContext("todo", "data", response.data);
      }
    };
    fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const getLatestPriority = (): number => {
    let maxPriority = 0;
    todoData &&
      todoData.forEach((item: ContainerType) => {
        if (item.priority > maxPriority) {
          maxPriority = item.priority;
        }
      });
    return maxPriority;
  };

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
            <div
              key={index}
              onDragOver={handleDragOver}
              onDrop={(e) => {
                if (isContainerDragging.status) {
                  handleContainerDrop(e, index, todoData);
                }
              }}
            >
              <Card
                item={item}
                index={index}
                groupId={item.id}
                items={item.item}
                name={item.group_name}
                isDragging={isDragging}
                handleDragEnd={handleDragEnd}
                handleDragOver={handleDragOver}
                handleItemDrag={handleItemDrag}
                isContainerDragging={isContainerDragging}
                handleContainerDrag={handleContainerDrag}
                handleContainerDragEnd={handleContainerDragEnd}
                handleContainerStartDragging={handleContainerStartDragging}
                handleItemDrop={(draggedData) => {
                  if (!isContainerDragging.status) {
                    handleItemDrop(draggedData, item.id, id);
                  }
                }}
              />
            </div>
          ))}
      </>
      <>
        {todoData && todoData.length !== 4 && (
          <button
            className="text-white card-container-button bg-black"
            aria-label="add todo"
            aria-labelledby="div"
            title="Add"
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
        <AddTodoContainer
          open={open}
          handleClose={() => setIsOpen(false)}
          priority={getLatestPriority()}
        />
      </>
    </TodoLayout>
  );
});

export default Todo;
