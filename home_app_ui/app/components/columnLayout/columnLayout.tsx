import React, { ReactNode } from "react";

interface ColumnLayoutPops {
  children: ReactNode;
  className?: string;
}

function ColumnLayout({ children, className = "" }: ColumnLayoutPops) {
  const noOfColumns = React.Children.toArray(children).filter(
    (child) =>
      React.isValidElement(child) && (child.type as any)._IS_COLUMN_ === true
  ).length;
  const columns = Math.max(1, noOfColumns);
  return (
    <div
      className={`grid ${className}`}
      style={{ gridTemplateColumns: `repeat(${columns}, minmax(0, 1fr))` }}
    >
      {children}
    </div>
  );
}

export default ColumnLayout;
