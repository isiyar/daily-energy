import { Button } from "@heroui/react";

import { nowDate } from "@/utils.ts";
import { ChatIcon, ProfileIcon } from "@/icons.tsx";
import { TabType } from "@/pages/Main.tsx";

export function Navbar({ onChange }: { onChange: (tab: TabType) => void }) {
  return (
    <div className="fixed bottom-0 left-0 right-0 py-[2dvh]">
      <nav className="w-[100]dvw flex justify-around items-center">
        <Button
          isIconOnly
          className="bg-transparent"
          onPress={() => onChange("TodayPlan")}
        >
          {nowDate()[0]}
          <br />
          {nowDate()[1]}
        </Button>
        <Button
          isIconOnly
          className="bg-transparent"
          onPress={() => onChange("Chat")}
        >
          <ChatIcon />
        </Button>
        <Button
          isIconOnly
          className="bg-transparent"
          onPress={() => onChange("Profile")}
        >
          <ProfileIcon />
        </Button>
      </nav>
    </div>
  );
}
