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
