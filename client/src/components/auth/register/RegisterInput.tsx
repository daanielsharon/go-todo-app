import React, { useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router-dom";
import { api } from "../../../api/api";
import { err } from "../../../types/err";
import axios from "axios";

const LoginInput = () => {
  const [err, setErr] = useState<err>({
    status: false,
    message: "",
  });
  const [inputValue, setInputValue] = useState<string>("");
  const nameRef = useRef<HTMLInputElement | null>(null);
  const navigateTo = useNavigate();

  useEffect(() => {
    nameRef.current?.focus();
  }, []);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (inputValue) {
      try {
        const response = await api.post("/users/register", {
          username: inputValue,
        });

        if (response.code === 200) {
          return navigateTo("/login");
        }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          console.info("error", error);
          if (error.response) {
            console.info("error", error.response.data.data);
            setErr({ status: true, message: error.response.data.data });
          }
        }
      }
    } else {
      setErr({ status: true, message: "username is empty" });
    }
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
            onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
              setErr({ status: false, message: "" });
              setInputValue(e.target.value);
            }}
          />
        </div>
        {err.status && (
          <p className="text-sm font-medium text-red-900 block mb-2 dark:text-gray-300 mt-2">
            error: {err.message}
          </p>
        )}
        <button className="w-full mt-5 text-white bg-gray-700 hover:bg-gray-800 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus:ring-gray-800">
          Sign up
        </button>
      </form>
    </div>
  );
};

export default LoginInput;
