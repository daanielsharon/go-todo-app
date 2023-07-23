import axios from "axios";

type isApiError = {
  isValid: boolean;
  response: string;
};

const isApiError = (error: unknown): isApiError => {
  if (axios.isAxiosError(error)) {
    console.info("error", error);
    if (error.response) {
      //   console.info("error", error.response.data.data);
      return {
        isValid: true,
        response: error.response.data.data,
      };
    }
  }

  return {
    isValid: false,
    response: "",
  };
};

export default isApiError;
