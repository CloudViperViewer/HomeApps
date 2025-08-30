/*Data type for bank data*/
//Casing is required for API conversions

export interface TransactionDT {
  TransactionID: number;
  TransactionTypeId: number;
  TransactionWith: string;
  DateTime: string;
  Value: number;
 OnOffBillId?: { Int16: number; Valid: boolean } | null;
RecurringPaymentId?: { Int16: number; Valid: boolean } | null;
  ViaPaypal: boolean;
}
