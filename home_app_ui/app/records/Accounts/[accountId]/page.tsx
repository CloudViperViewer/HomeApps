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
import { formatDateTime } from "@/app/functions/formatDateTime";
import { formatCurrency } from "@/app/functions/formatCurrency";

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
          value: [Number(params.accountId)],
        },
      ],
    },
    pagingInfo: { startIndex: 1, batchSize: 1 },
  };
  const accountCall = await post("api/select", accountsQuery);
  //Check for successful call
  {
    if (
      !accountCall.success ||
      !Array.isArray(accountCall.data) ||
      accountCall.data.length === 0
    ) {
      return (
        <Card className="bg-white">
          <div>Error occurred when retrieving account data</div>
        </Card>
      );
    }
  }
  //Dereference account
  const accountData: AccountDT = accountCall.data[0] as AccountDT;
  /*Get bank query*/
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

  /*Transaction query*/
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
          value: [Number(params.accountId)],
        },
      ],
    },
    pagingInfo: { startIndex: 1, batchSize: 25 },
  };
  /*Get bank and transaction data*/
  const [bankCall, transactionCall] = await Promise.all([
    post("api/select", bankQuery),
    post("api/select", transactionQuery),
  ]);
  const bankData: BankDT | undefined =
    bankCall.success && Array.isArray(bankCall.data)
      ? (bankCall.data[0] as BankDT)
      : undefined;

  /*Return account data*/
  return (
    <Card className="bg-white">
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
              <ReadOnlyTextField label="BSB" value={accountData.bsb} />
              <ReadOnlyTextField
                label="Account Number"
                value={accountData.accountNumber}
              />
              <ReadOnlyTextField
                label="Bank"
                value={bankData?.bankName ?? ""}
              />
            </Column>
            <Column>
              <ReadOnlyTextField
                label="Balance"
                value={formatCurrency(accountData.balance)}
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
                {(transactionCall.success && Array.isArray(transactionCall.data)
                  ? transactionCall.data
                  : []
                ).map((transaction: TransactionDT) => (
                  <tr
                    className="bg-white dark:bg-gray-800"
                    key={transaction.transactionId}
                  >
                    <td
                      scope="row"
                      className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
                    >
                      {transaction.transactionTypeId}
                    </td>
                    <td className="px-6 py-4">{transaction.transactionWith}</td>
                    <td className="px-6 py-4">
                      {formatDateTime(transaction.dateTime)}
                    </td>
                    <td className="px-6 py-4">
                      {formatCurrency(transaction.value)}
                    </td>
                    <td className="px-6 py-4">
                      {transaction.onOffBillId?.Valid
                        ? transaction.onOffBillId.Int16
                        : ""}
                    </td>
                    <td className="px-6 py-4">
                      {transaction.viaPaypal ? "Yes" : "No"}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </Card>
        </div>
      </>
    </Card>
  );
}

export default AccountSummary;
