import { useMutation, useQuery } from "react-query";
import { useNavigate } from "react-router-dom";

import { createUser, getUserByTgId, User } from "@/api/user.ts";

export function useUser(utgid: number) {
  return useQuery<User, Error>({
    queryKey: ["user"],
    queryFn: () => getUserByTgId(utgid),
    enabled: !!utgid,
  });
}

export function useRegister() {
  const navigate = useNavigate();

  return useMutation({
    mutationFn: createUser,
    onSuccess: () => {
      navigate("/");
    },
    onError: () => {
      navigate("/greet");
    },
  });
}
