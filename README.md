# Cricket
Zebi Cricket NFT Platform
Wrapper APIs :

status_code: 
200 - OK
400 - Bad Request
401 - Unauthorized
Headers: Username & Password

1. NFT denom creation:
Command to be executed in cricd:
Eg:- 
 cricd tx nft issue test123 --from=cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890 --name=testnft1 --chain-id Test-NFT-Cricket --fees=500.0cric --schema=<schema-content or path to schema.json>

Then it will ask for confirmation:
confirm transaction before signing and broadcasting [y/N]: y

Cmd line response:
{"height":"173490","txhash":"6FE6331F097EF092C72EFC24CC1279858B1723B81064E08E9EFC4E1C5F2FF8F3","codespace":"","code":0,"data":"0A0D0A0B69737375655F64656E6F6D","raw_log":"[{\"events\":[{\"type\":\"issue_denom\",\"attributes\":[{\"key\":\"denom_id\",\"value\":\"test123\"},{\"key\":\"denom_name\",\"value\":\"testnft1\"},{\"key\":\"creator\",\"value\":\"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"issue_denom\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"issue_denom","attributes":[{"key":"denom_id","value":"test123"},{"key":"denom_name","value":"testnft1"},{"key":"creator","value":"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890"}]},{"type":"message","attributes":[{"key":"action","value":"issue_denom"},{"key":"module","value":"nft"},{"key":"sender","value":"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890"}]}]}],"info":"","gas_wanted":"200000","gas_used":"57393","tx":null,"timestamp":""}

API1
Endpoint: ipaddress:portno/nft/createDenom
 
Request Body: 
{
“nft_denom_name”:”zebinft”,
“acct_addr”: “cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh”,
“nft_denom_id”:”abcd1234”,
“Chain_id”: “Test-NFT-Cricket”,
“Fees”:”500cric”,
“schema”:”schema-content or path to schema.json”
}

Expected Response: 
{
"status_code": 200,
"status": "success/fail",
"creator":"cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh",
"denom_id":"abcd1234",
"denom_name":"zebinft",
"Height":"173490",
"txhash":"6FE6331F097EF092C72EFC24CC1279858B1723B81064E08E9EFC4E1C5F2FF8F3"
}

2.NFT token minting

cricd tx nft mint test123 testtoken123 --recipient=cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890 --from=cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890 --chain-id Test-NFT-Cricket --fees=500.0cric
{"body":{"messages":[{"@type":"/zebi.cric.nft.MsgMintNFT","id":"testtoken123","denom_id":"test123","name":"","uri":"","data":"","sender":"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890","recipient":"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"cric","amount":"500"}],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
{"height":"174200","txhash":"679DF47FB65A4DC213D89D31FF0EF703FCC2C4D51B3DC166F1DA4364A9067812","codespace":"","code":0,"data":"0A0A0A086D696E745F6E6674","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"mint_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890\"}]},{\"type\":\"mint_nft\",\"attributes\":[{\"key\":\"token_id\",\"value\":\"testtoken123\"},{\"key\":\"denom_id\",\"value\":\"test123\"},{\"key\":\"token_uri\"},{\"key\":\"recipient\",\"value\":\"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"mint_nft"},{"key":"module","value":"nft"},{"key":"sender","value":"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890"}]},{"type":"mint_nft","attributes":[{"key":"token_id","value":"testtoken123"},{"key":"denom_id","value":"test123"},{"key":"token_uri","value":""},{"key":"recipient","value":"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890"}]}]}],"info":"","gas_wanted":"200000","gas_used":"60993","tx":null,"timestamp":""}


<denom-id> <token-id> --uri=<uri> --recipient=<recipient> --from=<key-name> --chain-id=<chain-id> --fees=<fee>

API2

Endpoint: ipaddress:portno/nft/mintToken
Request Body: 
{
“nft_token_id”:”zebinfttoken”,
“uri”:”path to the nft token”,
“from_acct”:”cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh”
“rec_acct”: “cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh”,
“nft_denom_id”:”abcd1234”,
“Chain_id”: “Test-NFT-Cricket”,
“Fees”:”500cric”
}
Expected Response: 
{
"status_code": 200,
"status": "success/fail",
"creator":"cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh",
"denom_id":"abcd1234",
"nft_token_id":"testtoken123",
"Height":"173490",
"txhash":"6FE6331F097EF092C72EFC24CC1279858B1723B81064E08E9EFC4E1C5F2FF8F3"
}
  
3.Burn
cricd tx nft burn <denom-id> <token-id> --from=<key-name> --chain-id=<chain-id> --fees=<fee>

cricd tx nft burn test123 testtoken123 --from=cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890 --chain-id Test-NFT-Cricket --fees=500.0cric
{"body":{"messages":[{"@type":"/zebi.cric.nft.MsgBurnNFT","id":"testtoken123","denom_id":"test123","sender":"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"cric","amount":"500"}],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
{"height":"189989","txhash":"A80EED83A2A66752452A8E5791EBD572BEA76CD78AF818D79B8B86E42F08440F","codespace":"","code":0,"data":"0A0A0A086275726E5F6E6674","raw_log":"[{\"events\":[{\"type\":\"burn_nft\",\"attributes\":[{\"key\":\"denom_id\",\"value\":\"test123\"},{\"key\":\"token_id\",\"value\":\"testtoken123\"},{\"key\":\"owner\",\"value\":\"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"burn_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"burn_nft","attributes":[{"key":"denom_id","value":"test123"},{"key":"token_id","value":"testtoken123"},{"key":"owner","value":"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890"}]},{"type":"message","attributes":[{"key":"action","value":"burn_nft"},{"key":"module","value":"nft"},{"key":"sender","value":"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890"}]}]}],"info":"","gas_wanted":"200000","gas_used":"55456","tx":null,"timestamp":""}

API3

Endpoint: ipaddress:portno/nft/burnToken
Request Body: 
{
“nft_token_id”:”zebinfttoken”,
“from_acct”:”cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh”
“nft_denom_id”:”abcd1234”,
“chain_id”: “Test-NFT-Cricket”,
“fees”:”500cric”
}
  
Expected Response: 
{
"status_code": 200,
"status": "success/fail",
"denom_id":"abcd1234",
"nft_token_id":"testtoken123",
"Height":"173490",
"txhash":"6FE6331F097EF092C72EFC24CC1279858B1723B81064E08E9EFC4E1C5F2FF8F3"
}

4. Edit nft token
cricd tx nft edit <denom-id> <token-id> --uri=<uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>


API4

Endpoint: ipaddress:portno/nft/editToken
 
Request Body: 


{
“nft_token_id”:”zebinfttoken”,
“from_acct”:”cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh”
“nft_denom_id”:”abcd1234”,
“chain_id”: “Test-NFT-Cricket”,
“uri”:”uri to token”
“fees”:”500cric”
}

Expected Response: 
{
"status_code": 200,
"status": "success/fail",
"denom_id":"abcd1234",
"nft_token_id":"testtoken123",
"Height":"173490",
"txhash":"6FE6331F097EF092C72EFC24CC1279858B1723B81064E08E9EFC4E1C5F2FF8F3"
}

5. NFT Transfer
tx nft transfer <recipient> <denom-id> <token-id> --uri=<uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>

 cricd tx nft transfer cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh test123 testtoken123 --from=cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890 --chain-id=Test-NFT-Cricket --fees=50.0cric
{"body":{"messages":[{"@type":"/zebi.cric.nft.MsgTransferNFT","id":"testtoken123","denom_id":"test123","name":"[do-not-modify]","uri":"[do-not-modify]","data":"[do-not-modify]","sender":"cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890","recipient":"cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"cric","amount":"50"}],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
{"height":"191171","txhash":"16E70CD827A333E5986940BCD55378156A117BAD869B891845537E6CB798ABEA","codespace":"nft","code":23,"data":"","raw_log":"failed to execute message; message index: 0: not found NFT: test123: unknown nft collection","logs":[],"info":"","gas_wanted":"200000","gas_used":"52263","tx":null,"timestamp":""}

API5

Endpoint: ipaddress:portno/nft/transferToken
 
Request Body: 
{
“nft_token_id”:”zebinfttoken”,
“recepient”:”cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh”
“from_acct”:”cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh”,
“nft_denom_id”:”abcd1234”,
“chain_id”: “Test-NFT-Cricket”,
“uri”:”uri to token”
“fees”:”500cric”
}
Expected Response: 
{
"status_code": 200,
"status": "success/fail",
"denom_id":"abcd1234",
"nft_token_id":"testtoken123",
"Height":"173490",
"txhash":"6FE6331F097EF092C72EFC24CC1279858B1723B81064E08E9EFC4E1C5F2FF8F3"
}

API6:
To get all the denom values
 Eg:- cricd query nft denoms
Cmd line Resp:
denoms:
- creator: cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh
  id: abcd1234234
  name: zebinft
  schema: ""
- creator: cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890
  id: test123
  name: testnft1
  schema: ""
- creator: cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890
  id: test1234
  name: testnft
  schema: ""
 pagination:
  next_key: null
  total: "0"

Get API
Endpoint: ipaddress:portno/nft/getDenoms
 

Expected Response:  array of jsons
[
 { “creator”: “cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh”
  “id”: “abcd1234234”
  “name”: “zebinft”
  “schema”: ""
},
{ “creator”: “cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890”
  “id”: “test123”
  “name”: “testnft1”
  “schema”: ""
},
{ “creator”: “cric1pdt5g5c2gjp04sy3xtcx3a5s5nwmf4cr93c890”
  “id”: “test1234”
  “name”: “testnft”
  “schema”: ""
}

]

API7:
To get the details of an NFT token 
 Eg:- cricd query nft denom abcd1234234
Cmd line response:
creator: cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh
id: abcd1234234
name: zebinft
schema: ""

Get API
Endpoint: ipaddress:portno/nft/getnftDetails

Expected Response: 

 { “creator”: “cric1mc3t0cye333achqg6enw2nkfmftxja4x3mt0zh”
  “id”: “abcd1234234”
  “name”: “zebinft”
  “schema”: ""
}
