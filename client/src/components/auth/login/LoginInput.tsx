import React, { useEffect, useRef } from "react";

const LoginInput = () => {
  const nameRef = useRef<HTMLInputElement | null>(null);
  const todoNameRef = useRef<HTMLInputElement | null>(null);

  useEffect(() => {
    nameRef.current?.focus();

    // logged in logic
  }, []);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const value = JSON.stringify({
      name: nameRef.current?.value,
      todoName: todoNameRef.current?.value,
    });
    console.log("value", value);

    // api logic

    localStorage.setItem("todo", value);
  };

  return (
    <div className="bg-white shadow-md border border-gray-200 rounded-lg w-auto p-4 sm:p-6 lg:p-8 dark:bg-gray-800 dark:border-gray-700">
      <form onSubmit={handleSubmit}>
        <div className="flex flex-col">
          <label
            htmlFor="name-id"
            className="text-md font-medium text-gray-900 block mb-2 dark:text-gray-300"
          >
            Username
          </label>
          <input
            ref={nameRef}
            className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm text-md rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
            type="text"
            name="input-name"
            id="name-id"
            placeholder="Input your name"
          />
        </div>
        <button className="w-full mt-5 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
          Sign in
        </button>
      </form>
    </div>
  );
};

export default LoginInput;
