import { useEffect, useRef, useState } from "react";
import { addTodo } from "../../context/todo/item";
import useAuth from "../../hooks/useAuth";
import service from "../../service";
import { err } from "../../types/err";
import { ItemType } from "../../types/todo";
import isApiError from "../../util/error";

type AddTodoModalProps = {
  open: boolean;
  handleClose: () => void;
  groupId: number;
};

const AddTodoModal = ({ open, handleClose, groupId }: AddTodoModalProps) => {
  const [err, setErr] = useState<err>({
    status: false,
    message: "",
  });
  const {
    user: { id },
  } = useAuth();

  const nameRef = useRef<HTMLInputElement | null>(null);

  useEffect(() => {
    nameRef.current?.focus();

    // logged in logic
  }, []);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const value = nameRef.current?.value;

    // api logic
    try {
      if (value) {
        const response = await service.todo.item.create(id, groupId, value);
        if (response.data) {
          addTodo(groupId, response.data as ItemType);
          handleClose();
        }
      }
    } catch (error) {
      const { isValid, response } = isApiError(error);
      if (isValid) {
        setErr({ status: true, message: response });
      }
    }
  };

  return (
    open && (
      <>
        {/* Main modal  */}
        <div className="bg-gray-200 bg-opacity-70 fixed z-20 w-full h-full top-0 left-0 flex items-center justify-center">
          <div className="relative">
            <button className="absolute right-4 top-4" onClick={handleClose}>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                viewBox="0 0 24 24"
              >
                <path d="M24 20.188l-8.315-8.209 8.2-8.282-3.697-3.697-8.212 8.318-8.31-8.203-3.666 3.666 8.321 8.24-8.206 8.313 3.666 3.666 8.237-8.318 8.285 8.203z" />
              </svg>
            </button>
            <div className="bg-white shadow-md border border-gray-200 rounded-lg w-auto p-4 sm:p-6 lg:p-8 dark:bg-gray-800 dark:border-gray-700">
              <form onSubmit={handleSubmit}>
                <div className="flex flex-col">
                  <label
                    htmlFor="name-id"
                    className="text-md font-medium text-gray-900 block mb-2 dark:text-gray-300"
                  >
                    Todo name
                  </label>
                  <input
                    ref={nameRef}
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm text-md rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                    type="text"
                    name="input-name"
                    id="name-id"
                    placeholder="Input todo name"
                  />
                </div>
                <div className="max-w-2xl">
                  {" "}
                  {err.status && (
                    <p className="text-sm font-medium text-red-900 block mb-2 dark:text-gray-300 mt-2">
                      error: {err.message}
                    </p>
                  )}
                </div>
                <button className="w-full mt-5 text-white bg-gray-700 hover:bg-gray-800 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus:ring-gray-800">
                  Add
                </button>
              </form>
            </div>
          </div>
        </div>
      </>
    )
  );
};

export default AddTodoModal;
