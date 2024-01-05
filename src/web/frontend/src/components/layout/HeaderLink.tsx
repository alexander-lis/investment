import { classNames } from "../../utils/helpers";

type HeaderLinkProps = {
    name: string
    href: string
    isCurrent: boolean
}

export function HeaderLink({
  name,
  href,
  isCurrent,
}: HeaderLinkProps) {
  return (
    <a
      key={name}
      href={href}
      className={classNames(
        isCurrent
          ? "bg-gray-900 text-white"
          : "text-gray-300 hover:bg-gray-700 hover:text-white",
        "rounded-md px-3 py-1 text-sm font-medium"
      )}
      aria-current={isCurrent ? "page" : undefined}
    >
      {name}
    </a>
  );
}
