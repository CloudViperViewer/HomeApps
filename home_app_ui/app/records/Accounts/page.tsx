import post, { Query } from "../../components/apis/post";
import Card from "../../components/card/card";
import { DisplayValue } from "../../functions/displayValue";
import { AccountDT } from "@/app/components/dataTypes/AccountDT";
import { BankDT } from "@/app/components/dataTypes/BankDT";
import ColumnLayout from "@/app/components/columnLayout/columnLayout";
import Column from "@/app/components/columnLayout/column";
import PageHeader from "@/app/components/headers/pageHeader";

async function Accounts() {
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
  //Call api
  const banks = await post("api/select", bankQuery);
  const accounts = await post("api/select", accountsQuery);
  const bankIds: number[] = banks.data.map((b: BankDT) => b.BankId);
  const bankNames: number[] = banks.data.map((b: BankDT) => b.BankName);

  return (
    <Card className="bg-white">
      <PageHeader headerText="Accounts" />
      {accounts.success === true ? (
        accounts.data.map((item: AccountDT) => (
          <Card
            key={item.accountId}
            className="mt-5"
            link={`/Records/Accounts/${item.accountId}`}
          >
            <ColumnLayout>
              <Column>
                <div>
                  <h1 className="underline">{item.accountName}</h1>
                  <h3>{DisplayValue(item.bankId, bankIds, bankNames, "")}</h3>
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
                    <b>${item.balance}</b>
                  </h1>
                </div>
              </Column>
            </ColumnLayout>
          </Card>
        ))
      ) : (
        <div>Failed</div>
      )}
    </Card>
  );
}

export default Accounts;
