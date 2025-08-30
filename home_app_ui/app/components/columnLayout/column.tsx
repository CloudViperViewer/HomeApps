/*Column Component to be used in Column Layout*/

import { ReactNode } from "react";

interface ColumnProps {
  children: ReactNode;
}

function Column({ children }: ColumnProps) {
  return <div>{children}</div>;
}

// Stable identity for ColumnLayout
Column.displayName = "Column";
Column._IS_COLUMN_ = true;

export default Column;
