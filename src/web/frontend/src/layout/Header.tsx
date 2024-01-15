import { Menu, Transition } from "@headlessui/react";
import { BellIcon, UserCircleIcon } from "@heroicons/react/24/outline";
import { Fragment } from "react";
import { classNames } from "../utils/helpers";
import { User } from "../types/types-common";
import { Link } from "react-router-dom";

type HeaderProps = {
  user?: User;
};

const actionsMenu = [
  { name: "Профиль", href: "#" },
  { name: "Настройки", href: "#" },
  { name: "Выход", href: "#" },
];

const loginMenu = [{ name: "Войти", href: "#" }];

export default function Header({ user }: HeaderProps) {
  const userMenu = user ? actionsMenu : loginMenu;

  return (
    <div className="flex h-12 px-4 border-b border-solid items-center bg-gray-100">
      {/* Logo */}
      <div className="flex flex-none hover:cursor-pointer">
        <Link to="/">
          <img src="../public/logo.png" title="logo" className="h-8 w-8"></img>
        </Link>
      </div>
      {/* Breadcrumbs */}
      <div className="flex flex-1 font-bold text-lg pl-3">Money Planner</div>
      {/* Searchbar */}
      <div className="flex flex-1"></div>
      {/* Buttons*/}
      <div className="flex flex-none">
        {/* Notifications dropdown */}
        <div className="flex">
          <button
            title="notifications"
            type="button"
            className="relative rounded-full text-gray-600 hover:text-gray-800 transition ease-in-out"
          >
            <BellIcon className="h-6 w-6" aria-hidden="true" />
          </button>
        </div>
        {/* Profile dropdown */}
        <Menu as="div" className="relative ml-1 flex">
          <Menu.Button className="relative rounded-full text-gray-600 hover:text-gray-800 transition ease-in-out">
            <UserCircleIcon className="h-6 w-6" aria-hidden="true" />
          </Menu.Button>
          <Transition
            as={Fragment}
            enter="transition ease-out duration-100"
            enterFrom="transform opacity-0 scale-95"
            enterTo="transform opacity-100 scale-100"
            leave="transition ease-in duration-75"
            leaveFrom="transform opacity-100 scale-100"
            leaveTo="transform opacity-0 scale-95"
          >
            <Menu.Items className="absolute right-0 z-10 mt-7 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
              {userMenu.map((item) => (
                <Menu.Item key={item.name}>
                  {({ active }) => (
                    <a
                      href={item.href}
                      className={classNames(
                        active ? "bg-gray-100" : "",
                        "block px-4 py-2 text-sm"
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
