import { SubHeaderLink } from "../components/layout/SubHeaderLink";

const navigation = [
  { name: "На квартиру", href: "#", isCurrent: true },
  { name: "На пенсию", href: "#", isCurrent: false },
];

export default function SubHeader() {
  return (
    <div className="flex items-center h-7 px-8 bg-gray-200 drop-shadow-md">
      {/* Buttons */}
      {navigation.map((item) => (
        <SubHeaderLink {...item} />
      ))}
    </div>
  );
}
