import { Response } from "../types/api";
import { http } from "./config";

class Api {
  public async get(url: string): Promise<Response> {
    const response = await http.get(url);
    return response.data;
  }

  public async post(url: string, data: undefined): Promise<Response> {
    const response = await http.post(url, data);
    return response.data;
  }

  public async put(url: string, data: undefined): Promise<Response> {
    const response = await http.put(url, data);
    return response.data;
  }

  public async delete(url: string, data: undefined): Promise<boolean> {
    const response = await http.delete(url, data);
    const result: Response = response.data;

    return result.code === 200 ? true : false;
  }
}

export const api = new Api();
