const TodoLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="w-[100vw] h-[100vh] flex flex-col justify-center items-center">
      {children && children[0 as keyof typeof children]}
      <div className="w-[75%] h-[50%]">
        <div className="cards">
          {children && children[1 as keyof typeof children]}
        </div>
      </div>
    </div>
  );
};

export default TodoLayout;
