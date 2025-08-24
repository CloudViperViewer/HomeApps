import { headers } from "next/headers";
import React from "react";

export interface PostResponse {
  data: any;
  success: boolean;
}

export interface filter {
  operator: string;
  field: string;
  value: any;
}

export interface logicalExpression {
  operator: string;
  filters: filter[];
  logicalExpressions: logicalExpression[];
}

export interface pagingInfo {
  startIndex: number;
  batchSize: number;
}

export interface Query {
  table: string;
  fields: string[];
  logicalExpression: logicalExpression;
  pagingInfo: pagingInfo;
}

const post = async (endPoint: string, query: Query) => {
  //Try call
  try {
    const req = await fetch("http://192.168.0.171:8081/" + endPoint, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        table: query.table,
        fields: query.fields,
        logicalExpression: query.logicalExpression,
        pagingInfo: query.pagingInfo,
      }),
    });

    //Check if request is okay
    if (!req.ok) {
      const text = await req.text().catch(() => "");
      const data: PostResponse = {
        success: false,
        data: null,
      };
      return data;
    }

    const data: PostResponse = await req.json();
    return data;
  } catch (err: any) {
    //Catch error
    const data: PostResponse = {
      success: false,
      data: null,
    };
    return data;
  }
};

export default post;
