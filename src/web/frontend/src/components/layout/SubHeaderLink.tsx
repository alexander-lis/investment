import { classNames } from "../../utils/helpers";

type SubHeaderLinkProps = {
    name: string
    href: string
    isCurrent: boolean
}

export function SubHeaderLink({
  name,
  href,
  isCurrent,
}: SubHeaderLinkProps) {
  return (
    <a
      key={name}
      href={href}
      className={classNames(
        isCurrent
          ? "bg-gray-800 text-white"
          : "text-gray-800 hover:bg-gray-700 hover:text-white",
        "px-3 py-1 text-sm font-medium"
      )}
      aria-current={isCurrent ? "page" : undefined}
    >
      {name}
    </a>
  );
}
