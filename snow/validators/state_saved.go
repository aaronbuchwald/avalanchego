package validators

import (
	"encoding/json"

	"github.com/ava-labs/avalanchego/ids"
)

// expected validator set from https://avalancheavax.slack.com/archives/C06C8SEJN5S/p1706063824354109
// this is the P-Chain validator set from a good node at height https://subnets.avax.network/p-chain/block/11269268
// I had to modify this to remove the "" around the weight values so it could be unmarshalled.
var ExpectedVdrSetQ map[ids.NodeID]*GetValidatorOutput

var j = `{
	"NodeID-12dyQ7nhRzsNSiFzEoW1RWK819Zkssf5g": {
	  "publicKey": null,
	  "weight": 7943179093219
	},
	"NodeID-12qs7z78gQu9yo5vR4acLEbWwKDfmQYQq": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-13XrVPjS5dVxBKaweeG94SjY1Q6yeM4EJ": {
	  "publicKey": null,
	  "weight": 8136763134383
	},
	"NodeID-15AuPME9h4AS8CAEDsokNE5KMjLovnphr": {
	  "publicKey": null,
	  "weight": 2100000000000
	},
	"NodeID-15ngGTze32Rsr3WgHyLHVgcdqnY2eyFo": {
	  "publicKey": null,
	  "weight": 2042668512762
	},
	"NodeID-168kTGVaSDD9CsPQzbqMpQg6zUvKatqr6": {
	  "publicKey": null,
	  "weight": 66842982030338
	},
	"NodeID-1A1GTyMcDcassSiTBZFVuoXs7AS5HcHu": {
	  "publicKey": null,
	  "weight": 2025200000000
	},
	"NodeID-1JZW5ZTiUmUioMEXV5fYLVpVZwBAJwc8": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-1Z67stQzn6v2hi1wD1Zd7nooPoqJE4es": {
	  "publicKey": null,
	  "weight": 510353962528201
	},
	"NodeID-1aA7BtLfTX4SRXaWR8HttP4z2UapE1R9": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-22k7HZj3D9DSAD7ujyvvuYn8XzWuwA1Hi": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-22wjev1roSt8jieZ4SW8rJLG3YiWyvbV3": {
	  "publicKey": null,
	  "weight": 10234413860681
	},
	"NodeID-23Xaok4Hr7SgTcCjjMnVsYJaN5cSZLR3Z": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-23jyZenUu8Fm26ebdyXEEVH6PDoEXJ9j2": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-24gzuDUUhGEsAGbxhMZNJJ2x5G1wDVHEo": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-25GWqcvqc8m2ZT5ULNWfan5SEJNFBZMcw": {
	  "publicKey": null,
	  "weight": 7009819906860
	},
	"NodeID-25RzatyYmpsm3QbYj9QdSWqdFQLKxdMR3": {
	  "publicKey": null,
	  "weight": 12501000000000
	},
	"NodeID-26VN5FUvCmFWEhHm2k85hHn4rDqyKsV3W": {
	  "publicKey": null,
	  "weight": 5674126229741
	},
	"NodeID-26jY6nu8uyVyQLQpi3SMdvSTbMTZeajHD": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-28wbL4a2ozWTgYXt7NekGT12h5wp6UyKg": {
	  "publicKey": null,
	  "weight": 12630370108516
	},
	"NodeID-2BSvxuZFVVSpjrezuABMPvC6cLMwH8TT8": {
	  "publicKey": null,
	  "weight": 8035303206451
	},
	"NodeID-2Bq98e7Q94vVWipc5tv9jCL4aMxL7pjxr": {
	  "publicKey": null,
	  "weight": 2070499130000
	},
	"NodeID-2EH9c6NwjHLZLWC2Ah3vekv3XUwcup62L": {
	  "publicKey": null,
	  "weight": 19942432667364
	},
	"NodeID-2F2gRAyyVkaNC7jtPzJ4zgcQh64QyMkvx": {
	  "publicKey": null,
	  "weight": 2513973373253
	},
	"NodeID-2FNBUKXo8ShvqRgGhEUF1FxYC95Rk5qtY": {
	  "publicKey": null,
	  "weight": 3653907014519
	},
	"NodeID-2HvkjoRiMzEcZgpWKNAruLKNv9r7Lb2pv": {
	  "publicKey": null,
	  "weight": 234909176470578
	},
	"NodeID-2KSietvSEq4mX2C8p1DAT83RTJPykxUqN": {
	  "publicKey": null,
	  "weight": 3766333811802
	},
	"NodeID-2KfgS6P7vf9L55fMRTbHPgS4ugVSDW3nj": {
	  "publicKey": null,
	  "weight": 11700000000000
	},
	"NodeID-2L2wFDTAv9Ti1wpqbFzqLbTa87dMw7QQM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2P5CC3b5NkbHVGvvnPDjn3eodiEco3Xjk": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2PBLMD1vY91yAoVjCxwiRtngTrL3cJpMj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2PZiqgTALZcdU1TUvQeypFARYEgjkQZPg": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2Q8DS1hgPgsaMB1y8VqmxeworYJ2h2Ubt": {
	  "publicKey": null,
	  "weight": 11085442428045
	},
	"NodeID-2THa3uLP7oEvBj19k4E4jm1tCpT4muhYi": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2UXMFjdXTw96iQHNL23PxH2hcLJUxnWic": {
	  "publicKey": null,
	  "weight": 8977001499907
	},
	"NodeID-2VJfSPqaevBmiJKSVDnw7GBwsdfgA1XuM": {
	  "publicKey": null,
	  "weight": 6773348031486
	},
	"NodeID-2WfvcQumS36rdyubDSioc8B5YinToKAaf": {
	  "publicKey": null,
	  "weight": 2625971877488127
	},
	"NodeID-2ZtHY1RPNrq1y5YuJLaCShK617g6CjRsU": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2aCXkeemxpBXiFui8BX1tVpyVsvhoZCU8": {
	  "publicKey": null,
	  "weight": 3867913362814
	},
	"NodeID-2bRK9nCjngVig5fMFjLhVgqf18L7dbTj6": {
	  "publicKey": null,
	  "weight": 2104398663100
	},
	"NodeID-2cH1ggEZ8Z6CoKoQsEeZudUy5kEihW4T2": {
	  "publicKey": null,
	  "weight": 2625613278242
	},
	"NodeID-2d3ZJcLUXLyZE11hw7FUi2YM3nUkMvvdP": {
	  "publicKey": null,
	  "weight": 3709000000000
	},
	"NodeID-2dSqwqZZEzAzrmLRUQFDjdGkcupgrp2PP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2eBu8RjbvCFTZMcxu1oh2oKtfHT7V4zyM": {
	  "publicKey": null,
	  "weight": 10000000000000
	},
	"NodeID-2eJ2XJDpUagguhE9wKFR9JUxYNaBk49Ko": {
	  "publicKey": null,
	  "weight": 70050000000000
	},
	"NodeID-2eraCsE3WijaEnkcFX77hbK494TneHXkY": {
	  "publicKey": null,
	  "weight": 148856814700234
	},
	"NodeID-2gRaAkSUa8fSrWBJ8i7muNepvyG6MPsSC": {
	  "publicKey": null,
	  "weight": 3832000000000
	},
	"NodeID-2hoNSdthtVU9RE6g4VJ7Y6tGdWm3s1PKN": {
	  "publicKey": null,
	  "weight": 707730553209342
	},
	"NodeID-2iVvryG4GyPR5XBQebYMMVTFYfsNKy9oB": {
	  "publicKey": null,
	  "weight": 3458999000000
	},
	"NodeID-2iWqUM3VWvrcTLyXi2KgBLVhunMvFW7vY": {
	  "publicKey": null,
	  "weight": 573632351047843
	},
	"NodeID-2iY9tRvYLGjeGtjnUGyGbMz7uoaBdZR58": {
	  "publicKey": null,
	  "weight": 2100000000000
	},
	"NodeID-2jnJ9jJrW2EULmCCNnZz88HHCAQYo1Bja": {
	  "publicKey": null,
	  "weight": 3729000000000
	},
	"NodeID-2mWa8ytKEtviCNwmHfSPmDvkTGPLH3kf7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2mZCDApt2CtfV6yJGu7zCa7XbCTC2iqpH": {
	  "publicKey": null,
	  "weight": 2034834059976
	},
	"NodeID-2nfH8Mmj5ZY6Tx5oD8ebst7prrEmbS4Ex": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2pN3EtqAUKWvJedQvYfPSgKeonNmFn8bA": {
	  "publicKey": null,
	  "weight": 248518103060353
	},
	"NodeID-2pNQqfaBqMqwWgeJiqPbmHZk1cUtWcjqb": {
	  "publicKey": null,
	  "weight": 6582618326200
	},
	"NodeID-2rpzBQ931ezzMPY8EKKivdZWfkiWtAyvZ": {
	  "publicKey": null,
	  "weight": 2009076077014
	},
	"NodeID-2rsWfeyECbxExQUEcrok2ff8d3tDz74PJ": {
	  "publicKey": null,
	  "weight": 60000000000000
	},
	"NodeID-2sNYafYrpqSjsGspW2emjd6TqNTtWz74S": {
	  "publicKey": null,
	  "weight": 134564754998530
	},
	"NodeID-2sspVTzGYqTkeiQeBsdFKv57gWpLg6Efg": {
	  "publicKey": null,
	  "weight": 13019418758004
	},
	"NodeID-2uA1cSeJhPtYgGtMUcnGgjesqAK7LzpQw": {
	  "publicKey": null,
	  "weight": 3000000000000
	},
	"NodeID-2vbdREFyZitz6LvSBca5gCe3eCPHZGq3b": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2ves7JhkWUcLfPWWtKVgh7uvJogz8HLnR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2wWroHMggzJvKh6t3tdPtJTTP9DNmdc4K": {
	  "publicKey": null,
	  "weight": 2707879206174
	},
	"NodeID-2xWeMjBPrnJejBo1vFiyNZzy2FcquZnae": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-2yC7fWX23kGjCd7ytqGzQQcRJf7xaj1N1": {
	  "publicKey": null,
	  "weight": 44997238080361
	},
	"NodeID-2yrtwg1pRGi9jB6ayaXgpNbMTeGHUyofw": {
	  "publicKey": null,
	  "weight": 886267530145986
	},
	"NodeID-2zzp8nVwqw7ssQsMXniESwjwpZYMzeFsa": {
	  "publicKey": null,
	  "weight": 2444808221259174
	},
	"NodeID-31xXC3YuXN5S9Bv9d63kH7ev1vuQKG3wR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-32EFwYSrpMmuxwDsgGf29ouKrcPquTQ5e": {
	  "publicKey": null,
	  "weight": 3206863360445
	},
	"NodeID-32aVnE9HYv1FesKKbvvUh1ZvFU3ARNtR2": {
	  "publicKey": null,
	  "weight": 250690623939877
	},
	"NodeID-32dhPQyEQArm4ybXTxAvAAWAap7DkZAPh": {
	  "publicKey": null,
	  "weight": 2106526325420
	},
	"NodeID-33miCHPn9eN8H9Yi4bCzEaL9Sc5BLeKMg": {
	  "publicKey": null,
	  "weight": 9376859340754
	},
	"NodeID-34HkMQ1oU1vf7wNpzY9xWCiazE2tGhyJs": {
	  "publicKey": null,
	  "weight": 264618096800761
	},
	"NodeID-34QP8d17f1XjB778dCTCsBfjL3T4PeMAf": {
	  "publicKey": null,
	  "weight": 280902928607238
	},
	"NodeID-35rq6ZWsLWs3coxC38LACEayLh99jxMav": {
	  "publicKey": null,
	  "weight": 7172142407501
	},
	"NodeID-36Vywf4J2tu9gA8xJCxZjwi2JABWSF7qo": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-36i6jQQTfB6Z8NXFPQJsY3FVukaSGKTDp": {
	  "publicKey": null,
	  "weight": 114975382096369
	},
	"NodeID-37ZMQ9ZZ4e7ZD1kmRg1WSTpRrEPSQ5LGT": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-38miyibfUUmFqzvmvuB9brEDE3Vm2FKmL": {
	  "publicKey": null,
	  "weight": 2004521368170
	},
	"NodeID-3AMVspPXZimHQ3yqWizDf1bcW3GjL2B2e": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3AUVAzMqQ66svzG1H7JVgURhUQHU8Y3r7": {
	  "publicKey": null,
	  "weight": 56821758146273
	},
	"NodeID-3BmiiYjMNvspTqKKo4fXJWbjDbaaWRNXE": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3CHc7PuHw5sGxtSVjujpHRxgiFDSFfucK": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3DwhPYMEQABuocceWDAZEpi8GMcLyvTYy": {
	  "publicKey": null,
	  "weight": 2258236260816004
	},
	"NodeID-3DyUWkRptB3CRUVHk39Ni6Dpr6QvGWXwA": {
	  "publicKey": null,
	  "weight": 38190152744731
	},
	"NodeID-3EzewMeb8MrVyWam8FhnHPHPoKsT7ERXM": {
	  "publicKey": null,
	  "weight": 2032124735980
	},
	"NodeID-3FhRFK6UxSfMED5EkK7jZT84pFZy9f17D": {
	  "publicKey": null,
	  "weight": 19997893708236
	},
	"NodeID-3HLAgdaA61zPrTx5yQ7Cc6waKWidsiqMT": {
	  "publicKey": null,
	  "weight": 6854349792249
	},
	"NodeID-3HvUXQy1siDNUBGWMYxwMfjufh8mxLtQY": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3K3PUAqo3cKxRoQyYto1EsXtuTHoDZ2B6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3KAxYX7JeLQgm1fwzVqbzjow6eNzSS9Aa": {
	  "publicKey": null,
	  "weight": 223371208291150
	},
	"NodeID-3KX9tgCEQcTHbC8W9yY4zY32Cj3ET9MNQ": {
	  "publicKey": null,
	  "weight": 5930094714349
	},
	"NodeID-3KyemFW2jou47TCfoJG4YQhtCDd6Si64q": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3M5i35f5u99QbBn6fP1FzeAto7NA4TFKt": {
	  "publicKey": null,
	  "weight": 2100420000000
	},
	"NodeID-3MFu1eLpGRrRcWksHJLBuk516Vk5PoYgC": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3NJgvio7B47MB7BZWm31LHzbPWVdkEiEP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3NXS6ZAzHeqV7w4saG4vXxAd4tcbpmxfk": {
	  "publicKey": null,
	  "weight": 2100790443475
	},
	"NodeID-3PJY4Rpb5BRR3xSUvuU5Dj3cyHcAmzfjD": {
	  "publicKey": null,
	  "weight": 14435540302979
	},
	"NodeID-3QvnmD9KrJt7BcoYytWfCsCD83TmKo16d": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3ShFK7JbJ2LN2gFT2W4iXp4N2NVuP4vZC": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3TWEGuLyjvyKYLQqXNvQ7GQqSNoSQrviN": {
	  "publicKey": null,
	  "weight": 6818199398138
	},
	"NodeID-3U6UcmBu6WVCgodXYsH72VwVKpgE4XomQ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3U945Ju6EUzVaiA25ea2BvaapV111iqcW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3V8B3h9bS1cQerQMF47sw8Tr9jmJz4uHG": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3ViTS5vVn4uLoQQ5d7Qs8mF5eYvgQcLQe": {
	  "publicKey": null,
	  "weight": 8298199267576
	},
	"NodeID-3W61mLt8G9fVa3VNnNiStsbqmkeeSPQ5U": {
	  "publicKey": null,
	  "weight": 34798997088103
	},
	"NodeID-3YkLrm1D9MqZ6K4YTwaQkt1NBb6wP9Ldx": {
	  "publicKey": null,
	  "weight": 264000000000000
	},
	"NodeID-3Zbddv2qkn6gWAfd12ysNN7N9EsDTjc9v": {
	  "publicKey": null,
	  "weight": 5732514315904
	},
	"NodeID-3b2dLL4mSiGK2gx9pmML5n2Za4T9pLHe3": {
	  "publicKey": null,
	  "weight": 4077151284322
	},
	"NodeID-3bciNCJdZcW5k8jMtHujT6msFbGMotkSF": {
	  "publicKey": null,
	  "weight": 10276451620905
	},
	"NodeID-3cqFtBUEe4LgkL94wHotgwtEMhboQqxr5": {
	  "publicKey": null,
	  "weight": 6000000000000
	},
	"NodeID-3dan8fxgCWWJZpNUkrBBEw1vhL8E85h4i": {
	  "publicKey": null,
	  "weight": 4000000000000
	},
	"NodeID-3douD9GDP69zSq1eMzRAYfKCUVtjEEUCt": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3fntmpjeW5JhEgbVfKT9fy8xeaMg4tquo": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3g8sfLNgojMQbeHKgeAMD93HYyqotfT1P": {
	  "publicKey": null,
	  "weight": 5113266850713
	},
	"NodeID-3iXKmHPAMgJa7z7pGcPpMC6pAmx8fjM1q": {
	  "publicKey": null,
	  "weight": 660000000000000
	},
	"NodeID-3iknBWGJowmNu2d63Qv7mRM52xoQ3mCoB": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3izZtcg2iecpF5rEkQF7Mwofd4puVGZwz": {
	  "publicKey": null,
	  "weight": 5637952626042
	},
	"NodeID-3jstEycDqqUdVbc5uH8Zst7KninBndazA": {
	  "publicKey": null,
	  "weight": 4009150205198
	},
	"NodeID-3kGMSCRnasr5CjsgNm1D8FbZnY4C5iNRE": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3kTMjX1jJvdH2S94GHLi6Qxy1DKyLhuc3": {
	  "publicKey": null,
	  "weight": 4421761096230
	},
	"NodeID-3kUUeZB7umtaENekUfWxa1bxrAU2eq1AD": {
	  "publicKey": null,
	  "weight": 832053861179197
	},
	"NodeID-3kn4Q8wTARwg4A1LuLixLVZtauWD5rsGU": {
	  "publicKey": null,
	  "weight": 3253102864960
	},
	"NodeID-3mvtMQ554k7VbrbYFUaNZR4JmeCoggD5P": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3nuVL5JFrrumBeA3xqAk7BqTkRpu5mJsh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3oJ5XYknQTmDrwTkBxLx4DCuFgFKFP2h6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3oaegcGaKyncpiWQCzEVd7vN1Vcsue3Wc": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3p3GjN8gxjGNpFCEKyErnWoXJGENDdj7W": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3p3zJHazEyiVrhRXKL6jJds5esWnwzbW4": {
	  "publicKey": null,
	  "weight": 12500000000000
	},
	"NodeID-3pSYSkZy99npP3NS1UfES7beUM1y84EjP": {
	  "publicKey": null,
	  "weight": 19934417899567
	},
	"NodeID-3pfvvxwJWQb1w7KrtGfGhd9ZtLZhrBrBc": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3qUq9jKXL43HwdMEov6KBsnCjcvWXmAoF": {
	  "publicKey": null,
	  "weight": 3381719746161
	},
	"NodeID-3rw6bDxFFNVoRnZmBAnTV4ZEXnkjooVq4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3s22Xmy4kPYnaUb21BtmCaTVh4uVuQCKv": {
	  "publicKey": null,
	  "weight": 16204544921219
	},
	"NodeID-3tD5VZzL9FeFTJNvon3qeoCdDyX77drQ8": {
	  "publicKey": null,
	  "weight": 2013792846892
	},
	"NodeID-3tTQWaeudKziFWpbmzbawLJTunDsQnjnu": {
	  "publicKey": null,
	  "weight": 9683501210330
	},
	"NodeID-3tjM63VDvTaHkhzxWxV3zrnJjAfHFAgfC": {
	  "publicKey": null,
	  "weight": 2013672641872
	},
	"NodeID-3txc7u47G8EyCyD5aSZdTxeVdhAWSbpoa": {
	  "publicKey": null,
	  "weight": 8620627011296
	},
	"NodeID-3vkjzH23PgBPmHixW9xpdUD9vndf6ffBZ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-3wd8cyGCDmhuoZYWmNDab2FhAVpnKYKJE": {
	  "publicKey": null,
	  "weight": 4657250366579
	},
	"NodeID-3yiw3g5Rer7repmzoJaATJpkf7GnWD5j9": {
	  "publicKey": null,
	  "weight": 966741216439575
	},
	"NodeID-41CFoAtQx8Bqivuq3tkg5oqnj6WpBz32S": {
	  "publicKey": null,
	  "weight": 2523162521126
	},
	"NodeID-41mvUJUdqhRjP7pX9BB7R89PXkDhC7f7i": {
	  "publicKey": null,
	  "weight": 48066903692424
	},
	"NodeID-42kmqmMDkmMyw2q6gS1uLi4wiXdC2NLwW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-43AKDBv1R4hvnt9bjDFA4SEcwFMzZqpXX": {
	  "publicKey": null,
	  "weight": 39486077332666
	},
	"NodeID-43LKNkQ9avhKgVj7KrHXjq1bYi6mvxQ2C": {
	  "publicKey": null,
	  "weight": 4082539774528
	},
	"NodeID-45hNv286MEyVSjGrtiYLt3qqnj8G7FecS": {
	  "publicKey": null,
	  "weight": 2004555454436
	},
	"NodeID-45p6WjZk3E9Je9Sw4q4SvEaagbYF7Jsud": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-47pbycgGRRWtAB4FZ5fJZApoeG4nzv41U": {
	  "publicKey": null,
	  "weight": 39938954940902
	},
	"NodeID-48tcs7C6Q7sBPrCMJbhWNAMj7bEnnDVNE": {
	  "publicKey": null,
	  "weight": 2914999000000
	},
	"NodeID-49LTjmBTcdjMyD33u7gKkfREPqEhhfPaj": {
	  "publicKey": null,
	  "weight": 461686365510801
	},
	"NodeID-4AW19ZAJMCyr64UKfFAUhZXuZDtVshQ36": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4AWHc6b817tKesbKJ22EAEsJa4GkrDuNE": {
	  "publicKey": null,
	  "weight": 110970443262880
	},
	"NodeID-4CJDfSDWT9X3hTosPsgui1ZRkBNRhX6jV": {
	  "publicKey": null,
	  "weight": 2150000000000
	},
	"NodeID-4DHwFAw2xZ8HuSu3jFzMU9cXvNhUHCkZ9": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4G2p1zWDewtnyzCvEuQ27fvbdaFa8zRcf": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4GFFG65jrNUU3X6vsZKgUhmz5iCG21nyS": {
	  "publicKey": null,
	  "weight": 61436434106471
	},
	"NodeID-4GcMxoKhvXDebqgeZq2zKWPgQZF5aDPm4": {
	  "publicKey": null,
	  "weight": 144030281230479
	},
	"NodeID-4GkgAWZzSWHi3hZZLjLGes5HwuJ5FXDuj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4JcnzK8FQKGHGMt5EQHRWgBVxxdA4sFy7": {
	  "publicKey": null,
	  "weight": 326292789346898
	},
	"NodeID-4K6rXew2J41T5LTWmTi8AjMbUFsidXGpE": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4KKUrWvuKKEwNQWvZj7x6NDhUEan1m44C": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4MqtiVkCYGD4TmoBhH3b6a6UTGNrPhvvW": {
	  "publicKey": null,
	  "weight": 2098413318435
	},
	"NodeID-4PKUhJMxeL5C7A5epar2X5TAAzXLqH44r": {
	  "publicKey": null,
	  "weight": 2054514539024
	},
	"NodeID-4Q7EdeK1p9JkHULo5nZ9KVwgWYjDjDf9F": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4Q97SD75d14ikayvi4C21CkbdoaboGonV": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-4R4zmBEWqY3dCtKHmvDd56Dupq97UAwtP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4RV8eRbw9andBLK2og4rtt3W7txBmSuxA": {
	  "publicKey": null,
	  "weight": 2842551982857
	},
	"NodeID-4RVd14QquiKdEXitdrnTuZiYpaBY1W6QM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4S2uvFvPmHh2Q4f2To1XznG2HMsyohuA1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4S5sLovGppHvP9uv4v6jZHV7JtASRpUUk": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4SvFyvJPHPYvkMnJqBaNcwr5yuU8sCsem": {
	  "publicKey": null,
	  "weight": 9516868089474
	},
	"NodeID-4TSV8FnyRHrVAmPALXfvLnGaHGspS5W2R": {
	  "publicKey": null,
	  "weight": 8035303206449
	},
	"NodeID-4TaEnGcWM77nvEityjYDxB4zYdLQ6LiZ1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4Ubqsj2vfwdGUUYNg1jtYpkYNNLugNBQ9": {
	  "publicKey": null,
	  "weight": 5427372899083
	},
	"NodeID-4V2KBeNd58jdBXej3ohL8E5d1bNTTp4tT": {
	  "publicKey": null,
	  "weight": 11207924257280
	},
	"NodeID-4VpJH4PkkQmg7KrJ9z3czh2Uipg16PRYt": {
	  "publicKey": null,
	  "weight": 2375673271612
	},
	"NodeID-4Vs1rw4jmL5QersqfLr3qB9HmyNtfeVeK": {
	  "publicKey": null,
	  "weight": 5061708358607
	},
	"NodeID-4WGNJ4vv6bH3FLCJgovqn3D7RCJ8rKBDR": {
	  "publicKey": null,
	  "weight": 2013698903798
	},
	"NodeID-4Xw1ekhVqhHqzZ4pLTPtofbTZxpkvcQNi": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4Zw4yFm9gJwULZHsgYuG7AbFKUJ7wuKMh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4bKCgF7VcG9ZtcfMwAVFfBJoh2xDgvz1j": {
	  "publicKey": null,
	  "weight": 7624000000000
	},
	"NodeID-4bajtcpjERHRChiaYpovKU8XE5qAE4usY": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4btZGj8TmrycK22kwgBK5wJEFighAFWiZ": {
	  "publicKey": null,
	  "weight": 44589959877833
	},
	"NodeID-4cJyBoufFfiMShtRhcd6WGTYdvUCE3Ai4": {
	  "publicKey": null,
	  "weight": 8497496514516
	},
	"NodeID-4cuTK1XYjm1VMTitd4MPBcwZ7LYyiqjfd": {
	  "publicKey": null,
	  "weight": 708103941176356
	},
	"NodeID-4cw926TqMXDNo7QyraCShjWMXNSCDZqMQ": {
	  "publicKey": null,
	  "weight": 3000000000000
	},
	"NodeID-4cwQT5hvhnhgYM7kJia4MK5kT8XLVZSNz": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4dJTYbWkF82oeT3gYAjoVNNbaga2zXsy4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4e1C32U5TBUFuQpykL6rSw46RijbjdwRz": {
	  "publicKey": null,
	  "weight": 4050371485512
	},
	"NodeID-4eaj7e8pXaR1TSuCF5n7tKDCHPPemZGky": {
	  "publicKey": null,
	  "weight": 4392568761504
	},
	"NodeID-4edDKb3xKovZjTSuFbhNzJZ3Y9KTAzDaQ": {
	  "publicKey": null,
	  "weight": 2172781898693
	},
	"NodeID-4f7B178dgkbrTAcLVZxcW3SeNdnpW4ijp": {
	  "publicKey": null,
	  "weight": 2029678473383
	},
	"NodeID-4fK92LkTyEUzPoDW44Bo9b5YvL5kJ7369": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4gTwepTF5fcacXB7gdYZLTtfSFsYh4faj": {
	  "publicKey": null,
	  "weight": 1359768423955180
	},
	"NodeID-4givb8yw6262YETrNnm4hTSRPeK4qEBfk": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4k9FLLYj63sJNrrGycj6MTRm6JqyacEav": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4kCLS16Wy73nt1Zm54jFZsL7Msrv3UCeJ": {
	  "publicKey": null,
	  "weight": 1004857941176455
	},
	"NodeID-4kRVv9q9bLrAVCdpgVXJfEBEUXSGkobMU": {
	  "publicKey": null,
	  "weight": 2221000000000
	},
	"NodeID-4kmJpUz7Eimok8fLcWeWMYSrZnPgEnc1M": {
	  "publicKey": null,
	  "weight": 2336073528314
	},
	"NodeID-4m1TLqY9ob5u23YrJx9zWDj11dz6DmV19": {
	  "publicKey": null,
	  "weight": 12555000000000
	},
	"NodeID-4nc9Pi32BZwWxcM21U95mVHsbExbPSmkq": {
	  "publicKey": null,
	  "weight": 3267541116050
	},
	"NodeID-4ngWFEcMdBhXKj38YwJA29WS3mczGNrNd": {
	  "publicKey": null,
	  "weight": 2688504675058
	},
	"NodeID-4oA58fddyvuKuRtqanZtp8V1Sz9mbr6sS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4ofBXitDMQ6QZi83yvPjCYn5LG6HBwqSn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4pZqGB9JdyanFMVQEeq5VZ6YG9yZNsCF8": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4ppYtgngzEnrgBZ7JUVw8bTbaBD76efcR": {
	  "publicKey": null,
	  "weight": 7843831236142
	},
	"NodeID-4qzj4sLxsLnmhhktyTyR3BWXu8nXTABnY": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4soc4KzELnfnTLLw4FTgpEefZpE6aSQbw": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-4uHXbHesQbAzcDBRqaFyJ5kpgovHpU576": {
	  "publicKey": null,
	  "weight": 2589754826995
	},
	"NodeID-519R4edrXr2qJBZnFX7X9dJvo2JjJcesT": {
	  "publicKey": null,
	  "weight": 6000000000000
	},
	"NodeID-51wXjrXuQkpHqgnyZ6pgQVbiBwTWgVQZi": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-56PBmfRBSmT2sYsPLiBPKVq2fQUVGEg9g": {
	  "publicKey": null,
	  "weight": 268529091656800
	},
	"NodeID-56RnA8AJBddFSt2rrGu5WmsZF8qW4zNeP": {
	  "publicKey": null,
	  "weight": 639303668237286
	},
	"NodeID-56Ua5fd6jovnFkCc9LfzSMCNdkviX3EY9": {
	  "publicKey": null,
	  "weight": 3075174317361
	},
	"NodeID-57VYS3UDdnkurTnK9WVzfpNT4H4V2fCmT": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-57o8JLPG29tydZsDVGuD8iBZ3rVorYeoB": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-57vP1ZP91HztWSRTHurGdrMSw3TWHYumo": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-58VvvNEcptC2Rw5p1yqwgjwSWGfotrVbq": {
	  "publicKey": null,
	  "weight": 5790613491521
	},
	"NodeID-599kjqvHr1rEhyvT2mFA2whFJZS2Yoexg": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5AzNzec73Pth9Acw6jVd2BYZWyY5KjV6y": {
	  "publicKey": null,
	  "weight": 2114335649119478
	},
	"NodeID-5B2Uyysf1nWcRbiCYdKwXzevVF5a7sN6T": {
	  "publicKey": null,
	  "weight": 3784888397462
	},
	"NodeID-5C5QEUprYutWMbxYicmQd8dUFhc4eh6TW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5CDfrTrcGndVpY6qtuWsKWLm3j93GXyey": {
	  "publicKey": null,
	  "weight": 2585179982175
	},
	"NodeID-5CjaXtU43Zar53KzzGHqZrmihayAdbFcb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5CwgaeHgQgWbeaMB2ERYZ1ntrhKCmSs5D": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5EuG66om4jxQDK8hRFyq89SsMqfdvESew": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5GPPVZdPNj8bicXQQcYGCtksCSReksKaw": {
	  "publicKey": null,
	  "weight": 3968653447181
	},
	"NodeID-5H8JzjMq3xei3gU2BJ3DZjxYxgc44GYZ5": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5HkCsXm4TQUb9uXxM6Cjy1WzvuDSspPui": {
	  "publicKey": null,
	  "weight": 250142234978953
	},
	"NodeID-5LGFWx2kfwMZXpyG52FkRtrfYfVRimjh5": {
	  "publicKey": null,
	  "weight": 2558004448540809
	},
	"NodeID-5PMCtewYFfdxWmE5gcxuZeXAQAxHqMhu3": {
	  "publicKey": null,
	  "weight": 8170939693507
	},
	"NodeID-5PMGUqdapvGYEATmbQ48hMJTwcYyKFNDg": {
	  "publicKey": null,
	  "weight": 7811598289527
	},
	"NodeID-5Pc2rnbGHgdmVtThh6fn8JqGVJavn9CrR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5PeQVPLrJH9xiSQZ8hcmeXWtUzhcArEhQ": {
	  "publicKey": null,
	  "weight": 6058823026624
	},
	"NodeID-5Pg156uQvovbZQ3F6JKUiyAg5MrdFseMP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5PvfcnfPd3MLnpq4MujxvBQHPwtdjYk8s": {
	  "publicKey": null,
	  "weight": 35000000000000
	},
	"NodeID-5PxpmHkfB3gNh3spZWuzEKis3DJ8WKxLt": {
	  "publicKey": null,
	  "weight": 21064838222838
	},
	"NodeID-5QV9Tg6tKyGVWcEQM7vmAvac5w3f9wtp9": {
	  "publicKey": null,
	  "weight": 585462590000000
	},
	"NodeID-5QzeGmasNDHSxzxBqiUQ9TBPqkkwzJWGe": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5RuEXmu7SJpf8bwuY17UpMpoEsGZRWnuj": {
	  "publicKey": null,
	  "weight": 65481797799246
	},
	"NodeID-5S729stbM7nFyWHdsoBeFzo5NMUXaTrjR": {
	  "publicKey": null,
	  "weight": 2103714972563
	},
	"NodeID-5SmDUGU8WwZkvxKMnjwKAYvfa9w2qEe1U": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5UW1pHAXDJi1tRFfGSEgujuExMztq6sCc": {
	  "publicKey": null,
	  "weight": 2026063177012
	},
	"NodeID-5Uf5RUZ89pWtYj4Sgc7pxf1mkq2f4EEMW": {
	  "publicKey": null,
	  "weight": 89160830927212
	},
	"NodeID-5WCpR3DKFt9665Wj9jDCTdncLNX7AbzfK": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5YX9uqiPm6hmJEmDy2fUyvFEjvSvSbiWE": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5ZXpg581dpjG8AdoJgTDeXXLnxrQc9Wtd": {
	  "publicKey": null,
	  "weight": 145278051490276
	},
	"NodeID-5ZYfA8hgaSvUmXbsLjmu9aCy66Ha31obb": {
	  "publicKey": null,
	  "weight": 60036996503864
	},
	"NodeID-5aCHiSvLejirNRt8Xgw6SNzd1dDq3XviL": {
	  "publicKey": null,
	  "weight": 2023089637204
	},
	"NodeID-5aQHgP4cTuyentoPz9KK3y2anCBYyG7tG": {
	  "publicKey": null,
	  "weight": 8546031913818
	},
	"NodeID-5bJYpPDsUq3JpGJrFjVRtr1GybXQDP1M1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5cL4dSEdWWKnzZxvg1rqen4M31YeZhFkf": {
	  "publicKey": null,
	  "weight": 1867835538512181
	},
	"NodeID-5eFuL1vXb4NSxW4ZF16WZsJM3gyjS7i2Z": {
	  "publicKey": null,
	  "weight": 2536597351366
	},
	"NodeID-5eVCCzU5VQXhj2iqpqneRxTUNSD8aCJki": {
	  "publicKey": null,
	  "weight": 2961597432400
	},
	"NodeID-5eogZcASu4bDhWxMuUF6h7VFK5SCt3msp": {
	  "publicKey": null,
	  "weight": 663692052852557
	},
	"NodeID-5gcdejFBQ3wMPFpo7qKUKBREJ8w3PRTM1": {
	  "publicKey": null,
	  "weight": 8808729963681
	},
	"NodeID-5gmMRqNob9UjcgroCKH67bQb9PVwtnoeD": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5h3GjwRs4cndBMxxfRpuDMyrX9SF9dPoe": {
	  "publicKey": null,
	  "weight": 5635187941828
	},
	"NodeID-5hNGcpQHUrBDd19vV5QcugACUAmwiWU3D": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5hSVoN9CgKTxGVj2M3pXBDBYtfZkjQF1D": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5idht8398Yg1AqYyBkTQf2kN7zFK1J7Qn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5mmVqScHrzTiSa2WVjo99g2yzL83EBS6W": {
	  "publicKey": null,
	  "weight": 1029692176470570
	},
	"NodeID-5nWWPwUZPzdTJYGXJoWxzWj3Xm63vLL7p": {
	  "publicKey": null,
	  "weight": 3519210946062
	},
	"NodeID-5naWZjifbRQdBoW5bDtAcvuFjKr3k6G2o": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5o4eWuhvE9cmScEoZHr8ngGebxa94rFLo": {
	  "publicKey": null,
	  "weight": 194857297570531
	},
	"NodeID-5ppnn5JSeWMAznssuPMJujyJhzHkXHN8E": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5qr8okF4mnAB2JhAvaJuwqMGXN3gi9QXD": {
	  "publicKey": null,
	  "weight": 19725719960584
	},
	"NodeID-5qxuSbJfb3xqgNQ2pFycD4ybuEH3ATg66": {
	  "publicKey": null,
	  "weight": 135000000000000
	},
	"NodeID-5tUAsBWWmhHrLnjmieom6twf8YREwfosQ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5uYQ6R4WF7kmGfraM9LtsUJG2CDmh78Lf": {
	  "publicKey": null,
	  "weight": 4032036983256
	},
	"NodeID-5wCTAXLJSc5i9RpksPG5fFZXemq4dm2A6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5wQr9SafAyQ6BKMjtTomLB7Bc4tg8iYD4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5wUKLtxUEckPyuzLQSozcPmqVzMRx6v7j": {
	  "publicKey": null,
	  "weight": 10781037934983
	},
	"NodeID-5wf3A9TPUALDHaRhufq3Ry4jfH3cA9SEe": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5xEL4zcSuMjZCEzY3WhQ81Sb1Pw4E5VL8": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-5xhFgBpc99AjBEXqfzNAG3W8mX8vnmXWZ": {
	  "publicKey": null,
	  "weight": 2004515538836
	},
	"NodeID-5yhiTAmJzTBpDLKdXSiZ7Puu5c6VSJtDc": {
	  "publicKey": null,
	  "weight": 112368095202882
	},
	"NodeID-5zENHPe3oP2SdU13oTAvxZA2cEupejYhD": {
	  "publicKey": null,
	  "weight": 431360548265913
	},
	"NodeID-5zP7p1nk2KnAL4SHvgFWcF9Xei3Q13CC3": {
	  "publicKey": null,
	  "weight": 8051089661226
	},
	"NodeID-5zWQQHoMSd29NfKUMziBuccvtt2aahnHH": {
	  "publicKey": null,
	  "weight": 2168561931221858
	},
	"NodeID-61REShk3NbB5RfPc6D8a6MLYsfFPLycNN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-61rJbupdhGqaumerueZFVVJT4Sbp4fciN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-62XYpHtuv92hHhJJXVd4wxJqPFv9rbWA8": {
	  "publicKey": null,
	  "weight": 6694236180024
	},
	"NodeID-62hEBGiabBK34vV52kGT64T6S9QVcdMVN": {
	  "publicKey": null,
	  "weight": 1317394100134047
	},
	"NodeID-62hKwt6g9MMngpDLHapNFJmxXRcDoe2sc": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-643Sjdbh4n7krQXPSNbSK491xfVbh3hqy": {
	  "publicKey": null,
	  "weight": 9496901926109
	},
	"NodeID-64Y4Dq4dwfnhTjNnZEFjCexmc2Wwi1fyM": {
	  "publicKey": null,
	  "weight": 17427167380731
	},
	"NodeID-64Zz8dh68ypXYWaUcfkXDm9UuG9VNXrk8": {
	  "publicKey": null,
	  "weight": 60938037476083
	},
	"NodeID-65zy7gGi1v2US2tkBFfjpwuDoWVPjPsZW": {
	  "publicKey": null,
	  "weight": 2009130300700
	},
	"NodeID-68ZngAKzsRbzQMq4CpuZ41J7AgCDEijPp": {
	  "publicKey": null,
	  "weight": 463366887624628
	},
	"NodeID-69vkfdsq14i5HZSd6m2kUwRAu59ePbeWw": {
	  "publicKey": null,
	  "weight": 1945261636570076
	},
	"NodeID-6AKFYXuxWErAD6nqqFkPnbwHzXB1JC4VW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6AXAfsfvp25QX5Tff6Ag6ERLXQ577wurM": {
	  "publicKey": null,
	  "weight": 660061347424838
	},
	"NodeID-6Aed6k5Zj6vffgNdvr29Z9VcdDxk5PoMq": {
	  "publicKey": null,
	  "weight": 350475620387154
	},
	"NodeID-6CGouDLcYyiuBY9xEi4rEaHCoBzH6UQrS": {
	  "publicKey": null,
	  "weight": 8000000000000
	},
	"NodeID-6DqFjxDUK7nh68zrFDE7iHfVsXvoM5yzt": {
	  "publicKey": null,
	  "weight": 8049116099812
	},
	"NodeID-6ETncdUXndB43iT4LijRYLHF711gWeAgJ": {
	  "publicKey": null,
	  "weight": 4328651468952
	},
	"NodeID-6Ef219qf1xYjFL5SCFQo29ZYKDrcNi6xT": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6F2LZ5hykkwmnRYYabqHH19MnzknJ7rij": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6FSSZuuLEZ4sHJZGfpY4Wuz8M3qnK6ec5": {
	  "publicKey": null,
	  "weight": 2002000000000
	},
	"NodeID-6FWuMoDyCe1gsmeQj2RsGBmYv1hoGUbzs": {
	  "publicKey": null,
	  "weight": 83493620132974
	},
	"NodeID-6GEno5sempCQdCZvTEuPZpDLqVvpN2JEB": {
	  "publicKey": null,
	  "weight": 2412930732640
	},
	"NodeID-6J3LY7ojkK7WCZzmArvxEozaESDMb42hX": {
	  "publicKey": null,
	  "weight": 9783397103676
	},
	"NodeID-6K9dR5Dx8YkPXzYeNvDRBionRmEo6HGE3": {
	  "publicKey": null,
	  "weight": 421304304283839
	},
	"NodeID-6KsZqkvobK4vUHJzo2VdwxWMsbHtXCZPb": {
	  "publicKey": null,
	  "weight": 2819755809320
	},
	"NodeID-6LHDq7hq3PS1xBj7cFQuvxozaM2hJXheF": {
	  "publicKey": null,
	  "weight": 2081091964951
	},
	"NodeID-6LRbs1K2iPGmhn1EBbX8u2udyPbUt5S5x": {
	  "publicKey": null,
	  "weight": 27519603786577
	},
	"NodeID-6PS1ECxRbQ5x3wj31utVWecd9Xkgcp8sb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6PjT9Qhfxp8iDamQMDjoMgYYVgoBwCdg8": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6QD1KNQkx6wj142Tv6demd3FYbghQvCM2": {
	  "publicKey": null,
	  "weight": 11066474351948
	},
	"NodeID-6S8zAjmFJ4JdZwPURyFGp8Q8tqQ5NcEbX": {
	  "publicKey": null,
	  "weight": 154928418592420
	},
	"NodeID-6SwnPJLH8cWfrJ162JjZekbmzaFpjPcf": {
	  "publicKey": null,
	  "weight": 66000000000000
	},
	"NodeID-6TG73ofz7EU4keEwv6jt2xLrpaFzqXnYR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6TzpKmXTYp2f1ok4rEzEAgP5b87FSQ9G": {
	  "publicKey": null,
	  "weight": 5377642284827
	},
	"NodeID-6UMb71Emubx62ZEsNBPhhuEzZdqgGpKZY": {
	  "publicKey": null,
	  "weight": 4099131190595
	},
	"NodeID-6UpNHJdkRQM9TRX1m1wb1sfNX32Ze7ZhA": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6WwTZqp76ms4iEUepaagpeEsf4N4saDGb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6XwiaBfuKwKG761Bc28jDAsiFJ2gDuVgQ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6YNW6QK3JDpCaEnVNsG5wsNA1SvxT1dhP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6aZ3KyGBPTcn6NF9K2jpvxTjhhDjks3Ev": {
	  "publicKey": null,
	  "weight": 88877492316564
	},
	"NodeID-6anRg13mVkUeZmpjf2Mm41sug56BH7Jof": {
	  "publicKey": null,
	  "weight": 8658786642678
	},
	"NodeID-6biB22M9yY6jfRUeKHvLvz7dyVVZ2XYNy": {
	  "publicKey": null,
	  "weight": 4191998421457
	},
	"NodeID-6cYVDNPxCAcSCBzYPaqxF4TiGHF8XKoLV": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6cem8uPoUwQ6o45VYS3D5mmFFzTuNC2M7": {
	  "publicKey": null,
	  "weight": 191638005508655
	},
	"NodeID-6eLQBH9yJUPLiZGXToPxQPj2gTWv15Lf7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6gAC5LaDoT9EWJezcQ7ijFjcsYa61Atdb": {
	  "publicKey": null,
	  "weight": 33000333182069
	},
	"NodeID-6gFsNyzYU474fgL7t7x43Efjbak7Vx7Wp": {
	  "publicKey": null,
	  "weight": 5710955438081
	},
	"NodeID-6gRx5vFhuMTDtaRYP82ZSGuBq7eR5s3jw": {
	  "publicKey": null,
	  "weight": 6173868716188
	},
	"NodeID-6ghBh6yof5ouMCya2n9fHzhpWouiZFVVj": {
	  "publicKey": null,
	  "weight": 912282705882294
	},
	"NodeID-6iLoGX7rEPJW2GZMekiNPEDNZKmsbn3iU": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6izbDmxXdSqHwnBdWkGWuJtxVuTh7AZyA": {
	  "publicKey": null,
	  "weight": 5532330961745
	},
	"NodeID-6kZ9aC6TPxRtXaWgDR7UHz1aWyfkU6o6h": {
	  "publicKey": null,
	  "weight": 2325920000000
	},
	"NodeID-6mRmWBQmMyFzSD75SJzjx4VmToenpMt5u": {
	  "publicKey": null,
	  "weight": 6836685207470
	},
	"NodeID-6mSXCB3r7oeP8Suy1AWroDY2KEF9hT9Mi": {
	  "publicKey": null,
	  "weight": 336755402630614
	},
	"NodeID-6na5rkzi37wtt5piHV62y11XYfN2kTsTH": {
	  "publicKey": null,
	  "weight": 9825261145658
	},
	"NodeID-6o65ccWJmdcMSDkDphL2i9ajLCZRsvAfj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6oBNaAUr67MFY9Gtu4eis7bsGY1RswvAR": {
	  "publicKey": null,
	  "weight": 7419000000000
	},
	"NodeID-6p7pSJ2hoYnpRAXUt14cr1tsA5kEm9vXE": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6pShVkG6mZsinZNWZr48xvQSuSnmnyh5o": {
	  "publicKey": null,
	  "weight": 4251617787574
	},
	"NodeID-6pXaVajr1G1nj8Z4rEXisGUajXGcFPh14": {
	  "publicKey": null,
	  "weight": 40000000000000
	},
	"NodeID-6pzGb1VVyQT4RLChXUQYfhvjLrqEmWiNi": {
	  "publicKey": null,
	  "weight": 2160394332638
	},
	"NodeID-6qPNz3b5VZpMWkVNjmLmpBrL1sYxr38bs": {
	  "publicKey": null,
	  "weight": 2783456452539
	},
	"NodeID-6rJMJqgEbTGJCv3hBzDaY3axQ3Fq98yFX": {
	  "publicKey": null,
	  "weight": 495000000000000
	},
	"NodeID-6rjd7h5dPJwVEytA9zm86ZBfqEfBuSjKi": {
	  "publicKey": null,
	  "weight": 66669135101382
	},
	"NodeID-6rtV2pPKXyf2Ek7nWnmuzdiXvk2Ma3ynn": {
	  "publicKey": null,
	  "weight": 5991000000000
	},
	"NodeID-6sLdnpLBDDMG2eEzgdiGjTkne7a1iETEu": {
	  "publicKey": null,
	  "weight": 6565300000955
	},
	"NodeID-6sywKjSw7tS2qgFAmLSyJR5D6oB2BHT3Z": {
	  "publicKey": null,
	  "weight": 2183127465430
	},
	"NodeID-6tkBfGVTgfpbaiv5HFX8LpfKdu386tfwr": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6upkG4FRNCoUQ8nz2by26m5vdVSthXCtK": {
	  "publicKey": null,
	  "weight": 9878798867830
	},
	"NodeID-6w8o4bbNdVLscimLv4X6BA7QuSqX37TFL": {
	  "publicKey": null,
	  "weight": 2278830384805009
	},
	"NodeID-6xPTxaGM5hVHVKKFuC2MjpCL7ZD6xT3wN": {
	  "publicKey": null,
	  "weight": 11521613598291
	},
	"NodeID-6xQ5oYAQ348ntvsshCMMSd7uduReey2wC": {
	  "publicKey": null,
	  "weight": 2025000000000
	},
	"NodeID-6xjnXHhLrLDafewh5t69uRqtHea75RSGv": {
	  "publicKey": null,
	  "weight": 2994015674813
	},
	"NodeID-6yS1J9HFdHgdX2ruW6XMouQvECDF2Fta4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6yf4tK8VmAKhQHC7tdQwy5FTw39i9mXVV": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6yi5sG94i8wSHhqFuuUip6RuZGkKDHzxy": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6yuuujhDsENasQ1raXaTo7mgLbqF8pc6H": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-6zcpV968BV1e3BrZoPmcKuoUNRcP6vjxt": {
	  "publicKey": null,
	  "weight": 2999100000000
	},
	"NodeID-6zytvHVN3xs7VRFG8RBhjWVmcvaYxWHDo": {
	  "publicKey": null,
	  "weight": 2206232805023
	},
	"NodeID-71ZzUv6td3aKrZG4LhGDDAnjsEcEchwDe": {
	  "publicKey": null,
	  "weight": 2113743915176
	},
	"NodeID-71taHoSLZsm1Aqe4spXKJvNXw3o7982Hz": {
	  "publicKey": null,
	  "weight": 21498909483443
	},
	"NodeID-72D6UYvQjfgKuSrSFtMfyCtHq1jZZ2gN6": {
	  "publicKey": null,
	  "weight": 3360863125214
	},
	"NodeID-72VBCaJrN7UqM1tQLjfhTuXwoGUxPCUvz": {
	  "publicKey": null,
	  "weight": 51767485661344
	},
	"NodeID-72ncx4NEgf4cziN5GPDRhn7i71rgwH1wV": {
	  "publicKey": null,
	  "weight": 16780000000000
	},
	"NodeID-74oLbJEptxh1cVzvh3QTwaykED3FXjPYK": {
	  "publicKey": null,
	  "weight": 2000999000000
	},
	"NodeID-75PnenZJE9PQbQycTys2zonb1atGLsbX6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-76GDFDnZW2ihXCsum843HEDmrvcVuMDZ5": {
	  "publicKey": null,
	  "weight": 21367005966775
	},
	"NodeID-76KcZZBhWfiWuFqrx21KyeBHdXxsaVyX2": {
	  "publicKey": null,
	  "weight": 7825000000000
	},
	"NodeID-76nqFaqDGakCDcm7kDsgP7GJn2VSM33xh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-77kWbGy45j1Scbvgh5z3zhoUBL1S3qjAi": {
	  "publicKey": null,
	  "weight": 311829994322276
	},
	"NodeID-7843EeyboY1mZSmdzjqdJodxHycNn7Kv2": {
	  "publicKey": null,
	  "weight": 3391710471885
	},
	"NodeID-793wJG3RHeLFVCXkNTQxsM1pSUHNyDyMv": {
	  "publicKey": null,
	  "weight": 595985533042985
	},
	"NodeID-799wg9bHu8gGe83SHdioqQ6hUg8QCYRv3": {
	  "publicKey": null,
	  "weight": 3727031613801
	},
	"NodeID-79NUUnbkpDfvGzGtq3Q9FArGK9D1Wzpgq": {
	  "publicKey": null,
	  "weight": 2026000000000
	},
	"NodeID-7AVurguoWJggFMe8WBEBuTzZjmN8AmN8y": {
	  "publicKey": null,
	  "weight": 149998416352840
	},
	"NodeID-7BvjRu2P26PSUXWzCSh8LLxLNdSpHEq8f": {
	  "publicKey": null,
	  "weight": 2916616189409934
	},
	"NodeID-7CK2DghfK6VgUpPr8vz3EHxBYhpfYGEpB": {
	  "publicKey": null,
	  "weight": 144501692863872
	},
	"NodeID-7CQwXe5Vnm7hzSQeWDRXqUsGDc5vcPBfZ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7EEYbCpedmA4bgHnnvhx4zXgMUpKe5Gkd": {
	  "publicKey": null,
	  "weight": 9287030421598
	},
	"NodeID-7GQQx6v1hhM4au2kLKrtbvau4N4xqte36": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7HgUf4o1UiEGh2Z2tWp2erbkn5JaXYYZG": {
	  "publicKey": null,
	  "weight": 8000000000000
	},
	"NodeID-7HsxnZQnGRWYG5jFVWDX2L26zEAYaUwAP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7JijGwmqUuditkaCwdcCAr4Xc6AAMy89A": {
	  "publicKey": null,
	  "weight": 40587253419370
	},
	"NodeID-7KKSDmFVJ4YXsdcnxjDnV4girXR32yRni": {
	  "publicKey": null,
	  "weight": 2044500000000
	},
	"NodeID-7LvgVzHX6jsppFQggdZmcgcvbp3pRm3ue": {
	  "publicKey": null,
	  "weight": 204202974726872
	},
	"NodeID-7P5SGZLFp95WbF2yv3A9WbQaQDzqfsAFF": {
	  "publicKey": null,
	  "weight": 2009120277338
	},
	"NodeID-7PJZJFNdWE59m5jrLoMNLYN4DdtjyZ3K6": {
	  "publicKey": null,
	  "weight": 2029521296996
	},
	"NodeID-7Pjqo8qr8Gd78Wmzgo5xqF85qcafKQ7W6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7Pt1CaFRMYN1azEC7SFoVjeXLmRF3s7EC": {
	  "publicKey": null,
	  "weight": 2126670000231
	},
	"NodeID-7RBkUCXxXbZ8m1dAH76jFxABESmVnzT9H": {
	  "publicKey": null,
	  "weight": 2013753896042
	},
	"NodeID-7Tv2ubT5xnAXWNRf8X61Eyw7VcU9J1w5b": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7ULxX37k6SSaTX3oF5RdrxL9r5aGZu75M": {
	  "publicKey": null,
	  "weight": 13309725443707
	},
	"NodeID-7UaxuvJw9C9FytVkUjRCQ3csmDiNzujGW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7XrrK8iL25SJfWrrkyz3gZSo2t5iEadvL": {
	  "publicKey": null,
	  "weight": 7318356765211
	},
	"NodeID-7a1GnsoFxviSiL8gJyFqrkazv29xY48mr": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7a5FbgWRUJeRWPzyidAyL3ENC8wQDJ1eK": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7aBosQAfHykUE3rmubktp5zZtZmavTzow": {
	  "publicKey": null,
	  "weight": 2396382757996
	},
	"NodeID-7bRrZrBdCveg1eXqS6NhKtqxXAQqteq34": {
	  "publicKey": null,
	  "weight": 2084519763692
	},
	"NodeID-7c866Md3Geveoz9G9XvP1TroT4LKdz5hL": {
	  "publicKey": null,
	  "weight": 2004521560442
	},
	"NodeID-7cwvfriG4LFWV4iF89E7GhYfLB14hFjMa": {
	  "publicKey": null,
	  "weight": 2036501440142
	},
	"NodeID-7cyp41vXvj62jBRh7y2j6VjLXhoJ6Hoab": {
	  "publicKey": null,
	  "weight": 20050000000000
	},
	"NodeID-7fJWXpi4Bj2XgeVXG7Tj3JkNNrGnnJhdp": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7g7qvfKvRXmPZAxhBMfj8XCGuWVVHexHe": {
	  "publicKey": null,
	  "weight": 106931922832667
	},
	"NodeID-7gLtM9D45daj6WJfqrT6uUje4muznwKwQ": {
	  "publicKey": null,
	  "weight": 2125087484629
	},
	"NodeID-7gqKK3aZ6m81v2138yfc4fsAdzFHsB3xb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7gu4gthY5YhpmpUpQvuAy5CDDvDs2yEVt": {
	  "publicKey": null,
	  "weight": 3303140020753
	},
	"NodeID-7hQFksPPW1Y5hZukNcQn6bTQjuTZPnj42": {
	  "publicKey": null,
	  "weight": 388397999999958
	},
	"NodeID-7haZar18iKUdnps7YZEYbJCSwnFw4KMY1": {
	  "publicKey": null,
	  "weight": 4233022159981
	},
	"NodeID-7hnK4EoWzm7V5qFPqJvpcRQti59au2BD8": {
	  "publicKey": null,
	  "weight": 9483261794100
	},
	"NodeID-7iFGXSHq1R8MvEM88EkS64rW5z7MwukYk": {
	  "publicKey": null,
	  "weight": 4915749284502
	},
	"NodeID-7iHpsauHYyQoCjMSZsXJkqXodq7J5eXTR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7idSkp4X9MwUacdNn32NkE9SXo1XWveLj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7jqH9PXuyYhoTaaQDAF1eVzLtCVSgmSpz": {
	  "publicKey": null,
	  "weight": 77781045352231
	},
	"NodeID-7oHSQLPBV4EzZurguKU4a2FUvPxShuVzf": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7oKBFksJjAcBCNEY3ZcRLQeXLX7iA9dNE": {
	  "publicKey": null,
	  "weight": 21025140178354
	},
	"NodeID-7oyvSG1pLLp3zf2ouSAT6rUeehDuQcNtQ": {
	  "publicKey": null,
	  "weight": 115948151087510
	},
	"NodeID-7p8akVY9Z56XkLykaohLpNGp9sU9PLTvh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7pDhCWQKNrpCjTL6tB7uNvBUnBQ5amgob": {
	  "publicKey": null,
	  "weight": 2779890909494963
	},
	"NodeID-7qrESfShxxqXfa35MfBC8D62t6enjm2RC": {
	  "publicKey": null,
	  "weight": 2181344216085
	},
	"NodeID-7rE4BjdFpwUk2igXEeauKaZT8mCjkt9Hs": {
	  "publicKey": null,
	  "weight": 51918864536020
	},
	"NodeID-7rXi72Jm76kAKUi6BiyJChnAm9xuKrud9": {
	  "publicKey": null,
	  "weight": 51164278522613
	},
	"NodeID-7sTpQZFCZVzhLTugv8PKhNE8tdQ394LRo": {
	  "publicKey": null,
	  "weight": 27169016745794
	},
	"NodeID-7seuWZpGnhEox3dL8T363sghDtCT5UNp4": {
	  "publicKey": null,
	  "weight": 66000000000000
	},
	"NodeID-7sitbRzmAsrFjKRitC7bub1soEaHXsgbA": {
	  "publicKey": null,
	  "weight": 2131755793650
	},
	"NodeID-7stntWreGt1wPeFuY9Z8bNURuHxW4JN3D": {
	  "publicKey": null,
	  "weight": 2253252862129
	},
	"NodeID-7t5yL3rJ5JgME7vxa8em5sUNTPeYfmMop": {
	  "publicKey": null,
	  "weight": 406714588235286
	},
	"NodeID-7t8gvYRY2ZpMzT5uFgmgXsqyHVks6Vh1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7tZXydqQqbjAQaNDK9MNEZFMsSFCp4xJW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7ts1Uhry9k5qkmpYZNfdivuzj2eBPemKa": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7v2CCw3NVjU9UxBhGzN1jhSmatHKE2mXZ": {
	  "publicKey": null,
	  "weight": 207251647058806
	},
	"NodeID-7vWb35jCnXGxAhLStVtKKZhWiw1ByhBdT": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7w5D1gxJvKekgHwtiYzF1Pf91ecxWEsGT": {
	  "publicKey": null,
	  "weight": 2004548549868
	},
	"NodeID-7wMUdN6T9awYChEoRgzQ4eg8qZj1BFrne": {
	  "publicKey": null,
	  "weight": 90000000000000
	},
	"NodeID-7wSVhYSmrwcts5o98Yg863PoHhJWWQcbt": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7wyu3EXASVzsJpRPoqhssRWdmTWa6Fycd": {
	  "publicKey": null,
	  "weight": 441404399322240
	},
	"NodeID-7x5HbUJqp763SUw8Sz27yU3bKWCdQiqAk": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7xCJtVM3wTprzgDBhXeXCXEUy4NEuqUvi": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7xcqn3xNHBxreF7mzRx7qEDbpJguaYg3i": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7yBDXjCcqHCnfqC8LfQ1oL9Yyx1se2Rj9": {
	  "publicKey": null,
	  "weight": 2500000000000
	},
	"NodeID-7yUhoXtjqGABwMeSgQG42cWX56r1DkBwh": {
	  "publicKey": null,
	  "weight": 17500000000000
	},
	"NodeID-7ynNyRtZx8re55eU1CwJ1Ga5cEaRMPJ5s": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7zA1m9YUevKKyKdyJK3snpSiAnXhUG12R": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-7zM8yAFcY3D4nA5Wz8voneD3JjBiU2SgL": {
	  "publicKey": null,
	  "weight": 2055414147091
	},
	"NodeID-7zQLzhnMhaj5SbRXVpUUsMP4c4r8yQAD1": {
	  "publicKey": null,
	  "weight": 45781490918473
	},
	"NodeID-7zeJ5VBLvTF7LvTn1VRZCJfdt5EsWUg3e": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-81U2wpq95De12KchbFhatf2fBt6KTFmAr": {
	  "publicKey": null,
	  "weight": 2013703614562
	},
	"NodeID-82qNDEtRpFygtdH5AL4Fu5ZV12k6WCLqL": {
	  "publicKey": null,
	  "weight": 3473243423878
	},
	"NodeID-83uKDYgau5PctbifticdDZUaTcDkakhFi": {
	  "publicKey": null,
	  "weight": 3776865775776
	},
	"NodeID-84RAL3MngQcrLpLgTgJwtLobhryyt6TiM": {
	  "publicKey": null,
	  "weight": 20600000000000
	},
	"NodeID-84XTAh1VM9QmfdewNnqFVhz2kbjrFvK4Y": {
	  "publicKey": null,
	  "weight": 2987730190013880
	},
	"NodeID-85QtYSfGWxkvQjbiFJzv8R6ajAutwFS4F": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-86HV89ZFehy2HHmnkpgBuUn6vRxg25if4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-86LjyKDK2eKVUpdesMPWJtvo9gadfW9m3": {
	  "publicKey": null,
	  "weight": 7973506295144
	},
	"NodeID-88fsGoe69sPGtRjqe21SK72KrJNezzz5u": {
	  "publicKey": null,
	  "weight": 22841251312250
	},
	"NodeID-88ik4ovrpTx4dTExqhXhSQ4JwAQnHbuGA": {
	  "publicKey": null,
	  "weight": 2049095000000
	},
	"NodeID-89ucgJfiRvHtGXynRbta7VcaX6MEcJ7K6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8ALyunjeYV85zBSrD2rjJdQ754z9jasvq": {
	  "publicKey": null,
	  "weight": 6826436398297
	},
	"NodeID-8ANrVm333EMbpyAdjYhUUACUTaU3tx38X": {
	  "publicKey": null,
	  "weight": 3338558296639
	},
	"NodeID-8DLefLvLwgb6fkcghkLgj8w91MSrnPdjh": {
	  "publicKey": null,
	  "weight": 33636355058416
	},
	"NodeID-8DVxsXms5opComNTHqFMgR9pRS99gt27r": {
	  "publicKey": null,
	  "weight": 3205833024822
	},
	"NodeID-8E3UaD6hjNK8aqK2BBUmqPXyrfB19WJJ2": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8Fiuf4Uucg4x3ijMfkyVvYXdfgdBFvsFk": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8Fn8YCu76VUHny7MHQTkm3mPSuGPkVeqn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8Gct3pDyuCEY5xH1F9RdYjvk9ry23cTtX": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8HHLzk7pF96soZ45bJ1tfHNcSp7moXiTB": {
	  "publicKey": null,
	  "weight": 2051000000000
	},
	"NodeID-8JYJVijgzBsSTK5EUsXitrMEYNGkqeba5": {
	  "publicKey": null,
	  "weight": 2239917238642
	},
	"NodeID-8Jg1Hs7CeqZpGHdZ3j5Vgny42bDQikJec": {
	  "publicKey": null,
	  "weight": 2740693624482
	},
	"NodeID-8Mk5Kpvp3oTpLbRSnwmCJmPbvRwsR48ra": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8Nr7TDSiK215oiNDY17SaDgF5nfty4KUi": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8PmuHd5fRb8VWpTTyNhytQuv3sy83P2jT": {
	  "publicKey": null,
	  "weight": 2421541228074
	},
	"NodeID-8Q1uX2BfyLbmCCor1t2h9qvn414WXfQ92": {
	  "publicKey": null,
	  "weight": 2026000000000
	},
	"NodeID-8QKheMcZ1jn5n9RVi5MFBkzGrKjFJNNAX": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8SJJiFHcGfH8KErq78txDKhNGARL1K9mp": {
	  "publicKey": null,
	  "weight": 2441230659967
	},
	"NodeID-8ST5juiAQzVfXhNAgER1UapEpNowZyMoj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8TArWpFgH3sazEH8qP4gUjtGtFMvjw1aR": {
	  "publicKey": null,
	  "weight": 251440204745412
	},
	"NodeID-8TXq3f5WiXqZuyQWF3Sa7ncoW3YRRaevh": {
	  "publicKey": null,
	  "weight": 10000000000000
	},
	"NodeID-8TubANvEGU3Zkz6nKeRx3ShmHyULLjD89": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8TviTXQLiqxcxkCikPTut9DwifYjFaMcT": {
	  "publicKey": null,
	  "weight": 83000000000000
	},
	"NodeID-8UynAEU8PuuGetmjiiPcbGmvD3EyycuN6": {
	  "publicKey": null,
	  "weight": 2923523643564380
	},
	"NodeID-8WowWhvPM6hwXBLGkuX7HvwApmXNDCyrH": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8WvUBZyrk6XyF7pvcg3NJZk8jtJGi1W9G": {
	  "publicKey": null,
	  "weight": 11221817653884
	},
	"NodeID-8Xt57NiwBsGKt3CjYZ1ymKx5h3nHU3vcX": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8Y4i7kpWyv3pvvwJSk4EsGx7frk5mU1Ah": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8YCf3nhXEnz8vthaxYRWR87ez2orwP7yW": {
	  "publicKey": null,
	  "weight": 2004521069336
	},
	"NodeID-8YNtmpa8fe2dt12PqDzhDCBtwcQDTXa9p": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8YqEZ2ufMNUFPXyNpUAvLJWRXLKqC6D2H": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8YrAcDNtKarb1DeZR26bHfTcGdKyZYDHr": {
	  "publicKey": null,
	  "weight": 2013684525988
	},
	"NodeID-8Zvnq1815CaYYegXmuyBD5BVtJnurtpC2": {
	  "publicKey": null,
	  "weight": 2044346782107
	},
	"NodeID-8a2ghUUrY7geFumraEVBbLPCDTgnwML8Q": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8bJ1oBBjCK8TaCKv4GqifdHDAFzHmFV1t": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8byZLbVuVJbedEvsAjZxTTSqorsjz7ThK": {
	  "publicKey": null,
	  "weight": 4364122548709
	},
	"NodeID-8cKy5FzzP2CHV1rcwdKq62vuYP9yUzYuG": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8dCZ9beUGmAWyCZPwdM8YZ5FDBRPXUxFf": {
	  "publicKey": null,
	  "weight": 9964246454793
	},
	"NodeID-8f27M6Ju9Dnh7YU5pHqpkBLhCmBq9EBBD": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8fLXgVYgby12VunVnSMrRrvcgEijVknbo": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8fq1H8oStMi9BZrEk8qiE8QAV7AtqGnXt": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8gLyc1WmUgnTJh6TAjg1nbqsqdF6xGSeL": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8gndDR9ULm3ywRNbh3FUveaSyo5g7V3ns": {
	  "publicKey": null,
	  "weight": 2077000000000
	},
	"NodeID-8ja1C24FM5FyQgUGjsm4qwpD6shiu71KH": {
	  "publicKey": null,
	  "weight": 7099224362177
	},
	"NodeID-8qkeNyB3cbs7LqzP4AB7rHcrcHLAYVDEs": {
	  "publicKey": null,
	  "weight": 2025010000000
	},
	"NodeID-8rEmSGNj5VW2HUjGYMyynpWTYJuDrMpFE": {
	  "publicKey": null,
	  "weight": 99999757266200
	},
	"NodeID-8sczBhZjP4BGoCqbn7SaRCnerh6cJ6jE8": {
	  "publicKey": null,
	  "weight": 2013716758786
	},
	"NodeID-8uDtidG6poHRfZfZJXXSVBwivoMogXQ66": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-8ugXJLrq8DW4m7m7vBTpkh9rCLn3sUmmc": {
	  "publicKey": null,
	  "weight": 2013698746504
	},
	"NodeID-8vuV1UF3KHF98EdqFyuFAMDJUzwkQWEA7": {
	  "publicKey": null,
	  "weight": 10475000000000
	},
	"NodeID-8z8rezdB9vXDxBurBeCYnNBfNcD1kVrJP": {
	  "publicKey": null,
	  "weight": 2004521070660
	},
	"NodeID-91W73TnEik6kseSpBHSAZzsNVRGk8Qfkm": {
	  "publicKey": null,
	  "weight": 50000000000000
	},
	"NodeID-92aEQicro9Pht1HRiexK9ado2BToyi7fY": {
	  "publicKey": null,
	  "weight": 107783668468301
	},
	"NodeID-93LVjgmjMcBmcJSByxYiTuyPNbdVw9q3n": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-93rEN3kFDyrRYoe8aJ2HB2dES8AyTY9oV": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-94CDXgCTFRPKPDmK55UcfG8KqVQfy4QsV": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-955GU1MqWL8yXAtoc8AsE7FNx4nGC9JyL": {
	  "publicKey": null,
	  "weight": 10145124379044
	},
	"NodeID-95rcDYyjGNKckKCF8PKuTyEV1wxkQvP3n": {
	  "publicKey": null,
	  "weight": 125000000000000
	},
	"NodeID-96tBuFHLkXJwyYug768fE2GDEFmebUUTy": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-96tjSzF58iKJUG1hDQve6HJdzjwLNMzFW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-972yzi4wAgqU7RGmcQGED4ueQeBXzccka": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-98JuUPhpQBgiY9ozqFTh8rEjx3jL9D3aJ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-99Dk6vNf7RaX4S9e4oAAbuhubzzRnrGxC": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-99hQvpQKJi1y7gTjD2icHKYwk1pNMHFhg": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-99sqQME296SfcnyYWTmGrLXQJxaGUKAA7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9AjYPCvDWKnwzFgRiy7jeTEKmVxK42LSz": {
	  "publicKey": null,
	  "weight": 3072451875500
	},
	"NodeID-9BCUhY6Q2Eji4J1u6MV9XKnPspciWWdLR": {
	  "publicKey": null,
	  "weight": 3310144111060
	},
	"NodeID-9Bb6E7B6GMd1MnAuxYVVYF9CSf1XiBCEY": {
	  "publicKey": null,
	  "weight": 4082539774528
	},
	"NodeID-9CkG9MBNavnw7EVSRsuFr7ws9gascDQy3": {
	  "publicKey": null,
	  "weight": 930199075602602
	},
	"NodeID-9CnrQBBFSkE2Xzfcz3Tk1e8iauq8iNR88": {
	  "publicKey": null,
	  "weight": 46562009833640
	},
	"NodeID-9D33RPjZwKx3gV6MVs77z4uPJcKgqy4ns": {
	  "publicKey": null,
	  "weight": 20807767719089
	},
	"NodeID-9D63iEsgkSKzSz1aBLpELbh47Y7aT5ujj": {
	  "publicKey": null,
	  "weight": 3334483222673
	},
	"NodeID-9DCqby2EWyUKFsKs1VNH5mX7V9FDkqUnL": {
	  "publicKey": null,
	  "weight": 2731491960392454
	},
	"NodeID-9DmoV3n9Sb1A2aLkYuv9wZV7uPJo5JhpG": {
	  "publicKey": null,
	  "weight": 2072070105632
	},
	"NodeID-9Ef44CXDWWfP63UaLr8XSCsmEWUFY5i2z": {
	  "publicKey": null,
	  "weight": 7111719921352
	},
	"NodeID-9FZQBRnW2jMU1PW99pqVTnxkPHF2Srfeb": {
	  "publicKey": null,
	  "weight": 2078900000000
	},
	"NodeID-9G5DhsoBgHjDA3j1FqFXKyEJH8AGPeB67": {
	  "publicKey": null,
	  "weight": 13429372680194
	},
	"NodeID-9HT8KtMkcxP5aWjmxaVVGBCMZ1r6jGmgz": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9Hp3rH2xzTzDoZj6JgYqhbmTFGVeS1BnL": {
	  "publicKey": null,
	  "weight": 4576000000000
	},
	"NodeID-9KjtB2xC9aehpcYRuqJR4myLGJaq1tVse": {
	  "publicKey": null,
	  "weight": 2526852158783
	},
	"NodeID-9Lyd45k1mXNkksF2C3nwgCJRX9YY42FT6": {
	  "publicKey": null,
	  "weight": 1018399700938603
	},
	"NodeID-9MCd77R8qxsNg1ajgHqFMsRYCaQ4KpGot": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9NEeo3Ayq5qGodBbMZ1gdCQ4bmGH59BBo": {
	  "publicKey": null,
	  "weight": 5319000000000
	},
	"NodeID-9P6CxSTDfAi7cRr4LUW4bbdA4F1nK89T1": {
	  "publicKey": null,
	  "weight": 2092449022684
	},
	"NodeID-9PzVetjSSSY4Xtk3VnNcGzQdEjZSSkR4g": {
	  "publicKey": null,
	  "weight": 58127586833133
	},
	"NodeID-9S8duheEvrp7gKW4qLLrWVCnC2wvo39w8": {
	  "publicKey": null,
	  "weight": 451801329877574
	},
	"NodeID-9SQzWXUh9SF4b4qDmYy7p6LcqWSgxtu6o": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9T7NXBFpp8LWCyc58YdKNoowDipdVKAWz": {
	  "publicKey": null,
	  "weight": 288400212230332
	},
	"NodeID-9TMnzP54UVfmYmYcxAzVuNC9rpVr8zvgB": {
	  "publicKey": null,
	  "weight": 2120615284366
	},
	"NodeID-9UWtBGvW7G2RVSwcR15MS8gesUvLk6f24": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9UgnxmEX92MdaAFQc7FG9SAfPECeT6Qp3": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9Uis3D51cjQauBaBPnqsvKjKg8Byj6Xwb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9W3QhxVhg9EncK9438UUepFBrCxdtFp2w": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9W5RaY7LA6so9hGCPGfLezBXMnWNRQb8V": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9YiwE69B2wcvXtEW986YcgQxGg9ctzmKZ": {
	  "publicKey": null,
	  "weight": 2004521561842
	},
	"NodeID-9YqDsggitZTgCi2WnnrKBz3jAcD1S7oin": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9Z7jK8jN5FsTXSqkG1fV2QNG7GktPU8az": {
	  "publicKey": null,
	  "weight": 2004515539220
	},
	"NodeID-9awX4ceeXavKKmccnAuDU7uvhqGfaEDSS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9azF97hpcwpzj3t7S9mD4VKY1S5eJDrcW": {
	  "publicKey": null,
	  "weight": 3058864748157
	},
	"NodeID-9bU9jwHLH6KxcTu8pBbqJQqkHYR4woY7L": {
	  "publicKey": null,
	  "weight": 238193700977377
	},
	"NodeID-9bmxQyFJCrENbZgmMdy6xXorVN1RL5qWY": {
	  "publicKey": null,
	  "weight": 6036759940558
	},
	"NodeID-9cazV2uY7efo1f9bjbixH87FFNG5wGyts": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9dN9NU5XCr3WQGYnFjGudKRScQ1HKGmBn": {
	  "publicKey": null,
	  "weight": 4028782770081
	},
	"NodeID-9eH3Zu7LN7kNwvUy3Am8mkgvvekrkxD6d": {
	  "publicKey": null,
	  "weight": 631798027160179
	},
	"NodeID-9g6keiiMr9vXhaSNP3oqdxhu42LEZ3pWu": {
	  "publicKey": null,
	  "weight": 4335294733735
	},
	"NodeID-9hW5Lb3t5JMgP945DB8wHUJP73uLuDLAZ": {
	  "publicKey": null,
	  "weight": 3247680902266
	},
	"NodeID-9hj2xhiHktPCpcD9koLLH17a1GKoWY7QQ": {
	  "publicKey": null,
	  "weight": 2013699959058
	},
	"NodeID-9hksERdA8ypkBykmHuch4mw24RR2sg5js": {
	  "publicKey": null,
	  "weight": 6599516921354
	},
	"NodeID-9iGXJ935QZWkp5j2v6KNSTr8b45a7Xbfs": {
	  "publicKey": null,
	  "weight": 629819292708414
	},
	"NodeID-9iiCATASsRcosMQ811KpxSMHsGdM38rPJ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9ipN26YdR2YZFyz4hdoyMSXFYNFMDgDfy": {
	  "publicKey": null,
	  "weight": 3820987590390
	},
	"NodeID-9iuMMyL7LZyUxMK9s735hcRng58bSj3gw": {
	  "publicKey": null,
	  "weight": 125176789568602
	},
	"NodeID-9izFTRjPnyph3GJu1sp4Cf7zHiwaK7KbS": {
	  "publicKey": null,
	  "weight": 2763712704432
	},
	"NodeID-9kNp7yMjHnDg1QLQboms9ACgQ7BGu8zpS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9kgmvmU19mNsTxPS4f2F7ZiSZSwZ7dfGk": {
	  "publicKey": null,
	  "weight": 2930041178250
	},
	"NodeID-9m519rUcX9um2PdgG3h8tYSZzCRPZ3Dwr": {
	  "publicKey": null,
	  "weight": 2004521517002
	},
	"NodeID-9mtTuxamXQyiePCNyheFEyjr4cAQo4Wbg": {
	  "publicKey": null,
	  "weight": 31614922389680
	},
	"NodeID-9ntqrySGMSMwqYR8yFdChFdU2ABj5zSEp": {
	  "publicKey": null,
	  "weight": 5493967713644
	},
	"NodeID-9oP7KmQjtxi7rK5GKmRXFNuNpPmr2vzfw": {
	  "publicKey": null,
	  "weight": 2001000000000
	},
	"NodeID-9odH3Jbw7uPuRrbaVkVoYeitdzwX6o69A": {
	  "publicKey": null,
	  "weight": 48452250254008
	},
	"NodeID-9okkd3AwB3eurfQxdP46mcSM5sVhD9jxm": {
	  "publicKey": null,
	  "weight": 2877089877109
	},
	"NodeID-9opRroLX5zG47LVmyw8hGnQJx6dxmfGsf": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9owgkqjtxYrUzskf8xM4US4NRSbayEzeE": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9pBxdzsJRoGFdAjEJzcdWbzKEEkdzyGCr": {
	  "publicKey": null,
	  "weight": 52635000000000
	},
	"NodeID-9pTnTcJZ714TcKA5ibKwjc4ztbxnt4YR4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9rD1LoMFV5AP5CbT9CRFqk4GoJrjTPkMx": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9tAAt3ADXpmStdjfpUKsgKCBowaPjXueE": {
	  "publicKey": null,
	  "weight": 3951117604975
	},
	"NodeID-9tJaF6mjrwczuEZk3oQRJsB2AzPk7QFHn": {
	  "publicKey": null,
	  "weight": 403569882352866
	},
	"NodeID-9tLM44jzwQqnsLmQTfhv6Q2GaAEkCdev6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9tWnniKF15vWEqJa6o2hjmL8p61jq4ZmW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9u5MGjm2Y2BvQSZBqwDkqjUrqFKffDund": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9ucNhNpdx1QhUnivTY35AtNEVpjY8AAK4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9wCVwfdLGJe14xss2yjAT8ntowavfe6Cu": {
	  "publicKey": null,
	  "weight": 2009079816744
	},
	"NodeID-9xTBvp3kCUiUGu8aM8XpBWPf2co3vJTJN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9yFkNXGSAUgPu3hB4UxkWGhbp7cXBLaJN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-9zPtXnScuWRvoiTDe498ZtjgoTXwTwxr9": {
	  "publicKey": null,
	  "weight": 642024250000000
	},
	"NodeID-A29C9b97WqpgPBZS8Kh4NEvf7RLRm1cdZ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-A2aPfvr1t99Nm1K82U2XyWtxse3CNmZ64": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-A3azVkFN8rEsnnw7yQoXHDze3gVSSF4YB": {
	  "publicKey": null,
	  "weight": 3042026502473
	},
	"NodeID-A45ZieXX9VhAjrz2jhACBGYYypn2HEjaS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-A49YH2FmGteFpSV46b9WDxZWDEebAnH4q": {
	  "publicKey": null,
	  "weight": 6316453796956
	},
	"NodeID-A4amNjJyWX5kd1wNhxZVqDFqWKomCndp2": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-A4ebHxQGYx1fUuJWf3uWzSaQttx4VwKaA": {
	  "publicKey": null,
	  "weight": 2031000000000
	},
	"NodeID-A58AWZohmoEPEG2FDjqdT8m5XeUL63FZc": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-A5CGvttTf2Bv3KvCeGwzUkxTd6Ux4MURA": {
	  "publicKey": null,
	  "weight": 2516923532459
	},
	"NodeID-A5SHWeGWKUiuvPZhRvNAXKyiEnyDhJzdh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-A5djhLPbKEsfiWZn9na9vFEHThJ6Ud6Hg": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-A5y781VL74i4yfFdcMLnjxr2qLbtjaAgi": {
	  "publicKey": null,
	  "weight": 2352930116919
	},
	"NodeID-A6onFGyJjA37EZ7kYHANMR1PFRT8NmXrF": {
	  "publicKey": null,
	  "weight": 684128823529362
	},
	"NodeID-A7GwTSd47AcDVqpTVj7YtxtjHREM33EJw": {
	  "publicKey": null,
	  "weight": 786846806750323
	},
	"NodeID-A8jypu63CWp76STwKdqP6e9hjL675kdiG": {
	  "publicKey": null,
	  "weight": 527475612766755
	},
	"NodeID-ABHKUic6BDS8Kg95tSWoYv71CXWeuS1Re": {
	  "publicKey": null,
	  "weight": 2027100000000
	},
	"NodeID-ABaXv5WwxLamczWpE6o58Dm5FyngUbFrg": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ACB5BBC2RNLSGZVEGkRRQdWg8hqXjnTpM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ACMpxS8rnwkAA4JfCUm4vNoWQaznNCqGV": {
	  "publicKey": null,
	  "weight": 17650665522203
	},
	"NodeID-AEtF29pZDpKEkZ625bg3reDrux4wyMjhh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AEtKZyYiqnWqzsrwSKEPWrciPHEAn2Q2T": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AEvhXczn6htbBHSuLQNG8KaL6w5YcGSU4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AFWbdZEPvWxEq6A5SThDPVydUj9FSHjCQ": {
	  "publicKey": null,
	  "weight": 4332161326200
	},
	"NodeID-AFZVPKaEHwRZgWvcqRpT6mTd4k8xDubGg": {
	  "publicKey": null,
	  "weight": 2212284907548
	},
	"NodeID-AFnD2mSUo1cB9HSMxGT9hQREHX6V2aCmm": {
	  "publicKey": null,
	  "weight": 5444471457151
	},
	"NodeID-AFtmZMo4p7AZVenJtT9D2Tbk1od5mrJge": {
	  "publicKey": null,
	  "weight": 9620831054745
	},
	"NodeID-AG9hZpGAPNtzpPJWKmJHDjyhFppqM16uc": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AGEnkaodYJksMTnwS8LHHD2X7MDQBSMJU": {
	  "publicKey": null,
	  "weight": 39698619248566
	},
	"NodeID-AGurvCTxa6JFJkk3553Wx8XPtgz6w4xMw": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AHTFPmvMrQBt1hcoGS7HugNCuu15Q5gY6": {
	  "publicKey": null,
	  "weight": 2244000000000
	},
	"NodeID-AJR4Z9bAx8tom9ygqv3PkpQtdxDV8JeLD": {
	  "publicKey": null,
	  "weight": 2009113829326
	},
	"NodeID-AMMTYTe6uZR7yGL1AbEMuwYayd4N1J4bL": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ANQn6yF5HqoipSTWowYpMjyiHifhewv5y": {
	  "publicKey": null,
	  "weight": 3493180817695
	},
	"NodeID-ANeoZC7jBSmoDSYVLip5uASJW48GwWLLG": {
	  "publicKey": null,
	  "weight": 5433301963403
	},
	"NodeID-AQT1i38PMJQxRqjpwrQ98QESW3EJzG3p7": {
	  "publicKey": null,
	  "weight": 5379244888461
	},
	"NodeID-AR44a1cyPrD8cCNZgyhv7Tc2Lvw3ryhEu": {
	  "publicKey": null,
	  "weight": 4365867201878
	},
	"NodeID-ASgm7N2wwZKYc5dZPiHExfDdNcfyZQx3T": {
	  "publicKey": null,
	  "weight": 42935000000000
	},
	"NodeID-AThE3ad1aiPLKsj1XfmhCsNbai8sMyPc": {
	  "publicKey": null,
	  "weight": 2029964663100
	},
	"NodeID-AUeAjRR7q12Pps6Qckyy1PUzFwP34pFij": {
	  "publicKey": null,
	  "weight": 7074525619737
	},
	"NodeID-AUhcXp1So1gL2rZ4oXYN5SABzkpVB57pL": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AV1uDZqWGSheRS3M6RHUbQJHDGXy6f8Y7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AVAXg2LFFdkWJjZi4554GBNprWUepqKxm": {
	  "publicKey": null,
	  "weight": 80238426069362
	},
	"NodeID-AVAXnRLmetkBadrUoqtLSVkohnoEpeSm4": {
	  "publicKey": null,
	  "weight": 6382645573941
	},
	"NodeID-AVp7QsomV2auUPqPXQesRcuKwxn4JmDRh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AWLRFc3R7oCqHd6qe95QsEaoBPPTQYpuv": {
	  "publicKey": null,
	  "weight": 4364122548708
	},
	"NodeID-AWMoYmhjbfpWphJ3nmgxGAPpVFzQbdp1G": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AWirVtsoiwvNdqhqQRwz7vQQEiEyvhq3y": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AX9Fkm7jKPuY9Mit6AyZoS73tFKWvTMF2": {
	  "publicKey": null,
	  "weight": 2039000655065
	},
	"NodeID-AXPrBC45azzYv7UKCwLYJhidPkKat2veF": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AZPULN7CsYdfSq1zepSS3CZNhxJsU3a4n": {
	  "publicKey": null,
	  "weight": 45452926287990
	},
	"NodeID-AZfAXX6qyLxVXBUJZPwEgR1iJYwjwozm8": {
	  "publicKey": null,
	  "weight": 2007600000000
	},
	"NodeID-AZiGXCQ7UJvsVjmBpUMmXAQ12GRj1iaz9": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AZszYbUsVw7PhcssGJYM5274dwoe3gQem": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Aa6HmCaYwZzNoNXRf2xd1aqgtaWdnnryn": {
	  "publicKey": null,
	  "weight": 40026999000000
	},
	"NodeID-AaxT2P4uuPAHb7vAD8mNvjQ3jgyaV7tu9": {
	  "publicKey": null,
	  "weight": 495000000000000
	},
	"NodeID-AbfdkS6u8F8kxh2pPuK32KASNK4sSr4QR": {
	  "publicKey": null,
	  "weight": 204734675683135
	},
	"NodeID-AcVHs6PgeipByAajqn3h7PHDk9cZXDZ2K": {
	  "publicKey": null,
	  "weight": 5775917024887
	},
	"NodeID-AcZuWDkVDZy32y8YSxMVU57sLtkyHbVFQ": {
	  "publicKey": null,
	  "weight": 2810023839293
	},
	"NodeID-Acigg3i7aU9j7coVrns9SEtyq9dox6cKz": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AdSMHubbQcGeJYMcY9nufaieoYjh57uzp": {
	  "publicKey": null,
	  "weight": 19012586453596
	},
	"NodeID-AejE54hVtRxSm2e8KW4zUHpBTgHz9fpEw": {
	  "publicKey": null,
	  "weight": 10326129458639
	},
	"NodeID-Afe4vZ4x36V7NZfRLLwCEurz35jfho3Be": {
	  "publicKey": null,
	  "weight": 2944687257710
	},
	"NodeID-Ag1ZLeovcUsR8V9C3EH6NQpGZoQBhT5Q6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AgfJSHN4yaVjgxanEWMfQXKGDNScwrVZL": {
	  "publicKey": null,
	  "weight": 9544794787956
	},
	"NodeID-AhFT8H88EZjd7K1dhrRPfhatXKuz2iBiY": {
	  "publicKey": null,
	  "weight": 21282513909836
	},
	"NodeID-Ahq7qT8wG6ufLoM8MrvB5G65SbEfZC85B": {
	  "publicKey": null,
	  "weight": 9451971426006
	},
	"NodeID-Ahy3pk9UZEeVLzks1BRSqsLL6zvJt4bN4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AiVA6rmmPhouiy3ERwF7k1eoLTrPErd2p": {
	  "publicKey": null,
	  "weight": 2428194908500
	},
	"NodeID-Aix93q3XbWsEcoNuuWjYuQM49ZeCX7NLS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AkWM2duSrGSfWNEdN1pmWm9ZLrwZDg5MN": {
	  "publicKey": null,
	  "weight": 2728619590154
	},
	"NodeID-AoGM9c6bKC46YpEecEe9tHG5HfcjqPxtb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AoHdNKxXdndkKrbQegnB2TCaJF8JDEXLX": {
	  "publicKey": null,
	  "weight": 15613457566601
	},
	"NodeID-AoU4gTzEVKVmopKx7s9hKXBbQ14VYyAcc": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AoZGrfarMpU4aCdLvnbd52icBe93QfdNJ": {
	  "publicKey": null,
	  "weight": 2276750402622
	},
	"NodeID-AqBVQzsBnqMjTuoMoCAQH7NgTz9vC5APv": {
	  "publicKey": null,
	  "weight": 2013676303082
	},
	"NodeID-Aua1pVxQsWLvTTf8bNVw2ZCA2NfogNqiE": {
	  "publicKey": null,
	  "weight": 4269081306250
	},
	"NodeID-AuqWFe8yzJxPgaZShBZpa3AhEnKT8SmEk": {
	  "publicKey": null,
	  "weight": 14888228054702
	},
	"NodeID-AvTYuKYvu8oB8dXCbAsLh7nTpCvih4fdd": {
	  "publicKey": null,
	  "weight": 246042940976379
	},
	"NodeID-AwBfjtuNZk2Cu8wmQjYsNipqWK6nYMnv2": {
	  "publicKey": null,
	  "weight": 45750000000000
	},
	"NodeID-AwmHrKdtkhnuEVY5TvbwxVksQ8dWd4a1P": {
	  "publicKey": null,
	  "weight": 760246370588224
	},
	"NodeID-Ax97guVgJnjpDYiZkHghBpjELMVrbhKLq": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AxSapUskp62c3KkAnzd2WSDammmNXaG53": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-AyByRGSrfbDaeW4njEYwCzhsnMyZ2KMcw": {
	  "publicKey": null,
	  "weight": 2025022111800
	},
	"NodeID-B39dR3Zj4ZtiqSHTPLxck66yqxTZ9pRmK": {
	  "publicKey": null,
	  "weight": 4534000000000
	},
	"NodeID-B4EU3Tk8BsVczMpkPtJAqS9FSYxNjYNBz": {
	  "publicKey": null,
	  "weight": 1023000000000000
	},
	"NodeID-B4Pubug4ct7TKPRuSAWGqS4Hq79v912VF": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-B5MJkFe7GoyrXY3MqrpjMrrTqj9J5oprQ": {
	  "publicKey": null,
	  "weight": 9356949494775
	},
	"NodeID-B6QdUpYpbcwLNJj55jVknfMnRT1QPaJRX": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-B89cLwRxEtJiSX46XzHA4sTG2ExKXtauz": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-B8RQVdW7nFrVR4Peh6oMfsdK8Muuwfeme": {
	  "publicKey": null,
	  "weight": 1900000000000000
	},
	"NodeID-B8rRoKwovNH4Hv4cgREBqjHcGSCN5Lrcq": {
	  "publicKey": null,
	  "weight": 2832212436418
	},
	"NodeID-B9dnZzwTmygKPxEP7GvyqDh1F6pQHkmmi": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-B9o1a3g4PFp4JWyMJxRv2nvnWHsq1rSvC": {
	  "publicKey": null,
	  "weight": 57564751550044
	},
	"NodeID-BDFwWjNUNeLYrrfGmYqBMqZSGZQtmCcxx": {
	  "publicKey": null,
	  "weight": 267168832683683
	},
	"NodeID-BDnEyGooDS9w6bk2Ty7UncMyQ8t1iLLqh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BE6YSS6Dw6tsEv4fj8smNNkHpqQvAf242": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BFgsRa6TLCKcHZhN9iuufSqRengbXuxQy": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BGEvsCNRi5gr4ga83cVx1E7PTSiBZjyhK": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BHPTN8HM5xA9HNiwUPceNa314kh3YnDpp": {
	  "publicKey": null,
	  "weight": 321908848668346
	},
	"NodeID-BHsQhBq2NYN4YYB9t9VxoxcL2xkCD4zPj": {
	  "publicKey": null,
	  "weight": 5962458442006
	},
	"NodeID-BLjWtGBQU3PxgBSMvhiGxHGZ6Scn1gyyB": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-BMaQ42mVapxbCcX2RLPRTeRW3th91XdaX": {
	  "publicKey": null,
	  "weight": 80873903647645
	},
	"NodeID-BPsD5nsuqTKR7ToqUMDFqanYioFnAVo8C": {
	  "publicKey": null,
	  "weight": 2004591010738
	},
	"NodeID-BQ6uwBhrSU3mwDpguw4rGPXa7kVkk6FyM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BQEo5Fy1FRKLbX51ejqDd14cuSXJKArH2": {
	  "publicKey": null,
	  "weight": 478323529411722
	},
	"NodeID-BQjBjtF6gz6AXUqRbEvN1M5EnFL3Hcnv": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BQpbmskmST3wkcSbypVuFZu7YULXG6pKR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BQsr3wEWirFH6C81M1LYsYagHXVxXRFSS": {
	  "publicKey": null,
	  "weight": 26076940995292
	},
	"NodeID-BR8ysKTDMKpwbgH9fsfCdVQWMsniMQBWp": {
	  "publicKey": null,
	  "weight": 1013267864839910
	},
	"NodeID-BRRp1FdWN4PepbmQNUksWPsRJaZ1mFP7q": {
	  "publicKey": null,
	  "weight": 35516190000000
	},
	"NodeID-BRrFzQrJs8nnmeSZxWwr9yCRcq4U7resP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BSornzAY2Y4SmDcQw4Pgsf5zH4Wdv6d1S": {
	  "publicKey": null,
	  "weight": 10587380317875
	},
	"NodeID-BTUXn6xfxjtA8P4nXYHiwEJpso9Q9t8FZ": {
	  "publicKey": null,
	  "weight": 63809536991255
	},
	"NodeID-BTjbJK42vkCLo8LnJfXHHtVeDrZJgfTbw": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BTybEYu3y8aW54EK9SSivLUYn2fiCosjD": {
	  "publicKey": null,
	  "weight": 945284117646996
	},
	"NodeID-BURrZL8SidbaotZ1cNxvPxNH5DGvXS7gN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BVSFJU3T86CWXQxSTc2Y7cHo7nYJx1Jn8": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BXTBUqX8gitUDtVam4fhRWGD1SfeHGoBx": {
	  "publicKey": null,
	  "weight": 5828997243407
	},
	"NodeID-BYGnq6ZKg5ncN419eBsxd2YQcxoewxfrQ": {
	  "publicKey": null,
	  "weight": 2229521931945
	},
	"NodeID-BaPYDXQYeS9aCFHv76r38gUhYrtfpzq6A": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BaTDuWRvZkEy8gg4FRYmNYXASWL4QKzFD": {
	  "publicKey": null,
	  "weight": 2252665374232
	},
	"NodeID-BaaUYiEUD9ogi7dDikVZAzaVu5yxAHjGs": {
	  "publicKey": null,
	  "weight": 10339930723702
	},
	"NodeID-Baof3ssuxMwMHxd3RXpNwiWmRdskL3hLD": {
	  "publicKey": null,
	  "weight": 116511057043730
	},
	"NodeID-BcfnNieXDpsvkMApb8FcCmGAoQdtyu8jH": {
	  "publicKey": null,
	  "weight": 2265756385145
	},
	"NodeID-BdD2fp6PXxbSdGshiVrHDzarWpgLFEMF9": {
	  "publicKey": null,
	  "weight": 2069890000000
	},
	"NodeID-BeFK3PyWKSnVPjQPRo9t7EBgkrr88M2Uf": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Bf6sxEpjBCbnoouk6sdrPNPfHfxfNumDB": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BfeJfXUfFveZPbS5tFaJar7X2uQQ7S7so": {
	  "publicKey": null,
	  "weight": 2351618894254
	},
	"NodeID-Bg5GpEVehBRU9HDf9e66vXmrRJ39BW1hG": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-Bh9vBw87fdqVzMzrFrPiTQqA4irjDP6G4": {
	  "publicKey": null,
	  "weight": 2517000000000
	},
	"NodeID-BjoNuhod6yNfMJJDJgyfNWWRBUG38ZPCs": {
	  "publicKey": null,
	  "weight": 311892705882328
	},
	"NodeID-BmKvEFG1hXNQ52rPBTXwQ3ZxEmcCmxXZx": {
	  "publicKey": null,
	  "weight": 4138331447021
	},
	"NodeID-Bn4LPbETjf8D9pZE45abJgcCk5KSQCbo7": {
	  "publicKey": null,
	  "weight": 6000000000000
	},
	"NodeID-Bn8ZcnL352FHc4utgVPVDjD3o1HfYwtdA": {
	  "publicKey": null,
	  "weight": 5550010273068
	},
	"NodeID-BoN8r3MqeJiqVcvrQx1VubZMGDQE5pJi7": {
	  "publicKey": null,
	  "weight": 4747374952367
	},
	"NodeID-Bom5dUsayGwdagVLoNNxe1t1FFDZXGGK": {
	  "publicKey": null,
	  "weight": 881968588235130
	},
	"NodeID-Bp4nG1LrP7JCg9eSQpfBE91mBSd8biThQ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-BqtSwNPmkfSi5W3Zz7CRffeLmYrUYuEmJ": {
	  "publicKey": null,
	  "weight": 8097650874533
	},
	"NodeID-Br6QnUhtZGAC4CymYLPRSMCrsNyhLMmH7": {
	  "publicKey": null,
	  "weight": 2162319197280
	},
	"NodeID-BrexWemEB1Vgpbpnnriy3k8e2CyRvxsPR": {
	  "publicKey": null,
	  "weight": 2083217044342
	},
	"NodeID-BtPauDeR52DaiZeBXuJSvRWaFnNhVFMSf": {
	  "publicKey": null,
	  "weight": 1403850772136490
	},
	"NodeID-BtcHfm5o3A4WhAmkTm6YPWRhPCChDPLMM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Btr5hkRibMyYztC9jQg5T3fpKNeYf3Qdq": {
	  "publicKey": null,
	  "weight": 2250000000000
	},
	"NodeID-BvDJUkrDJotTxgZi96zsEe6iYQBLtz958": {
	  "publicKey": null,
	  "weight": 10309421480165
	},
	"NodeID-BvpPGydG4inQCbkwRZ3nPMp1zAGToUTxz": {
	  "publicKey": null,
	  "weight": 87275225018572
	},
	"NodeID-Bw6s9DYKi2V9jRqNmaheGGg2Wz2Dv7DQo": {
	  "publicKey": null,
	  "weight": 3931665886944
	},
	"NodeID-BwsGZ6YpSqF2j3QRRiQJDf9JC4FmCGXq5": {
	  "publicKey": null,
	  "weight": 2091957047379
	},
	"NodeID-BxzE2rHWw4hBSTVqP17p9earG4qYKpCwv": {
	  "publicKey": null,
	  "weight": 79666836734474
	},
	"NodeID-C1oRu2mGfJHUD3UyNCWVJnJ4T6AWXhj5Q": {
	  "publicKey": null,
	  "weight": 8671633315720
	},
	"NodeID-C3aMR9tsKqTNQra8FbqpFkMgHV5DtuJGx": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-C4UwtB1v86mBctdWgoQZGCCUk9EUtjHSh": {
	  "publicKey": null,
	  "weight": 28940998067910
	},
	"NodeID-C4emGUFHEeCEZgoeMDHZddr5nyZBgEUKf": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-C5Gqdho8U8fTWHhBvymcyRVBU43D1WMpc": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-C6b58cQGYSqm1FeQyvXKXUgZ3R7Q9D9es": {
	  "publicKey": null,
	  "weight": 3252991461000
	},
	"NodeID-C6fd6GFYMnbD44LpZR6AuddMDRXTvR5Nn": {
	  "publicKey": null,
	  "weight": 5728976743274
	},
	"NodeID-C6poeYWySHhDuKWMbckvcny24gTGKEtEK": {
	  "publicKey": null,
	  "weight": 3000000000000
	},
	"NodeID-C6zVv1ab5JKn2j8DDCGWGai9jgWhmpLVA": {
	  "publicKey": null,
	  "weight": 4038196247536
	},
	"NodeID-C7VsWUVsE7uFmSgkRQpcjAJWB6gRdm9Pc": {
	  "publicKey": null,
	  "weight": 8485466024557
	},
	"NodeID-C7YyYDva6rvbpY5GEu2Su653eLBNt49TA": {
	  "publicKey": null,
	  "weight": 10687072702549
	},
	"NodeID-C8b7aGd2caAe7PgtP2vnJcWta1wUezAAa": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-C8zmJazq2HwujDgsjKpmYJsvSuNMD1kwt": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-C9Bcji6rXUp22qyWhvYKm4v45hJAkzvpU": {
	  "publicKey": null,
	  "weight": 1113708352941126
	},
	"NodeID-C9CJ3oVhdeLf2Db71JPuzvEUKjy3BdoAA": {
	  "publicKey": null,
	  "weight": 85613663860371
	},
	"NodeID-C9FY2J7xE6f6sGSaThsmcWva6FU5y5Qi5": {
	  "publicKey": null,
	  "weight": 19418239509235
	},
	"NodeID-C9eewJtL6WMZi2ce8YoPxhHTCYemhvnKm": {
	  "publicKey": null,
	  "weight": 50000000000000
	},
	"NodeID-CAy5q5U4AnPR35i225a2QXM7ftYvbZEEg": {
	  "publicKey": null,
	  "weight": 124263792565796
	},
	"NodeID-CBPbPrYDbghbAP4RLSCQr7bHQqHWh1c9Y": {
	  "publicKey": null,
	  "weight": 8735998409966
	},
	"NodeID-CEHgpbVVJh676muJTUNKzP3jmZ6Cs2zar": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CF8WKgeNBHxWFmijN7tysmmRSHfJFEKD9": {
	  "publicKey": null,
	  "weight": 3099727077561
	},
	"NodeID-CFoWfGyYw9HPuDK79fjor9d96S9wLAbpR": {
	  "publicKey": null,
	  "weight": 63369188512700
	},
	"NodeID-CGrPauwXiEvvY3yLBxR9HM25B8HpNTVfV": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CHxae1CeKzs6DNxur9rHcjSAecpdN1p9Q": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CL3N5kkR7HsmVtKavcuofwQGgMNjppY4N": {
	  "publicKey": null,
	  "weight": 9923597065833
	},
	"NodeID-CM5KJnEvKQfjFirjbUzHa7WbanH6dLHva": {
	  "publicKey": null,
	  "weight": 2572790471477
	},
	"NodeID-CMFSQ9wzby52sJqTBeajH2MaJq8HwnZPL": {
	  "publicKey": null,
	  "weight": 2031200000000
	},
	"NodeID-CNPUVCHqT8Jpjpix7SJfXhh184EjPSsu6": {
	  "publicKey": null,
	  "weight": 2733979087618
	},
	"NodeID-CNyGDr2EvPqxfybfNLkpGWqMQHiZthJ3Q": {
	  "publicKey": null,
	  "weight": 1679707793892103
	},
	"NodeID-CPWtTaJjKj3uiQnNKeKVk5yKfEiE5mok8": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CPkJH4zTnyhERzo7H1usRiSmMVDeUszm": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CSmSE2veuHYG3dJD9WuqobvsAdJ8iF5pF": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CTry2Hj2aXeA1rfWq1s6pDL8SzaCFGLMC": {
	  "publicKey": null,
	  "weight": 2153245270221
	},
	"NodeID-CU7T6w5aj2hPEtfHSf1GWXpXKLWnN8uWG": {
	  "publicKey": null,
	  "weight": 87347358150846
	},
	"NodeID-CUTxQEQariSZiwv88qRoGtmUMWabjJoce": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CVJCnaM2gVKr6WCmJxKSPjEA6GFtmBknY": {
	  "publicKey": null,
	  "weight": 2605834340577
	},
	"NodeID-CVXpUPtPZnsEM4mSnJbctBpL6Sgj1rCRt": {
	  "publicKey": null,
	  "weight": 99763142993834
	},
	"NodeID-CWTULAJzqphbmedDkRpcHuCvGiHARvodz": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CWuWBhX2WaAY1pBTAmLRTGsjX9y3EqWAe": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CY9BmJwkVatHX86Pt3Fkxts1QoKTovyAs": {
	  "publicKey": null,
	  "weight": 4715796458450
	},
	"NodeID-CYdCC4uTyAmEHW2CNB6oivzBtzqJn55mS": {
	  "publicKey": null,
	  "weight": 110892914332344
	},
	"NodeID-CZRwtJDY1KeQspzxqmHfSdwvhA8xAeV9K": {
	  "publicKey": null,
	  "weight": 259255058823498
	},
	"NodeID-CZVcr8ZP7njdQynukR8B1RcCcG9qxSdD5": {
	  "publicKey": null,
	  "weight": 12412576988618
	},
	"NodeID-CaNQUnkzWJFcsYNhdx4iyRk31McfqtKhX": {
	  "publicKey": null,
	  "weight": 6721400180075
	},
	"NodeID-CadcH2YHzeVWDSBT2LR9bfqHvMVfH6bTs": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CbFmga7as6xEZXwhgCTmcvU8hix4Kf7ru": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CciiJQMXRWQ5967qBS5XncuEiyor1Png1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Ccx7u3SwfPt52GxWbNc74kjVJvqktTkn8": {
	  "publicKey": null,
	  "weight": 37500000000000
	},
	"NodeID-CeaG1icxFWcMkTzsqQjpwt82nePJrk5gZ": {
	  "publicKey": null,
	  "weight": 17196975103240
	},
	"NodeID-CfM5DQx47iD6AWMsFdHGfWUnt97Sbt9n7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CgXPRNG5FeVSz8qC31u5UDpZ51pJ2mN2X": {
	  "publicKey": null,
	  "weight": 6029634297368
	},
	"NodeID-CgbgB1UFpUESETV1CVGDM17bE9JA4tVHK": {
	  "publicKey": null,
	  "weight": 46491599779256
	},
	"NodeID-ChjgtCzFxwhCYQGsUzGHUUTi9nFYrLsXh": {
	  "publicKey": null,
	  "weight": 10072494427229
	},
	"NodeID-Chy78FPagNNU7oAj1QWzrdPFjJgZP3Scu": {
	  "publicKey": null,
	  "weight": 2041341996259
	},
	"NodeID-CiKdcSyNH27re2W17ygscpZ4xG7474E5U": {
	  "publicKey": null,
	  "weight": 37278139297157
	},
	"NodeID-CjX7mnJM1qasMtuyyxeuKoKy6b9fdSSbE": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CjooWTHxne1o2w1hetMGgDL5Et8JzKMiw": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Ck4VEAv4JTMgnijLQDeGCWST8Zdgxe9Q6": {
	  "publicKey": null,
	  "weight": 2327733155156
	},
	"NodeID-Cmp153iz82jTqJNoNFSvGr8wsaVyV93S2": {
	  "publicKey": null,
	  "weight": 2988999000000
	},
	"NodeID-Cn4ocmPe1Hqo29RyYCC1QXr8b763Wpw92": {
	  "publicKey": null,
	  "weight": 29033335591858
	},
	"NodeID-CoKgGkWBUGouNZTqKCYLBNJ9dw4yJXUEW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CoRQHCd3aSKe7FrhHSUnvNVZzW52xu2YA": {
	  "publicKey": null,
	  "weight": 2030000000000
	},
	"NodeID-Conr75BBUMBFNXD2VP6gvEbCRihCvjfej": {
	  "publicKey": null,
	  "weight": 132000000000000
	},
	"NodeID-CpS69JexgMfH1sy7a77fiwQgN3c5QtrXa": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CqbkQ1txNkMjeroDyJpAtvUAbJf2VbV6x": {
	  "publicKey": null,
	  "weight": 2059999000000
	},
	"NodeID-Cqc4sovCY4ypZww5envqKgtrvd6zrRPTz": {
	  "publicKey": null,
	  "weight": 5771313642521
	},
	"NodeID-Crfqdz3qPavSYvQUWSXMawWjUXxnr3Gg5": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CsGzEmegg1sVokBzrSKkMygKy9YCVEzkq": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CskPetRMvtH5Xr6gLa5cwfY4hR34UgkM5": {
	  "publicKey": null,
	  "weight": 5904317527812
	},
	"NodeID-CtagCY3NrdLi1C3usFLhzRXQbxeVWCULX": {
	  "publicKey": null,
	  "weight": 2004521554634
	},
	"NodeID-CuEMaXrtX3tdegH3BPJowHr7cMUyf7P19": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Cve7X2y1c9WG3AvnWrMGBLkorLdmQTrL4": {
	  "publicKey": null,
	  "weight": 3453094877500
	},
	"NodeID-CvzmyAWqVXtrEcJWfXDP742oE9PsnCyoB": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-CzW1jKjZwFXgkB2Ria2nNPvah4bMq85bH": {
	  "publicKey": null,
	  "weight": 2658976199849
	},
	"NodeID-D12qpTYhB3feogxvMBzS1QPvwyKyhRK8G": {
	  "publicKey": null,
	  "weight": 164569288638166
	},
	"NodeID-D1H3GmiDVvoMsp3foTPmWi9YPXweYv9fn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-D1rgZbjdrevpXXDF1orhAvnUmojVfeZwM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-D2hsJkh6oAf5X9WEuNgbChFvkeUgC99Ns": {
	  "publicKey": null,
	  "weight": 5015755677814
	},
	"NodeID-D2iAoisJf33uGaaKh9uCUMrEjN9HG5dib": {
	  "publicKey": null,
	  "weight": 2004521569870
	},
	"NodeID-D3onstuMsGRctDjbksXU6BV3rrCbWHkB9": {
	  "publicKey": null,
	  "weight": 888243049026222
	},
	"NodeID-D47UrYCDdkfrDrv9gUvYKLUyT44TXYCJZ": {
	  "publicKey": null,
	  "weight": 2117461348490
	},
	"NodeID-D5aUrNtb3vKUBKYpdapiaGEaBTmUz3Jr2": {
	  "publicKey": null,
	  "weight": 6653212480142
	},
	"NodeID-D5wkgZDwkAQVpFNYtNZmGUJGnUYRB97RJ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-D5yA7HbEr8AJ4yzwVZXsDtb4xhKkTRz73": {
	  "publicKey": null,
	  "weight": 100000000000000
	},
	"NodeID-D7JFYevRv4N7LXa8ihCwdzVtc5szT9Q7j": {
	  "publicKey": null,
	  "weight": 2025000000000
	},
	"NodeID-D8vyyYopHECxjFCQ2K4X1sttJjC1B7zva": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-D98ujhSFuA5pXc8bnKPPpbP1VNFnfESVk": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-D99fW2bCQpU6QCN1pu4UccDLZTQdVEwMj": {
	  "publicKey": null,
	  "weight": 2009105776440
	},
	"NodeID-DAA9XwXQnbGUS7aDEecWA5sH9rQJKDsCp": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DAMYbaM35qcXWvuEKt6spGK6EzVKDJUkT": {
	  "publicKey": null,
	  "weight": 703001158957707
	},
	"NodeID-DAZNmxJPADAnk53993F5wntS7LUhxNH2i": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DAtCoXfLT6Y83dgJ7FmQg8eR53hz37J79": {
	  "publicKey": null,
	  "weight": 292166705882327
	},
	"NodeID-DBB8Ve4B6dXejYtFhwcQ8ZMFbmZu3VpSC": {
	  "publicKey": null,
	  "weight": 2004517200612
	},
	"NodeID-DCUpFFgVWQBseGis77N45cBAZudSbaZyg": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DDzbSLsiD9qMPr9eaUhxwrChfbbXW8Zu1": {
	  "publicKey": null,
	  "weight": 33869724953284
	},
	"NodeID-DEa3uaa1iztrxytaMqaP3PdydErzjWEgh": {
	  "publicKey": null,
	  "weight": 6825858119250
	},
	"NodeID-DFH5SbPENeErMmeys9FcQTMvdKkYrdLU2": {
	  "publicKey": null,
	  "weight": 19752822784321
	},
	"NodeID-DFUyhAzXyYi3TUaQYBvm5cezyEaREeSYB": {
	  "publicKey": null,
	  "weight": 11707596365384
	},
	"NodeID-DFuM24ytt31NZquBMCS6X3EHmrbfJXMa8": {
	  "publicKey": null,
	  "weight": 2329247646163
	},
	"NodeID-DHUX369UfTJb2vVuWSRdJr9s27owdMbkS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DJ6nZjHBeTwdQiikyuz1foaQ6dWLLYm4u": {
	  "publicKey": null,
	  "weight": 11701000639111
	},
	"NodeID-DJbw4TrSRu9uovTKqo6w12axPLHWHsjJD": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DKSq9hmU318ceXbU4LBEcSacDPZ5oSNAf": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DKgP4pe8GtAwQ3oVgvHJFG3WexnLjkkSt": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DKxvuUSqd8xtyWejxXC2JvQfP79SAsPvG": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DLFjLpQupsWf5UdaKRFiGELdeY2ViJfwT": {
	  "publicKey": null,
	  "weight": 2146850909718
	},
	"NodeID-DLwFM6gfY5wWfHHMf5tC1NjttQqxZq7gr": {
	  "publicKey": null,
	  "weight": 6000000000000
	},
	"NodeID-DMn1E4EyXZu8GZGKHV3p1naTmpKkpf4fM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DN3fqQmgTBvk5JDxKjCFzmbs2XSVXTKny": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DQiR1rDxgjLJPPgdfDoDxDvFjLgbq1tNB": {
	  "publicKey": null,
	  "weight": 4126279677204
	},
	"NodeID-DUmZowZh6nxKskHUybupqNAhnXVA5yQ1r": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DVh35F3R11HAAxzeSTbk3eR9w2BzPrJQt": {
	  "publicKey": null,
	  "weight": 803994240000000
	},
	"NodeID-DY2Tk5ZqUir1NpSiTwXVGdETCGGEV8wHe": {
	  "publicKey": null,
	  "weight": 5716117432372
	},
	"NodeID-DaKj4jkyyCN4B5zVW65krECqpjxYGUFzX": {
	  "publicKey": null,
	  "weight": 1275908541166417
	},
	"NodeID-DcCkk2wqS6L9XpNX41DVN5RZnwgixCYqE": {
	  "publicKey": null,
	  "weight": 9899130784126
	},
	"NodeID-DceQxAzA4nCGFm9BKa1PEJ4c1zwx3YLuu": {
	  "publicKey": null,
	  "weight": 9074202399131
	},
	"NodeID-Dcs4UwSwQe7vKaopXi9AcM891ybKatAvp": {
	  "publicKey": null,
	  "weight": 36380039364488
	},
	"NodeID-Dd6BPwCsmEQaAgx9hpck23LYe1SopA3p9": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DdJNZBekdB5EVLEe2vE9LusaHjbGmsMYh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Ddm8dYWWMgiZv5dyNWuzgkDprubqB1wkN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DeT8Fw4Tp875GUQNxfcfKoBSPjKpGi6eR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DhNYzUwYHPF24Wfz9MmmVyqLBUF6tkEpb": {
	  "publicKey": null,
	  "weight": 111651058233870
	},
	"NodeID-DhuG2YhqQy26zabyLqTqNr85T9XscsWpY": {
	  "publicKey": null,
	  "weight": 5362035036215
	},
	"NodeID-Do5ZUHCK8DHm194KGqdcwFc5wWE4fCQ6s": {
	  "publicKey": null,
	  "weight": 2035000000000
	},
	"NodeID-DoEC8uw9qRNi33EiYrPoK6Vu9pnsH2rnW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DoTTJmp8cmERizJdYQGFsPVFgZPJxfcis": {
	  "publicKey": null,
	  "weight": 125000000000000
	},
	"NodeID-DomnwLG65RUJM41RgCJ2tRdR4iZEkVCkk": {
	  "publicKey": null,
	  "weight": 6397718916639
	},
	"NodeID-Doy4qhTgun22byNDyHxNSRceygqwhwXCn": {
	  "publicKey": null,
	  "weight": 5331387720133
	},
	"NodeID-DpXBTEP6ZCmm5bgBbv2rb6MzdRovARVhi": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-Dpba952dAUnayPmf52GJzQj1PLoqtXahS": {
	  "publicKey": null,
	  "weight": 8922895437491
	},
	"NodeID-DpwJ1uaLSpSqKu8ZTa7XpmFp5BuBYN4Ns": {
	  "publicKey": null,
	  "weight": 4436266224035
	},
	"NodeID-DqqpR7tQ3goBbudggakvzYs2K9TzR7daR": {
	  "publicKey": null,
	  "weight": 27900281133045
	},
	"NodeID-DrijBGc9SonJrGUwex3vXtzzEGW3cB5DY": {
	  "publicKey": null,
	  "weight": 2995511701712474
	},
	"NodeID-Drv1Qh7iJvW3zGBBeRnYfCzk56VCRM2GQ": {
	  "publicKey": null,
	  "weight": 1013490647058813
	},
	"NodeID-Ds3ri3EksrtyjtAJ3RbX8Wku19LRdKTti": {
	  "publicKey": null,
	  "weight": 439268175129258
	},
	"NodeID-DsMTbudbrnZtFyexvZihGfb9qETaxFBm2": {
	  "publicKey": null,
	  "weight": 9220587519670
	},
	"NodeID-DsZ2yLdYvfHZptxMZqmHKgs4T5ehRTx35": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DtCpt2Bi6gvmQqwgKE59fanof5F3bepnd": {
	  "publicKey": null,
	  "weight": 5945098242433
	},
	"NodeID-Du7SAKQLR3j5iK7tGv7NvTSxiEVTfYfeS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DueWyGi3B9jtKfa9mPoecd4YSDJ1ftF69": {
	  "publicKey": null,
	  "weight": 153356337806361
	},
	"NodeID-Dv4KWhKB56So4dZNxRRi3J2V74TaUXWMg": {
	  "publicKey": null,
	  "weight": 6088515945171
	},
	"NodeID-DvCdC1yWp8a9NNLBR7agvTV4YNmJY7zb4": {
	  "publicKey": null,
	  "weight": 6212101990140
	},
	"NodeID-DvKgKdhkDk6CUf68iqA5Kwaa44RkBa4zN": {
	  "publicKey": null,
	  "weight": 19174729312237
	},
	"NodeID-DvMo3a7eJYQexjLrR1rykrc1jsNuVwnPn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-DvVgVUfihpm93kUJs381BdFcS6goazwUn": {
	  "publicKey": null,
	  "weight": 2800000000000
	},
	"NodeID-DvwWC5ed3j5iG6e4Y5yxAUacPkXvrxLFT": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Dw7tuwxpAmcpvVGp9JzaHAR3REPoJ8f2R": {
	  "publicKey": null,
	  "weight": 309189264642590
	},
	"NodeID-DxdedkGTUuVoRS4VJA1qYXuFU6MerKKyX": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-Dy47rtrYr4eiyRTSnbaFi8KZVnR56XrUC": {
	  "publicKey": null,
	  "weight": 192927531357070
	},
	"NodeID-DyXd5SxgwWWadT9PdEpieSPtjcKxYCMjE": {
	  "publicKey": null,
	  "weight": 2100000000000
	},
	"NodeID-E1c6FEFavQ1RiaNARpxGx4sWnXAGh2YHN": {
	  "publicKey": null,
	  "weight": 12440180280340
	},
	"NodeID-E4KwVKcddwozth12cNvNYApyHhVbvnby6": {
	  "publicKey": null,
	  "weight": 7276271488319
	},
	"NodeID-E4pVcU7hgYn962iz8sQRSf2R7jNKDQewK": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-E5DL2H2CxUUdxe6TcMW7uTHC5mFujF1eD": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-E6GJmPWBnGKad7hPfxCK9GRwsuAKc2MJQ": {
	  "publicKey": null,
	  "weight": 6538621800498
	},
	"NodeID-E6uwzybtqQ8H7aArfxxVJWVsEcaCPLT7t": {
	  "publicKey": null,
	  "weight": 2009075996882
	},
	"NodeID-E8rXzDkKySREDECoXPqaPPhQA7QuxL9M2": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-E9cBEoCDBGi6qFGcW6bfFoyhWpq6Vc57Q": {
	  "publicKey": null,
	  "weight": 3500000000000
	},
	"NodeID-EA7WyeKRdVx34ixXFCvfWbv5P9DSNsj1E": {
	  "publicKey": null,
	  "weight": 4325000000000
	},
	"NodeID-EBF7imfP93Sc4cv3cUAK21maYDvaFMS5j": {
	  "publicKey": null,
	  "weight": 12174140421393
	},
	"NodeID-EBHJ9gHUbjtNQMnNNcJEsqojFpAbLrnVS": {
	  "publicKey": null,
	  "weight": 14497827738175
	},
	"NodeID-ECNKxnabxDwevTW7bLgP2Meaid1j8ztVn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ED4xkYFhZPNNj3WXDQiogcTULeBpDVjSY": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ED9nabFu9bXzdDizCeNytDisKqkdEk83Y": {
	  "publicKey": null,
	  "weight": 8554637543711
	},
	"NodeID-EEipCNkeRgfvwiDHbaBdQ2WeRnsRVSH7P": {
	  "publicKey": null,
	  "weight": 9084987883998
	},
	"NodeID-EFd4jseGxHxQmYX1ATcJUSZoCnE7z6dFe": {
	  "publicKey": null,
	  "weight": 4724993270050
	},
	"NodeID-EG6gFe5McNc2EyFzARxuvTApN995FdBZj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EHJoNS9sgSnjg5gkvKL4AF5vyi8c4qTvg": {
	  "publicKey": null,
	  "weight": 4050371485513
	},
	"NodeID-EHpoaNccurjGvMCc4fFHKXvWKCzCeFsCw": {
	  "publicKey": null,
	  "weight": 5013681396986
	},
	"NodeID-EKH5gkuABvZWiZGHvWuZLgGar4Zts6qX6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ELNdDSKVpamvTWSQvqqBWzuAKFiJLzoT7": {
	  "publicKey": null,
	  "weight": 2312130520324
	},
	"NodeID-ELcEoPFQPwAQmnndhM7QXtVc6sQiJiTHu": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ELjBQTBvaEAMk7Qx7qJNJsEh2DYC4nxUn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ELx4yRvGiDRcG8r8Tow1LUAhMXkVYVzHo": {
	  "publicKey": null,
	  "weight": 2541122251166
	},
	"NodeID-ELxAoLLT4rkHwrAxxo6Ubkg11P4ye32kL": {
	  "publicKey": null,
	  "weight": 5451716361858
	},
	"NodeID-EM1NUe17RtPAJ4ofro1qPc8vEfcvEhjZM": {
	  "publicKey": null,
	  "weight": 2120998317880
	},
	"NodeID-EMETD23g2Bne8CuQKyR8HB15KWomaAuFe": {
	  "publicKey": null,
	  "weight": 916331747193007
	},
	"NodeID-EN2aDGGjPAqm4Npzk9wCjCdGfg5hVrWgp": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ENc7M77QRhgtpDojDQY5nqjndjiYCWR4i": {
	  "publicKey": null,
	  "weight": 8446592848686
	},
	"NodeID-EPqYJkZ7knGWY2TNe4rxRtNgkqc6spyDC": {
	  "publicKey": null,
	  "weight": 274030718451584
	},
	"NodeID-ESRpVRqLwts8M93Ly9QPm4ZSdMF4wEpbt": {
	  "publicKey": null,
	  "weight": 2367379089801
	},
	"NodeID-ESVMWUoEkKz8uAEDfKcEpu5UUCkA6GeyC": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ETwxZVMHXcikZmKsKfjiZuVG84KsZDdqz": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EVchqoADmKDHejumCGQDgZKc1RoqAWnFY": {
	  "publicKey": null,
	  "weight": 4039392099154
	},
	"NodeID-EWNvJxWq2UBBsn2UH35soQugQJn8CVKdS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EWmdnUhsVGSfZPnbuSodLUTA6ReJcXz8v": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EY4iiXyj2TfoHspEHBHFfBYrHxXESnSN5": {
	  "publicKey": null,
	  "weight": 7863493133306
	},
	"NodeID-EYDcoNgbX22og1oyEn5s7iH1BYSmR5yF1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EYgNPuFCT4Qy7YhvsM2s8WKTT74QECvTV": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EZ38CcWHoSyoEfAkDN9zaieJ5Yq64YePY": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-EZ7FxvAP8MCsQgCz36DMBUGA1tcHL2PG7": {
	  "publicKey": null,
	  "weight": 2879609820783
	},
	"NodeID-EaEUnvUyfTr19uMW8g51bvFXfi7r9xL1W": {
	  "publicKey": null,
	  "weight": 261301235294055
	},
	"NodeID-EbwYUq8FZAK4jf65nZ4dcb5kk5WZ7RBgb": {
	  "publicKey": null,
	  "weight": 13890040680639
	},
	"NodeID-EbzFVG1XWNeioXm9KKPCvidgkZELPrj86": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Ec6W7Mq2fnotBFoiD6rwQSsZq2XKVEyTF": {
	  "publicKey": null,
	  "weight": 11577997823636
	},
	"NodeID-EcoA7iNFXhc2DwrmWxiEvjvsDvDmQrEKK": {
	  "publicKey": null,
	  "weight": 2531460778756
	},
	"NodeID-Ed7NMisAoxzWRjxQVrWnBvMShLNcydRa6": {
	  "publicKey": null,
	  "weight": 4746298840806
	},
	"NodeID-Ee6MdWL7mFKdPAXbJTWZGvwEUsi17869c": {
	  "publicKey": null,
	  "weight": 7657482830786
	},
	"NodeID-Eeg2Huipo4K5oURLwd6EJNm286b1js6rL": {
	  "publicKey": null,
	  "weight": 866907256027648
	},
	"NodeID-EekL2aJ7TdkbiXDzSsPo5PKyaaQuQqVg7": {
	  "publicKey": null,
	  "weight": 6486522641272
	},
	"NodeID-EfnXCVeLUtW24v8gwW1VYiNuqDeRPEVYC": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EgAUHBUiTj1QTvNmw7cdqWDe7W7P8B2PH": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EgoLXZqsn7pgCbT8Y4VRi2QpJVvCTN6YJ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EgzctrMz5jVPbfUEWGufKxf9CXUMLBRGG": {
	  "publicKey": null,
	  "weight": 7522140543438
	},
	"NodeID-Eht1FjvmFncRbagWk1PGMyE7hVrCS9Acp": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EibWFSUHeLNNjWZrEznmuxH7ZWNsavRgh": {
	  "publicKey": null,
	  "weight": 2231352602056
	},
	"NodeID-EiyGqDavdQ8dMUtN5wjAezDXAWwVBdSBK": {
	  "publicKey": null,
	  "weight": 14938044261750
	},
	"NodeID-Ek1pare46s7nXPF9JtsoyM7PcaFnCBms5": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EmPwabyobnM3jYDvQuxZdLBTut5V5pq2n": {
	  "publicKey": null,
	  "weight": 5304565934972
	},
	"NodeID-EmVPjC4ePzg8perYoF3c3ut6fHHDDfAqF": {
	  "publicKey": null,
	  "weight": 577937765640403
	},
	"NodeID-EmWrMTi9PcnYH3ZPCXky3FDDTatB7bfuk": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EmhfRnDPSzpEMhGy3KSC4BWhPdtoxmqqV": {
	  "publicKey": null,
	  "weight": 2483458931115
	},
	"NodeID-EnANSJ3hrErv6v7ZGSi8p5RCHhdhbGrb": {
	  "publicKey": null,
	  "weight": 12499346233165
	},
	"NodeID-Eo31HyvvA1Madeagm6U8CM376cEfgS4Bj": {
	  "publicKey": null,
	  "weight": 792000000000000
	},
	"NodeID-EoJirncHtvkCuiLac47FracheuCHmMxnL": {
	  "publicKey": null,
	  "weight": 393587738714932
	},
	"NodeID-EoNXC39QyT1eHFPXb5PgvDNigBN8tMeCC": {
	  "publicKey": null,
	  "weight": 3474067228350
	},
	"NodeID-EpU1DGKMMvDRtag9u9G7uwZHRsJF5NodL": {
	  "publicKey": null,
	  "weight": 2004521068470
	},
	"NodeID-EqBR9F3QBBrTQTN7CDN7sDEUsDGDdMYps": {
	  "publicKey": null,
	  "weight": 1441098346500934
	},
	"NodeID-EqauaQBwVXnpamJvLi84aeZLTQra4GDep": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Ess2uf91CDTfKsLezaTLKqqsWcpEHGvJv": {
	  "publicKey": null,
	  "weight": 4338030079243
	},
	"NodeID-EtFG3SrbbudeFQCaWRxhwvv28wHpo8VRq": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-EvY8p6SYho6VwicRvVasSBBv1MtanH19D": {
	  "publicKey": null,
	  "weight": 50367197887483
	},
	"NodeID-EvZXX2K2eVEV7n6Y8cnmeSRgj4JhkCVad": {
	  "publicKey": null,
	  "weight": 233113411764702
	},
	"NodeID-EvdPJG2AAVGCP2s3543KDx1LZ1g5SiH2Y": {
	  "publicKey": null,
	  "weight": 233670240237782
	},
	"NodeID-EzKkVMd9y6XqzQsvRDLuaaiXqXK8fFipU": {
	  "publicKey": null,
	  "weight": 6590950252828
	},
	"NodeID-Ezk1hTiN5bogETxygJ5NciRvHXgXvUVVM": {
	  "publicKey": null,
	  "weight": 200034419782535
	},
	"NodeID-F1K4jks15y71PthvXVNks8pgutgBwdHZH": {
	  "publicKey": null,
	  "weight": 75863989808967
	},
	"NodeID-F1KpunDAMisiGKLX3R6k9uBoJoeLwzJoU": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-F1qvuNMn9XBXSjv9iWyyMGfV2PQ2Na7KM": {
	  "publicKey": null,
	  "weight": 7538138550031
	},
	"NodeID-F3SZA2ZNdRjTBe3GYyRQFDaCXB3DyaZQQ": {
	  "publicKey": null,
	  "weight": 2341000000000
	},
	"NodeID-F3b5mj8PWnMP53hFbfGNxkiXpbK6tPyDU": {
	  "publicKey": null,
	  "weight": 3230000000000
	},
	"NodeID-F3r4EqguPQgupLCYrTG3SeJSCi5qtc3x4": {
	  "publicKey": null,
	  "weight": 2206491740663
	},
	"NodeID-F4Q66oFkqCmRnBjDUAh28iXvWx6WxbxJJ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-F4SmFmnXuFBkW5fgFgh5RT2T5w8KseGRN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-F4kKmjMhZJWkYwsWyPSLFxtNcj7BhELmY": {
	  "publicKey": null,
	  "weight": 63186049381646
	},
	"NodeID-F5sBP4yNA2YXHbD6FJUjsZJ5zhoMqMi6g": {
	  "publicKey": null,
	  "weight": 2373992900679
	},
	"NodeID-F8HvQKYDbonSeJ3x42Ynb6ibxiJa3FWVL": {
	  "publicKey": null,
	  "weight": 4364122548710
	},
	"NodeID-F8VWioW7dC64159mMhswhzpkHxnB34cjS": {
	  "publicKey": null,
	  "weight": 5276066939455
	},
	"NodeID-F8Yq4R4Mx1LHGbQ7g8s3beNB4LGnCUBpv": {
	  "publicKey": null,
	  "weight": 2890189911851
	},
	"NodeID-FBtCkPRmkfDry8GGpNenBW2fn3fD3qoKn": {
	  "publicKey": null,
	  "weight": 84009727008586
	},
	"NodeID-FDMe6hokunP6wTY3GFQ9u2uVPMtrfXGFw": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FDTRdm7eS7sBuzAnQqMrXUmNiuNHP7VM3": {
	  "publicKey": null,
	  "weight": 8875203043348
	},
	"NodeID-FDZXWAQeuMzZVq8tv53b1NdiwyJDUHGpB": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FELjhENFRLHTgVrwiKTdL7aSQbU8mBd5Q": {
	  "publicKey": null,
	  "weight": 3876186911930
	},
	"NodeID-FEUSZzFmjF2JyrniKXksGfRjkQUH5R4eH": {
	  "publicKey": null,
	  "weight": 12139410002210
	},
	"NodeID-FEtgjYGndfWGzLZELTJJoikpSobjVKTjk": {
	  "publicKey": null,
	  "weight": 290530654819994
	},
	"NodeID-FF2Szx54cN5iys83j9H5xaEqHiz3RipkG": {
	  "publicKey": null,
	  "weight": 3499842184895
	},
	"NodeID-FFWS6gV9FrKecNWCYJQuyR67t7GY6UiVq": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FGRoKnyYKFWYFMb6Xbocf4hKuyCBENgWM": {
	  "publicKey": null,
	  "weight": 1001165117647040
	},
	"NodeID-FGt2J8WREPFbbTXgVU9YC5NjdERhoMa7E": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FHGwnPLGdU3FBzbwnhayx1sDjxiiyb9sD": {
	  "publicKey": null,
	  "weight": 63966006999110
	},
	"NodeID-FHHEXLYRNrnQhKFneEwDPX8TZV8WtUpQY": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FJ7WdKoGBkqb78mMGrr3r6kSL45R5Vspp": {
	  "publicKey": null,
	  "weight": 2274838779758
	},
	"NodeID-FJfF4TPzajcf4QNHjHUi4ctGPq2AxkGJt": {
	  "publicKey": null,
	  "weight": 2250000000000
	},
	"NodeID-FKCbGm3jmceEEpfSdA4uQUVANpmqFLAty": {
	  "publicKey": null,
	  "weight": 3875301174520
	},
	"NodeID-FKGhEFYnHUdFadoR2ePvdTTM5Sz8tgjHp": {
	  "publicKey": null,
	  "weight": 7534914295785
	},
	"NodeID-FLQ8ifj1DwXuRWKShjck624SAgYXVpQzz": {
	  "publicKey": null,
	  "weight": 2280600416050
	},
	"NodeID-FNfrZNBD1bxGj1VmdgSrEKT8oMePC448e": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FPwgrsmgeX4DSkG4abJ28zZQhjcqirPYi": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FQKJ8yATVmwGuy3LN2ExsURDEtZNWgr3v": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FS8PXhYXxoo5GoTHoki5Y3Fqk5zzdMY6E": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FUpfjb5sPRnT3XxMbtephYBc1MHZE2RYL": {
	  "publicKey": null,
	  "weight": 4513432658071
	},
	"NodeID-FVgPdpTafzx8jazSNoxvQmJNxoEcAMmxm": {
	  "publicKey": null,
	  "weight": 3060560108794
	},
	"NodeID-FVgkUobMuUBABuMA2KngEsm2SG9XLGvdW": {
	  "publicKey": null,
	  "weight": 2220965636777074
	},
	"NodeID-FXzsVMpT3ipDikUzabDtmeLrXW8JYyNZp": {
	  "publicKey": null,
	  "weight": 2004521559624
	},
	"NodeID-FYv1Lb29SqMpywYXH7yNkcFAzRF2jvm3K": {
	  "publicKey": null,
	  "weight": 503791245788051
	},
	"NodeID-FapbxinG8eCCXzZQKCjxicKcXijmnKzD3": {
	  "publicKey": null,
	  "weight": 15000000000000
	},
	"NodeID-Fd4DMc88ELLiTHHSLVZ7L4MB3Z7jbDEAy": {
	  "publicKey": null,
	  "weight": 9459947489300
	},
	"NodeID-Fe3121hoArBshDTJEEeMeJTNW6HcLXxT1": {
	  "publicKey": null,
	  "weight": 550755964239958
	},
	"NodeID-FeddPuWb8xtJeuXCDa4CAbjgv3n4PBQXV": {
	  "publicKey": null,
	  "weight": 3827954149193
	},
	"NodeID-Ff4Fx8mcbfBVZSpMFwQKDUURzufmLjLgS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Ff4s5gpeGMf5HaBKNxBX1uRxB2hxRx5By": {
	  "publicKey": null,
	  "weight": 4568109557985
	},
	"NodeID-FfRKiU9BpTzMx2KGmrN4YyiwftTAXZ6FH": {
	  "publicKey": null,
	  "weight": 97259598281971
	},
	"NodeID-Ffh3UfBYLouumUQgGuFbSTBbD5wLXn7xG": {
	  "publicKey": null,
	  "weight": 8886250780183
	},
	"NodeID-Fg3L1eo53UWpf3J9Y45PK2ksQ1zyqajHL": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-FhH7R1HX2ipEetm1eH7u9aKfBCtCQyuQM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Fip1D8AHnD6XbWpFYHskGN4hhMN8Kym57": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FjPxKjpGqPNj8NyEVVdJgGT2bXFVkNyW5": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Fjgs4ALe7yajXjZYmcRAvaQBGrFE61QME": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FjpAr4XXwanzuhLMajEVdm6bwdcGgtZNv": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FkHyxMuq6DGNLdQtMKiLzhQQuwjSCqwHj": {
	  "publicKey": null,
	  "weight": 43935852859553
	},
	"NodeID-FoVyB5QV1hmvAxsNunHRRzonCBJGkLvtj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Fp9D4hAMGrJFnTNrVmGxKBb95MJueQUan": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FpA1n8w3xHmJ1QK3iDAS49vbLJokejfMA": {
	  "publicKey": null,
	  "weight": 4150656229676
	},
	"NodeID-FqiffXvatKSLPRhkkc5y2MkALrZ8EFN91": {
	  "publicKey": null,
	  "weight": 2009076628946
	},
	"NodeID-FrGVKPnD3xVmoMK1RjtbD6evNGk3wrpv7": {
	  "publicKey": null,
	  "weight": 5119816140522
	},
	"NodeID-FrPwbz7L7qvxbxR63g6CMpWPDEyHiiqrD": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FrTRZzNgZZoHMGvC5t6XtJKAv7E2LvYYU": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FrWaLbhUVQHEzXfeekv6MCbkmLfx74HrP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FrhMaTxsagqjYCetCyxLxpbcpkN3cjVLH": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FroHhWhFsUAh3s6Y4hjh7mx3E79mBUZGj": {
	  "publicKey": null,
	  "weight": 2376461948277
	},
	"NodeID-FrrgfMhhK4dn5pD7iKD1JdUjabTe7qWxx": {
	  "publicKey": null,
	  "weight": 2162000000000
	},
	"NodeID-FtJRoxeiXq4D4hsM9PCAZkoerJDTqPEW7": {
	  "publicKey": null,
	  "weight": 4391528397052
	},
	"NodeID-Fv3t2shrpkmvLnvNzcv1rqRKbDAYFnUor": {
	  "publicKey": null,
	  "weight": 48332764705871
	},
	"NodeID-FvAAKfRTbxB5ANcapWNtQ2VGMFWxX3LCh": {
	  "publicKey": null,
	  "weight": 7979406074149
	},
	"NodeID-FvFBsrTf6BhCcS1nFPJtRrWZ5Nosb47h6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FwGJCx9bXYP9VqBxGrwyXQ2UiDhChcZfR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-FxZHg1hxvYyBWFCHzqQPhFYsXte84UT9A": {
	  "publicKey": null,
	  "weight": 6015397310683
	},
	"NodeID-FxaeyWwirpBBbBwRgpytNrtR3NY6xTiwT": {
	  "publicKey": null,
	  "weight": 26169686914884
	},
	"NodeID-FyTHAQeNPFhoDetc7pznBMJ5BRj5s6FJr": {
	  "publicKey": null,
	  "weight": 21060483340496
	},
	"NodeID-FzJA5i97oBDNiKeePSfHc2t3zt5cRHnjJ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-G1KHG6SL2B9jnfbzdTudhjRPa9hp5tuRQ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-G36WdNdkGY4aSh77e1DziUy9SXTgohYCE": {
	  "publicKey": null,
	  "weight": 76390133983174
	},
	"NodeID-G36seLQVfXnaUgaNobbSWg8anRta2WUwE": {
	  "publicKey": null,
	  "weight": 136475261184099
	},
	"NodeID-G4dtkvboFCTYpnLyHYMVPq8GLxns99NPq": {
	  "publicKey": null,
	  "weight": 2420310000000
	},
	"NodeID-G6FCgKDoufiETLSBsfC5sENYYt3FApMkM": {
	  "publicKey": null,
	  "weight": 2036700000000
	},
	"NodeID-G6nXgtucmeKn4nkCp8YTDdM4btVSnW4ad": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-G6oS8Nbb74RyJatPe34iwhTUPDqJNtuwD": {
	  "publicKey": null,
	  "weight": 5908029020646
	},
	"NodeID-G7aPjXRfR2XtrB4PGATDS8NoqxmqZmryS": {
	  "publicKey": null,
	  "weight": 2310005833626
	},
	"NodeID-G7kG9MRuLaJuMLopdfnWEKUPQp54qc2T3": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-G8uVCikgvakai25byQwYx1z56N33unmj": {
	  "publicKey": null,
	  "weight": 3240775052653
	},
	"NodeID-G8yv8mWQy8FLdFfAhJ4v294T5Z2DWDVEH": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-G9aa4VXB19wnUzrmVq2SheaxnKEieUW3S": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-G9ucCcZoHdXC6mrVpKVqKzHsMrnoVZLU8": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GAEZpCd2Lg5RSrPtQhGvxw6Kid5tAL6bw": {
	  "publicKey": null,
	  "weight": 2025298000000
	},
	"NodeID-GBvezLieTFYRqHsd8QzWqH1JUhuwD21mJ": {
	  "publicKey": null,
	  "weight": 13141169546173
	},
	"NodeID-GCL85wyq71aw2FMRTct2PJ2F4qGQL7WH8": {
	  "publicKey": null,
	  "weight": 2207292250677992
	},
	"NodeID-GD3XH39S9feuK74XEh8AhhNEapSt2hqHo": {
	  "publicKey": null,
	  "weight": 2964419396127552
	},
	"NodeID-GDmLiXrEtmFvBiE2s1ZGDYdZZXdkkLW5F": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GE2jerBVdHtvfKkqfNHwBr2MCkDf6tLir": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GFw4C6SgNLkbVq5EeaGtZTjGKwZzmWbUz": {
	  "publicKey": null,
	  "weight": 6967194363100
	},
	"NodeID-GGHpri3tVbRgLUEFKqxNivRbegTKBAnyG": {
	  "publicKey": null,
	  "weight": 4005592547479
	},
	"NodeID-GGfhdfGbqenvgChEXpdngjkKYAByo9di7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GHBthsGgAzrPLXuXk7vbq4qsc9CvSbwXS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GHY8D7pCGPY2qScDmtLp2PAvpbnAiVib7": {
	  "publicKey": null,
	  "weight": 5004611792287
	},
	"NodeID-GHZZWcg26qxY8K9uDLXu6Gc9Gh3P11dK4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GHd5dmrKvYkAbxii5V7wmJw8sRf4JqYAi": {
	  "publicKey": null,
	  "weight": 163568596494672
	},
	"NodeID-GHomCXMMRTSPRZrrxEULYPFHHXEhb7qvY": {
	  "publicKey": null,
	  "weight": 12898819479870
	},
	"NodeID-GKBykzpxeALzYee1fGQjFn6oZiA18W26m": {
	  "publicKey": null,
	  "weight": 9614179161447
	},
	"NodeID-GLSSsNA2oHSaebdYDBQ3u98kvHfH7zjYf": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GMJcCQwhNFDGQy4NAkKM563BMudjmHpXb": {
	  "publicKey": null,
	  "weight": 229607623193700
	},
	"NodeID-GMoxCucPFt2LPSpTh5hPbDPAp71fiV648": {
	  "publicKey": null,
	  "weight": 19208640149345
	},
	"NodeID-GPYurKNwtkhJ7M7UQoDinFz5P7Rsu9yUo": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GR11oP3kLHf82Xbi3H2kewqoM7ydHF5ES": {
	  "publicKey": null,
	  "weight": 600140753989114
	},
	"NodeID-GRReUAPgHbGdXcv5Btpm4ByGvLrD1NgXG": {
	  "publicKey": null,
	  "weight": 69720850851209
	},
	"NodeID-GRrpJze5AEaBezaYp5QRtVcLAF6AzdJSy": {
	  "publicKey": null,
	  "weight": 2218015062729
	},
	"NodeID-GS68HENWev9AWcUrhLqxYvsv2opJBZ7km": {
	  "publicKey": null,
	  "weight": 3954656796890
	},
	"NodeID-GSgaA47umS1px2ohVjodW9621Ks63xDxD": {
	  "publicKey": null,
	  "weight": 1009744212103785
	},
	"NodeID-GTu4yYzFRv4QGXyyvcorGDUMWjzzMWuqN": {
	  "publicKey": null,
	  "weight": 6198881030357
	},
	"NodeID-GU7boimY7cxeqdDmnPRyYrwLiRmwiHxhn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GYAMwLv3XViragqqJUQyTNRcUtoRWbi5F": {
	  "publicKey": null,
	  "weight": 176122057261226
	},
	"NodeID-GYYGoe6y5TtL8vNBVkvTqk8sPTVVqbGaR": {
	  "publicKey": null,
	  "weight": 100297446939444
	},
	"NodeID-GZpKzj11mC7eg52BWpL6x5YpQHBAhKPaD": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GbPDdvSVXhGRjzExg5Fp4fFQAiRWeM8EJ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Ges52zCMZimxqri3ia1J32mUAKyUsvcpV": {
	  "publicKey": null,
	  "weight": 2025000000000
	},
	"NodeID-Gey3HmUZyAXuvawThT1McvEjrofAwrGi6": {
	  "publicKey": null,
	  "weight": 2327356073123
	},
	"NodeID-Gf2f3hWkqxbYtdwPHFRJXQwf6PFzvFigx": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GfDdBdRqzEKK55gCKqThE2LVxLdLhJY3A": {
	  "publicKey": null,
	  "weight": 15500000000000
	},
	"NodeID-GfRY31feBaTknS7uvpF7XHromUfckDeWh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Gfeb3fWPCSpA4g9edG68NS5aJcYpXn6T3": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Gg8SEQhx4vU2KWzQS8ESCrSpcjzEqQ7kY": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GhG3ac1x2g1Fc4xPWXxXRoTLmMuT2asMV": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GhUSE8wcdouEgd3xeCvmYhmevkdNetfXk": {
	  "publicKey": null,
	  "weight": 35566592772119
	},
	"NodeID-GiEUKDy7YfYRdxD5rA7t6CgEQxyrWP9sj": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-GjYkDpBkCueF4vTkZVdtq3yFiMJha2XpF": {
	  "publicKey": null,
	  "weight": 509975294117594
	},
	"NodeID-GjjemGsR2kXtPLSwv4oNV3JFt579oJUnn": {
	  "publicKey": null,
	  "weight": 110349396008988
	},
	"NodeID-GkiJ8CRTibePNf3Y2F7LbGTinHm71HvnF": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Gmitibwg27b6WqHLjqmGiXZAh1kmZVq5P": {
	  "publicKey": null,
	  "weight": 9672394316189
	},
	"NodeID-GmzoRvWfUdTTUZBmWFM8rswgmwMKkcDAB": {
	  "publicKey": null,
	  "weight": 3334286253666
	},
	"NodeID-Gn7jLTVPuvdrx6mF5hxSbazRoKPWsM14s": {
	  "publicKey": null,
	  "weight": 151287437160683
	},
	"NodeID-GnKLq2ipT4D2CMcrrGYndSCuQL58yhCR7": {
	  "publicKey": null,
	  "weight": 4314244983456
	},
	"NodeID-GpDhyVHYVaL8qXFB2a1QPBsXyUMZjiXLF": {
	  "publicKey": null,
	  "weight": 2119513563100
	},
	"NodeID-GpEnAaYG182MkURWh8KY4wL5KgWB7cBw8": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Gq6cTyqtHRGHSHELPNdwXCjQJH6vwUkBc": {
	  "publicKey": null,
	  "weight": 2012037023644
	},
	"NodeID-GqRuMoELZus5EYMwBpiBVtWFDnvLmaJj7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Gr9xzzL6emX5jWv1LoMKLpCebYR2V2nYs": {
	  "publicKey": null,
	  "weight": 399925000000000
	},
	"NodeID-GrkiQsGvbv9PSqJf9r3eE79fUgbZqydmK": {
	  "publicKey": null,
	  "weight": 3972848271290
	},
	"NodeID-GryC4h9okVKe5e9iZAWFxPWtbtMSMervo": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GsQoMs5GAL8inkh88XPVJWKMwmJqm5T8N": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-GuTDsR1tBFgQUbKqCf7J29yBo8ofudHEZ": {
	  "publicKey": null,
	  "weight": 2322724062400
	},
	"NodeID-GuWMRuFkQPvPjdnuegSQYDmTMbj3UtPBm": {
	  "publicKey": null,
	  "weight": 2009076200410
	},
	"NodeID-GujHWE263qNN3MjEYhEeAQdPpQdhHquLJ": {
	  "publicKey": null,
	  "weight": 4055482208351
	},
	"NodeID-GvEoakPRjaB9jTKup6nAMpFSF3HiEeocP": {
	  "publicKey": null,
	  "weight": 10193937924805
	},
	"NodeID-GvQdXQD6qFzfAYBvnvxiep7WaXVj2FSnJ": {
	  "publicKey": null,
	  "weight": 6386031431605
	},
	"NodeID-GwcsDtmQrD8thUHiQWgrpRvhuMjxTqZP6": {
	  "publicKey": null,
	  "weight": 3214248993941
	},
	"NodeID-Gx4oY9nZmAzJrHUfobrmuK9LrJhwejykC": {
	  "publicKey": null,
	  "weight": 47534940040919
	},
	"NodeID-GxUaibkyR88sErviW2WK3ukDJsT3DFpoU": {
	  "publicKey": null,
	  "weight": 8934579234575
	},
	"NodeID-GzSFF2tfNMc7YdJzGeNeNxfSWjDvENJUt": {
	  "publicKey": null,
	  "weight": 48859424988976
	},
	"NodeID-GzhNocnYWSCziBTP5Jj4xAKn29xDKW9Yr": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-H1R3mBHFDeBoQzVYyHzX6RGnVV4nMTrWZ": {
	  "publicKey": null,
	  "weight": 6168838402088
	},
	"NodeID-H1YSzbspeyANTzErprmBQgV2Crq5hXSW8": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-H1yZ17sZyZe12kT52ywW3QnVBySqUHLET": {
	  "publicKey": null,
	  "weight": 4314620450651
	},
	"NodeID-H4mD37NM384RtLGjz5aERQoNCrtGEjgK4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-H94Q5zdkBByZVWvqrB5ByiX4dZdG1J45y": {
	  "publicKey": null,
	  "weight": 2819678708380
	},
	"NodeID-HAYp8Bi1yZ4dy4JeBQRBydGC8NqPEURXC": {
	  "publicKey": null,
	  "weight": 2034247730264
	},
	"NodeID-HBZn6SECcmcGkec36VCqLFsNjzaUZmpBn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HCw7S2TVbFPDWNBo1GnFWqJ47f9rDJtt1": {
	  "publicKey": null,
	  "weight": 500000000000000
	},
	"NodeID-HDXjAjk4rG42HXey7gzQPx6dMBfpWeh5k": {
	  "publicKey": null,
	  "weight": 3975000000000
	},
	"NodeID-HDmmLUSLtok2sGQEruUsYSMjBeQbUrwpH": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HFFiNVtYmv2YF2WHCZWHSCigo25RxgnFz": {
	  "publicKey": null,
	  "weight": 2025950000000
	},
	"NodeID-HGGdFL79J8hGBoMMoimh48CcKRpCxHcb": {
	  "publicKey": null,
	  "weight": 117574754644496
	},
	"NodeID-HGtZapUapv9uNq2JyUA3njmdFCPJbjdqR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HHQ4BFhNMVJwMz8b5tenWtkzdcSYvonmi": {
	  "publicKey": null,
	  "weight": 42977102454429
	},
	"NodeID-HHV1onrMrN68JdRoKfQGeZphebgRrcADT": {
	  "publicKey": null,
	  "weight": 2888195819894
	},
	"NodeID-HKskHZG5zu629RsEkScAaGrpfGQABqep7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HM1vYJWnjWhJkJi8dErcu4p5UDxwUKbp9": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HQ4Nfe891SxSJ2FTx3e9AdabywKG2ufYF": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HRZAiPM3BaZJpAq9jy3UME2EGsrv1ZMaa": {
	  "publicKey": null,
	  "weight": 285128235294114
	},
	"NodeID-HRmKWRK8HjqqTKWvYMBNv9hcddSYrhqnF": {
	  "publicKey": null,
	  "weight": 2000998000000
	},
	"NodeID-HRsnjv1s8VzZz8Ybuiyx9MgfDevUcPeYF": {
	  "publicKey": null,
	  "weight": 2066964948138
	},
	"NodeID-HS1PKAMjZczSb5ZQRastn1WG6eJeNmahd": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HTAhjx3CNb84CJUMsGj4SbpyvEsuehcWC": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HTF8ugVvxeY2snCGE1UcoTbQC9DUYeDob": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HTG7W4VHoGJBUTip1gKXzXGtEBhiW1DXT": {
	  "publicKey": null,
	  "weight": 10898635519880
	},
	"NodeID-HUNYARt6cNSXgANtsSakokNpRCUKCFoVS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HWwYYkiTKnXgYycNN5gXQfxhhA8TZdBvs": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HX7H9nBZX3nZRvLVUxCDBesVX6n8Uv4a": {
	  "publicKey": null,
	  "weight": 2004515538692
	},
	"NodeID-HYBD7r66XFw2Kx5UydvFNAb3w3nnJVtY2": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HYwzPrs3tDR6yyMtKZo22fq8grJsP1DVF": {
	  "publicKey": null,
	  "weight": 2000825602671
	},
	"NodeID-HZnBEWRUGocZ4HYX8EuMP2YEY94NtySWR": {
	  "publicKey": null,
	  "weight": 4849580631160
	},
	"NodeID-Ha5u99aoKUdNdTdomp8bQcWHUw8PSF7em": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-HampaaUVCoGtR32oic9Jx1rFE4zqK6ZWB": {
	  "publicKey": null,
	  "weight": 7401209159770
	},
	"NodeID-HauVjq2twoqYvMdNe2dCqzSJoyvG3DJc1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HaxdTdeEr1sWMRZvXwRN2iotd6RYd9XYM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HbmSiec8zTKiPWxaBwStg58a3N5KePhC9": {
	  "publicKey": null,
	  "weight": 4919834649736
	},
	"NodeID-HcFuyTmh2wPEDBCQc7T7YqUS3N3qj4oRe": {
	  "publicKey": null,
	  "weight": 41488165447782
	},
	"NodeID-HcmVRX8VUbHUpax33WivuLTQkDpmEjat3": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Hcq3gpdMqAEiVwq19E5DZhsjggbkxFx2L": {
	  "publicKey": null,
	  "weight": 8228563046692
	},
	"NodeID-HdG8EdGWAy8LR8YPA8hndoa1uwwgGv1si": {
	  "publicKey": null,
	  "weight": 4000000000000
	},
	"NodeID-Hfa96c7a4xF9vYHuMh3cRsKUHa6YkDsYN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Hgs8rDmSTASnKigsVfz7k3yZ3DgXP5s1j": {
	  "publicKey": null,
	  "weight": 20085563200902
	},
	"NodeID-HhEvRMAeXAuXqULRysWKNLK2XZKzYwCnK": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Hhkq3c3UzzX3UpwC2kq9TcbJdguht164A": {
	  "publicKey": null,
	  "weight": 2625000000000
	},
	"NodeID-HiFv1DpKXkAAfJ1NHWVqQoojjznibZXHP": {
	  "publicKey": null,
	  "weight": 159010846593170
	},
	"NodeID-HkE9eyj5U8saYexeGWuPcQQbiu5MyRHHa": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HkEQ2sJogxEps4tmqFRDuVG9hyfxe8vMg": {
	  "publicKey": null,
	  "weight": 3761374462601
	},
	"NodeID-HkpgaE4YB9Z7KmueQ2mz2cZFyfttvo6GP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Hn7YWxiVC9JufWW5n8u1mvaYgNgqWG8VT": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HoC4UxP4UuJyda7MZs5jUkZRk8Rfj3pzD": {
	  "publicKey": null,
	  "weight": 365791434955388
	},
	"NodeID-Hp82btBefk1ffBD7d1qtoBL3DuuNUv6qa": {
	  "publicKey": null,
	  "weight": 3714865114268
	},
	"NodeID-HpEw7munmgQtTNd8PWqp8T4GNU738Jvc2": {
	  "publicKey": null,
	  "weight": 103099123988371
	},
	"NodeID-HpMSfYT2ox1vkr1hFNMwmdWuq8QWhvH2u": {
	  "publicKey": null,
	  "weight": 996665102336382
	},
	"NodeID-Hq2aEQ2RYhEo7BkVR83c78n6wbVLPrbqJ": {
	  "publicKey": null,
	  "weight": 2004546006054
	},
	"NodeID-Hq5vdoGnZLMfasw4ykMzdM6PhUFDmp7FL": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HqPSHk6fxtJW1EsF9Y8RPH7KvPvdRpnzy": {
	  "publicKey": null,
	  "weight": 13321383000327
	},
	"NodeID-Hr78Fy8uDYiRYocRYHXp4eLCYeb8x5UuM": {
	  "publicKey": null,
	  "weight": 231836705882340
	},
	"NodeID-HrGKtuSRaRj9b4CNy2JeLfv1WAULmhwtF": {
	  "publicKey": null,
	  "weight": 6803538860733
	},
	"NodeID-HrPfqGPogaKpTTcBhNBhmYfLRHZhiMWyR": {
	  "publicKey": null,
	  "weight": 2248836482043
	},
	"NodeID-HrbPcibo4ijcFChRoBLUnskyXHtjuaLob": {
	  "publicKey": null,
	  "weight": 93015438576872
	},
	"NodeID-HsBEx3L71EHWSXaE6gvk2VsNntFEZsxqc": {
	  "publicKey": null,
	  "weight": 872525100134184
	},
	"NodeID-HsSzvb9mMLFcidMWoQYXwHrnG2gZKQtPC": {
	  "publicKey": null,
	  "weight": 8489123755189
	},
	"NodeID-HuPCfRaLN2cXoXDYyi8GYvjATLDmNJMHs": {
	  "publicKey": null,
	  "weight": 794319812880300
	},
	"NodeID-HvxZoKdC58RkD4ZVTYwjVZ3TsiQoTESEP": {
	  "publicKey": null,
	  "weight": 7169622919490
	},
	"NodeID-HwD5fu2NCooHXCPc2GMuNHrhfsThpsQYY": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-HwTU12fiGyZoyqaShTAFgTadY2w4DnazT": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Hwo6hDmWvSH3rQGrxCoujGvsDWzC9NgyG": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Hxx8vTs8kH3stQWgPw1XdHKqpbuD2nMRi": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Hyr1U8Tbw1Q9wgJYRfraYsEp9EzMugRte": {
	  "publicKey": null,
	  "weight": 2229998663100
	},
	"NodeID-HysvJteG6zyScbEu5WKiqAetFXHNrnJHA": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Hzfy8c2s9PyC7CPiUB1PmPckWYs7nj9Tm": {
	  "publicKey": null,
	  "weight": 231000000000000
	},
	"NodeID-J1CD2Nw1eX7iNUVw3iXjFJkUgrSm9qbm6": {
	  "publicKey": null,
	  "weight": 5352779431354
	},
	"NodeID-J2RTGUXRNDSx2kBU8w7Am6Y7k3LABR5iU": {
	  "publicKey": null,
	  "weight": 4314244983454
	},
	"NodeID-J4NNvkEUScCLYG4aN7eW5RYx15vSCHrui": {
	  "publicKey": null,
	  "weight": 258722996326200
	},
	"NodeID-J5aCjRo3ZSLqhcVLE7tutSS4PdAZcicGu": {
	  "publicKey": null,
	  "weight": 2733397363100
	},
	"NodeID-J7GwUKfHXhxTqgFbuQd4us6yKYdwahHHx": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-J7PV3mKZgSJRDy3atopJhh15Y1TqYsySk": {
	  "publicKey": null,
	  "weight": 6760490054578
	},
	"NodeID-J7PjfVJgpLL3mBMUwpicPqiqDFAmNWhKA": {
	  "publicKey": null,
	  "weight": 4314244983455
	},
	"NodeID-J7gmHAAEPkSPbNeurMqhg6bknPXWgJgS2": {
	  "publicKey": null,
	  "weight": 2833955420372
	},
	"NodeID-J7n7bUvog8Zdgj4bCNxRzjzrSKAirk7vi": {
	  "publicKey": null,
	  "weight": 564723909380763
	},
	"NodeID-J7vWeKgzhtKNy14NLT82Vkz5N6L9sqB38": {
	  "publicKey": null,
	  "weight": 2062434089135
	},
	"NodeID-JAkZxmCTCkBmscBpSisz5xPFDWcrBbkUC": {
	  "publicKey": null,
	  "weight": 3304154181959
	},
	"NodeID-JAqrskRRwjmmU8fRp6RAf5wDoEtjXFyXu": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JCT5WWZVaKqcMNDXuuYm3Gpgah9TpKevV": {
	  "publicKey": null,
	  "weight": 7204251613482
	},
	"NodeID-JDH5gzcB2DTfbEokaFU7EvpLN7E1RHzpH": {
	  "publicKey": null,
	  "weight": 27696890675154
	},
	"NodeID-JDS6pq6tZxZxePaDc77TvrmNnY5HBUNX4": {
	  "publicKey": null,
	  "weight": 10000000000000
	},
	"NodeID-JDYxVGLrUb8SHwuu5jZBXNC9rqYWd29PG": {
	  "publicKey": null,
	  "weight": 4039896552589
	},
	"NodeID-JE8JBF8Kb7xWQ4U59wKp3PcZMBfzvKo8h": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JE8niwmtgpDARKNUqTEdXPQY41uHtb1PN": {
	  "publicKey": null,
	  "weight": 2025000000000
	},
	"NodeID-JFMBphfETbxfarEkQBHw1RSpfoFhdhUSR": {
	  "publicKey": null,
	  "weight": 9999000000000
	},
	"NodeID-JKCgdbWtCa8BreNP3xftfUuoWGXoy8hWX": {
	  "publicKey": null,
	  "weight": 2001000000000
	},
	"NodeID-JKxZwnEvK2MAsL9SQ46fyCnTo1ZDkz22D": {
	  "publicKey": null,
	  "weight": 4644097919717
	},
	"NodeID-JPSnjdvcPjD2NQNaTgC8YP6dhcGeKhse8": {
	  "publicKey": null,
	  "weight": 2572864465819
	},
	"NodeID-JPvThCax26E5wEh4HSpjhfbfGn8LnvGx8": {
	  "publicKey": null,
	  "weight": 693099000000000
	},
	"NodeID-JR96tvzbG8L6xdoGLUWUWhQiGkCfNFP7z": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JR9oQgVHfNBybNVMm64AB2AGvHFCfaaxA": {
	  "publicKey": null,
	  "weight": 2335970093203
	},
	"NodeID-JRCLnbo62xRaNLhNeNd3HzSHK4dW5iMx4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JRNiPWVdV7Vq7MzBmAQB4jekNTfvRAedj": {
	  "publicKey": null,
	  "weight": 78834937521276
	},
	"NodeID-JReo9xxdEAX82pAt3berdoi59Pn9km8nV": {
	  "publicKey": null,
	  "weight": 2349859332967
	},
	"NodeID-JSKsSxV74dyp94CvBKqEKUcEEKBXyNbge": {
	  "publicKey": null,
	  "weight": 3902388490063
	},
	"NodeID-JUncQvtGY9EVfsfHNHrdsiv4kRfAQSsvu": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JVTZ5F1RvjDd4dA4Th7hiV3furipAaKTF": {
	  "publicKey": null,
	  "weight": 2039379230985
	},
	"NodeID-JWBfuYPeFXRkuL1AeW6p2fquqoQTVJsV2": {
	  "publicKey": null,
	  "weight": 3101099429534
	},
	"NodeID-JWRFL3w3x2xjxYZGun74vRjNofTw4Gpd5": {
	  "publicKey": null,
	  "weight": 11290000000000
	},
	"NodeID-JWwGze2HXQq7THFQj5zWRmBa5XaMj55jZ": {
	  "publicKey": null,
	  "weight": 13709102135489
	},
	"NodeID-JX6Akjfs5Qyk7m6o6AL7c82BvkPtXzZ72": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JXT4uEeTEuFcEEgRZYFusfRm8c5iH3rqj": {
	  "publicKey": null,
	  "weight": 8328346982219
	},
	"NodeID-JYoS8akGmcoeLq4kxpxfsPJxwi3UWeXu2": {
	  "publicKey": null,
	  "weight": 4682442798127
	},
	"NodeID-JZg54UPtzhHhXoVkX4r6WSTc74cAgdFjY": {
	  "publicKey": null,
	  "weight": 124423411764672
	},
	"NodeID-JZoTeECdroSJ8sDK3nwpeBz2cYCEPPjiu": {
	  "publicKey": null,
	  "weight": 5383634861179
	},
	"NodeID-JaS7Df16bwZUSL87wersMaXez7836EH2X": {
	  "publicKey": null,
	  "weight": 2125210864405
	},
	"NodeID-JbsnKEwxa9J6arKtZ52j4obF7FwvH634k": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JbyG4cewqSwjatqGtg7mDYbcvb299Cb8Z": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JcVg17s5KwKNv2yJtYfBjEFfB6GdHDmXN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JdwZvLUi1hnwuidkpJnR2ZfryBb3FbkGG": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Je9mYMihjzv2oDF3ebvzDKs96zfaXDLHF": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Jfx9FcVMQuy1FXGd53EuigXnCEgRQUUMg": {
	  "publicKey": null,
	  "weight": 69796941176436
	},
	"NodeID-JgEayrUnvk1tvMFS5Qvr6LPNSE9wNYNik": {
	  "publicKey": null,
	  "weight": 2309768150940
	},
	"NodeID-JgdYonZgwW2LRYGMwcDcHywQ4GmvjJoe2": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JgfyD94R76Wj1KBXWkGpnsSAXLM6f6dnN": {
	  "publicKey": null,
	  "weight": 4275722807656
	},
	"NodeID-JgnwfnQC9MeT7Tba3i62KGhZiCpSNNQUe": {
	  "publicKey": null,
	  "weight": 2386293199867
	},
	"NodeID-JjXsWF3RPpFJStQz2biYCMCTbQD4fAnKd": {
	  "publicKey": null,
	  "weight": 2950087181013
	},
	"NodeID-JkjTJNKyJwjEfxdfgQhvrUuf9x3ovpQBs": {
	  "publicKey": null,
	  "weight": 56284694256928
	},
	"NodeID-Jm1k2q2WAkH99w4ZBEgzrmm6Kte39qCWJ": {
	  "publicKey": null,
	  "weight": 2290000000000
	},
	"NodeID-Jmvu5u8svgTa716m6xBEcrB6wweP9tX5h": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JnREhSP5nvBcrCY1L3it6odDaHKZDryA6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Jp4dLMTHd6huttS1jZhqNnBN9ZMNmTmWC": {
	  "publicKey": null,
	  "weight": 17166712732455
	},
	"NodeID-Jpad4xARtJKgtzo199VoEaGbmYdBYbN39": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JpgUnEssGrh9NBagctXHVP2cXhYZHxgwf": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JqHkFFzz1tYoGHzPuTsueqF1yKifPiBEG": {
	  "publicKey": null,
	  "weight": 4773897272856
	},
	"NodeID-JqVwhr995DJVkh8z6u79nRknt34pmdjNY": {
	  "publicKey": null,
	  "weight": 4786860085166
	},
	"NodeID-JqXD8Tsj1nKXCBEp65ESmFhBrEhPL4WCs": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JrHvS97Avz8uqX3E1RgG6mCSaiedjWf5E": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JrVPY7PhRPecPFebosd8b5ztJnqWH8oZy": {
	  "publicKey": null,
	  "weight": 2027343730254
	},
	"NodeID-Js3ahWihoTLJWTrWzpth9mttniJNuyoki": {
	  "publicKey": null,
	  "weight": 7196520459810
	},
	"NodeID-JsNnvXMTu37NctoU7tHx8mt4f5oswJk93": {
	  "publicKey": null,
	  "weight": 86025919697803
	},
	"NodeID-JwiikR8FhmnYxxGjywAZeJMFjxfpfD3L3": {
	  "publicKey": null,
	  "weight": 4275722807655
	},
	"NodeID-JxJxzCNp5c2p447FF24jEDPdHkbB2mhJY": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-JxxUAnbxUYQXzTQ8eNeE4jWQt9fTzWd8o": {
	  "publicKey": null,
	  "weight": 7096931686103
	},
	"NodeID-Jy9zeyBRu9XvAxXfSGQ7pdZ1opTHeERgq": {
	  "publicKey": null,
	  "weight": 2025000000000
	},
	"NodeID-JyCGTjJaPELmLZdtn5J1HwKHZzki6Ae9W": {
	  "publicKey": null,
	  "weight": 5610830720228
	},
	"NodeID-Jym9FJcjAbqezV5HUbFT3eG3NraSQSLMZ": {
	  "publicKey": null,
	  "weight": 2683825453801
	},
	"NodeID-Jz2nuAFAfyvgni96rviZWsXxyYYBgSmXf": {
	  "publicKey": null,
	  "weight": 3910684627203
	},
	"NodeID-Jz9EUhGBC19bzUE7kurEPyoKTRjVMgi3v": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-K1E6w1DSaHzZky2aUM4qyQfPYMxF6Mixi": {
	  "publicKey": null,
	  "weight": 2251168370944
	},
	"NodeID-K1JfVJnvPQ9ThN4FUqU2iyAQywjx1Gf5K": {
	  "publicKey": null,
	  "weight": 20000000000000
	},
	"NodeID-K2fyNZcrWA4DBUNyFme4pNshAgSw2YSSL": {
	  "publicKey": null,
	  "weight": 1025430242920091
	},
	"NodeID-K3gyeR7GXuuf4coJXD4CxviWpRopysX4j": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-K3unRwiJFD4p6dXDHJcdeQn6CaoFiqo2": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-K3wvYyc46fLE7sSy3G81uta4wDKVkdKbZ": {
	  "publicKey": null,
	  "weight": 5086511760302
	},
	"NodeID-K484QXg69NxbC2BUsB7Bgc511t22E6MAs": {
	  "publicKey": null,
	  "weight": 3177730772336
	},
	"NodeID-K5ecmXoCnxQPFTwrDxqL63kfdh8X4Z3oJ": {
	  "publicKey": null,
	  "weight": 2073099000000
	},
	"NodeID-K6qqr2J2JcYZnSMRSqpbVv5ysZfnwMreQ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-K7ZsFPBybHPuD3YbzBpp1fe9QTj81WYSU": {
	  "publicKey": null,
	  "weight": 343479532361397
	},
	"NodeID-K8EaudGE2PS93rgD6aayqS4dPRfrseWdi": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-K8tzccU8StAHuSZfS5trxHKbvA17wixN2": {
	  "publicKey": null,
	  "weight": 10066954328798
	},
	"NodeID-K8zr3RfHUQuupNJdPDKnbhFhwcD3TTbuD": {
	  "publicKey": null,
	  "weight": 1020218206246674
	},
	"NodeID-KALbip9YkRerN81YxU32ySYk5PbGrkysz": {
	  "publicKey": null,
	  "weight": 5729843487215
	},
	"NodeID-KAjUQ36dNGmMpw35iavX7i8Cr5XGXwby7": {
	  "publicKey": null,
	  "weight": 4266955020336
	},
	"NodeID-KB5e5rq9ioGR5uicmKdmgjR5hU7iRUFTz": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KBcYcnQmw47qpb49rKtRr5tZgsxnsdkrZ": {
	  "publicKey": null,
	  "weight": 2004516364630
	},
	"NodeID-KBvpDcnJ9ztfqA4jiAoMwARqfDaS91xqx": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KCQARVq1pjYw31EPSaW9571giKoFH498X": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-KCX4L5dr83TbmBnE4UeFtzGeE2NSLskt8": {
	  "publicKey": null,
	  "weight": 272279679093174
	},
	"NodeID-KDHVpdM7PeMtoVm2mmNNXjFfs93cABSut": {
	  "publicKey": null,
	  "weight": 69517492316564
	},
	"NodeID-KFhoZ7aZMtAReBrb8KgcKG322q27akhN4": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-KGQWaR3xNXewSaQhp4zAu8ZkZv6DAdipn": {
	  "publicKey": null,
	  "weight": 6978683585143
	},
	"NodeID-KGYVGbz1pbyuJpujtJAkXhSY3AtLqBG2Z": {
	  "publicKey": null,
	  "weight": 663993857213936
	},
	"NodeID-KH8r5RoMBBpy4XHHb9DMrbMLc2a1WJKWC": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KH9WnxfcrJLod4JWXHxESehVTAqE5DoLy": {
	  "publicKey": null,
	  "weight": 1348802852171131
	},
	"NodeID-KHRscCPBynyUHXSz4nXrxmFmoMgUPZBTn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KHaBbtVbSMHmkqs5SKsdCPCgRbouNcwy1": {
	  "publicKey": null,
	  "weight": 3251590880463
	},
	"NodeID-KHqznMCanU1tA9DcNLBwJQ7zUkRHC9yDW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KJLkmhbvSYUqdMybmRScnyo7aDY6X2mYa": {
	  "publicKey": null,
	  "weight": 7849011451580
	},
	"NodeID-KJfoYJrDRXdmLaaaaRMjNbNKAwqSj3pWT": {
	  "publicKey": null,
	  "weight": 25495748832141
	},
	"NodeID-KK6ifknXca7VGbqB3d5ccUzKta1zJiVM7": {
	  "publicKey": null,
	  "weight": 168371275561525
	},
	"NodeID-KKz1wXitLQWsfKSSUPE5JL6Wxw35S2218": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KLPimX5rN51nsSg9UQQahoRvKapDLD1pc": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KLiUDfG672VocEHwbAwGxxFdCHovg1Zp5": {
	  "publicKey": null,
	  "weight": 10000000000000
	},
	"NodeID-KLsfWtM1cJLWNG21LdwPxaLk83pZo99oV": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KNLkh3KVKFFhBWujmcZ5P3p2fJc3BbdNA": {
	  "publicKey": null,
	  "weight": 2957504497625141
	},
	"NodeID-KNPCT1vLFcbVsFvLzgQqrbCS2kvEyZKBz": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KNm1ZqZXdAMJxggBGJkSs7A4tSwYLAjEb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KQfghSuT1CskYLvgrQAt6oxUQU7ask2QS": {
	  "publicKey": null,
	  "weight": 793895294117634
	},
	"NodeID-KQfmZVTmDMvuSEW2kts1rmVNmspyzUwVq": {
	  "publicKey": null,
	  "weight": 2062637088018
	},
	"NodeID-KQocfzJn3D4V4gBGPYsUB9Es7nT28Atvq": {
	  "publicKey": null,
	  "weight": 48125587640904
	},
	"NodeID-KRAcashX1MWv1pWqHrNzrU6PWV7nZcKpm": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-KRhtytNVPjiSeZMJrqqE1vdHTCoMTHCer": {
	  "publicKey": null,
	  "weight": 10316453149091
	},
	"NodeID-KUgJ4Tiwfx9di8h6oPCkTEcrBBDiDq4hg": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KVB8F5Gy8adXA2zCg3i3m3qJ23NoCDDiA": {
	  "publicKey": null,
	  "weight": 94422556778059
	},
	"NodeID-KVF3A8uB916AuexeZWS6eeeMdK2ep2vi2": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KVXQMie2X9quytg2ZdRq8rsadxfQCvNHF": {
	  "publicKey": null,
	  "weight": 8949312121499
	},
	"NodeID-KWnYciu66jjGTmiqs5U8hyeexiHEUa1eu": {
	  "publicKey": null,
	  "weight": 2047921326747
	},
	"NodeID-KXuvoHe9KgzW6KfVNHXW4CUNunzWDDC4M": {
	  "publicKey": null,
	  "weight": 662272911094205
	},
	"NodeID-KZLSjb8umvkj4nUVdARnY55NBKQcYfVV8": {
	  "publicKey": null,
	  "weight": 1021155559494017
	},
	"NodeID-KZiZi4sS44KZjmGtkeF6r1vs2MNN7p2kZ": {
	  "publicKey": null,
	  "weight": 3207269771983
	},
	"NodeID-KZnxZckzgz9eQWzB3iuvM5GC2QdJLE5yv": {
	  "publicKey": null,
	  "weight": 1023000000000000
	},
	"NodeID-KbqaQBQVzCb44wByiZX4ZChg7CtmX1pVK": {
	  "publicKey": null,
	  "weight": 2116052251385
	},
	"NodeID-KdEdrTAZpr9rdfqcpXjJxEbHC1oSWnvZP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KdxMzeB5ubHBCbk8SqgesC4BAUx7fCT6q": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Keh97V9b6k4QvYt4JewbkabSmGoQ4sA3Q": {
	  "publicKey": null,
	  "weight": 5138027041461
	},
	"NodeID-KeyQeQwf1D4bzdcHVP34xXXWhJzFjVTSj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KfPXTUa3haWHCr1z91ULmY9Pit4VwUQ9C": {
	  "publicKey": null,
	  "weight": 53190152744731
	},
	"NodeID-Kfqu9wo9FhcXG1jK9nVwbTrBV3d3kzmKR": {
	  "publicKey": null,
	  "weight": 152211752686857
	},
	"NodeID-KhvVVixdSrBD6u9q25YYNbayxBnNsw9h": {
	  "publicKey": null,
	  "weight": 2260000000000
	},
	"NodeID-Kk1UZNLJiyf57AL7vYPHntjPYLBdmTyu6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KkCG4Sktn8zxkaFRV9bi5RaJxZNu1XyMq": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KnSegLyDRx7V29WU6PgTAy5FnWwQwdz2x": {
	  "publicKey": null,
	  "weight": 2004521522520
	},
	"NodeID-KqKMF5JYhxjkMVpYWQz8zTiv1N1qm2nSf": {
	  "publicKey": null,
	  "weight": 6059253877971
	},
	"NodeID-KqmFvyCFfsaTcs2XDquCerCfWx3S6Zhj9": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KsSwTshRyF2xV75YDNXcEB5psLaBFHKfr": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KsUkk3D4hSKJGM1WBmeu1yj8Ab7h512X1": {
	  "publicKey": null,
	  "weight": 2161719826122
	},
	"NodeID-KtJJDa5PVQYChXaNd1fKVmF5LSsBB8NHP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-KtJLagoYnWfw2oBnRki63NcguiGjuA2Cr": {
	  "publicKey": null,
	  "weight": 558464461790320
	},
	"NodeID-KtKf9AuKyBvofsQE7Jn4eAzAndKHu1f4M": {
	  "publicKey": null,
	  "weight": 10934948034905
	},
	"NodeID-KtSsjRubDzNuGP6FsXphHmtjWx62KCqvf": {
	  "publicKey": null,
	  "weight": 2281198298979
	},
	"NodeID-KtyNKMikjPbRi6bx4AeDC8WLad2XKapgx": {
	  "publicKey": null,
	  "weight": 2013788128190
	},
	"NodeID-KuStLqzCMZmLv3NigAtKPt2BEE12q71wQ": {
	  "publicKey": null,
	  "weight": 6419100000000
	},
	"NodeID-KvBcQpCDXbb3w8x7ktMhsx6rCq8eSYcd3": {
	  "publicKey": null,
	  "weight": 181098548266269
	},
	"NodeID-KvN79EtxbHdT86o7w9GBQEVeRWZ5woAJx": {
	  "publicKey": null,
	  "weight": 2381412680245
	},
	"NodeID-KwWsJneoQZaVzUaUjjkKDj44BxJqoXT7D": {
	  "publicKey": null,
	  "weight": 107317066379075
	},
	"NodeID-L1RcwxtQjb72GiCABuonnV3PWFZKPTV4P": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-L1u2bQdMywnBLHVkrMiAoTSCYJTGSwE9k": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-L3JrbBLmihsqtaXAPXg7acKQbGa9tjhEc": {
	  "publicKey": null,
	  "weight": 4246894992682
	},
	"NodeID-L3SyWrZH6uorqgExkkmGuXPx9neqELqXA": {
	  "publicKey": null,
	  "weight": 89317106784246
	},
	"NodeID-L6MwdnCDs69eZBXJJZ4jw4ZHW3CyZahyK": {
	  "publicKey": null,
	  "weight": 39822801942760
	},
	"NodeID-L9HsJPQ9TF5aXjdx3KnQrh1eWaHJaNXNj": {
	  "publicKey": null,
	  "weight": 74786635614729
	},
	"NodeID-L9eaKWJQNkywXWVg7FKMZNWFno52oh3J6": {
	  "publicKey": null,
	  "weight": 2562391813261
	},
	"NodeID-LAk2MapcDFFxS7HZ9KeuwxPhdCD3YtZDm": {
	  "publicKey": null,
	  "weight": 5507651107506
	},
	"NodeID-LBA1YYYBbxwYMMAA9odGhR69d9SCbyq8v": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LCQBGoE9MtECsDjNLohwAP6L4rXDbSooL": {
	  "publicKey": null,
	  "weight": 2867234993418
	},
	"NodeID-LCVUCE4bvtPopRuYvoWwTopRScmpLHTcg": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LCvxVSwts7BQAS5mVuXH8XYGuVJZSRtCn": {
	  "publicKey": null,
	  "weight": 6771034956426
	},
	"NodeID-LDgwboNU4iVyy5MQgArg2rcDPssiAe6jT": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LF7ktpNDXo1fjHYN6MYHnitMPrHwN3pup": {
	  "publicKey": null,
	  "weight": 4207205782267
	},
	"NodeID-LFNBojkAbg1az3ys8htBJHmyQ96hYhFzP": {
	  "publicKey": null,
	  "weight": 7397082440977
	},
	"NodeID-LG2NYC4RxxzCG4zFGWoAwSo4U8mjMux1f": {
	  "publicKey": null,
	  "weight": 9800373847239
	},
	"NodeID-LGrCunBRJPWSqZLEw2QTHruuQtZai5P64": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LJBZ5j8beyjdcnVdNk5fwtThNn4LokYpV": {
	  "publicKey": null,
	  "weight": 74635918493019
	},
	"NodeID-LK64Kz2PKbHZBnhxL61AhiLXWeZdVY4jS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LKPvVELtRehiZPQo3aSWuch123JpA7H48": {
	  "publicKey": null,
	  "weight": 8389488940214
	},
	"NodeID-LKbxWKUAE7s5dtGWhm4Cr8hsosDxLsNzu": {
	  "publicKey": null,
	  "weight": 3873547892074
	},
	"NodeID-LL3tKdbQ9qCDRMm4cif2pGsxGDXsNNzBp": {
	  "publicKey": null,
	  "weight": 2349361274655587
	},
	"NodeID-LLVMxBRYjgvczhD9u9RJDSmZ48bbXdWx5": {
	  "publicKey": null,
	  "weight": 2025100000000
	},
	"NodeID-LN7tsaTsJaPDTJHjUPUX4ePk2HFBWZVh9": {
	  "publicKey": null,
	  "weight": 49958029404822
	},
	"NodeID-LP9FxV6iBCayBiupHa4QiPxtqDWNoBoVx": {
	  "publicKey": null,
	  "weight": 50000000000000
	},
	"NodeID-LPf42gVqz97N6bogZGXJNTWhgBABEe4qE": {
	  "publicKey": null,
	  "weight": 2025000000000
	},
	"NodeID-LQP4aci528q8z4KMe8Ug3Gh1aHj4yojPX": {
	  "publicKey": null,
	  "weight": 2001000000000
	},
	"NodeID-LQhwkBnuj2vjE786WcgsGneFVWcjiH6KA": {
	  "publicKey": null,
	  "weight": 2414924563057
	},
	"NodeID-LRBqzoj2mB8eD4AHE2oCTBziNdssTfxuf": {
	  "publicKey": null,
	  "weight": 142543529411724
	},
	"NodeID-LS7NEp4LEVzAsjFRaRm6RFKDDL5tWARVo": {
	  "publicKey": null,
	  "weight": 2004521517270
	},
	"NodeID-LTM3DRAzRtWoWHVtB58s5CohXKYexiqVp": {
	  "publicKey": null,
	  "weight": 5635948083493
	},
	"NodeID-LUFJGzAfqkmUaTDBVScjvdWtrruDkbFvr": {
	  "publicKey": null,
	  "weight": 9371571527572
	},
	"NodeID-LUhB7mVaTMnkYLQsqf2RV2PUerJfsq2wW": {
	  "publicKey": null,
	  "weight": 9613414480639
	},
	"NodeID-LUqMj414dv39Z5SNRXRb3bz9YhLJ2oWrS": {
	  "publicKey": null,
	  "weight": 2779348317544
	},
	"NodeID-LVqsakFrQLyqNiChxcYqsBzKaKFRMiBsN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LXCSfcmd7p7LroJzJtYoSAhbsawe47CNk": {
	  "publicKey": null,
	  "weight": 131298524675635
	},
	"NodeID-LXVb4LZEU53c7XCk2v4bU2qacG7ux5dUB": {
	  "publicKey": null,
	  "weight": 22276096526362
	},
	"NodeID-LXpULpbU1A4AobEzCSBy6wYLEbogwsMK1": {
	  "publicKey": null,
	  "weight": 2025000000000
	},
	"NodeID-LZMN1BDS3u6pe7T5Umjuhg1Z1MahXkK9C": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Lai2VTTYk897ae9uq6cGk9FbhKD1KHvFS": {
	  "publicKey": null,
	  "weight": 652497500000000
	},
	"NodeID-Lb2X2hJpi6krux6aRpYSG7xcNvCXNg6M7": {
	  "publicKey": null,
	  "weight": 7828547306358
	},
	"NodeID-LbAkNKxyv3kCs1C2ngkb91STbKpYHKs1q": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LcBGn2kscipZRgqhKqDCA5dq4WmHuuvai": {
	  "publicKey": null,
	  "weight": 9292306158558
	},
	"NodeID-LcNseW1dfEL8FnXQB5HidFpzNu452vuGc": {
	  "publicKey": null,
	  "weight": 2001000000000
	},
	"NodeID-Le5rVEBBPdgYEcEJpvYeKo36hRs75WE5S": {
	  "publicKey": null,
	  "weight": 2565222925902045
	},
	"NodeID-LeZUTECdpTLux94oAXKruxYgz5ppx7tfr": {
	  "publicKey": null,
	  "weight": 4444492537844
	},
	"NodeID-LeyuRqXk3UxoDMZWjgqSY8gR9wDVg4ht6": {
	  "publicKey": null,
	  "weight": 8786785131845
	},
	"NodeID-Lg6HaVBjndKUjyXiq8ZjHookjqzkDxH6Y": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Lga7mB6EfqgSJv4VTPaUwDjEtq6utiNQH": {
	  "publicKey": null,
	  "weight": 35516162049372
	},
	"NodeID-Lgjda988AdujVYMYdwuEM4Tc8aopF6DRk": {
	  "publicKey": null,
	  "weight": 1900029898514317
	},
	"NodeID-LhjS7bvahdBn2BomUUBHHFNbnzLRjri6D": {
	  "publicKey": null,
	  "weight": 2173163110118
	},
	"NodeID-LhkzuEyUkHR9SW4cMxB55QQRbbGF83zzc": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LjxLvskdhxQ5jAhBKSzKrxU9NQD1DPuKR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LkDLSLrAW1E7Sga1zng17L1AqrtkyWTGg": {
	  "publicKey": null,
	  "weight": 6952080282764
	},
	"NodeID-LnR5bqTLkmeGjsDKiPdJNSbRPr4pjAURR": {
	  "publicKey": null,
	  "weight": 4034455660888
	},
	"NodeID-LqLgyDEyMbd1uXfRa6kf8EPjnRhpF4KCF": {
	  "publicKey": null,
	  "weight": 5139965009259
	},
	"NodeID-LqUHsZZeoCvhiwunLtAJMSXx1GxYpozzq": {
	  "publicKey": null,
	  "weight": 8265404099292
	},
	"NodeID-LuAPt3GG24RsuzmjPL5neGBwca7r1h31E": {
	  "publicKey": null,
	  "weight": 120980000000000
	},
	"NodeID-LvKMfPzfWT1VfAPLef6denZ8hTwSAFMGY": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-LvWcDhi9d75T8dwerYcQRoVWD6yDtywFN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LvioWzV19k7FxhFZC4Z2Xq4qxnswkL46C": {
	  "publicKey": null,
	  "weight": 2013671685366
	},
	"NodeID-LvpYyT4UY4UAgP4quynTFzffYmjoxqZtR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-LwiNyRJeqRfAGS5phAQvvKhX9xj8LbQ86": {
	  "publicKey": null,
	  "weight": 2727811338575
	},
	"NodeID-LxrJtyQXpdweZH7LR3ujcEQZJ9VDTW3tC": {
	  "publicKey": null,
	  "weight": 739084558598472
	},
	"NodeID-LydNf31SmNLztbKL7iNsHXPhnKxz7g3Cw": {
	  "publicKey": null,
	  "weight": 898446352941164
	},
	"NodeID-M17yakji8RzPBFJNPXnksjoeabg7BXUJH": {
	  "publicKey": null,
	  "weight": 9201476759474
	},
	"NodeID-M1SrBrn1szYR6jQJkyrJzd6YXsbAeWSnj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-M1Zsz7o9AiDgUqA6KFE8hrCkvybN6EqcP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-M33qUGE35HbVRppTTMY1qyyfK5T52id2a": {
	  "publicKey": null,
	  "weight": 2004522802906
	},
	"NodeID-M3VrNcHHd3LGpwm6U6Gw1vfuJHoJzfG7g": {
	  "publicKey": null,
	  "weight": 3475651819624
	},
	"NodeID-M4AhBqhjwLxGixtjxaLvenUpTe8mxNfhg": {
	  "publicKey": null,
	  "weight": 2896463611991
	},
	"NodeID-M4ojoWJoEmsVD3vf7KxZ2BnfR5Ue2wr41": {
	  "publicKey": null,
	  "weight": 63235778255966
	},
	"NodeID-M4psvzu2dS7hQ8gHgCUzXfipVKuFEsoFU": {
	  "publicKey": null,
	  "weight": 218180230055197
	},
	"NodeID-M51d2sDyE92NjgRaiTnSjSLaG3U9eJ3BD": {
	  "publicKey": null,
	  "weight": 5429523737163
	},
	"NodeID-M5wUvCymq18DUFjXk9s2UCZmyWzgo9ZtJ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-M64wor3rGStFTBWuzvkGPYYsWUu2dX7S2": {
	  "publicKey": null,
	  "weight": 2095165045057
	},
	"NodeID-M77VJM2GfeAzueJ6DjumoM5enZma86Zop": {
	  "publicKey": null,
	  "weight": 37685646785422
	},
	"NodeID-M96ypSznK7ECF4EGd5KfxcxCiboyAZnAQ": {
	  "publicKey": null,
	  "weight": 12354298707836
	},
	"NodeID-M9Q4Fc2StCtFCofjppsy6DC7JDGmtRfUA": {
	  "publicKey": null,
	  "weight": 4097373258502
	},
	"NodeID-MBuKQwYxNAttaiugN4eMt1aV7vhLUo7Qh": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MBucJmfSNuMy7CiXWb9Zzd1Gf1yCoKwGy": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MEnWKkj5iUC3Z2rNhKuDTm5EVDu2ZUZZo": {
	  "publicKey": null,
	  "weight": 9391142075987
	},
	"NodeID-MFp5wwuKzSfCNESpdD3avbZBkZ3KMkEeK": {
	  "publicKey": null,
	  "weight": 2438999000000
	},
	"NodeID-MFxpZAcCZZagmDn9LtJoidGktKPpoZoBt": {
	  "publicKey": null,
	  "weight": 7413680055581
	},
	"NodeID-MGCmqCMxtsm3WkLzq9gB7uLLzgt7tyanR": {
	  "publicKey": null,
	  "weight": 2349635440062
	},
	"NodeID-MGJcpxBEXDn2Yk94uTXQwxgmYtuX6Ky99": {
	  "publicKey": null,
	  "weight": 9400000000000
	},
	"NodeID-MHBrAomZpJhaYjqWKWXSS3YusVo7afBn6": {
	  "publicKey": null,
	  "weight": 9989888272641
	},
	"NodeID-MJvk5Su2t8dFhcGe6z58Fx8A2QAAZTMGd": {
	  "publicKey": null,
	  "weight": 2133153663100
	},
	"NodeID-MK6HYMwXj3k3HZfzjdegdnd9p9LQ1dk8x": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MLDFMQwHyrvXrYtJvhuXDkHAS3kRxydrU": {
	  "publicKey": null,
	  "weight": 58677000000000
	},
	"NodeID-MLRgfkWbezKNtcH2ghSyE5J3C4Rcxg5Sy": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MLkE5eQrTGa4BoQSR9QLEZKQB9yPJEnuN": {
	  "publicKey": null,
	  "weight": 40833273817006
	},
	"NodeID-MMemzZih83vtU6TRK5oiJt6bJpNSWm1uE": {
	  "publicKey": null,
	  "weight": 4878800229925
	},
	"NodeID-MN3pkk9h9anNMuCpUuXPVxGKw1rUw177u": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MNm7zKhCHwqQ3a3oqqEqDSXM9uVFMcEA7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MP8NsoqJxPMUfCbw1izND6n4i5pTJJkhS": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MPLzWkWAbT3937j7U1uWfxdTMg3iw3BqY": {
	  "publicKey": null,
	  "weight": 4853158576127
	},
	"NodeID-MR5cJw9vYb2cA6TExPE4ZXioGV11zErUq": {
	  "publicKey": null,
	  "weight": 3676355599908
	},
	"NodeID-MSbDAksgjGPoKE4azGEgvrDj1EK2GG8T7": {
	  "publicKey": null,
	  "weight": 7386574963702
	},
	"NodeID-MSrtdRQ9bEWLhi5d8YLQtnB9E2biVkkrJ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MTEbgxTNBdL6GWh8SdQMfMXvEU1jC9aws": {
	  "publicKey": null,
	  "weight": 18420000000000
	},
	"NodeID-MTQbFuBcjcQ7TuVQjvoXC7hiWAWPkouzA": {
	  "publicKey": null,
	  "weight": 4050371485512
	},
	"NodeID-MTmtdGgEg6gSEVhG5ShyfReR8ubtcZhbb": {
	  "publicKey": null,
	  "weight": 106710893842765
	},
	"NodeID-MV2a235PcwYscp7qq8N9M4pYQBqMnfZ6T": {
	  "publicKey": null,
	  "weight": 3085559556240
	},
	"NodeID-MVRcFZmxaMJWedF4oVqsxT2YYkX3kSuGG": {
	  "publicKey": null,
	  "weight": 916103534354928
	},
	"NodeID-MVUMPzSBAtrHwTD5SZSNPRVTLkYnDPDz4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MVayMsZd2VrMvEwFMDreTrauJHGBZu3Xa": {
	  "publicKey": null,
	  "weight": 2076700000000
	},
	"NodeID-MWqZRLNFJzDPKnEuUHJSpEg56ejrzY56V": {
	  "publicKey": null,
	  "weight": 6921432823217
	},
	"NodeID-MWwstjwXDrzkd3Zqnfaoa9vWtCpb32Fhd": {
	  "publicKey": null,
	  "weight": 2004877182504
	},
	"NodeID-MXxYFnzdP24DpJmdtUKqzMEN7nh69aYqD": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MXzyYAXy4JPVDjmXapcJXW6HUq9Yxmnr1": {
	  "publicKey": null,
	  "weight": 78465013799124
	},
	"NodeID-MZV1PpqVuFPa56t7TmfA8VVdXxDSHnQW7": {
	  "publicKey": null,
	  "weight": 5913500000000
	},
	"NodeID-MZxjrRvxrjvAdaz6pDS82v65xLwckmiti": {
	  "publicKey": null,
	  "weight": 2121771358169
	},
	"NodeID-Ma3Ztm2A48bCVjfiSmoZS5MKNzKVTRN7j": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MdQZNCCZN5gTEMVaJkSekNYCZkvp5xp4L": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Mdpo4kYtg8eqn61AhjTrZ51exnHLJmcpA": {
	  "publicKey": null,
	  "weight": 865360227550601
	},
	"NodeID-Mf3rZ5w3aEtM1bzRBwFqhZYZMwks6FYF5": {
	  "publicKey": null,
	  "weight": 14300000000000
	},
	"NodeID-MfVDQRZ6jSUmke5aofBuNXruEQunxZ5F7": {
	  "publicKey": null,
	  "weight": 3817119327425
	},
	"NodeID-MhTrFwGJdGnHHf3PTm2wz7GDBNssJHqF9": {
	  "publicKey": null,
	  "weight": 2004525374168
	},
	"NodeID-Mihpu9zfzMJuEZWirJhyVVHJv7fd2bWyj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MikkG849wcdpBdeUJ2HTV2gTEDfeXWseg": {
	  "publicKey": null,
	  "weight": 838810159104260
	},
	"NodeID-MkBd2PF2EQQxdbgbLvTX2FWY1XbCm7feu": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Mm8qGQPb5tHF96ENH1MuYsYwhA5Nj2DHj": {
	  "publicKey": null,
	  "weight": 2602844050558
	},
	"NodeID-MmTfUacXsPdKryR3WkVteMqoFvyxEdYf4": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Mn2isVXFR85bfjMHPM1DT8EY8wig7YhTW": {
	  "publicKey": null,
	  "weight": 6613202857760
	},
	"NodeID-MnQoJXMCgPk3zxKQ7kkYdmoZwj2dcDcmb": {
	  "publicKey": null,
	  "weight": 18195330855526
	},
	"NodeID-MogFiRpYxhVVcgNvb7DS6rmdTHd5bPzj2": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Moz5Jcjj2JoDgay9ScwcRDQReSqX8sjyz": {
	  "publicKey": null,
	  "weight": 202603856768325
	},
	"NodeID-Mq7hjS5ySKApHMKCGw16YFo1tW7xA5DXm": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MrGm83LnY7pRQCQz9mG5TTpRWSBgNhDr": {
	  "publicKey": null,
	  "weight": 12737942948066
	},
	"NodeID-Mrh4UKAifNZtd8RJXqokDzajQnAsZuw8Q": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Mss1z2VZYsNSa4LyqD99vXeNd8h3hUMDA": {
	  "publicKey": null,
	  "weight": 11886535646370
	},
	"NodeID-Mu35k31HAXFbEp1SqJr1uUrue9nwqVFe6": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-MuEFvXvbZ964sk1rdQKWNtUc5PJAGkWeL": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Mv3qUS893ju2jtQE7VBG4GyiZr5CtCuoF": {
	  "publicKey": null,
	  "weight": 2050155748538
	},
	"NodeID-MvfTRb81PAx9AHYxbz5xACaoG84PM7YQe": {
	  "publicKey": null,
	  "weight": 10000000000000
	},
	"NodeID-MyaTMxwSKbtrFRRypqr1nWKJ3EK6di3dF": {
	  "publicKey": null,
	  "weight": 4150656229675
	},
	"NodeID-MydZ4Ju3KA959gXeERXvqxN88EtUrAaGc": {
	  "publicKey": null,
	  "weight": 10000000000000
	},
	"NodeID-MyeSQZnEZDzTfFFQvxZ5qqfqVDwg15NPB": {
	  "publicKey": null,
	  "weight": 363000000000000
	},
	"NodeID-N1aFXom4C5ALmH33bXHuRpPBo7Nwqeu8t": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-N21EWyi4Y5D8MNapwZ5tXs9D5A3rrfPi1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-N2nTLuaVwKk8EF3NBsSygCVmpGHwNz2Nz": {
	  "publicKey": null,
	  "weight": 3218904783296
	},
	"NodeID-N2t1CAS75972obgtRPHwVanMnRF1rRo1B": {
	  "publicKey": null,
	  "weight": 785901210812677
	},
	"NodeID-N35BjCN6ARikprxERf4WiMeUnh2ZsfwJJ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-N3WnyXmV89jbtQgmjqpoxBxpi5Rah86az": {
	  "publicKey": null,
	  "weight": 14966279040004
	},
	"NodeID-N3e9W3EngjabGnTZVqyZwunVcbCdrY5Qy": {
	  "publicKey": null,
	  "weight": 7522568575840
	},
	"NodeID-N46bdSibDXA2hx5pFFTuVG117Khi8j8db": {
	  "publicKey": null,
	  "weight": 102419397224156
	},
	"NodeID-N4eXezpVzti1ssWvyQKqkuGB15dJNJsp5": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-N4r8sLfeDVa2c9eDFpE4wuxJccUku3LKu": {
	  "publicKey": null,
	  "weight": 2001000000000
	},
	"NodeID-N4vaumNTt2kixxUTtWEBgy4vwsxnuahc2": {
	  "publicKey": null,
	  "weight": 2194408025543253
	},
	"NodeID-N5FWBAn5j1UVxb3nH45N9DyXEAaeZiU4q": {
	  "publicKey": null,
	  "weight": 2449414570689
	},
	"NodeID-N6z9WSBMoSxLSRtdPDz8jZvasAskYcCrd": {
	  "publicKey": null,
	  "weight": 2415356015239
	},
	"NodeID-N7FjoUy3LrDi1kSYeaBwZjja6H2Y2NLsN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-N7WPJ9DXvEx8xLJdtDswNh1JdNmCKejSq": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-N7iZFQdtUfPRcWGs8DB3JQRZt3Pq33eoZ": {
	  "publicKey": null,
	  "weight": 2056000000000
	},
	"NodeID-N8tijV4qwWMep4STmLzWJEkqbg6bSdCmn": {
	  "publicKey": null,
	  "weight": 4022378948823
	},
	"NodeID-N95GoeyiprvxNj6E3G8QvYg3gpQ7x8NKE": {
	  "publicKey": null,
	  "weight": 72382692342106
	},
	"NodeID-N9S6JJT2VKRZ48vA5vtWgboDnfzXXycQn": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-N9c2YTRzGpRNp2J6qyPEEiAvyrViL4cGH": {
	  "publicKey": null,
	  "weight": 2094007766081
	},
	"NodeID-NBjvikZcadcyCJ2hEJsUn9EcCjys5hav9": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NCVrEpmYjJBHqj9vx9r5usFiMd1PooCfZ": {
	  "publicKey": null,
	  "weight": 9117088810612
	},
	"NodeID-NDMPc8h9L6keJw8F82NBobc874zZ6bDa6": {
	  "publicKey": null,
	  "weight": 2517868244218
	},
	"NodeID-NG15fF5qJuD737tg4u3PUw2hyZ8RZYNNU": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NGspMWkA7FgT6DnDghg4B95pQ468j7DK8": {
	  "publicKey": null,
	  "weight": 6141428130217
	},
	"NodeID-NGxhFFwufwUSXGHjxdCUZW1hNpodfTrmj": {
	  "publicKey": null,
	  "weight": 6778545472037
	},
	"NodeID-NHBv1LqcCdrcwatzQ9pGbGXPZQTi462Ur": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NHEESrJJNnyCPn5c5Uyp9EZYHCCkZwzuW": {
	  "publicKey": null,
	  "weight": 3751523270971
	},
	"NodeID-NJZWC8HxGtLkxjQk5F9KSG44WC3gCm8a9": {
	  "publicKey": null,
	  "weight": 55929063373472
	},
	"NodeID-NLBgQfsmCGFzXDehvwvmzjyfmkWQndx9L": {
	  "publicKey": null,
	  "weight": 2794247140318
	},
	"NodeID-NM6RqsshepCVojpp7R8747TGBLRvYpLLB": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NMfBW4PDVkVZhJzJa6SiCDZ4M7wAkXbAc": {
	  "publicKey": null,
	  "weight": 19986303379447
	},
	"NodeID-NN5j35c8DsuMydM5icoxhKePqc7AjwH1K": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NN6iut3SAWZqM56kbg8EhtpTU2NkMJgWV": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NNWrYMQgNFRNbvwAUA8grXiAVVE9VEyV6": {
	  "publicKey": null,
	  "weight": 8523818998938
	},
	"NodeID-NNrbnhLBjgzBvhi3JXK48ckVFcg6FXVnn": {
	  "publicKey": null,
	  "weight": 2098011718464
	},
	"NodeID-NPexpo2AXQ7RZC3wCKugj9q4y599pSfrb": {
	  "publicKey": null,
	  "weight": 2165330215756
	},
	"NodeID-NQWgSd85RNm2um4AywRNozgczjAwpmzbp": {
	  "publicKey": null,
	  "weight": 2905631258198
	},
	"NodeID-NSAi8z5XtoagdADeMNAPuWxfZpzUxacYB": {
	  "publicKey": null,
	  "weight": 2053000000000
	},
	"NodeID-NSHmxWXYXAvHLdHqNyKZjP8MgSSu1dNnW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NSNWfQTGX38pRWaky2MGe4ShCSkuLnqip": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NST51yBNtgVkXWuofyacbR1MbqYJEYdor": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NSUM9tzmeZ1DujarhTwjXAcnCeAH273v3": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NSa8Z4kZo8KnHzKHsVUkV9gW83NoNc6uR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NSngpSyxjXz3ErfjsFoc6CpfVPLxApJen": {
	  "publicKey": null,
	  "weight": 22492962300086
	},
	"NodeID-NTTNYBYdkf3n5U4zrZknt1nst9rHzpuCj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NTvJb2gZdvYU2kaMUrMbtu79KUMgyJNPW": {
	  "publicKey": null,
	  "weight": 7525708627518
	},
	"NodeID-NTvmKq2wmTvqoccqpxW7fQd1iGCiGhAsj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NUgiF4juieHZmcR6ekc9EDaRyG8qHUyjC": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NUur4HBKXj8GS5TX11jZhtPHLH6zMa6L5": {
	  "publicKey": null,
	  "weight": 2038675226757
	},
	"NodeID-NVSSRYMjjYyMkUpt7AHXWnKNTTJCd8UbS": {
	  "publicKey": null,
	  "weight": 2372984219612
	},
	"NodeID-NVkHd9WCbVAr79r3gtA7AmKBkVkNL9EXb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NWKQQXERMEAN4YZZsJ2KdsgfkRAHKmkRA": {
	  "publicKey": null,
	  "weight": 3003701000917
	},
	"NodeID-NWwQRYczsB8gf2tiiLLrk1foSgNUyrFLe": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NYjgvA2mJy4kTCQeEHnSuycnu5Kgo9gEq": {
	  "publicKey": null,
	  "weight": 148928473903518
	},
	"NodeID-NZEnXQpa55wAhwpcUZFgSAREXW1FLdkzQ": {
	  "publicKey": null,
	  "weight": 272492668255577
	},
	"NodeID-NZGsXrgoEuExZdJn8WxmyrLt4Cxb7LFyD": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NZYwtzqeQFeNZHjYgdukRvsZb4xQ43Fdj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NbYMJFDrhNhubNus3qsw3kJBy6HNEV9UU": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NcSLm9skkUrfA9mKGjdtmT163iDjqAw6H": {
	  "publicKey": null,
	  "weight": 4296451975943
	},
	"NodeID-NcZtrWEjPY7XDT5PHgZbwXLCW3LGBjxui": {
	  "publicKey": null,
	  "weight": 3329224415437
	},
	"NodeID-NfCjnjUUh22jvDYKUSWgX3J47VgxCEyWS": {
	  "publicKey": null,
	  "weight": 2004521518494
	},
	"NodeID-NfJFVKNj7kCUUWCRYFJULe1AR82f8a5ox": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Nfz4ZGQ8ZPvNQ3yHoRNMT7aTCiFYjMnSA": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NgkksDCWH73BnfgN2ALUBykc2kftsquff": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NhL9H8VSfhrktdjmQQ3P47wYXzRfQwBiy": {
	  "publicKey": null,
	  "weight": 11166644447089
	},
	"NodeID-NiW2qTm2fSKHGuMqeo1twtM8j7LoAGzHR": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NkXpMv916WuB4Np5P31BZmByna2poFAby": {
	  "publicKey": null,
	  "weight": 2009116016342
	},
	"NodeID-NkYLNRp4S6exbWamVvMzUUpXvHeVEzLR6": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-NoNP7kXuQMi5ZCMDsMuWETj7KhFfXhvo1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Nom39tjDi1Vf51VxuaZvJiDeQHwbyg6MU": {
	  "publicKey": null,
	  "weight": 81746158818531
	},
	"NodeID-NovHTxGVasNU2y33YPwmQ39DbYDkXiEfe": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Nr584bLpGgbCUbZFSBaBz3Xum5wpca9Ym": {
	  "publicKey": null,
	  "weight": 385917411764654
	},
	"NodeID-NrVSfY4zCbPjYTe4XPjmQB7qLJhYW3V6h": {
	  "publicKey": null,
	  "weight": 122743031235697
	},
	"NodeID-Nrs246dXUZmALw7a2HRTANx6U1LeawsKq": {
	  "publicKey": null,
	  "weight": 4630199971973
	},
	"NodeID-Ns1eDN3K9HTnavgiBQhtpzRxUCxJ55Xhc": {
	  "publicKey": null,
	  "weight": 1350000000000000
	},
	"NodeID-NsVdtyjprgWXLsvmuWvYe749tx3tdbY55": {
	  "publicKey": null,
	  "weight": 12813430837928
	},
	"NodeID-Nsdn6eemD7KUpKEaGYjYFuJjxC36c3qXC": {
	  "publicKey": null,
	  "weight": 2170512944632
	},
	"NodeID-NtDMBzQGPnp7suCbyCMcdE2CFU4JXcvwX": {
	  "publicKey": null,
	  "weight": 5160426838381
	},
	"NodeID-NtX4Q64xT4Yad8wdhQ4erSG6mi3DPUwzX": {
	  "publicKey": null,
	  "weight": 718724635651507
	},
	"NodeID-NtgR3kDx7Mm2szCiDTs3rtWaqUDWRMDyY": {
	  "publicKey": null,
	  "weight": 260380210429323
	},
	"NodeID-NuJhcXgmQMbFJwm9yTTmdvGN86qygnXwP": {
	  "publicKey": null,
	  "weight": 2009129960650
	},
	"NodeID-NudTZgq7DKGQsY1Gkt1v3ZG7pgaWbbBuj": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NvsbNmirter6SLhta1kkvoEjcgvaHxEt9": {
	  "publicKey": null,
	  "weight": 8505982192341
	},
	"NodeID-NwZ4wmDE8DrbaDY7V2nZscmLBKdRUtUyN": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Nwr26Q5CPcRfSHQKgQ5y3ayVPgjziWkkT": {
	  "publicKey": null,
	  "weight": 250000000000000
	},
	"NodeID-NxgyyDw6YGHprTJjZv7qh9X1yoj3yZN7u": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-NxzZe138M8B3MQh7N9u4qetqMS5ViAgqV": {
	  "publicKey": null,
	  "weight": 2469005334175
	},
	"NodeID-Nz3N9ywEYBpRPZU9yjrQtG2iNdnUqcLX7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Nz93c8UB78eEVVtfxpcecxHmiy2gZ4iAi": {
	  "publicKey": null,
	  "weight": 336953029475004
	},
	"NodeID-NzQnRDwTYZJPwYMcDkyNGCnhH7S1HQZgy": {
	  "publicKey": null,
	  "weight": 2132981681778
	},
	"NodeID-Nza7rHanhSFgJ6m9D8fFWgrx7fqbxp7q9": {
	  "publicKey": null,
	  "weight": 14995700677935
	},
	"NodeID-NzaVh4SaZh7zaT5uFmHyAT99zRbQhzvvd": {
	  "publicKey": null,
	  "weight": 2004522595316
	},
	"NodeID-P2EM7MDMhKVo4LLgSqtxYSyc39wfPe2TL": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-P2Jp7sd5mv4ZGBqThYfjHkz1EyJUXoMcE": {
	  "publicKey": null,
	  "weight": 3109302799075
	},
	"NodeID-P2xEUig1YeBTwrUF3Xj52pxaNH6gSHjsb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-P3EhUmKXJ5F2BiZw1TwNftAC1y4XyGufz": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-P3budG9yRu9ph2UedCGdf3KaLQKQLehy2": {
	  "publicKey": null,
	  "weight": 2080433261329
	},
	"NodeID-P4tSY1ZLb4q2xTdYGXCkgwXWeQN7mUUWq": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-P5hTsAQgfSLJumMDVDK2773U4gjyPtYnr": {
	  "publicKey": null,
	  "weight": 3866900265077
	},
	"NodeID-P5q5numCvCNo6GmD8vhDqq1tU25Yok4Ky": {
	  "publicKey": null,
	  "weight": 10000000000000
	},
	"NodeID-P5zxQ1Y22ztxzjunvJU1y8BiSjdP4RUwY": {
	  "publicKey": null,
	  "weight": 794812417181165
	},
	"NodeID-P6M41urqadxF7NBcxEHNFxKkGPkHjC95v": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-P9FGdfCp7sHRcUwHERp2hzo98ifUaCZ5x": {
	  "publicKey": null,
	  "weight": 4150656229676
	},
	"NodeID-P9kf9VSRrgD8B1MQKMsjravHvMHVxvoaZ": {
	  "publicKey": null,
	  "weight": 2319000000000
	},
	"NodeID-PD54wi24ENhbBnKKhoD68MGXCDc3A3iA8": {
	  "publicKey": null,
	  "weight": 2925742005842997
	},
	"NodeID-PDZYy4rfhPZWskL32EMt45zuLD2rWjDLs": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PERaNX2PDtAr8zkXgbZvwXHfjq4VKtbew": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PF9fY3h6ZMMqPmTZZUwc7wyLnZ1rSoz5u": {
	  "publicKey": null,
	  "weight": 4150656229676
	},
	"NodeID-PFgqACUqhkFJ5hpEK8bMhwRGqrBaunmvH": {
	  "publicKey": null,
	  "weight": 6066122524697
	},
	"NodeID-PGonFPDqHgyAWCFaoLxQvAQEz6uJYzZpW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PH5DLzGXEhe6RMHgod3CH9T3s3WVfQm8Z": {
	  "publicKey": null,
	  "weight": 4607032663100
	},
	"NodeID-PJA7yH2ZVtccSj4RrB94a8MtHboS9xozZ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PKiUH3xw8SmVNVJsAqFV1mXwDXToFVATi": {
	  "publicKey": null,
	  "weight": 97970760709939
	},
	"NodeID-PL7xq98yQ9kTYyGDvVQg3efowKeHug3HN": {
	  "publicKey": null,
	  "weight": 50962478955319
	},
	"NodeID-PLs2JZXeTtfZvYx4DJqkA3sKXz8ynvKLM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PMGyPdRXPKRPRiWXAqR3XqbAHke2gNDUe": {
	  "publicKey": null,
	  "weight": 2004524111400
	},
	"NodeID-PP76vCKnSoCicpRZXRfjiuV2HWJhdcYjD": {
	  "publicKey": null,
	  "weight": 2004521521620
	},
	"NodeID-PPH8k8RGmrTuc6Dun92atkbDkgs3sj6ws": {
	  "publicKey": null,
	  "weight": 2441543867648
	},
	"NodeID-PPYetyVA4dk3bYdNtvTsy5Cqg4L31EUze": {
	  "publicKey": null,
	  "weight": 6142806630311
	},
	"NodeID-PPoo9VFqdSWsrZZHrnrBBfi3uE2cw9JgW": {
	  "publicKey": null,
	  "weight": 2084775898757
	},
	"NodeID-PRmhTtYtzQyLMcJqak6ibefP1bkbU39mS": {
	  "publicKey": null,
	  "weight": 2020000000000
	},
	"NodeID-PRxbx3LZSnwvLqFa5audJBsSqAJe9rmsP": {
	  "publicKey": null,
	  "weight": 6454465550401
	},
	"NodeID-PT45awwLjTTiEqvFufS4nCkVN7wtEsDmf": {
	  "publicKey": null,
	  "weight": 3063375824168
	},
	"NodeID-PUw5p37RZKVjpK8kxBa12TBtZZdhykGUP": {
	  "publicKey": null,
	  "weight": 15000000000000
	},
	"NodeID-PVm9UudEC7N6HJ3jus5KJ1Rufck88g1Dt": {
	  "publicKey": null,
	  "weight": 2006906359280616
	},
	"NodeID-PXZcW18N1BcYiDy7iWgHLGKwiQMQuBSyg": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PY3QRSGTdYzmPTQEwrH2AmYUMBnrcge9Y": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PYFkkThWT79Sc7QqHijAc1yZjSqfXu6Eb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PYddUczdgvPidT1AS8RKPSLUyHBPuwmMz": {
	  "publicKey": null,
	  "weight": 2037325313491
	},
	"NodeID-PYecKsWbPTayLsHDdopKjB1Mw5NP3kDUJ": {
	  "publicKey": null,
	  "weight": 1751561608427146
	},
	"NodeID-PYmftFnrKf6y9k6kXXuB5GWwjLtkP7sLB": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PYwMEam56vM4Nix3AG1vufHZbMvdv8zwy": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PZ66baswBL8xsSHaFVG2Yx7qH8MkBfXSK": {
	  "publicKey": null,
	  "weight": 265349754297454
	},
	"NodeID-PZvapthv5iRB7txuVNTrKCo2fWovBExLh": {
	  "publicKey": null,
	  "weight": 5000443226200
	},
	"NodeID-PaSvEqWq5ZESHSuNGTZeQL6sybzsQzoZ5": {
	  "publicKey": null,
	  "weight": 33599396008988
	},
	"NodeID-PaYLWdJGvpxC63eK3RWXLLS8w37442aV8": {
	  "publicKey": null,
	  "weight": 100000000000000
	},
	"NodeID-PcDGg5a4P3XK1E7PXS5YedYL7vi8vSXvB": {
	  "publicKey": null,
	  "weight": 2113799904976
	},
	"NodeID-PcT74x6X77AhoTQrdKDrGTyX1dVZLoGaq": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PczCXQXPxNgeKy3P4jtAt1ycvrjvghVw1": {
	  "publicKey": null,
	  "weight": 9643522355997
	},
	"NodeID-Pd7Hexv94PQNVnEbYXWUttEAAybG34Yiy": {
	  "publicKey": null,
	  "weight": 9320155943608
	},
	"NodeID-PdT66xrYRTStCRewvEjiHRo8Vg6hyFXaa": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Pe1pUMoFAts11s2kA3F7EPmzLyZwj6Msv": {
	  "publicKey": null,
	  "weight": 2001000000000
	},
	"NodeID-Pe5xmD9DhTrSCDykMqFdG8CazG3QyJxos": {
	  "publicKey": null,
	  "weight": 5199194247015
	},
	"NodeID-PfY5X2GWsSaj8y6usEji4HZmokHHdWsqi": {
	  "publicKey": null,
	  "weight": 361106746905775
	},
	"NodeID-Pfj9qXo6LGR3goGzcLSJsCZwYNFFndV75": {
	  "publicKey": null,
	  "weight": 84184957767796
	},
	"NodeID-PhB7iV3nDxiotJnAMetF6puzg6fxTqZCP": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PhDmCZH4R8VFjVzQLRBquKMVss5sKukw3": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Phds7RN13xWBmpe9MTPhJde3hajVpUB38": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-PiaUvMvdFCWXwhRcrb2K1iPn4h7NugKjk": {
	  "publicKey": null,
	  "weight": 2009083732200
	},
	"NodeID-PmH5X1xavW8PjgvkW1u4K9qTHAbMx36hq": {
	  "publicKey": null,
	  "weight": 4527354224153
	},
	"NodeID-PmJQ7UjZo5db9XL494EbbEufAbYxqWswU": {
	  "publicKey": null,
	  "weight": 14785197046282
	},
	"NodeID-PnuVZDXvdSNRbC92ibq5w1FtPvbTFvwmM": {
	  "publicKey": null,
	  "weight": 2382243427567
	},
	"NodeID-PpYToPR9XKswScUX1ypXKjDUP7G7utRDz": {
	  "publicKey": null,
	  "weight": 2518235603971
	},
	"NodeID-PqL7ZLMh4fFQNxKrKt9uWuZKa3UGsdcdQ": {
	  "publicKey": null,
	  "weight": 3345513230457
	},
	"NodeID-PqPwxtYAt6AmikigAwkRzTQCaeMqPvpff": {
	  "publicKey": null,
	  "weight": 5344011399917
	},
	"NodeID-PrPRAipDzRmPYPSDhm1dKnomZNR1YjLwu": {
	  "publicKey": null,
	  "weight": 2394033597197
	},
	"NodeID-PuYX3u6gfLje9bt6fCDFHkSwqaL8VSpv6": {
	  "publicKey": null,
	  "weight": 52500000000000
	},
	"NodeID-Puv7ksUTmwsxqEx7XLikRicKtKss6QNNi": {
	  "publicKey": null,
	  "weight": 9497197002529
	},
	"NodeID-Pv5X1dBk1wfZJwRqm8Aj9QecxvA87r31j": {
	  "publicKey": null,
	  "weight": 7322528308663
	},
	"NodeID-PvZnYX1Nz8jp1GjQFixs1TAX9C9QXwrwL": {
	  "publicKey": null,
	  "weight": 2627000000000
	},
	"NodeID-PzQafFPgbCLkvMgrP7nh7aCG1eKypdNi2": {
	  "publicKey": null,
	  "weight": 8668692244648
	},
	"NodeID-Pzx6GZGZeYkSAQ3DNCUPqWani5F75DeLd": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Q3iFBGVU6gka1yX9EMjpNuezkocDTq1MB": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Q44XVjVEhgxacoq2ETj7yLBuUra9ru2BA": {
	  "publicKey": null,
	  "weight": 10000000000000
	},
	"NodeID-Q4mUNVKvwEHF3tkxQcCYbrQA8XWdybLzv": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Q5rNbr5yFwotdbeRPGeTagHP1hbHYuaJg": {
	  "publicKey": null,
	  "weight": 383114823529386
	},
	"NodeID-Q5wrexzsrQ94xJcXsqRQmyvRHCuYbrZof": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-Q5xvrPQSjHhJX8eLKbBnZMWbS648NqnsZ": {
	  "publicKey": null,
	  "weight": 2721997326200
	},
	"NodeID-Q6ZLZ1EifMyNd6evH83FJrVGyauaUSP9S": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-QBowbP1jz8zA86CtSzRwsFsxr5q1rY6pB": {
	  "publicKey": null,
	  "weight": 4300000000000
	},
	"NodeID-QCfTXZ1cdPt5XT1jiJfycHKTFo6Yg164L": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-QCkFTaWJg5ixVhZoNbePxzs9HWrTutVy4": {
	  "publicKey": null,
	  "weight": 1288448473003229
	},
	"NodeID-QDLrtBngMA6mFveLNZTXxKQVEYx518fG8": {
	  "publicKey": null,
	  "weight": 2034715788886
	},
	"NodeID-QDfLjUuR9D5a6G5ZAQdPPbggJxtLnbPj3": {
	  "publicKey": null,
	  "weight": 3772818587637
	},
	"NodeID-QHX9w6oLs8zjHdBvbzx1ZadHvy9aUNLr5": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-QJMNxwTLjdg9C1vY2BKNtapWUuoeCgHFj": {
	  "publicKey": null,
	  "weight": 10134519715545
	},
	"NodeID-QKGoUvqcgormCoMj6yPw9isY7DX9H4mdd": {
	  "publicKey": null,
	  "weight": 768000570722447
	},
	"NodeID-QKn4ujiEcP1n8iFbEJqNsKZUYbUG4fHAc": {
	  "publicKey": null,
	  "weight": 5878919505790
	},
	"NodeID-QisfDcwSCUGaw6CwsBLA4mjJmTDto3he": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-QmTPCR33deSWVwZ4Ri2jUw6cDAhoUwZP": {
	  "publicKey": null,
	  "weight": 9481412958134
	},
	"NodeID-TJ6G6FU4QtQFVNjzdjhjknqek5SZZGcy": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-TgaMGAEpkXKAisumnnmzzRzVkexbSkB7": {
	  "publicKey": null,
	  "weight": 3000000000000000
	},
	"NodeID-TzsHrNrY8qaAkg9MPAmJd7LfzCUDjvzi": {
	  "publicKey": null,
	  "weight": 23083038066325
	},
	"NodeID-UC1qmKMoMcfShU6idcreEAsxSHs3BDBj": {
	  "publicKey": null,
	  "weight": 2039014018762
	},
	"NodeID-V3fYZv2vHtLB5gQF548JRnSyk4tZfZ2o": {
	  "publicKey": null,
	  "weight": 3620420707794
	},
	"NodeID-V5R4XjWDtjs8Zi2g3TRNzngnGJiXm2ic": {
	  "publicKey": null,
	  "weight": 7900597540000
	},
	"NodeID-VT3YhgFaWEzy4Ap937qMeNEDscCammzG": {
	  "publicKey": null,
	  "weight": 30005000000000
	},
	"NodeID-WAHX8MXekynX6xZSFiSewY9Qxnhq5HZh": {
	  "publicKey": null,
	  "weight": 222220000000000
	},
	"NodeID-WrL5pKhhp8e2YLp6h1e62NBUPhuR6pGa": {
	  "publicKey": null,
	  "weight": 8773621542900
	},
	"NodeID-XYAkZN1GfMwcYRRyGLiZsGV2thZbqDGT": {
	  "publicKey": null,
	  "weight": 2503742376879
	},
	"NodeID-YdCKQdXecSPWqABWcRtPZvBaLd3MPPR8": {
	  "publicKey": null,
	  "weight": 6059046197220
	},
	"NodeID-ZVLPagHUUrQ2vB7mjsWQ44wNGf5SVKKW": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-ZwShn4JT1Bg91QphD3VBLryA88GFnAg5": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-aFRSyUykQjScZiWASVqmdTGzBfV9KKj1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-axbVjo1hwZgpUr1wBapsz5XHmMLFP3wQ": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-bTSBn4p38MofpJvhDgfNxuKAQkKxdqEi": {
	  "publicKey": null,
	  "weight": 20288620845495
	},
	"NodeID-bx2UN8z5nfLMTFphQqnYeyngQnccL6tN": {
	  "publicKey": null,
	  "weight": 2165937473524691
	},
	"NodeID-cmcBnV62GZxrLLmebZdeofd3AesueXac": {
	  "publicKey": null,
	  "weight": 2088409786534
	},
	"NodeID-cme6KToKFLCwqj27F7a7Ys9iC4cLrqMz": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-duaNgYtghtCRzj2rkiRDDs51BPX8CCjn": {
	  "publicKey": null,
	  "weight": 5200000000000
	},
	"NodeID-fr4FFtJ5PF5UZpyA3RyoxhNks6Yfhu7u": {
	  "publicKey": null,
	  "weight": 4150656229676
	},
	"NodeID-i9uQjR8P5Wzd1wTLeALkZ2oXqKjMKQbs": {
	  "publicKey": null,
	  "weight": 6194988424204
	},
	"NodeID-jB3tEp5oBcHYRLLmNaMgP4v59dyr1a9E": {
	  "publicKey": null,
	  "weight": 2051367026794
	},
	"NodeID-jjTi6fViUUkPwZQbE9N8Bm7AThmVyjGW": {
	  "publicKey": null,
	  "weight": 2004521553668
	},
	"NodeID-jz1mR5zZukizoZEZ1Wa17h18Kou1Rfzt": {
	  "publicKey": null,
	  "weight": 3090379930674
	},
	"NodeID-k9Qr9tpMZuBxJ3wtW22CQxW92hMrQC4A": {
	  "publicKey": null,
	  "weight": 3292366007735
	},
	"NodeID-kCLme7NNnAVrbQGZadzgWwtoPJpkQcJR": {
	  "publicKey": null,
	  "weight": 4500000000000
	},
	"NodeID-kZNuQMHhydefgnwjYX1fhHMpRNAs9my1": {
	  "publicKey": null,
	  "weight": 657199764705822
	},
	"NodeID-mEbcN58z2esHFnnXHeR3T2UTfsR5H6WJ": {
	  "publicKey": null,
	  "weight": 50031758438622
	},
	"NodeID-mRrTBjdn3s4oSyxjKEzLG2ibWU6V6Vzm": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-mfUNFSM9ak3NuCvaiFCtcmEcHF9ptfrU": {
	  "publicKey": null,
	  "weight": 2009113812404
	},
	"NodeID-mwty8tviuGBvDEG8e6vTQByDZN2MKnNo": {
	  "publicKey": null,
	  "weight": 5900955789300
	},
	"NodeID-mzhghLeMzv8qrRrK76wABwrf3LrL6mP1": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-nJm8Ltnf7MDCiuP3PzU4hv3RqXXKmeqd": {
	  "publicKey": null,
	  "weight": 2009075172088
	},
	"NodeID-nT4KHRTgD9Lq44epZo7GPhjQTEFcGbT7": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-oZyu1tEJrBeFkquoqZ5rYpvsEAkfFiwD": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-qyJZPRspVKhwXuaimMA7Ysvb5KYLeMjB": {
	  "publicKey": null,
	  "weight": 2040102964212
	},
	"NodeID-rA57yc1JzDdo6Fc4gND6Mwz8hQ4RBxJM": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-rKCpkJKNAZeE8sKEHduq8Co7GwrAdTJb": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-rS4TWt8iuAbQTsg6ez6u5tXvUzicRzuq": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-rroLHLixQkZMTsTnePdWhDrGpWenYX7t": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-sAyLfHTC8ZQRn5HtfjNKxXDbnZTaJUqa": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-t9a4uF4ymJ5rmXeW3RDbvv86nJYztoLF": {
	  "publicKey": null,
	  "weight": 2000000000000
	},
	"NodeID-tHkVEqdhWPdX2vJGbqDWebw2VPrho4Aq": {
	  "publicKey": null,
	  "weight": 6364523839016
	},
	"NodeID-vV5S1LUtAxCWq9ASiDSjZpX88y2WcFGM": {
	  "publicKey": null,
	  "weight": 726000000000000
	},
	"NodeID-w6kJAMpdT4hB9jE4NnEfBm47PzD1UXfW": {
	  "publicKey": null,
	  "weight": 2054500000000
	},
	"NodeID-xQU1Ntg4uYkkSgXivV2ELwr8s8FovRCS": {
	  "publicKey": null,
	  "weight": 24994112359329
	},
	"NodeID-xVVbmFL3eww4dmBkgQ2akBxcnEy4VyLh": {
	  "publicKey": null,
	  "weight": 7140120000000
	},
	"NodeID-z3rX3CnWLi5KBXQ8nVB4W8NnLNdUBcPD": {
	  "publicKey": null,
	  "weight": 22139889575955
	},
	"NodeID-zbFM8qH7MnQ8uo6rm4tZaq6vhbF4cpgt": {
	  "publicKey": null,
	  "weight": 890344773985347
	}
  }`

func init() {
	res := make(map[ids.NodeID]*GetValidatorOutput, 0)

	if err := json.Unmarshal([]byte(j), &res); err != nil {
		panic(err)
	}
	ExpectedVdrSetQ = res
}
