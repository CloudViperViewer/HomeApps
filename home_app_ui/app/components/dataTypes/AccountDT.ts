/*Data type for account data*/
//Casing is required for API conversions

export interface AccountDT {
  accountId: number;
  bankId: number;
  balance: number | string;
  accountName: string;
  bsb: string;
  accountNumber: string;
}
