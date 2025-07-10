import { Button } from "@heroui/button";
import { Progress } from "@heroui/react";
import { useState } from "react";

import { QUESTIONS } from "@/constants";
import { BackIcon } from "@/icons";
import { Name } from "@/layouts/Name";
import { Target } from "@/layouts/Target";
import { Weight } from "@/layouts/Weight.tsx";
import { Height } from "@/layouts/Height.tsx";
import { Gender } from "@/layouts/Gender.tsx";
import { Birthday } from "@/layouts/Birthday.tsx";
import { Activity } from "@/layouts/Activity.tsx";
import { User } from "@/api/user.ts";
import { useRegister } from "@/hooks/user.ts";

export interface FormProops {
  onChange?: (
    value: ((prevState: Partial<User>) => Partial<User>) | Partial<User>,
  ) => void;
  increaseQuestionId: () => void;
}

export function Register() {
  const [userData, setUserData] = useState<Partial<User>>({
    utgid: window.Telegram.WebApp.initDataUnsafe.user.id,
  });
  const [questionId, setQuestionId] = useState(0);
  const { mutate } = useRegister();

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

  function sendUserData() {
    mutate({ user: userData, initData: window.Telegram.WebApp.initData });
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
        {questionId === 0 && (
          <Name
            increaseQuestionId={increaseQuestionId}
            onChange={setUserData}
          />
        )}
        {questionId === 1 && (
          <Gender
            increaseQuestionId={increaseQuestionId}
            onChange={setUserData}
          />
        )}
        {questionId === 2 && (
          <Target
            increaseQuestionId={increaseQuestionId}
            onChange={setUserData}
          />
        )}
        {questionId === 3 && (
          <Weight
            increaseQuestionId={increaseQuestionId}
            onChange={setUserData}
          />
        )}
        {questionId === 4 && (
          <Height
            increaseQuestionId={increaseQuestionId}
            onChange={setUserData}
          />
        )}
        {questionId === 5 && (
          <Birthday
            increaseQuestionId={increaseQuestionId}
            onChange={setUserData}
          />
        )}
        {questionId === 6 && (
          <Activity
            increaseQuestionId={increaseQuestionId}
            sendUserData={sendUserData}
            onChange={setUserData}
          />
        )}
        {/*<Button*/}
        {/*  className={`float-right text-[6dvw] p-[5dvw] bg-[#F08629] text-white ${questionId === 6 ? "mt-[7dvh]" : "mt-[15dvh]"}`}*/}
        {/*  color="warning"*/}
        {/*  size="lg"*/}
        {/*  onPress={increaseQuestionId}*/}
        {/*>*/}
        {/*  {questionId === QUESTIONS.length - 1 ? "Создать план" : "→"}*/}
        {/*</Button>*/}
      </main>
    </div>
  );
}
