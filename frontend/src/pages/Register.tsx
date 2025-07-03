import { Button } from "@heroui/button";
import { Progress } from "@heroui/react";

import { BackIcon } from "@/icons";

export function Register() {
	return (
		<div>
			<header className="flex align-middle justify-center">
				<Button isIconOnly variant="light">
					<BackIcon />
				</Button>
				<Progress />
			</header>
		</div>
	);
}
