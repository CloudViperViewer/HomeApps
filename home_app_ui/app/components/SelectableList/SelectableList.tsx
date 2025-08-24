"use client";

import React from "react";
import SelectableListRow from "./SelectableListRow";

import post, { Query } from "../apis/post";

async function SelectableList() {
  const query: Query = {
    table: "Bank",
    fields: ["BankID", "BankName"],
    logicalExpression: {
      operator: "AND",
      filters: [
        {
          operator: "=",
          field: "IsActive",
          value: [true],
        },
      ],
      logicalExpressions: null,
    },
    pagingInfo: { startIndex: 1, batchSize: 10 },
  };
  //Call api
  const data = await post("api/select", query);

  return (
    <div className="p-2 shadow-xl/10 bg-white">
      <h1>Banks</h1>
      <ul>
        {data.data.map((item: any, index) => (
          <SelectableListRow
            Index={index}
            RowId={item.BankId}
            RowName={item.BankName}
            selected={index === selectedRow}
            onSelectRow={UpdateSelectedRow}
          />
        ))}
      </ul>
    </div>
  );
}

export default SelectableList;
