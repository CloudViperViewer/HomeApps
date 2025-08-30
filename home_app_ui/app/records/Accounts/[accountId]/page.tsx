import Card from "@/app/components/card/card";
import { Query } from "@/app/components/apis/post";
import post from "@/app/components/apis/post";
import PageHeader from "@/app/components/headers/pageHeader";
import { AccountDT } from "@/app/components/dataTypes/AccountDT";
import ColumnLayout from "@/app/components/columnLayout/columnLayout";
import Column from "@/app/components/columnLayout/column";
import ReadOnlyTextField from "@/app/components/readOnlyFields/readOnlyTextField";
import { BankDT } from "@/app/components/dataTypes/BankDT";
import { TransactionDT } from "@/app/components/dataTypes/TransactionDT";
import { FormateDateTime } from "@/app/functions/formatDateTime";

type Params = { accountId: string };

async function AccountSummary({ params }: { params: Params }) {
  /*Get Account Details*/
  const accountsQuery: Query = {
    table: "Account",
    fields: [
      "BankID",
      "AccountName",
      "Balance",
      "AccountID",
      "AccountTypeID",
      "AccountNumber",
      "BSB",
    ],
    logicalExpression: {
      operator: "AND",
      filters: [
        {
          operator: "=",
          field: "AccountID",
          value: [params.accountId],
        },
      ],
    },
    pagingInfo: { startIndex: 1, batchSize: 1 },
  };
  const accountCall = await post("api/select", accountsQuery);
  const accountData: AccountDT = accountCall.data.map(
    (account: AccountDT) => account
  )[0];
  /*Get bank details*/
  const bankQuery: Query = {
    table: "Bank",
    fields: ["BankName"],
    logicalExpression: {
      operator: "AND",
      filters: [
        {
          operator: "=",
          field: "BankID",
          value: [accountData.bankId],
        },
      ],
    },
    pagingInfo: { startIndex: 1, batchSize: 10 },
  };
  const bankCall = await post("api/select", bankQuery);
  const bankData: BankDT = bankCall.data.map((bank: BankDT) => bank)[0];
  /*Get Transaction Details*/
  const transactionQuery: Query = {
    table: "Transaction",
    fields: [
      "TransactionID",
      "TransactionTypeId",
      "Value",
      "RecurringPaymentId",
      "OnOffBillId",
      "ViaPaypal",
      "DateTime",
      "TransactionWith",
    ],
    logicalExpression: {
      operator: "AND",
      filters: [
        {
          operator: "=",
          field: "AccountId",
          value: [params.accountId],
        },
      ],
    },
    pagingInfo: { startIndex: 1, batchSize: 25 },
  };
  const transactionCall = await post("api/select", transactionQuery);

  return (
    <Card className="bg-white">
      {accountCall.success ? (
        <>
          <PageHeader headerText={accountData.accountName} />
          <div className="p-3">
            <ColumnLayout className="mt-5">
              <Column>
                <ReadOnlyTextField
                  label="Account Name"
                  value={accountData.accountName}
                />
                <ReadOnlyTextField label="Account Type" value="Place Holder" />
                <ReadOnlyTextField label="BSB" value={accountData.BSB} />
                <ReadOnlyTextField
                  label="Account Number"
                  value={accountData.accountNumber}
                />
                <ReadOnlyTextField
                  label="Bank"
                  value={bankCall.success && bankData.BankName}
                />
              </Column>
              <Column>
                <ReadOnlyTextField
                  label="Balance"
                  value={`$${accountData.balance}`}
                />
              </Column>
            </ColumnLayout>
            <Card className="mt-5">
              <table className="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
                <thead className="text-xs text-gray-700 uppercase bg-gray-100 dark:bg-gray-700 dark:text-gray-400">
                  <tr>
                    <th scope="col" className="px-6 py-3 rounded-s-lg">
                      Transaction Type
                    </th>
                    <th className="px-6 py-3 rounded-s-lg">Transaction With</th>
                    <th className="px-6 py-3 rounded-s-lg">
                      Transaction Date Time
                    </th>
                    <th className="px-6 py-3 rounded-s-lg">Value</th>
                    <th className="px-6 py-3">Bill Reference</th>
                    <th className="px-6 py-3 rounded-e-lg">Paypal</th>
                  </tr>
                </thead>
                <tbody>
                  {transactionCall.data.map((transaction: TransactionDT) => (
                    <tr
                      className="bg-white dark:bg-gray-800"
                      key={transaction.TransactionID}
                    >
                      <td
                        scope="row"
                        className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
                      >
                        {transaction.TransactionTypeId}
                      </td>
                      <td className="px-6 py-4">
                        {transaction.TransactionWith}
                      </td>
                      <td className="px-6 py-4">
                        {FormateDateTime(transaction.DateTime)}
                      </td>
                      <td className="px-6 py-4">{transaction.Value}</td>
                      <td className="px-6 py-4">
                        {transaction.OnOffBillId.Int16}
                      </td>
                      <td className="px-6 py-4">{transaction.ViaPaypal}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </Card>
          </div>
        </>
      ) : (
        <div>Error occured when retrieving account data</div>
      )}
    </Card>
  );
}

export default AccountSummary;
