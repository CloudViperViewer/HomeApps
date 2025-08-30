/*Main Page for record list*/

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
      <nav aria-label="Records">
        <ul>
          {content.map((item, index) => (
            <li key={index} className="mt-5">
              <Card link={item.link}>{item.name}</Card>
            </li>
          ))}
        </ul>
      </nav>
    </Card>
  );
}

export default Records;
