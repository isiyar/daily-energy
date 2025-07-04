import { Button } from "@heroui/button";
import { Progress } from "@heroui/react";
import { useState } from "react";

import { QUESTIONS } from "@/constants";
import { BackIcon } from "@/icons";
import { Name } from "@/layouts/Name";
import { Target } from "@/layouts/Target";

export function Register() {
	const [questionId, setQuestionId] = useState(0);

	function increaseQuestionId() {
		setQuestionId(questionId + 1);
	}

	function decreaseQuestionId() {
		setQuestionId(questionId - 1);
	}

	return (
		<div className="p-[3dvh]">
			<header className="flex align-middle items-center">
				<Button isIconOnly variant="light" onPress={decreaseQuestionId}>
					<BackIcon />
				</Button>
				<Progress
					className="ml-[3dvw]"
					color="warning"
					value={(100 / 6) * (questionId + 1)}
				/>
			</header>
			<h1 className="text-white font-[400] text-[7dvw] mt-[2dvh] ml-[2dvw]">
				{QUESTIONS[questionId]}
			</h1>
			<main>
				{questionId === 0 && <Name />}
				{questionId == 1 && <Target />}
				<Button
					className="float-right text-[30px] p-[2dvw] bg-[#F08629] text-white mt-[15dvh]"
					color="warning"
					size="lg"
					onPress={increaseQuestionId}
				>
					{questionId === QUESTIONS.length - 1 ? "Создать план" : "→"}
				</Button>
			</main>
		</div>
	);
}
