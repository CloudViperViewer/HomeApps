import { ReactNode } from "react";

interface columnProps {
  children: ReactNode;
}

function column({ children }: columnProps) {
  return <div>{children}</div>;
}

export default column;
