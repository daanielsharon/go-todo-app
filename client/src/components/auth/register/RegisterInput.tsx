import React, { useEffect, useRef } from "react";
import classes from "./registerInput.module.css";

const RegisterInput = () => {
  const nameRef = useRef<HTMLInputElement | null>(null);
  const todoNameRef = useRef<HTMLInputElement | null>(null);

  useEffect(() => {
    nameRef.current?.focus();
  }, []);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const value = JSON.stringify({
      name: nameRef.current?.value,
      todoName: todoNameRef.current?.value,
    });
    console.log("value", value);

    localStorage.setItem("todo", value);
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className={classes.wrapper}>
        <div className={classes.inputs}>
          <label htmlFor="name-id" className={classes.labels}>
            Name
          </label>
          <input
            ref={nameRef}
            type="text"
            name="input-name"
            id="name-id"
            placeholder="Input your name"
          />
        </div>
        <div className={classes.inputs}>
          <label htmlFor="todoName-id" className={classes.labels}>
            Todo Name
          </label>
          <input
            ref={todoNameRef}
            type="text"
            name="input-name"
            id="todoName-id"
            placeholder="Input your todo name"
          />
        </div>
      </div>
      <button>Sign Up</button>
    </form>
  );
};

export default RegisterInput;
