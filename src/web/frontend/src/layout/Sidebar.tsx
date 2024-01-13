import { NavLink } from "react-router-dom";
import { classNames } from "../utils/helpers";
import { NavigationItem } from "../utils/types-common";

type SidebarProps = {
  navigation: NavigationItem[];
};

export default function Sidebar({ navigation }: SidebarProps) {
  return (
    <div className="w-full px-4 pt-4">
      <div className="mx-auto w-full p-2 rounded-lg bg-white border border-solid border-gray-200">
        <div className="px-3 pt-1 pb-2 font-bold">Навигация</div>
        <div>
          {navigation.map((item) => (
            <NavLink key={item.name} to={item.href}>
              {({ isActive }) => (
                <div
                  className={classNames(
                    isActive ? "bg-gray-50" : "",
                    "flex rounded-lg px-3 py-1.5 mt-1 hover:bg-gray-100 cursor-pointer items-center"
                  )}
                >
                  <span className="h-6 w-6">{item.icon}</span>
                  <span
                    className={classNames(
                      isActive ? "font-semibold" : "font-medium",
                      "pl-1.5"
                    )}
                  >
                    {item.name}
                  </span>
                </div>
              )}
            </NavLink>
          ))}
        </div>
      </div>
    </div>
  );
}
