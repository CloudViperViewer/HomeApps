//Landing page for the bank Record list
import post, { Query } from "../../components/apis/post";
import Card from "../../components/card/card";
import { BankDT } from "@/app/components/dataTypes/BankDT";
import PageHeader from "@/app/components/headers/pageHeader";

//Types for safety
type BankRow = { bankName: string; bankId: number };

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
        data.data.map((raw: BankRow) => {
          //convert to bank dt
          const item: BankDT = {
            bankId: raw.bankId,
            bankName: raw.bankName,
          };
          //Construct Bank Card
          return (
            <Card key={item.bankId} className="mt-5">
              {item.bankName}
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
