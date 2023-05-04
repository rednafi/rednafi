/* Import the latest 5 blog posts from rss feed */

import fetch from "node-fetch";
import { Parser } from "xml2js";
import { promises } from "fs";

const rssUrl = "https://rednafi.github.io/index.xml";
const outputFile = "README.md";
const parser = new Parser();

// Define an async function to get and parse the rss data
async function getRssData() {
  try {
    const res = await fetch(rssUrl);
    const data = await res.text();
    return await parser.parseStringPromise(data);
  } catch (err) {
    console.error(err);
  }
}

// Define an async function to write the output file
async function writeOutputFile(output) {
  try {
    await promises.writeFile(outputFile, output);
    console.log(`Saved ${outputFile}`);
  } catch (err) {
    console.error(err);
  }
}

// Call the async functions
getRssData()
  .then((result) => {
    // Get the first five posts from the result object
    const posts = result.rss.channel[0].item.slice(0, 5);

    // Initialize an empty output string
    let output = "";

    // Add a title to the output string
    output += `Healthcare hacker by day and OSS necromancer by night. SWE at [Dendi](https://dendisoftware.com/) and writing here on my [blog](rednafi.github.io). Find me on Twitter [@rednafi]((https://twitter.com/rednafi)).\n\n`;

    // Add a header row to the output string
    output += "| Title | Published On |\n";
    output += "| ----- | ------------ |\n";

    // Loop through the posts and add a row for each post to the output string
    for (let post of posts) {
      // Strip the time from the pubDate
      const date = post.pubDate[0].slice(0, 16);
      output += `| [${post.title}](${post.link}) | ${date} |\n`;
    }
    // Call the writeOutputFile function with the output string
    writeOutputFile(output);
  })
  .catch((err) => {
    // Handle the error
    console.error(err);
  });