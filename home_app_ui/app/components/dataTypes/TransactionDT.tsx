export interface TransactionDT {
  TransactionID: number;
  TransactionTypeId: number;
  TransactionWith: string;
  DateTime: Date;
  Value: number;
  OnOffBillId: {
    Int16: number;
  };
  RecurringPaymentId: number;
  ViaPaypal: boolean;
}
