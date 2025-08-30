/**
 * Takes a value finds its index in the first array and replaces it with the same index in second array
 * Value to look for
 * Array to search for value
 * Array to find replacement
 * default if no replacement found
 * Returns either the default or the replacement value
 */
export function DisplayValue<T, U, D extends U | null | undefined = U>(
  value: T,
  arrayToSearch: ReadonlyArray<T>,
  arrayToFindIndex: ReadonlyArray<U>,
  defaultValue: D
): U | D {
  const index = arrayToSearch.indexOf(value);
  if (index === -1 || index >= arrayToFindIndex.length) {
    return defaultValue;
  }
  return arrayToFindIndex[index];
}
