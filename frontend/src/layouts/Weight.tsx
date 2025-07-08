import { useState } from "react";
import { Button } from "@heroui/button";

import { NumberSlider } from "@/components/NumberSlider.tsx";
import { FormProops } from "@/pages/Register.tsx";

export function Weight({ onChange, increaseQuestionId }: FormProops) {
  const [weight, setCurrentWeight] = useState(0);

  function handleAnswer(weight: number) {
    if (onChange) {
      onChange((prev) => ({ ...prev, weight }));
    }
    increaseQuestionId();
  }

  return (
    <form>
      <NumberSlider
        numbers={Array.from({ length: 171 }, (_, i) => i + 30)}
        setValue={setCurrentWeight}
      />
      <Button
        className={`float-right text-[6dvw] p-[5dvw] bg-[#F08629] text-white mt-[15dvh]`}
        color="warning"
        size="lg"
        onPress={() => handleAnswer(weight)}
      >
        →
      </Button>
    </form>
  );
}
