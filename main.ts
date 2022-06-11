import { ky, dayjs } from "./deps.ts";
import { READ_USER_TOKEN, WEBHOOK_URL } from "./env.ts";

const query = `
query($userName:String!, $from:DateTime, $to:DateTime) {
  user(login: $userName){
    contributionsCollection(from: $from, to: $to) {
      contributionCalendar {
        totalContributions
      }
    }
  }
}
`;

const from = dayjs().format("YYYY-MM-DDT00:00:00");
const to = dayjs().format("YYYY-MM-DDT23:59:59");
const username = "yasudanaoya"
const variables = `
{
  "userName": "${username}",
  "from": "${from}",
  "to": "${to}"
}`;

const url = "https://api.github.com/graphql";
const json = { query, variables };

const { data } = await ky.post(url, {
  headers: {
    Authorization: `Bearer ${READ_USER_TOKEN}`,
  },
  json,
}).json()

const totalContributions: number = data.user.contributionsCollection.contributionCalendar.totalContributions;

if (totalContributions === 0) {
  ky.post(WEBHOOK_URL, {
    headers: {
      "Content-Type": "application/json",
    },
    json: {
      text: '今日は草生やしてません。',
   },
  })
}
