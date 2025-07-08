import { Button } from "@heroui/button";
import { useEffect, useState } from "react";

import {
  HighActivityIcon,
  LowActivityIcon,
  MiddleActivityIcon,
} from "@/icons.tsx";
import { FormProops } from "@/pages/Register.tsx";
import { PhysicalActivity } from "@/api/user.ts";

interface FormProopsWithSendData extends FormProops {
  sendUserData: () => void;
}

export function Activity({
  onChange,
  increaseQuestionId,
  sendUserData,
}: FormProopsWithSendData) {
  const [pendingUpdate, setPendingUpdate] = useState<PhysicalActivity | null>(
    null,
  );

  const handleAnswer = (physical_activity: PhysicalActivity) => {
    if (onChange) {
      onChange((prev) => ({ ...prev, physical_activity }));
    }
    increaseQuestionId();
    setPendingUpdate(physical_activity);
  };

  useEffect(() => {
    if (pendingUpdate !== null) {
      sendUserData();
      setPendingUpdate(null);
    }
  }, [pendingUpdate]);

  return (
    <form className="mt-[6dvh] flex flex-col gap-[4dvh] ml-[5dvw] mr-[5dvw]">
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] py-[5dvh]"
        color="default"
        onPress={() => handleAnswer("Low")}
      >
        <div className="flex w-full">
          <LowActivityIcon />
          <div className="ml-[3dvw] flex flex-col">
            <p className="font-[300] text-[5dvw] text-left">Низкая</p>
            <p className="font-[200] text-[4dvw] mt-auto">
              Сидячий образ жизни
            </p>
          </div>
        </div>
      </Button>
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] py-[5dvh]"
        color="default"
        onPress={() => handleAnswer("Medium")}
      >
        <div className="flex w-full">
          <MiddleActivityIcon />
          <div className="ml-[3dvw] flex flex-col">
            <p className="font-[300] text-[5dvw] text-left">Умеренная</p>
            <p className="font-[200] text-[4dvw] mt-auto">
              Тренировки 2-4 раза в нед.
            </p>
          </div>
        </div>
      </Button>
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] py-[5dvh]"
        color="default"
        onPress={() => handleAnswer("High")}
      >
        <div className="flex w-full">
          <HighActivityIcon />
          <div className="ml-[3dvw] flex flex-col">
            <p className="font-[300] text-[5dvw] text-left">Интенсивная</p>
            <p className="font-[200] text-[4dvw] mt-auto">
              Тренировки 5-7 раз в нед.
            </p>
          </div>
        </div>
      </Button>
    </form>
  );
}
