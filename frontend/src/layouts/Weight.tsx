import { useState } from "react";

import { NumberSlider } from "@/components/NumberSlider.tsx";

export function Weight() {
  const [weight, setCurrentWeight] = useState(0);

  return (
    <form>
      <NumberSlider
        numbers={[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}
        setValue={setCurrentWeight}
      />
    </form>
  );
}
