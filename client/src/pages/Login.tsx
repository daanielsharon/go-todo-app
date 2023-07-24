import { useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router";
import LoginInput from "../components/auth/login/LoginInput";
import AuthLayout from "../components/layout/Auth";
import useAuth from "../hooks/useAuth";
import service from "../service";
import { err } from "../types/err";
import isApiError from "../util/error";

const Login = () => {
  const [err, setErr] = useState<err>({
    status: false,
    message: "",
  });
  const [inputValue, setInputValue] = useState<string>("");
  const nameRef = useRef<HTMLInputElement | null>(null);
  const navigateTo = useNavigate();
  const {
    user: { username },
    setSession,
  } = useAuth();

  useEffect(() => {
    if (username) return navigateTo("/todo");

    nameRef.current?.focus();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (inputValue) {
      try {
        const response = await service.auth.login({ username: inputValue });

        if (response.data) {
          setSession(response.data);
          return navigateTo("/todo");
        }
      } catch (error) {
        const { isValid, response } = isApiError(error);
        if (isValid) {
          setErr({ status: true, message: response });
        }
      }

      return;
    }

    setErr({ status: true, message: "username is empty" });
  };

  return (
    <AuthLayout>
      <LoginInput
        err={err}
        setErr={setErr}
        setInputValue={setInputValue}
        nameRef={nameRef}
        handleSubmit={handleLogin}
      />
    </AuthLayout>
  );
};

export default Login;
