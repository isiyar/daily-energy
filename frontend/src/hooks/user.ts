import { useQuery } from "react-query";

import { getUserByTgId, User } from "@/api/user.ts";

export function useUser(utgid: number) {
  return useQuery<User, Error>({
    queryKey: ["user"],
    queryFn: () => getUserByTgId(utgid),
    enabled: !!utgid,
  });
}
