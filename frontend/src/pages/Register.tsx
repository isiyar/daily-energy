import { Button } from "@heroui/button";
import { Progress } from "@heroui/react";
import { useState } from "react";

import { QUESTIONS } from "@/constants";
import { BackIcon, BottomArrowIcon } from "@/icons";
import { Name } from "@/layouts/Name";
import { Target } from "@/layouts/Target";
import { Weight } from "@/layouts/Weight.tsx";
import { Height } from "@/layouts/Height.tsx";
import { Gender } from "@/layouts/Gender.tsx";
import { Birthday } from "@/layouts/Birthday.tsx";
import { Activity } from "@/layouts/Activity.tsx";

export function Register() {
  const [questionId, setQuestionId] = useState(0);

  function increaseQuestionId() {
    if (questionId < QUESTIONS.length - 1) {
      setQuestionId(questionId + 1);
    }
  }

  function decreaseQuestionId() {
    if (questionId > 0) {
      setQuestionId(questionId - 1);
    }
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
          value={(100 / 7) * (questionId + 1)}
        />
      </header>
      <h1 className="text-white font-[500] text-[7dvw] mt-[2dvh] ml-[2dvw]">
        {QUESTIONS[questionId]}
      </h1>
      <main>
        {questionId === 0 && <Name />}
        {questionId === 1 && <Gender />}
        {questionId === 2 && <Target />}
        {questionId === 3 && <Weight />}
        {questionId === 4 && <Height />}
        {questionId === 5 && <Birthday />}
        {questionId === 6 && <Activity />}
        <Button
          className={`float-right text-[6dvw] p-[5dvw] bg-[#F08629] text-white ${questionId === 6 ? "mt-[7dvh]" : "mt-[15dvh]"}`}
          color="warning"
          size="lg"
          onPress={increaseQuestionId}
        >
          {questionId === QUESTIONS.length - 1 ? "Создать план" : "→"}
        </Button>
        {questionId === 6 && (
          <div className="absolute bottom-[5dvh]">
            <BottomArrowIcon />
          </div>
        )}
      </main>
    </div>
  );
}
