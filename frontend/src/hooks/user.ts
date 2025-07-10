import { useMutation, useQuery } from "react-query";
import { useNavigate } from "react-router-dom";

import { createUser, getUserByTgId, User } from "@/api/user.ts";

export function useUser(utgid: number, initData: string) {
  return useQuery<User, Error>({
    queryKey: ["user"],
    queryFn: () => getUserByTgId(utgid, initData),
    enabled: !!utgid,
  });
}

export function useRegister() {
  const navigate = useNavigate();

  return useMutation({
    mutationFn: (params: { user: Partial<User>; initData: string }) =>
      createUser(params.user, params.initData),
    onSuccess: () => {
      navigate("/");
    },
    onError: () => {
      navigate("/greet");
    },
  });
}
