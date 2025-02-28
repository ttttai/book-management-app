export const BOOK_STATUS_NOT_PURCHASED = 0;
export const BOOK_STATUS_PURCHASED = 1;
export const BOOK_STATUS_READING = 2;
export const BOOK_STATUS_READ_COMPLETED = 3;

export const BOOK_STATUS_LABELS = {
  [BOOK_STATUS_NOT_PURCHASED]: "未購入",
  [BOOK_STATUS_PURCHASED]: "購入済",
  [BOOK_STATUS_READING]: "読書中",
  [BOOK_STATUS_READ_COMPLETED]: "読了",
} as const;

export const BOOK_STATUS_COLORS = {
  [BOOK_STATUS_NOT_PURCHASED]: "bg-gray-500",
  [BOOK_STATUS_PURCHASED]: "bg-yellow-500",
  [BOOK_STATUS_READING]: "bg-blue-500",
  [BOOK_STATUS_READ_COMPLETED]: "bg-green-500",
} as const;
