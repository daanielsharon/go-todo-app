import { PropsWithChildren } from "react";
import classes from "./auth.module.css";

const AuthLayout = ({ children }: PropsWithChildren) => {
  return <div className={classes.wrapper}>{children}</div>;
};

export default AuthLayout;
