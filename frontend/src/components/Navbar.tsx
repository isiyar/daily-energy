import { Button } from "@heroui/react";

import { nowDate } from "@/utils.ts";
import { ChatIcon, ProfileIcon } from "@/icons.tsx";
import { TabType } from "@/pages/Main.tsx";

export function Navbar({
  currTab,
  onChange,
}: {
  currTab: TabType;
  onChange: (tab: TabType) => void;
}) {
  return (
    <div className="fixed bottom-0 left-0 right-0 py-[2dvh] bg-[#252525]">
      <nav className="w-[100]dvw flex justify-around items-center">
        <Button
          isIconOnly
          className={`bg-transparent ${currTab === "TodayPlan" || currTab === "AddActivity" ? "text-white" : "text-[#858585]"}`}
          onPress={() => onChange("TodayPlan")}
        >
          {nowDate()[0]}
          <br />
          {nowDate()[1]}
        </Button>
        <Button
          isIconOnly
          className={`bg-transparent ${currTab === "Chat" ? "stroke-white" : "stroke-[#858585]"}`}
          onPress={() => onChange("Chat")}
        >
          <ChatIcon />
        </Button>
        <Button
          isIconOnly
          className={`bg-transparent ${currTab === "Profile" ? "stroke-white" : "stroke-[#858585]"}`}
          onPress={() => onChange("Profile")}
        >
          <ProfileIcon />
        </Button>
      </nav>
    </div>
  );
}
