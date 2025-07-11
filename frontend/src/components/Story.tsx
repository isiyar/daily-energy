import { ActionType } from "@/api/actions.ts";
import { Action } from "@/components/Note.tsx";

export function Story({ actions }: { actions: ActionType[] }) {
  return (
    <section className="flex flex-col gap-2 min-h-[20dvh] max-h-[20dvh] overflow-y-auto">
      {actions.length > 0 &&
        actions.map((action: ActionType) => (
          <Action key={action.id} action={action} />
        ))}
      {actions.length === 0 && <p>Тут пока ничего)</p>}
    </section>
  );
}
