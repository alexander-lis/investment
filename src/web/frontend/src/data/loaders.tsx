import { User } from "../utils/types-common";
import { Portfolio } from "../utils/types-stock";

export async function initialLoader() {
  const user: User = {
    Name: "alexander",
  };
  const portfolios: Portfolio[] = [
    {
      ID: 1,
      Name: "На квартиру",
    },
    {
      ID: 2,
      Name: "На квартиру",
    },
  ];
  return { user, portfolios };
}

// sum.js
export function sum(a: number, b: number) {
  return a + b
}