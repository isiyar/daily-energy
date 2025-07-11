export function formatDate(date: Date): string {
  const day = String(date.getDate()).padStart(2, "0");
  const month = String(date.getMonth() + 1).padStart(2, "0");

  return `${day}.${month}`;
}

export function nowDate(): string[] {
  const date: Date = new Date(Date.now());
  const day = date.getDate();
  const month = date.toLocaleString("en-US", { month: "short" });

  return [day.toString(), month];
}

export function getTodayAndTomorrowTimestamps(): {
  start_at: number;
  finish_at: number;
} {
  const now = new Date();

  const today = new Date(now.getFullYear(), now.getMonth(), now.getDate());
  const todayTimestamp = Math.floor(today.getTime() / 1000);

  const tomorrow = new Date(today);

  tomorrow.setDate(tomorrow.getDate() + 1);
  const tomorrowTimestamp = Math.floor(tomorrow.getTime() / 1000);

  return { start_at: todayTimestamp, finish_at: tomorrowTimestamp };
}

export function formatTime(date: Date): string {
  const hours = date.getHours().toString().padStart(2, "0");
  const minutes = date.getMinutes().toString().padStart(2, "0");

  return `${hours}:${minutes}`;
}

export function getPercentage(
  part: number,
  total: number,
  decimals = 0,
): number {
  if (total === 0) {
    return 0;
  }
  const raw = (part / total) * 100;

  return parseFloat(raw.toFixed(decimals));
}
