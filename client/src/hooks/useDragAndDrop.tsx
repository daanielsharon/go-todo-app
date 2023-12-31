import React, { useState } from "react";
import { ContainerDrag, ContainerType, ItemType } from "../types/todo";
import { swapContainerPosition, updateTodo } from "../context/todo/item";
import service from "../service";

const useDragAndDrop = (data: ContainerType[] = []) => {
  const [isDragging, setIsDragging] = useState<boolean>(false);
  const [isContainerDragging, setIsContainerDragging] = useState<ContainerDrag>(
    {
      status: false,
      containerIndex: 0,
    }
  );

  const handleContainerDragEnd = (): void => {
    setIsContainerDragging((prev) => ({ ...prev, status: false }));
  };

  const handleContainerUpdate = async (
    containerOrigin: ContainerType,
    containerTarget: ContainerType,
    priorityDestination: number
  ): Promise<void> => {
    const originPriority = containerOrigin.priority;
    const res = await service.todo.container.update(
      containerOrigin.id,
      priorityDestination,
      containerTarget.id,
      originPriority
    );
    if (res.data) {
      swapContainerPosition(
        containerOrigin,
        originPriority,
        containerTarget,
        priorityDestination
      );
    }
  };

  const handleContainerDrag = (
    e: React.DragEvent<HTMLDivElement>,
    index: number,
    data: ContainerType
  ): void => {
    setIsContainerDragging({
      containerIndex: index,
      status: true,
    });
    e.dataTransfer.setData("containerOrigin", JSON.stringify(data));
  };

  const handleContainerDrop = (
    e: React.DragEvent<HTMLDivElement>,
    indexTarget: number,
    todo: ContainerType[]
  ): void => {
    const containerOrigin = JSON.parse(
      e.dataTransfer.getData("containerOrigin")
    );

    // priority starts from 1, index starts from 0, so that's why it's added 1
    const priorityDestination = indexTarget + 1;
    const targetId = todo[indexTarget];

    handleContainerUpdate(containerOrigin, targetId, priorityDestination);
  };

  const handleItemUpdate = async (
    draggedData: ItemType,
    containerTarget: number,
    userId: number
  ): Promise<void> => {
    const res = await service.todo.item.update(
      draggedData.id,
      draggedData.name,
      containerTarget,
      userId
    );

    if (res.data) {
      updateTodo(draggedData.id, containerTarget, draggedData);
    }
  };

  const handleItemDrag = (
    e: React.DragEvent<HTMLDivElement>,
    data: ItemType | null
  ): void => {
    e.dataTransfer.setData("todo", `${JSON.stringify(data)}`);
    setIsDragging(true);
  };

  const handleItemDrop = (
    e: React.DragEvent<HTMLDivElement>,
    containerTarget: number,
    userId: number
  ): void => {
    e.preventDefault();
    setIsDragging(false);
    handleItemUpdate(
      JSON.parse(e.dataTransfer.getData("todo")),
      containerTarget,
      userId
    );
  };

  const handleDragEnd = (): void => setIsDragging(false);
  const handleDragOver = (e: React.DragEvent<HTMLDivElement>): void => {
    e.preventDefault();
  };

  const handleContainerStartDragging = (index: number): void =>
    setIsContainerDragging((prev) => ({
      containerIndex: index,
      status: !prev.status,
    }));

  return {
    data,
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
  };
};

export default useDragAndDrop;
