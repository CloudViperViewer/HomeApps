import React, { Children, ReactNode } from "react";

interface columnLayoutPop {
  children: ReactNode;
  className?: string;
}

function ColumnLayout({ children, className = "" }: columnLayoutPop) {
  const noOfColumns = React.Children.toArray(children).filter(
    (child) => React.isValidElement(child) && child.type.name === "column"
  ).length;
  return (
    <div className={`grid grid-cols-${noOfColumns} ${className}`}>
      {children}
    </div>
  );
}

export default ColumnLayout;
