import { useState } from "react";

import { Plan } from "@/layouts/Plan.tsx";
import { Navbar } from "@/components/Navbar.tsx";

export type TabType = "TodayPlan" | "Chat" | "Profile";

export function Main() {
  const [currTab, setCurrTab] = useState<TabType>("TodayPlan");

  return (
    <div>
      {currTab === "TodayPlan" && <Plan date={new Date()} />}
      {currTab === "Chat" && <h1>чат</h1>}
      {currTab === "Profile" && <h1>профиль</h1>}
      <Navbar onChange={setCurrTab} />
    </div>
  );
}
