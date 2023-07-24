import React, { useState } from "react";
import { ContainerType, ItemType } from "../types/todo";

const useDragAndDrop = () => {
  const [isDragging, setIsDragging] = useState<boolean>(false);
  const [data, setData] = useState<ContainerType[]>([]);

  const getContainerIndex = (draggedData: ItemType): number => {
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

  const getDraggedDataIndex = (
    currentContainerIndex: number,
    draggedData: ItemType
  ): number => {
    let index = 0;
    data[currentContainerIndex].item.forEach((item, idx) => {
      if (JSON.stringify(item) == JSON.stringify(draggedData)) {
        index = idx;
      }
    });

    return index;
  };

  const handleUpdateList = (
    draggedData: ItemType,
    containerTarget: number
  ): void => {
    const currentContainerIndex = getContainerIndex(draggedData);
    const draggedDataIndex = getDraggedDataIndex(
      currentContainerIndex,
      draggedData
    );
    const newData: ContainerType[] = [...data];

    // remove dragged data from currentContainerIndex
    newData[currentContainerIndex].item.splice(draggedDataIndex, 1);

    // add dragged data to the container target
    newData[containerTarget].item.push(draggedData);

    setData(newData);
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
    containerTarget: number
  ): void => {
    e.preventDefault();
    setIsDragging(false);
    handleUpdateList(
      JSON.parse(e.dataTransfer.getData("todo")),
      containerTarget
    );
  };

  const handleChange = (data: ContainerType[]) => {
    setData(data);
  };

  return {
    data,
    isDragging,
    handleChange,
    handleDragStart,
    handleDragEnd,
    handleDragOver,
    handleDrop,
  };
};

export default useDragAndDrop;
