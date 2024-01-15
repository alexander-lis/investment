import React from "react";
import { Asset, Group, Sector } from "../../types/types-portfolio";
import SectorNode from "./SectorNode";

type GroupProps = {
  group: Group;
  sectors: Sector[];
  assets: Asset[];
};

export default function GroupNode({ group, sectors, assets }: GroupProps) {
  const groupSectors = sectors.filter((s) => s.groupId == group.id);
  const groupAssets = assets.filter(
    (a) => groupSectors.filter((s) => a.sectorId == s.id).length > 0
  );

  return (
    <React.Fragment key={group.id}>
      <div className="col-span-full text-center font-semibold bg-red-200">
        {group.name}
      </div>
      <div
        className="col-start-1 bg-red-500"
        style={{ gridRow: GetRowSpan(groupSectors, groupAssets) }}
      >
        21412
      </div>
      {groupSectors.map((s) => (
        <SectorNode key={s.id} sector={s} assets={groupAssets}></SectorNode>
      ))}
    </React.Fragment>
  );
}

function GetRowSpan(groupSectors: Sector[], groupAssets: Asset[]) {
  const rows = Math.max(groupSectors.length, groupAssets.length);
  return `auto / span ${rows}`;
}
