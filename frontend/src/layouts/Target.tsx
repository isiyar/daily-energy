import { Button } from "@heroui/button";

import { Goal } from "@/api/user.ts";
import { FormProops } from "@/pages/Register.tsx";

export function Target({ onChange, increaseQuestionId }: FormProops) {
  function handleAnswer(goal: Goal) {
    if (onChange) {
      onChange((prev) => ({ ...prev, goal }));
    }
    increaseQuestionId();
  }

  return (
    <form className="mt-[6dvh] flex flex-col gap-[4dvh] font-[300] ml-[10dvw] mr-[10dvw]">
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] py-[4dvh] text-[5dvw]"
        color="default"
        onPress={() => handleAnswer("LoseWeight")}
      >
        <div className="flex w-full">
          <div>📈</div>
          <div className="ml-[3dvw] flex items-center">Похудеть</div>
        </div>
      </Button>
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] py-[4dvh] text-[5dvw]"
        color="default"
        onPress={() => handleAnswer("Maintain")}
      >
        <div className="flex w-full">
          <div>😊</div>
          <div className="ml-[3dvw] flex items-center">Поддерживать вес</div>
        </div>
      </Button>
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] py-[4dvh] text-[5dvw]"
        color="default"
        onPress={() => handleAnswer("GainMuscleMass")}
      >
        <div className="flex w-full">
          <div>💪</div>
          <div className="ml-[3dvw] flex items-center">Набрать вес</div>
        </div>
      </Button>
    </form>
  );
}
