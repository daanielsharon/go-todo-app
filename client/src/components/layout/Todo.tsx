const TodoLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="w-[100vw] h-[100vh] flex flex-col justify-center items-center">
      {children && children[0 as keyof typeof children]}
      <div className="max-h-min flex flex-row">
        <div className="cards">
          {children && children[1 as keyof typeof children]}
        </div>
        <div className="add-card">
          {children && children[2 as keyof typeof children]}
        </div>
      </div>
    </div>
  );
};

export default TodoLayout;
