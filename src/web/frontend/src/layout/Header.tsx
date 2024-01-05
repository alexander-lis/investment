import { Menu, Transition } from "@headlessui/react";
import { BellIcon, UserCircleIcon } from "@heroicons/react/24/outline";
import { Fragment } from "react";
import { classNames } from "../utils/helpers";
import { HeaderLink } from "../components/layout/HeaderLink";

const navigation = [
  { name: "Личный бюджет", href: "#", isCurrent: true },
  { name: "Инвестиции", href: "#", isCurrent: false },
];

const userNavigation = [
  { name: "Профиль", href: "#" },
  { name: "Настройки", href: "#" },
  { name: "Выход", href: "#" },
];

export default function Header() {
  return (
    <div className="flex items-center h-12 px-4 bg-gradient-to-l from-gray-800 via-gray-700 to-gray-800">
      {/* Buttons */}
      <div className="flex-1 flex space-x-3">
        {navigation.map((item) => (
          <HeaderLink {...item} />
        ))}
      </div>
      {/* Controls */}
      <div className="flex-1 flex space-x-1 justify-end">
        {/* Notifications dropdown */}
        <button
          type="button"
          className="relative rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
        >
          <span className="absolute -inset-1.5" />
          <span className="sr-only">View notifications</span>
          <BellIcon className="h-6 w-6" aria-hidden="true" />
        </button>

        {/* Profile dropdown */}
        <Menu as="div" className="relative">
          <div>
            <Menu.Button className="relative rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800">
              <span className="absolute -inset-1.5" />
              <span className="sr-only">View notifications</span>
              <UserCircleIcon className="h-6 w-6" aria-hidden="true" />
            </Menu.Button>
          </div>
          <Transition
            as={Fragment}
            enter="transition ease-out duration-100"
            enterFrom="transform opacity-0 scale-95"
            enterTo="transform opacity-100 scale-100"
            leave="transition ease-in duration-75"
            leaveFrom="transform opacity-100 scale-100"
            leaveTo="transform opacity-0 scale-95"
          >
            <Menu.Items className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
              {userNavigation.map((item) => (
                <Menu.Item key={item.name}>
                  {({ active }) => (
                    <a
                      href={item.href}
                      className={classNames(
                        active ? "bg-gray-100" : "",
                        "block px-4 py-2 text-sm text-gray-700"
                      )}
                    >
                      {item.name}
                    </a>
                  )}
                </Menu.Item>
              ))}
            </Menu.Items>
          </Transition>
        </Menu>
      </div>
    </div>
  );
}
