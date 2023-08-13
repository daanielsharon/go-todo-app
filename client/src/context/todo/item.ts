import { ObservablePrimitiveChildFns, observable } from "@legendapp/state";
import { ContainerType, ItemType } from "../../types/todo";

export const todoState = observable({
  data: undefined as unknown as ContainerType[],
});

export const addTodo = (groupId: number, data: ItemType) => {
  const newData = [...todoState.data];
  newData.forEach((item, index) => {
    if (
      item.id === (groupId as unknown as ObservablePrimitiveChildFns<number>)
    ) {
      todoState.data[index].item.push(data);
    }
  });
};

export const updateTodo = (id: number, group_id: number, data: ItemType) => {
  removeTodo(id);
  addTodo(group_id, data);
};

export const removeTodo = (id: number) => {
  const newData = [...todoState.data];
  newData.forEach((item, idx) => {
    item.item.forEach((element, itemIdx) => {
      if (
        element.id === (id as unknown as ObservablePrimitiveChildFns<number>)
      ) {
        todoState.data[idx].item.splice(itemIdx, 1);
      }
    });
  });
};

export const swapContainerPosition = (
  containerOrigin: ContainerType,
  originPriority: number,
  containerTarget: ContainerType,
  // indexTarget: number,
  priorityDestination: number
) => {
  // swap container position
  const newData = [...todoState.data];
  newData.forEach((item, index) => {
    if (JSON.stringify(item) === JSON.stringify(containerOrigin)) {
      item.priority =
        priorityDestination as unknown as ObservablePrimitiveChildFns<number>;
      todoState.data.splice(index, 1);
      todoState.data.splice(index, 0, item as unknown as ContainerType);
    }

    if (JSON.stringify(item) === JSON.stringify(containerTarget)) {
      item.priority =
        originPriority as unknown as ObservablePrimitiveChildFns<number>;
      todoState.data.splice(index, 1);
      todoState.data.splice(index, 0, item as unknown as ContainerType);
    }

    // if (index === indexTarget) {
    //   return (item.priority =
    //     originPriority as unknown as ObservablePrimitiveChildFns<number>);
    // }
  }) as unknown as ContainerType[];

  todoState.data.sort((a, b) => {
    if (a.priority < b.priority) {
      return -1;
    }
    if (a.priority > b.priority) {
      return 1;
    }

    return 0;
  });
};
