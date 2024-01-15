import Header from "./Header";
import Footer from "./Footer";
import { Outlet, useLoaderData } from "react-router-dom";
import Sidebar from "./Sidebar";
import { LayoutPageData, NavigationItem } from "../types/types-common";
import { BanknotesIcon, BriefcaseIcon } from "@heroicons/react/24/outline";

const navigation: NavigationItem[] = [
  {
    name: "Бюджет",
    href: "budget",
    icon: <BanknotesIcon />,
  },
  {
    name: "Инвестиции",
    href: "portfolios",
    icon: <BriefcaseIcon />,
  },
];

export default function Layout() {
  const { user } = useLoaderData() as LayoutPageData;

  return (
    <div className="h-screen w-screen flex flex-col bg-white text-gray-700 text-sm ">
      <Header user={user} />
      <div className="flex flex-row flex-1">
        <div className="flex flex-none w-1/5 min-w-52">
          <Sidebar navigation={navigation} />
        </div>
        <div className="flex flex-1">
          <Outlet />
        </div>
      </div>
      <Footer />
    </div>
  );
}
