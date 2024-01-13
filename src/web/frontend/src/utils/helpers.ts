// Concates array of class names into one string.
export function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(" ");
}
