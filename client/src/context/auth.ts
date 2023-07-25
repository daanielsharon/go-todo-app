import { observable } from "@legendapp/state";
import { persistObservable } from "@legendapp/state/persist";
import { ObservablePersistLocalStorage } from "@legendapp/state/persist-plugins/local-storage";

export const authState = observable({
  username: undefined as unknown as string,
});

persistObservable(authState, {
  local: "todo",
  persistLocal: ObservablePersistLocalStorage,
});
