import crypto from "crypto-js";

const useAuth = () => {
  const sessionInfo = sessionStorage.getItem("todo");

  if (sessionInfo) {
    const data = JSON.parse(sessionInfo);
    if (data) {
      return JSON.parse(
        crypto.AES.decrypt(
          data.username,
          import.meta.env.VITE_PASSWORD
        ).toString(crypto.enc.Utf8)
      );
    }
  }

  return "";
};

export default useAuth;
