import ky from "https://cdn.skypack.dev/ky@0.28.5?dts";
import dayjs from "https://cdn.skypack.dev/dayjs@1.10.4";
import { READ_USER_TOKEN } from "./env.ts";

const query = `
query($userName:String!, $date:DateTime) {
  user(login: $userName){
    contributionsCollection(from: $date, to: $date) {
      contributionCalendar {
        totalContributions
      }
    }
  }
}
`;

const now = dayjs().toISOString();
const username = "yasudanaoya"
const variables = `
{
  "userName": "${username}",
  "date": "${now}"
}`;

const url = "https://api.github.com/graphql";
const json = { query, variables };

const { data } = await ky.post(url, {
  headers: {
    Authorization: `Bearer ${READ_USER_TOKEN}`,
  },
  json,
}).json()

console.log(data.user.contributionsCollection.contributionCalendar.totalContributions);
