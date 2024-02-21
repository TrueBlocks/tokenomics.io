Location of Giveth API:

https://givback.develop.giveth.io/api-docs/#/default/get_eligible_donations

curl -X 'GET' \
  'https://givback.develop.giveth.io/eligible-donations?startDate=2021%2F01%2F01-00%3A00%3A00&endDate=2022%2F04%2F27-00%3A00%3A00&download=yes' \
  -H 'accept: application/json'

Fields:
"amount","currency","createdAt","valueUsd","giverAddress","txHash","network","source","giverName","giverEmail","projectLink","info"


Purple List:

curl -X 'GET' 'https://givback.develop.giveth.io/purpleList' -H 'accept: application/json'

List of addresses in JSON

