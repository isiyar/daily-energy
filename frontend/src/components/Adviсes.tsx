export function Advices({ advices }: { advices: string[] }) {
  return (
    <section className="text-black rounded-medium bg-[#D9D9D9] w-full max-h-[15dvh] p-[2dvw] mt-[2dvh]">
      <h2 className="font-[500]">Совет от ИИ-помощника:</h2>
      <ul className="overflow-y-auto max-h-[8dvh]">
        {advices.map((advice, idx) => (
          <li key={idx}>• {advice}</li>
        ))}
      </ul>
    </section>
  );
}
