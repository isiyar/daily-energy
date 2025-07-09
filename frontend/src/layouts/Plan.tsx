import { PlanHeader } from "@/components/PlanHeader.tsx";

interface PlanProops {
  date: Date;
}

export function Plan({ date }: PlanProops) {
  return (
    <div className="p-[3dvh]">
      <PlanHeader date={date} />
    </div>
  );
}
