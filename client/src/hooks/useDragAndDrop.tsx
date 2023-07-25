import React, { useState } from "react";
import { ContainerType, ItemType } from "../types/todo";
import { updateTodo } from "../context/todo";
import service from "../service";

const useDragAndDrop = (data: ContainerType[] = []) => {
  const [isDragging, setIsDragging] = useState<boolean>(false);

  const handleUpdateList = async (
    draggedData: ItemType,
    containerTarget: number,
    userId: number
  ): Promise<void> => {
    const res = await service.todo.update(
      draggedData.id,
      draggedData.name,
      containerTarget,
      userId
    );

    if (res.data) {
      updateTodo(draggedData.id, containerTarget, draggedData);
    }
  };

  const handleDragStart = (
    e: React.DragEvent<HTMLDivElement>,
    data: ItemType | null
  ): void => {
    // console.info("data", data);
    e.dataTransfer.setData("todo", `${JSON.stringify(data)}`);
    setIsDragging(true);
  };

  const handleDragEnd = (): void => setIsDragging(false);
  const handleDragOver = (e: React.DragEvent<HTMLDivElement>): void => {
    e.preventDefault();
  };

  const handleDrop = (
    e: React.DragEvent<HTMLDivElement>,
    containerTarget: number,
    userId: number
  ): void => {
    e.preventDefault();
    setIsDragging(false);
    handleUpdateList(
      JSON.parse(e.dataTransfer.getData("todo")),
      containerTarget,
      userId
    );
  };

  return {
    data,
    isDragging,
    handleDragStart,
    handleDragEnd,
    handleDragOver,
    handleDrop,
  };
};

export default useDragAndDrop;
