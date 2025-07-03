import { Button } from "@heroui/react";

import arrowSrc from "../assets/arrow.png";
import logoSrc from "../assets/logo.png";

export function Greet() {
	return (
		<div className="min-w-max min-h-max flex flex-col items-center">
			<img alt="logoSrc" className="w-[60dvw] mt-[2dvh]" src={logoSrc} />
			<div className="text-white mt-[2dvh]">
				<p className="text-[7dvw] font-[400]">
					Привет!
					<br />
					Мы - Daily Energy
				</p>
				<p className="text-[5dvw] mt-[1dvh] font-[300]">
					Твой персональный гид в<br />
					мире здоровья.
				</p>
				<img
					alt="arrowSrc"
					className="w-[18dvh] mt-[3dvh] ml-[5dvw]"
					src={arrowSrc}
				/>
			</div>
			<Button
				className="ml-auto mr-[5dvw] text-[30px] p-[2dvw] bg-[#F08629] text-white"
				color="warning"
				size="lg"
			>
				GO
			</Button>
		</div>
	);
}
