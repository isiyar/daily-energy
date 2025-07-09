import React, { useCallback, useRef, useState } from "react";

export function NumberSlider({
  numbers,
  setValue,
}: {
  numbers: number[];
  setValue: (value: number) => void;
}) {
  const initialIndex =
    numbers.length % 2 === 0 ? numbers.length / 2 : (numbers.length - 1) / 2;
  const [currentIndex, setCurrentIndex] = useState(initialIndex);
  const [startX, setStartX] = useState(0);
  const [isDragging, setIsDragging] = useState(false);
  const sliderRef = useRef<HTMLDivElement>(null);

  const SWIPE_THRESHOLD = 30;
  const ANIMATION_DURATION = 300;

  const goNext = useCallback(() => {
    setCurrentIndex((prev) => (prev + 1) % numbers.length);
  }, [numbers.length]);

  const goPrev = useCallback(() => {
    setCurrentIndex((prev) => (prev - 1 + numbers.length) % numbers.length);
  }, [numbers.length]);

  const handleTouchStart = (e: React.TouchEvent) => {
    setStartX(e.touches[0].clientX);
    setIsDragging(true);
  };

  const handleTouchMove = (e: React.TouchEvent) => {
    if (!isDragging) return;
    const x = e.touches[0].clientX;
    const diff = startX - x;

    if (diff > SWIPE_THRESHOLD) {
      goNext();
      setStartX(x);
    } else if (diff < -SWIPE_THRESHOLD) {
      goPrev();
      setStartX(x);
    }
  };

  const handleTouchEnd = () => {
    setIsDragging(false);
  };

  const getVisibleNumbers = useCallback(() => {
    setValue(numbers[currentIndex]);

    return [
      numbers[(currentIndex - 1 + numbers.length) % numbers.length],
      numbers[currentIndex],
      numbers[(currentIndex + 1) % numbers.length],
    ];
  }, [currentIndex, numbers, setValue]);

  const visibleNumbers = getVisibleNumbers();

  const handleButtonClick = useCallback(
    (direction: "prev" | "next") => {
      return (e: React.MouseEvent) => {
        e.preventDefault();
        e.stopPropagation();
        direction === "prev" ? goPrev() : goNext();
      };
    },
    [goPrev, goNext],
  );

  return (
    <div className="flex flex-col w-full align-middle items-center mt-[6dvh] pl-[5dvw] pr-[5dvw]">
      <div className="relative w-full overflow-hidden">
        <button
          className="absolute left-0 top-1/2 -translate-y-1/2 z-10 p-2 text-white rounded-full font-[300] text-[10dvw]"
          onClick={handleButtonClick("prev")}
          onTouchStart={(e) => e.preventDefault()}
        >
          &lt;
        </button>
        <button
          className="absolute right-0 top-1/2 -translate-y-1/2 z-10 p-2 text-white rounded-full font-[300] text-[10dvw]"
          onClick={handleButtonClick("next")}
          onTouchStart={(e) => e.preventDefault()}
        >
          &gt;
        </button>

        <div
          ref={sliderRef}
          className="flex justify-center items-center h-[20dvh] select-none touch-pan-x"
          style={{ transition: `transform ${ANIMATION_DURATION}ms ease` }}
          onTouchEnd={handleTouchEnd}
          onTouchMove={handleTouchMove}
          onTouchStart={handleTouchStart}
        >
          {visibleNumbers.map((num, idx) => (
            <div
              key={`${num}-${idx}`}
              className={`flex-shrink-0 h-12 flex items-center justify-center mx-2 transition-all duration-300 font-[300] ${
                idx === 1
                  ? "text-[8dvw] text-white mb-[3dvh] mx-[5dvw]"
                  : "text-[8dvw] text-gray-500"
              }`}
            >
              {num}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
