/**
 * Format a date/time into "dd/mm/yyyy, hh:mm" (en-AU).
 */
export function formatDateTime(input: Date | string | number): string {
  const d = new Date(input);
  if (Number.isNaN(d.getTime())) return "";
  return d.toLocaleString("en-AU", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}
