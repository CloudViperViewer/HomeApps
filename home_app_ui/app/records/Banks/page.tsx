//Landing page for the bank Record list

import React from "react";
import post, { Query } from "../../components/apis/post";
import Card from "../../components/card/card";
import { BankDT } from "@/app/components/dataTypes/BankDT";
import PageHeader from "@/app/components/headers/pageHeader";

async function Banks() {
  /*Bank query*/
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
        data.data.map((raw: any) => {
          //convert to bank dt
          const item: BankDT = {
            BankId: raw.BankId ?? raw.BankID,
            BankName: raw.BankName,
          };
          //Construct Bank Card
          return (
            <Card key={item.BankId} className="mt-5">
              {item.BankName}
            </Card>
          );
        })
      ) : (
        <div>Failed</div>
      )}
    </Card>
  );
}

export default Banks;
