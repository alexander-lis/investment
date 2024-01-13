import React from "react";
import { Asset, Sector } from "../../utils/types-portfolio";

type SectorProps = {
  sector: Sector;
  assets: Asset[];
};

export default function SectorNode({ sector, assets }: SectorProps) {
  const sectorAssets = assets.filter((a) => a.sectorId == sector.id);

  return (
    <React.Fragment>
      <div
        className="col-start-2 bg-pink-200"
        style={{ gridRow: GetRowSpan(sectorAssets) }}
      >
        23141
      </div>
      <div
        className="col-start-3 bg-green-200"
        style={{ gridRow: GetRowSpan(sectorAssets) }}
      >
        {sector.name}
      </div>
      {sectorAssets.map((a, i) => (
        <React.Fragment key={a.id}>
          <div key={a.id} className="col-start-4 col-span-7 bg-yellow-300">
            {a.ticker}
          </div>
          {i == 0 && (
            <div
              className="col-start-11 col-span-full bg-orange-500"
              style={{ gridRow: GetRowSpan(sectorAssets) }}
            >
              {sector.name}
            </div>
          )}
        </React.Fragment>
      ))}
    </React.Fragment>
  );
}

function GetRowSpan(sectorAssets: Asset[]) {
  const rows = Math.max(1, sectorAssets.length);
  return `auto / span ${rows}`;
}
