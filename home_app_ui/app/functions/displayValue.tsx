/*
 * Takes a value finds its index in the first array and replaces it with the same index in second array
 * Value to look for
 * Array to search for value
 * Array to find replacement
 * default if no replacement found
 * Returns either default or replacement value
 */
export function DisplayValue<T, U>(
  value: T,
  arrayToSearch: T[],
  arrayToFindIndex: U[],
  Default: any
) {
  const index = arrayToSearch.indexOf(value);
  if (index === -1 || index >= arrayToFindIndex.length) {
    return Default;
  }
  return arrayToFindIndex[index];
}
