import React, { useState } from "react";
import { ItemType } from "../types/item";
import { ContainerType } from "../types/container";

const useDragAndDrop = (api: ContainerType[]) => {
  const [isDragging, setIsDragging] = useState<boolean>(false);
  const [data, setData] = useState<ContainerType[]>(api);

  const getContainerIndex = (draggedData: ItemType) => {
    let index = 0;
    data.forEach((item, idx) => {
      item.item.forEach((element) => {
        if (JSON.stringify(element) === JSON.stringify(draggedData)) {
          index = idx;
        }
      });
    });

    return index;
  };

  const handleUpdateList = (
    draggedData: ItemType,
    currentContainer: ContainerType
    // containerTarget: number
  ) => {
    const containerIndex = getContainerIndex(draggedData);

    console.log("containerIndex", containerIndex);
  };

  const handleDragStart = (
    e: React.DragEvent<HTMLDivElement>,
    data: ItemType | null
  ) => {
    // console.info("data", data);
    e.dataTransfer.setData("todo", `${JSON.stringify(data)}`);
    setIsDragging(true);
  };

  const handleDragEnd = () => setIsDragging(false);
  const handleDragOver = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
  };

  const handleDrop = (
    e: React.DragEvent<HTMLDivElement>,
    currentContainer: ContainerType
  ) => {
    e.preventDefault();
    setIsDragging(false);
    handleUpdateList(
      JSON.parse(e.dataTransfer.getData("todo")),
      currentContainer
    );
  };

  return {
    isDragging,
    handleDragStart,
    handleDragEnd,
    handleDragOver,
    handleDrop,
    data,
  };
};

export default useDragAndDrop;
