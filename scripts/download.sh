#!/usr/bin/bash
curl 'https://backend.takeuforward.org/api/sheets/double/strivers_a2z_sheet' \
  -H 'accept: */*' \
  -H 'accept-language: en-US,en;q=0.9,ms;q=0.8' \
  -H 'cookie: _ga=GA1.1.614272531.1720652522; mp_0e970f61db79f122cdae972d86ce5e20_mixpanel=%7B%22distinct_id%22%3A%20%22%24device%3A1909ee1e407698-00b21356ea5357-11462c6f-100200-1909ee1e407698%22%2C%22%24device_id%22%3A%20%221909ee1e407698-00b21356ea5357-11462c6f-100200-1909ee1e407698%22%2C%22%24initial_referrer%22%3A%20%22%24direct%22%2C%22%24initial_referring_domain%22%3A%20%22%24direct%22%2C%22__mps%22%3A%20%7B%7D%2C%22__mpso%22%3A%20%7B%22%24initial_referrer%22%3A%20%22%24direct%22%2C%22%24initial_referring_domain%22%3A%20%22%24direct%22%7D%2C%22__mpus%22%3A%20%7B%7D%2C%22__mpa%22%3A%20%7B%7D%2C%22__mpu%22%3A%20%7B%7D%2C%22__mpr%22%3A%20%5B%5D%2C%22__mpap%22%3A%20%5B%5D%7D; _ga_51P1R4XNJ0=GS1.1.1733313555.16.1.1733313588.0.0.0' \
  -H 'dnt: 1' \
  -H 'if-none-match: W/"5731e-38NMStkYS09MhL0pUD04T9cYOvw"' \
  -H 'origin: https://takeuforward.org' \
  -H 'priority: u=1, i' \
  -H 'referer: https://takeuforward.org/' \
  -H 'sec-ch-ua: "Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-site' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36' > $1
