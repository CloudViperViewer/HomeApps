/*Customisable card component*/

import { ReactNode } from "react";
import Link from "next/link";

interface CardProps {
  children: ReactNode;
  link?: string;
  className?: string;
}

function Card({ children, link = "", className = "" }: CardProps) {
  const baseStyle =
    "rounded-xl border border-gray-200 p-6 shadow transition outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2";

  const clickableStyle = "hover:shadow-lg cursor-pointer";
  const disabledStyle = "cursor-default";

  const content = (
    <div
      className={`${baseStyle} ${
        link ? clickableStyle : disabledStyle
      } ${className}`}
    >
      {children}
    </div>
  );

  return link ? (
    <Link href={link} className="block">
      {content}
    </Link>
  ) : (
    content
  );
}

export default Card;
