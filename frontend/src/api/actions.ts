import axios from "axios";

import { PlanType } from "@/layouts/Plan.tsx";
import { API_URL } from "@/constants.ts";

export type ActionType = {
  id: string;
  utgid: number;
  date: number;
  activity_name: string;
  calories: number;
  type: PlanType;
};

export async function getActionsByUtgid(
  utgid: number,
  initData: string,
  type: PlanType,
  start_at: number,
  finish_at: number,
): Promise<ActionType[]> {
  const { data } = await axios.get<ActionType[]>(
    `${API_URL}/api/users/${utgid}/actions`,
    {
      headers: {
        initData: initData,
      },
      params: {
        type: type,
        start_at: start_at,
        finish_at: finish_at,
      },
    },
  );

  return data;
}
