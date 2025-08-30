import React from "react";
import Card from "../components/card/card";
import PageHeader from "../components/headers/pageHeader";

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
    <Card className="bg-white">
      <PageHeader headerText="Records" />
      {content.map((item, index) => (
        <Card key={index} link={item.link} className="mt-5">
          {item.name}
        </Card>
      ))}
    </Card>
  );
}

export default Records;
