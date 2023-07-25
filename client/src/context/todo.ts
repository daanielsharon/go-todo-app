import { observable } from "@legendapp/state";
import { ContainerType, ItemType } from "../types/todo";

export const todoState = observable({
  data: undefined as unknown as ContainerType[],
});

export const addTodo = (groupId: unknown, data: ItemType) => {
  const newData = [...todoState.data];
  newData.forEach((item, index) => {
    if (item.id === groupId) {
      todoState.data[index].item.push(data);
    }
  });
};

export const updateTodo = (id: unknown, group_id: unknown, data: ItemType) => {
  // const newData = [...todoState.data];
  // newData.forEach((item, index) => {
  //   if (item.id === group_id) {
  //     todoState.data[index].item.push({
  //       id: id as number,
  //       name: name as string,
  //     });
  //   }
  //   item.item.forEach((el, idx) => {
  //     if (el.id === id) {
  //       todoState.data[index].item.splice(idx, 1);
  //     }
  //   });
  // });

  removeTodo(id);
  addTodo(group_id, data);
};

export const removeTodo = (id: unknown) => {
  const newData = [...todoState.data];
  newData.forEach((item, idx) => {
    item.item.forEach((element, itemIdx) => {
      if (element.id === id) {
        todoState.data[idx].item.splice(itemIdx, 1);
      }
    });
  });
};
