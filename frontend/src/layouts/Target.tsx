import { Button } from "@heroui/button";

export function Target() {
  return (
    <form className="mt-[6dvh] flex flex-col gap-[4dvh] font-[300] ml-[10dvw] mr-[10dvw]">
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] py-[4dvh] text-[5dvw]"
        color="default"
      >
        <div className="flex w-full">
          <div>📈</div>
          <div className="ml-[3dvw] flex items-center">Похудеть</div>
        </div>
      </Button>
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] py-[4dvh] text-[5dvw]"
        color="default"
      >
        <div className="flex w-full">
          <div>😊</div>
          <div className="ml-[3dvw] flex items-center">Поддерживать вес</div>
        </div>
      </Button>
      <Button
        className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] py-[4dvh] text-[5dvw]"
        color="default"
      >
        <div className="flex w-full">
          <div>💪</div>
          <div className="ml-[3dvw] flex items-center">Набрать вес</div>
        </div>
      </Button>
    </form>
  );
}
