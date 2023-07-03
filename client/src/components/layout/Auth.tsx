import { PropsWithChildren } from "react";

const AuthLayout = ({ children }: PropsWithChildren) => {
  return (
    <div className="w-[100vw] h-[100vh] flex justify-center items-center">
      {children}
    </div>
  );
};

export default AuthLayout;
