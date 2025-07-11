import { RecommendationType } from "@/api/plan.ts";

export function Advices({ advices }: { advices: RecommendationType[] }) {
  return (
    <section className="text-black rounded-medium bg-[#D9D9D9] w-full max-h-[15dvh] p-[2dvw] mt-[2dvh]">
      <h2 className="font-[500]">Совет от ИИ-помощника:</h2>
      <ul className="overflow-y-auto max-h-[8dvh]">
        {advices.length > 0 &&
          advices[0].recommendation.split("\\n").map((line: string, index: number) => (
            <li key={index}>
              {line}
              <br />
            </li>
          ))}
        {advices.length === 0 && <li>Рекомендаций нет</li>}
      </ul>
    </section>
  );
}
