import { authState } from "./auth";
import { todoState } from "./todo";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type stateType = Record<string, any>;

const state$: stateType = {
  auth: authState,
  todo: todoState,
};

const getContext = (stateName: string, content: string) => {
  return state$[stateName][content].get();
};

const setContext = (stateName: string, content: string, replace: unknown) => {
  return state$[stateName][content].set(replace);
};

export default {
  getContext,
  setContext,
};
