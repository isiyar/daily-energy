import axios from "axios";

import { API_URL } from "@/constants.ts";

type Gender = "Male" | "Female";
type Goal = "LoseWeight" | "GainMuscleMass";
type PhysicalActivity = "Low" | "Medium" | "High";

export type User = {
  utgid: number;
  gender: Gender;
  weight: number;
  height: number;
  goal: Goal;
  physical_activity: PhysicalActivity;
  name: string;
};

export async function getUserByTgId(utgid: number) {
  const { data } = await axios.get<User>(`${API_URL}/api/users/${utgid}`);

  return data;
}
