/*
 * Takes a value finds its index in the first array and replaces it with the same index in second array
 * DateTime Date and time to format to string
 * Returns string

 */
export function FormateDateTime(DateTime: Date) {
  return new Date(DateTime).toLocaleDateString("en-AU", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}
