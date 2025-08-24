import React from "react";

interface SelectableListRowProps {
  Index: number;
  RowId: number;
  RowName: string;
  selected: boolean;
  onSelectRow: (item: number) => void;
}

const SelectableListRow = ({
  Index,
  RowId,
  RowName,
  selected,
  onSelectRow,
}: SelectableListRowProps) => {
  return (
    <li
      key={RowId}
      className={
        selected ? "hover:cursor-pointer bg-accent-100" : "hover:cursor-pointer"
      }
      onClick={() => {
        onSelectRow(Index);
      }}
    >
      {RowName}
    </li>
  );
};

export default SelectableListRow;
