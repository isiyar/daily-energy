import { Input } from "@heroui/input";
import { useState } from "react";
import { Button } from "@heroui/button";

import { FormProops } from "@/pages/Register.tsx";

export function Name({ onChange, increaseQuestionId }: FormProops) {
  const [name, setName] = useState<string>("");

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    setName(e.target.value);
  }

  function handleAnswer() {
    if (onChange) {
      onChange((prev) => ({ ...prev, name }));
    }
    increaseQuestionId();
  }

  return (
    <form className="mt-[6dvh]">
      <Input
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] rounded-xl font-[300] text-[5dvw]"
        color="default"
        label="Имя"
        radius="md"
        size="lg"
        type="text"
        onChange={handleInputChange}
      />
      <Button
        className={`float-right text-[6dvw] p-[5dvw] bg-[#F08629] text-white mt-[15dvh]`}
        color="warning"
        size="lg"
        onPress={handleAnswer}
      >
        →
      </Button>
    </form>
  );
}
