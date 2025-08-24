import { ReactNode } from "react";
import Link from "next/link";

interface CardProps {
  children: ReactNode;
  link?: string;
  className?: string;
}

function Card({ children, link = "", className = "" }: CardProps) {
  const baseStyle =
    "rounded-xl border border-gray-200 bg-white p-6 shadow transition";

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
