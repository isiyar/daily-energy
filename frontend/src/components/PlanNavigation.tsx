import { Button } from "@heroui/react";

import { PlanType } from "@/layouts/Plan.tsx";

export function PlanNavigation({
  onChange,
}: {
  onChange: (prev: PlanType) => void;
}) {
  return (
    <nav className="flex gap-[6dvw] mt-[2dvh] mb-[3dvh]">
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] text-[3dvw]"
        onPress={() => onChange("Activity")}
      >
        Активность
      </Button>
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] text-[3dvw]"
        onPress={() => onChange("Food")}
      >
        Питание
      </Button>
    </nav>
  );
}
