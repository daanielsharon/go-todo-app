import crypto from "crypto-js";
import context from "../context";

const useAuth = () => {
  // falsy values
  let user = { username: "", id: 0 };
  const sessionInfo = localStorage.getItem("todo");

  if (sessionInfo) {
    const username = JSON.parse(sessionInfo);
    if (username) {
      user = JSON.parse(
        crypto.AES.decrypt(
          username.username,
          import.meta.env.VITE_PASSWORD
        ).toString(crypto.enc.Utf8)
      );
    }
  }

  const setSession = (data: unknown) => {
    const userInfo = JSON.stringify(data);

    const encryptedUsername = crypto.AES.encrypt(
      userInfo,
      import.meta.env.VITE_PASSWORD
    );

    context.setContext("auth", "username", encryptedUsername.toString());
    // authState.username.set(encryptedUsername.toString());
  };

  return {
    user,
    setSession,
  };
};

export default useAuth;
