import { ActionType } from "@/api/actions.ts";
import { formatTime } from "@/utils.ts";

export function Action({ action }: { action: ActionType }) {
  return (
    <article className="flex justify-between items-center px-[2dvw]">
      <div>{formatTime(new Date(action.date))}</div>
      <div>{action.activity_name}</div>
      <div>{action.calories} ккал</div>
    </article>
  );
}
