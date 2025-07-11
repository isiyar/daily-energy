import axios from "axios";

import { API_URL } from "@/constants.ts";
import { PlanType } from "@/layouts/Plan.tsx";

export type RecommendationType = {
  id: string;
  utgid: number;
  date: number;
  calories_to_consume: number;
  calories_to_burn: number;
  recommendation: string;
  type: PlanType;
};

export async function getRecommendationsByTgId(
  utgid: number,
  initData: string,
  type: PlanType,
  start_at: number,
  finish_at: number,
): Promise<RecommendationType[]> {
  const { data } = await axios.get<RecommendationType[]>(
    `${API_URL}/api/users/${utgid}/plans`,
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
