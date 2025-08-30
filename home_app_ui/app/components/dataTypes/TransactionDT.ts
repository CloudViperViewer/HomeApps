/*Data type for bank data*/
//Casing is required for API conversions

export type NullInt16 = { Int16: number; Valid: boolean };

export interface TransactionDT {
  transactionId: number;
  transactionTypeId: number;
  transactionWith: string;
  dateTime: string;
  value: number;
 onOffBillId?: NullInt16 | null;
recurringPaymentId?: NullInt16 | null;
  viaPaypal: boolean;
}
