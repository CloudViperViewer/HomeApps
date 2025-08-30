import React from "react";
import post, { Query } from "../../components/apis/post";
import Card from "../../components/card/card";
import { BankDT } from "@/app/components/dataTypes/BankDT";
import PageHeader from "@/app/components/headers/pageHeader";

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
    },
    pagingInfo: { startIndex: 1, batchSize: 10 },
  };
  //Call api
  const data = await post("api/select", query);
  return (
    <Card className="bg-white">
      <PageHeader headerText="Banks" />
      {data.success === true ? (
        data.data.map((item: BankDT) => (
          <Card key={item.BankId} className="mt-5">
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
