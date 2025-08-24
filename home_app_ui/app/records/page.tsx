import React from "react";
import Card from "../components/card/card";

function Records() {
  const content = [
    {
      name: "Banks",
      link: "/Records/Banks",
    },
    {
      name: "Accounts",
      link: "/Records/Accounts",
    },
  ];
  return (
    <>
      <h1>Records</h1>
      {content.map((item, index) => (
        <Card key={index} link={item.link} className="mt-5">
          {item.name}
        </Card>
      ))}
    </>
  );
}

export default Records;
