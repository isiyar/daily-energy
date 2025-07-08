import { useState } from "react";

import { DatePicker } from "@/components/DatePicker.tsx";

export function Birthday() {
  const [currentDate, setCurrentDate] = useState(new Date("01.01.2000"));

  return (
    <form>
      <DatePicker initialDate={currentDate} onChange={setCurrentDate} />
    </form>
  );
}
