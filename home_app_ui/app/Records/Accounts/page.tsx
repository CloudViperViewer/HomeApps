import React from "react";
import post, { Query } from "../../components/apis/post";
import Card from "../../components/card/card";
import { DisplayValue } from "../../functions/displayValue";

async function Accounts() {
  //   const bankQuery: Query = {
  //     table: "Bank",
  //     fields: ["BankID", "BankName"],
  //     logicalExpression: {
  //       operator: "AND",
  //       filters: [
  //         {
  //           operator: "=",
  //           field: "IsActive",
  //           value: [true],
  //         },
  //       ],
  //       logicalExpressions: null,
  //     },
  //     pagingInfo: { startIndex: 1, batchSize: 10 },
  //   };
  const accountsQuery: Query = {
    table: "Account",
    fields: ["BankID", "AccountName", "Balance", "AccountID", "AccountTypeID"],
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
  //   const banks = await post("api/select", bankQuery);
  const accounts = await post("api/select", accountsQuery);
  return (
    <Card>
      <h1>Accounts</h1>
      {accounts.success === true ? (
        accounts.data.map((item) => (
          <Card key={item.accountId} className="mt-5">
            <div className="grid grid-cols-2">
              <div>
                <h1 className="underline">{item.accountName}</h1>
              </div>
              <div>
                <h1 className="text-right text-2xl">
                  <b>${item.balance}</b>
                </h1>
              </div>
            </div>
          </Card>
        ))
      ) : (
        <div>Failed</div>
      )}
    </Card>
  );
}

export default Accounts;
