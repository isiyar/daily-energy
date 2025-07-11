import { useQuery } from "react-query";

import { ActionType, getActionsByUtgid } from "@/api/actions.ts";
import { PlanType } from "@/layouts/Plan.tsx";

export function useActions(
  utgid: number,
  initData: string,
  type: PlanType,
  timestamps: {
    start_at: number;
    finish_at: number;
  },
) {
  return useQuery<ActionType[], Error>({
    queryKey: ["actions_by_utgid"],
    queryFn: () =>
      getActionsByUtgid(
        utgid,
        initData,
        type,
        timestamps["start_at"],
        timestamps["finish_at"],
      ),
    enabled: !!utgid,
  });
}
