import { useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router";
import RegisterInput from "../components/auth/register/RegisterInput";
import AuthLayout from "../components/layout/Auth";
import service from "../service";
import { err } from "../types/err";
import isApiError from "../util/error";

const Register = () => {
  const [err, setErr] = useState<err>({
    status: false,
    message: "",
  });
  const [inputValue, setInputValue] = useState<AuthType>({
    username: "",
    password: "",
  });
  const nameRef = useRef<HTMLInputElement | null>(null);
  const navigateTo = useNavigate();

  useEffect(() => {
    nameRef.current?.focus();
  }, []);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (inputValue) {
      try {
        const response = await service.auth.register({
          username: inputValue.username,
          password: inputValue.password,
        });

        if (response.code === 200) {
          return navigateTo("/login");
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
      <RegisterInput
        nameRef={nameRef}
        err={err}
        setErr={setErr}
        setInputValue={setInputValue}
        handleSubmit={handleSubmit}
      />
    </AuthLayout>
  );
};

export default Register;
