import { useState } from "react";
import { Button } from "@heroui/button";

import { DatePicker } from "@/components/DatePicker.tsx";
import { FormProops } from "@/pages/Register.tsx";

export function Birthday({ onChange, increaseQuestionId }: FormProops) {
  const [currentDate, setCurrentDate] = useState(new Date("01.01.2000"));

  function handleAnswer(date_of_birth: number) {
    if (onChange) {
      onChange((prev) => ({ ...prev, date_of_birth }));
    }
    increaseQuestionId();
  }

  return (
    <form>
      <DatePicker initialDate={currentDate} onChange={setCurrentDate} />
      <Button
        className={`float-right text-[6dvw] p-[5dvw] bg-[#F08629] text-white mt-[15dvh]`}
        color="warning"
        size="lg"
        onPress={() => handleAnswer(currentDate.getTime())}
      >
        →
      </Button>
    </form>
  );
}
