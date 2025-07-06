import React, { useRef, useState } from "react";

export function NumberSlider({
  numbers,
  setValue,
}: {
  numbers: number[];
  setValue: (value: number) => void;
}) {
  const [currentIndex, setCurrentIndex] = useState(
    numbers.length % 2 === 0 ? numbers.length / 2 : (numbers.length + 1) / 2,
  );
  const [startX, setStartX] = useState(0);
  const [isDragging, setIsDragging] = useState(false);
  const sliderRef = useRef<HTMLDivElement>(null);

  const SWIPE_THRESHOLD = 30;
  const ANIMATION_DURATION = 300;

  const goNext = () => {
    setCurrentIndex((prev) => (prev + 1) % numbers.length);
  };

  const goPrev = () => {
    setCurrentIndex((prev) => (prev - 1 + numbers.length) % numbers.length);
  };

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

  const getVisibleNumbers = () => {
    setValue(numbers[currentIndex]);

    return [
      numbers[(currentIndex - 1 + numbers.length) % numbers.length],
      numbers[currentIndex],
      numbers[(currentIndex + 1) % numbers.length],
    ];
  };

  const visibleNumbers = getVisibleNumbers();

  return (
    <div className="flex flex-col items-center mt-[10dvh]">
      <div className="relative w-full max-w-xs overflow-hidden">
        <button
          className="absolute left-0 top-1/2 -translate-y-1/2 z-10 p-2 text-white rounded-full font-[300] text-[10dvw]"
          onClick={goPrev}
        >
          &lt;
        </button>

        <button
          className="absolute right-0 top-1/2 -translate-y-1/2 z-10 p-2 text-white rounded-full font-[300] text-[10dvw]"
          onClick={goNext}
        >
          &gt;
        </button>

        <div
          ref={sliderRef}
          className="flex justify-center items-center h-20 select-none touch-pan-x"
          style={{ transition: `transform ${ANIMATION_DURATION}ms ease` }}
          onTouchEnd={handleTouchEnd}
          onTouchMove={handleTouchMove}
          onTouchStart={handleTouchStart}
        >
          {visibleNumbers.map((num, idx) => (
            <div
              key={idx}
              className={`flex-shrink-0 w-12 h-12 flex items-center justify-center mx-2 transition-all duration-300 font-[300] ${
                idx === 1
                  ? "text-[12dvw] text-white mb-[3dvh]"
                  : "text-[12dvw] text-gray-500"
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
