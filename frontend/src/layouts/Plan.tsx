import { useState } from "react";
import { Button } from "@heroui/react";

import { PlanHeader } from "@/components/PlanHeader.tsx";
import { PlanNavigation } from "@/components/PlanNavigation.tsx";
import { ActivityIcon, FoodIcon, PlusIcon } from "@/icons.tsx";
import { SemiCircularProgressBar } from "@/components/SemiCircularProgressBar.tsx";
import { Advices } from "@/components/Adviсes.tsx";
import { Story } from "@/components/Story.tsx";
import { useActions } from "@/hooks/actions.ts";
import { getTodayAndTomorrowTimestamps } from "@/utils.ts";
import { TabType } from "@/pages/Main.tsx";
import { useRecommendations } from "@/hooks/plan.ts";

interface PlanProops {
  date: Date;
  onChange: (value: TabType) => void;
}

export type PlanType = "Activity" | "Food";

export function Plan({ date, onChange }: PlanProops) {
  const [active, setActive] = useState<PlanType>("Activity");
  const { data, error, isLoading, refetch } = useActions(
    window.Telegram.WebApp.initDataUnsafe.user.id,
    window.Telegram.WebApp.initData,
    active,
    getTodayAndTomorrowTimestamps(),
  );

  const { recData, recRefetch } = useRecommendations(
    window.Telegram.WebApp.initDataUnsafe.user.id,
    window.Telegram.WebApp.initData,
    active,
    getTodayAndTomorrowTimestamps(),
  );

  function handleChangeActive(tab: PlanType) {
    setActive(tab);
    refetch();
    recRefetch();
  }

  if (isLoading) {
    return <div className="p-[3dvh]">Загрузка...</div>;
  }

  if (error) {
    return <div className="p-[3dvh]">Ошибка!</div>;
  }

  return (
    <div className="p-[3dvh]">
      <PlanHeader date={date} />
      <PlanNavigation onChange={handleChangeActive} />
      <div className="bg-[#272727] rounded-medium px-[3dvw] py-[5dvw]">
        <div className="flex justify-around items-center">
          {active === "Activity" ? <ActivityIcon /> : <FoodIcon />}
          {recData === undefined && (
            <SemiCircularProgressBar
              circleColor="text-[#666666]"
              curr={0}
              end={0}
              progressColor="text-[#D9D9D9]"
              showText={true}
              size={115}
              strokeWidth={12}
              textColor="text-white"
            />
          )}
          {recData !== undefined && (
            <SemiCircularProgressBar
              circleColor="text-[#666666]"
              curr={
                data!!.length > 0
                  ? data!!.reduce((sum: number, item: any) => sum + item.calories, 0)
                  : 0
              }
              end={
                active === "Activity"
                  ? recData!![0].calories_to_burn
                  : recData!![0].calories_to_consume
              }
              progressColor="text-[#D9D9D9]"
              showText={true}
              size={115}
              strokeWidth={12}
              textColor="text-white"
            />
          )}
        </div>
        <Advices advices={recData !== undefined ? recData : []} />
        <div className="flex w-full items-center justify-between mt-[1dvh]">
          <h2 className="text-[#858585] text-[5dvw]">История</h2>
          <Button
            isIconOnly
            className="bg-transparent"
            onPress={() => onChange("AddActivity")}
          >
            <PlusIcon />
          </Button>
        </div>
        <Story actions={data!!} />
      </div>
    </div>
  );
}
