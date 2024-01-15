import { useLoaderData } from "react-router-dom";
import { PortfolioPageData } from "../types/types-portfolio";
import { format } from "date-fns";
import GroupNode from "../components/portfolio/GroupNode";

export default function PortfolioPage() {
  const { portfolio, groups, assets, sectors } =
    useLoaderData() as PortfolioPageData;

  return (
    <div className="container pt-4 pr-4">
      <div className="p-2 rounded-lg bg-green border border-solid border-gray-200">
        <div className="px-3 pt-1 pb-2 font-bold text-lg text-center">
          {portfolio.name +
            " (" +
            format(portfolio.startDate, "dd.MM.yyyy") +
            " - " +
            format(portfolio.endDate, "dd.MM.yyyy") +
            ")"}
        </div>
        <div className="grid grid-cols-12 items-stretch">
          {/* Header */}
          <div className="col-span-2 text-base text-center font-semibold">
            Факт
          </div>
          <div className="col-span-8 text-base text-center font-semibold">
            Активы
          </div>
          <div className="col-span-2 text-base text-center font-semibold">
            План
          </div>
          {/* Risk Groups */}
          {groups.map((g) => (
            <GroupNode
              key={g.id}
              sectors={sectors}
              assets={assets}
              group={g}
            ></GroupNode>
          ))}
        </div>
      </div>
    </div>
  );
}
