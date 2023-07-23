import axios from "axios";

type isApiError = {
  isValid: boolean;
  response: string;
};

const isApiError = (error: unknown): isApiError => {
  if (axios.isAxiosError(error)) {
    if (error.response) {
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
