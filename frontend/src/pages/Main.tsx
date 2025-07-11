import { useState } from "react";

import { Plan } from "@/layouts/Plan.tsx";
import { Navbar } from "@/components/Navbar.tsx";
import { AddActivity } from "@/layouts/AddActivity.tsx";

export type TabType = "TodayPlan" | "Chat" | "Profile" | "AddActivity";

export function Main() {
  const [currTab, setCurrTab] = useState<TabType>("TodayPlan");

  return (
    <div>
      {currTab === "TodayPlan" && (
        <Plan date={new Date()} onChange={setCurrTab} />
      )}
      {currTab === "Chat" && <h1>чат</h1>}
      {currTab === "Profile" && <h1>профиль</h1>}
      {currTab === "AddActivity" && <AddActivity />}
      <Navbar currTab={currTab} onChange={setCurrTab} />
    </div>
  );
}
