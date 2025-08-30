const AUD_FORMATTER = new Intl.NumberFormat("en-AU", {
  style: "currency",
  currency: "AUD",
});

/**
 * Format a value as AUD currency.
 */
export function formatCurrency(
  value: number | string | null | undefined
): string {
  const n = typeof value === "string" ? Number(value) : value ?? NaN;
  if (!Number.isFinite(n)) return "";

  return AUD_FORMATTER.format(n);
}
