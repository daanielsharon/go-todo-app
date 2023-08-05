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
  const todoData = context.getContextUpdate("todo", "data");

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
    </TodoLayout>
  );
});

export default Todo;
