import { Route, Routes, useNavigate } from "react-router-dom";
import { useEffect } from "react";
import axios from "axios";

import { Register } from "./pages/Register";

import { Greet } from "@/pages/Greet.tsx";
import { useUser } from "@/hooks/user.ts";
import { Main } from "@/pages/Main.tsx";

function App() {
  const navigate = useNavigate();
  const { data, error } = useUser(
    window.Telegram.WebApp.initDataUnsafe.user.id,
    window.Telegram.WebApp.initData,
  );

  useEffect(() => {
    if (data) {
      navigate("/");
    } else if (
      error &&
      axios.isAxiosError(error) &&
      error.response?.status !== 200
    ) {
      navigate("/greet");
    }
  }, [data, error]);

  return (
    <Routes>
      <Route element={<Greet />} path="/greet" />
      <Route element={<Register />} path="/register" />
      <Route element={<Main />} path="/" />
    </Routes>
  );
}

export default App;
