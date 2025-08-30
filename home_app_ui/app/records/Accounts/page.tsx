/*Accounts landing page lising accounts*/

import post, { Query } from "../../components/apis/post";
import Card from "../../components/card/card";
import { DisplayValue } from "../../functions/displayValue";
import { AccountDT } from "@/app/components/dataTypes/AccountDT";
import { BankDT } from "@/app/components/dataTypes/BankDT";
import ColumnLayout from "@/app/components/columnLayout/columnLayout";
import Column from "@/app/components/columnLayout/column";
import PageHeader from "@/app/components/headers/pageHeader";
import { FormateCurrency } from "@/app/functions/formatCurrency";

async function Accounts() {
  /*Get bank data for references*/
  const bankQuery: Query = {
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
  /*Get Accounts*/
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
    },
    pagingInfo: { startIndex: 1, batchSize: 10 },
  };

  //Call apis
  const [banks, accounts] = await Promise.all([
    post("api/select", bankQuery),
    post("api/select", accountsQuery),
  ]);
  //Check for successful call
  {
    if (
      !accounts.success ||
      !Array.isArray(accounts.data) ||
      accounts.data.length === 0
    ) {
      return (
        <Card className="bg-white">
          <div>Error occurred when retrieving account data</div>
        </Card>
      );
    }
  }
  //Get bank ref data
  const bankIds: number[] = banks.success
    ? banks.data.map((b: BankDT) => b.BankId)
    : [];
  const bankNames: string[] = banks.success
    ? banks.data.map((b: BankDT) => b.BankName)
    : [];

  //List the accounts in a card list  layout
  return (
    <Card className="bg-white">
      <PageHeader headerText="Accounts" />
      {accounts.data.map((item: AccountDT) => (
        <Card
          key={item.accountId}
          className="mt-5"
          link={`/Records/Accounts/${item.accountId}`}
        >
          <ColumnLayout>
            <Column>
              <div>
                <h1 className="underline">{item.accountName}</h1>
                <h3>{DisplayValue(item.bankId, bankIds, bankNames, null)}</h3>
              </div>
            </Column>
            <Column>
              <div>
                <h1
                  className={
                    item.balance > 0
                      ? "text-right text-2xl text-green-500"
                      : "text-right text-2xl text-red-500"
                  }
                >
                  <b>{FormateCurrency(item.balance)}</b>
                </h1>
              </div>
            </Column>
          </ColumnLayout>
        </Card>
      ))}
    </Card>
  );
}

export default Accounts;
