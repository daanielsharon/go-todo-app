import { useRef, useEffect } from "react";

type AddTodoModalProps = {
  open: boolean;
  handleClose: () => void;
};

const AddTodoModal = ({ open, handleClose }: AddTodoModalProps) => {
  const nameRef = useRef<HTMLInputElement | null>(null);

  useEffect(() => {
    nameRef.current?.focus();

    // logged in logic
  }, []);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const value = nameRef.current?.value;
    console.log("value", value);

    // api logic
  };

  return (
    open && (
      <>
        {/* Main modal  */}
        <div className="bg-gray-200 bg-opacity-70 fixed w-full h-full top-0 left-0 flex items-center justify-center">
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
