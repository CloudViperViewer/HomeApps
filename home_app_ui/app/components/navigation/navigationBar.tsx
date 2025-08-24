import Link from "next/link";
import React from "react";

function NavigationBar() {
  const links = [
    {
      name: "Records",
      link: "/Records",
    },
  ];
  return (
    <div className="bg-primary-100 w-50 text-blue-100 p-4">
      <Link href="/" className="underline text-2xl">
        Home Apps
      </Link>
      <h3 className="mb-7">Budget</h3>
      {links.map((item, index) => (
        <Link key={index} href="/Records">
          <div className="text-lg hover:bg-primery-50">{item.name}</div>
        </Link>
      ))}
    </div>
  );
}

export default NavigationBar;
