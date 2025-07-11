import { getPercentage } from "@/utils.ts";

interface SemiCircularProgressBarProps {
  size?: number;
  strokeWidth?: number;
  circleColor?: string;
  progressColor?: string;
  showText?: boolean;
  textColor?: string;
  curr: number;
  end: number;
}

export function SemiCircularProgressBar({
  size = 120,
  strokeWidth = 10,
  circleColor = "bg-gray-200",
  progressColor = "bg-blue-500",
  showText = true,
  textColor = "text-gray-700",
  curr,
  end,
}: SemiCircularProgressBarProps) {
  const normalizedProgress = Math.min(
    Math.max(getPercentage(curr, end), 0),
    100,
  );
  const radius = (size - strokeWidth) / 2;
  const circumference = radius * Math.PI;
  const strokeDashoffset =
    circumference - (normalizedProgress / 100) * circumference;

  return (
    <div
      className="relative flex flex-col items-center"
      style={{ width: size, height: size / 2 }}
    >
      <svg
        className="absolute top-0 left-0"
        height={size / 2}
        overflow="visible"
        viewBox={`0 0 ${size} ${size / 2}`}
        width={size}
      >
        <path
          className={circleColor}
          d={`M ${strokeWidth / 2} ${size / 2} 
              A ${radius} ${radius} 0 0 1 ${size - strokeWidth / 2} ${size / 2}`}
          fill="none"
          stroke="currentColor"
          strokeWidth={strokeWidth}
        />
      </svg>

      <svg
        className="absolute top-0 left-0"
        height={size / 2}
        overflow="visible"
        viewBox={`0 0 ${size} ${size / 2}`}
        width={size}
      >
        <path
          className={progressColor}
          d={`M ${strokeWidth / 2} ${size / 2} 
              A ${radius} ${radius} 0 0 1 ${size - strokeWidth / 2} ${size / 2}`}
          fill="none"
          stroke="currentColor"
          strokeDasharray={circumference}
          strokeDashoffset={strokeDashoffset}
          strokeLinecap="round"
          strokeWidth={strokeWidth}
        />
      </svg>

      {showText && (
        <div
          className={`absolute text-[2dvw] bottom-0 text-center ${textColor}`}
        >
          {curr} ккал
          <br />
          из {end}
        </div>
      )}
    </div>
  );
}
