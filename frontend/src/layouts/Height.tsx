import { useState } from "react";
import { Button } from "@heroui/button";

import { NumberSlider } from "@/components/NumberSlider.tsx";
import { FormProops } from "@/pages/Register.tsx";

export function Height({ onChange, increaseQuestionId }: FormProops) {
  const [height, setCurrentHeight] = useState(0);

  function handleAnswer(height: number) {
    if (onChange) {
      onChange((prev) => ({ ...prev, height }));
    }
    increaseQuestionId();
  }

  return (
    <form>
      <NumberSlider
        numbers={Array.from({ length: 291 }, (_, i) => i + 10)}
        setValue={setCurrentHeight}
      />
      <Button
        className={`float-right text-[6dvw] p-[5dvw] bg-[#F08629] text-white mt-[15dvh]`}
        color="warning"
        size="lg"
        onPress={() => handleAnswer(height)}
      >
        →
      </Button>
    </form>
  );
}
