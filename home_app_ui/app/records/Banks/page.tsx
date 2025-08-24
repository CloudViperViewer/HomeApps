import React from "react";
import post, { Query } from "../../components/apis/post";
import Card from "../../components/card/card";
import Link from "next/link";

async function Banks() {
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
    <Card>
      <h1>Banks</h1>
      {data.success === true ? (
        data.data.map((item, index) => (
          <Card key={index} className="mt-5">
            {item.BankName}
          </Card>
        ))
      ) : (
        <div>Failed</div>
      )}
    </Card>
  );
}

export default Banks;
