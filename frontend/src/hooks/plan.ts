import { useQuery } from "react-query";

import { getRecommendationsByTgId, RecommendationType } from "@/api/plan.ts";
import { PlanType } from "@/layouts/Plan.tsx";

export function useRecommendations(
  utgid: number,
  initData: string,
  type: PlanType,
  timestamps: {
    start_at: number;
    finish_at: number;
  },
) {
  const { data, error, isLoading, refetch } = useQuery<
    RecommendationType[],
    Error
  >({
    queryKey: ["recommendations"],
    queryFn: () =>
      getRecommendationsByTgId(
        utgid,
        initData,
        type,
        timestamps["start_at"],
        timestamps["finish_at"],
      ),
    enabled: !!utgid,
  });

  return {
    recData: data,
    recIsLoading: isLoading,
    recError: error,
    recRefetch: refetch,
  };
}
