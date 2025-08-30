/*
 * Takes a value and formats it into the correct currency format
 * number number to format info currency
 * Returns string

 */
export function FormateCurrency(value: number) {
  return new Intl.NumberFormat("en-AU", {
    style: "currency",
    currency: "AUD",
  }).format(value);
}
