import { Link, useLoaderData } from "react-router-dom";
import { PortfoliosPageData } from "../types/types-portfolio";

export default function PortfoliosPage() {
  const { portfolios } = useLoaderData() as PortfoliosPageData;

  return (
    <div>
      {portfolios.map((p) => (
        <Link to={p.id.toString()}>
          {p.name}
          <br />
        </Link>
      ))}
    </div>
  );
}