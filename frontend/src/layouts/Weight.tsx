import { useState } from "react";

import { NumberSlider } from "@/components/NumberSlider.tsx";

export function Weight() {
  const [weight, setCurrentWeight] = useState(0);

  return (
    <form>
      <NumberSlider
        numbers={Array.from({ length: 171 }, (_, i) => i + 30)}
        setValue={setCurrentWeight}
      />
    </form>
  );
}
