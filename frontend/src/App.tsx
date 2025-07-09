import { Route, Routes, useNavigate } from "react-router-dom";
import { useEffect } from "react";
import axios from "axios";

import { Register } from "./pages/Register";

import { Greet } from "@/pages/Greet.tsx";
import { useUser } from "@/hooks/user.ts";
import { Plan } from "@/layouts/Plan.tsx";

function App() {
  const navigate = useNavigate();
  const { data, error } = useUser(
    window.Telegram.WebApp.initDataUnsafe.user.id,
  );

  useEffect(() => {
    if (data) {
      navigate("/");
    } else if (
      error &&
      axios.isAxiosError(error) &&
      error.response?.status === 404
    ) {
      navigate("/greet");
    }
  }, [data, error]);

  return (
    <Routes>
      <Route element={<Greet />} path="/greet" />
      <Route element={<Register />} path="/register" />
      <Route element={<Plan date={new Date()} />} path="/" />
    </Routes>
  );
}

export default App;
