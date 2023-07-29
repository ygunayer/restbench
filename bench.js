const autocannon = require('autocannon');
const fs = require('fs');
const path = require('path');

const outputDir = path.join(__dirname, 'results');

const testName = process.argv[2];

if (!testName) {
  throw `Pleace specify a test name`;
}

const outputFilename = path.join(outputDir, `bench-${testName}-${Date.now()}.json`);

autocannon({
  url: 'http://0.0.0.0:4000/api/v1/auth/register',
  method: 'POST',
  idReplacement: true,
  connections: 64,
  workers: 64,
  headers: {
    'content-type': 'application/json',
  },
  body: JSON.stringify({
    name: 'Yalin Gunayer',
    email: `yalin.gunayer+[<id>]@gmail.com`,
    password: '12345678',
    password_confirmation: '12345678',
  }, null, 0)
}, (err, result) => {
  if (err) {
    console.error(err);
    process.exit(-1);
  }

  fs.writeFileSync(outputFilename, JSON.stringify(result, null, 4), 'utf-8');
  console.log(`Wrote output to ${outputFilename}`);
})
