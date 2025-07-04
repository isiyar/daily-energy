import { Button } from "@heroui/button";

import { GraphIcon, SmileIcon, StrongIcon } from "@/icons";

export function Target() {
	return (
		<form className="mt-[6dvh] flex flex-col gap-[4dvh] font-[300] ml-[10dvw] mr-[10dvw]">
			<Button
				className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] p-[4dvh] text-[5dvw]"
				color="default"
				startContent={<GraphIcon />}
			>
				Похудеть
			</Button>
			<Button
				className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] p-[4dvh] text-[5dvw]"
				color="default"
				startContent={<SmileIcon />}
			>
				Поддерживать вес
			</Button>
			<Button
				className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] p-[4dvh] text-[5dvw]"
				color="default"
				startContent={<StrongIcon />}
			>
				Набрать вес
			</Button>
		</form>
	);
}
