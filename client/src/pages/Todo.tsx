import { useState } from "react";
import TodoLayout from "../components/layout/Todo";
import Card from "../components/todo/Card";
import useAuth from "../hooks/useAuth";
import fake from "../data/card.json";

const Todo = () => {
  const { username } = useAuth();
  const capitalizedUsername =
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    username!.charAt(0).toUpperCase() + username!.slice(1) + "'s todo list";

  const [data] = useState(fake);

  return (
    <TodoLayout>
      <>
        <h1 className="text-5xl font-bold leading-normal mt-2 mb-2 text-gray-900">
          {capitalizedUsername}
        </h1>
      </>
      <>
        {data.data.map((item) => (
          <Card name={item.name} items={item.item} />
        ))}
      </>
    </TodoLayout>
  );
};

export default Todo;
