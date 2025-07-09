import React, { useCallback, useMemo, useState } from "react";
import {
  addDays,
  addMonths,
  addYears,
  subDays,
  subMonths,
  subYears,
} from "date-fns";

import { MONTHS } from "@/constants.ts";

export function DatePicker({
  initialDate,
  onChange,
}: {
  initialDate: Date;
  onChange: (date: Date) => void;
}) {
  const [touchStartY, setTouchStartY] = useState(0);
  const [activeColumn, setActiveColumn] = useState<number | null>(null);
  const SWIPE_THRESHOLD = 30;

  const safeDate = useMemo(() => {
    const date = new Date(initialDate);

    return isNaN(date.getTime()) ? new Date() : date;
  }, [initialDate]);

  const getDaysInMonth = useCallback((year: number, month: number): number => {
    try {
      return new Date(year, month + 1, 0).getDate();
    } catch {
      return 30;
    }
  }, []);

  const getDisplayValues = useCallback(
    (date: Date) => {
      const day = date.getDate();
      const month = date.getMonth();
      const year = date.getFullYear();

      const daysInMonth = getDaysInMonth(year, month);
      let days: number[];

      if (day > 1 && day < daysInMonth) {
        days = [day - 1, day, day + 1];
      } else if (day === 1) {
        days = [getDaysInMonth(year, month - 1), day, day + 1];
      } else {
        days = [day - 1, day, 1];
      }

      const months = [
        MONTHS[(month - 1 + 12) % 12],
        MONTHS[month],
        MONTHS[(month + 1) % 12],
      ];

      const years = [year - 1, year, year + 1];

      return [days, months, years];
    },
    [getDaysInMonth],
  );

  const displayData = useMemo(
    () => getDisplayValues(safeDate),
    [safeDate, getDisplayValues],
  );

  const navigate = useCallback(
    (direction: number, index: number) => {
      let newDate: Date;

      try {
        switch (index) {
          case 0:
            newDate =
              direction > 0 ? addDays(safeDate, 1) : subDays(safeDate, 1);
            break;
          case 1:
            newDate =
              direction > 0 ? addMonths(safeDate, 1) : subMonths(safeDate, 1);
            break;
          case 2:
            newDate =
              direction > 0 ? addYears(safeDate, 1) : subYears(safeDate, 1);
            break;
          default:
            newDate = safeDate;
        }

        if (!isNaN(newDate.getTime())) {
          onChange(newDate);
        }
      } catch (error) {}
    },
    [safeDate, onChange],
  );

  const handleTouchStart = useCallback((e: React.TouchEvent, index: number) => {
    setTouchStartY(e.touches[0].clientY);
    setActiveColumn(index);
  }, []);

  const handleTouchMove = useCallback(
    (e: React.TouchEvent) => {
      if (activeColumn === null) return;

      const touchY = e.touches[0].clientY;
      const diff = touchStartY - touchY;

      if (Math.abs(diff) > SWIPE_THRESHOLD) {
        navigate(diff > 0 ? 1 : -1, activeColumn);
        setActiveColumn(null);
      }
    },
    [activeColumn, touchStartY, navigate],
  );

  const handleTouchEnd = useCallback(() => {
    setActiveColumn(null);
  }, []);

  return (
    <div className="flex justify-center items-center touch-none mt-[8dvh]">
      <div className="flex gap-8 select-none font-[400] text-[6dvw]">
        {displayData.map((column, index) => (
          <div
            key={index}
            className="flex flex-col items-center justify-center overflow-hidden relative"
            onTouchEnd={handleTouchEnd}
            onTouchMove={handleTouchMove}
            onTouchStart={(e) => handleTouchStart(e, index)}
          >
            {/* Эффект затухания сверху */}
            <div className="absolute top-0 left-0 right-0 h-[20%] bg-gradient-to-b from-[#212121] to-transparent pointer-events-none z-10" />

            <div className="flex flex-col items-center transition-transform duration-200">
              <div className="">
                {typeof column[0] === "number" ? column[0] : String(column[0])}
              </div>
              <div className="my-[2dvh]">
                {typeof column[1] === "number" ? column[1] : String(column[1])}
              </div>
              <div className="">
                {typeof column[2] === "number" ? column[2] : String(column[2])}
              </div>
            </div>

            {/* Эффект затухания снизу */}
            <div className="absolute bottom-0 left-0 right-0 h-[20%] bg-gradient-to-t from-[#212121] to-transparent pointer-events-none z-10" />
          </div>
        ))}
      </div>
    </div>
  );
}
