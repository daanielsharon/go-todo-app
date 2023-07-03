const useAuth = () => {
  const username = localStorage.getItem("todo");

  return username ? { username } : { username: null };
};

export default useAuth;
