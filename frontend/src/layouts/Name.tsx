import { Input } from "@heroui/input";

export function Name() {
	return (
		<form className="mt-[6dvh]">
			<Input
				className="[box-shadow:0_0_5px_5px_rgba(240,134,41,0.3)] rounded-xl"
				color="default"
				label="Имя"
				radius="md"
				size="lg"
				type="text"
			/>
		</form>
	);
}
