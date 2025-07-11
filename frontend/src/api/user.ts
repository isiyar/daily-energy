import axios from "axios";

import { API_URL } from "@/constants.ts";

export type Gender = "Male" | "Female";
export type Goal = "LoseWeight" | "Maintain" | "GainMuscleMass";
export type PhysicalActivity = "Low" | "Medium" | "High";

export type User = {
  utgid: number;
  gender: Gender;
  date_of_birth: number;
  weight: number;
  height: number;
  goal: Goal;
  physical_activity: PhysicalActivity;
  name: string;
};

export async function getUserByTgId(
  utgid: number,
  initData: string,
): Promise<User> {
  const { data } = await axios.get<User>(`${API_URL}/api/users/${utgid}`, {
    headers: {
      initData: initData,
    },
  });

  return data;
}

export async function createUser(
  user: Partial<User>,
  initData: string,
): Promise<number> {
  const response = await axios.post(`${API_URL}/api/users`, user, {
    headers: {
      initData: initData,
    },
  });

  return response.status;
}
