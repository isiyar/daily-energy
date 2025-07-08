import { useState } from "react";

import { NumberSlider } from "@/components/NumberSlider.tsx";

export function Height() {
  const [height, setCurrentHeight] = useState(0);

  return (
    <form>
      <NumberSlider
        numbers={Array.from({ length: 291 }, (_, i) => i + 10)}
        setValue={setCurrentHeight}
      />
    </form>
  );
}
