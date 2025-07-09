import { Button } from "@heroui/react";

import { formatDate } from "@/utils.ts";
import { CalendarIcon } from "@/icons.tsx";

export function PlanHeader({ date }: { date: Date }) {
  return (
    <header className="flex w-full items-center justify-between">
      <h1 className="text-white font-[500] text-[7dvw] mt-[2dvh] ml-[2dvw]">
        {date.toDateString() === new Date().toDateString()
          ? "Сегодняшний план"
          : `История за ${formatDate(date)}`}
      </h1>
      <Button isIconOnly className="bg-transparent">
        <CalendarIcon />
      </Button>
    </header>
  );
}
