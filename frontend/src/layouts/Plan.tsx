import { useState } from "react";

import { PlanHeader } from "@/components/PlanHeader.tsx";
import { PlanNavigation } from "@/components/PlanNavigation.tsx";
import { ActivityIcon, FoodIcon } from "@/icons.tsx";
import { SemiCircularProgressBar } from "@/components/SemiCircularProgressBar.tsx";
import { Advices } from "@/components/Adviсes.tsx";

interface PlanProops {
  date: Date;
}

export type PlanType = "Activity" | "Food";

export function Plan({ date }: PlanProops) {
  const [active, setActive] = useState<PlanType>("Activity");

  return (
    <div className="p-[3dvh]">
      <PlanHeader date={date} />
      <PlanNavigation onChange={setActive} />
      <div className="bg-[#272727] rounded-medium px-[3dvw] py-[5dvw]">
        <div className="flex justify-around items-center">
          {active === "Activity" ? <ActivityIcon /> : <FoodIcon />}
          <SemiCircularProgressBar
            circleColor="text-[#666666]"
            progress={20}
            progressColor="text-[#D9D9D9]"
            showText={true}
            size={115}
            strokeWidth={12}
            textColor="text-white"
          />
        </div>
        <Advices advices={["Турники", "Отжимания", "Бег"]} />
      </div>
    </div>
  );
}
