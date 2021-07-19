package providers

import "testing"

var p = `{
	"latest": {
	  "release": "1.17.1",
	  "snapshot": "1.17.1"
	},
	"versions": [
	  {
		"id": "1.17.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/8b976413591b4132fc4f27370dcd87ce1e50fb2f/1.17.1.json",
		"time": "2021-07-06T12:05:05+00:00",
		"releaseTime": "2021-07-06T12:01:34+00:00",
		"sha1": "8b976413591b4132fc4f27370dcd87ce1e50fb2f",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17.1-rc2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/16eaa3c82ff61349c340cad27c40478114a7136b/1.17.1-rc2.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-07-05T12:58:01+00:00",
		"sha1": "16eaa3c82ff61349c340cad27c40478114a7136b",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17.1-rc1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/950d7a337f8ae61137961bd46349d7645baa9e80/1.17.1-rc1.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-07-01T15:23:37+00:00",
		"sha1": "950d7a337f8ae61137961bd46349d7645baa9e80",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17.1-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9762f411b1a3b23fdf3b39ec13252bebcdccbcfb/1.17.1-pre3.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-06-30T15:43:16+00:00",
		"sha1": "9762f411b1a3b23fdf3b39ec13252bebcdccbcfb",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17.1-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2ed1208227537e5b8acd2f85ad007716d0649ae5/1.17.1-pre2.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-06-29T15:14:12+00:00",
		"sha1": "2ed1208227537e5b8acd2f85ad007716d0649ae5",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17.1-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/cd62ef85c8dd4b28d9c198b95a5532848429b488/1.17.1-pre1.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-06-18T12:24:40+00:00",
		"sha1": "cd62ef85c8dd4b28d9c198b95a5532848429b488",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/c64ee5e617d7af2543acdec234cbedab3b273b7b/1.17.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-06-08T11:00:40+00:00",
		"sha1": "c64ee5e617d7af2543acdec234cbedab3b273b7b",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17-rc2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/dde3abea71da8051912186864f11d02cd1eac06a/1.17-rc2.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-06-07T11:46:28+00:00",
		"sha1": "dde3abea71da8051912186864f11d02cd1eac06a",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17-rc1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a18641fca6dd394319f6d192d387b360886bb919/1.17-rc1.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-06-04T13:24:48+00:00",
		"sha1": "a18641fca6dd394319f6d192d387b360886bb919",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17-pre5",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/27da71336e94b16af1169cb7da7ce05807479f58/1.17-pre5.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-06-03T17:01:28+00:00",
		"sha1": "27da71336e94b16af1169cb7da7ce05807479f58",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/41aa7ba935b09064933240b78febf52a1bca5cfb/1.17-pre4.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-06-02T16:15:43+00:00",
		"sha1": "41aa7ba935b09064933240b78febf52a1bca5cfb",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d99bfcd30f793391c8e0b78cd056f03cc42b46fc/1.17-pre3.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-06-01T15:43:46+00:00",
		"sha1": "d99bfcd30f793391c8e0b78cd056f03cc42b46fc",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/579a8bb200431d5c3cf34fe21b171f241554c692/1.17-pre2.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-05-31T15:54:05+00:00",
		"sha1": "579a8bb200431d5c3cf34fe21b171f241554c692",
		"complianceLevel": 1
	  },
	  {
		"id": "1.17-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2de8c49d573fbf6e4ed63525f938080fbb67f86a/1.17-pre1.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-05-27T09:39:21+00:00",
		"sha1": "2de8c49d573fbf6e4ed63525f938080fbb67f86a",
		"complianceLevel": 1
	  },
	  {
		"id": "21w20a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5c581950ee5d0eb9967c0f08be13b4ae54e2fd11/21w20a.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-05-19T15:22:02+00:00",
		"sha1": "5c581950ee5d0eb9967c0f08be13b4ae54e2fd11",
		"complianceLevel": 1
	  },
	  {
		"id": "21w19a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0d846a767dfec2f4f16579cb6fad7926a888eb29/21w19a.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-05-12T11:19:15+00:00",
		"sha1": "0d846a767dfec2f4f16579cb6fad7926a888eb29",
		"complianceLevel": 1
	  },
	  {
		"id": "21w18a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2ce07891d7199eeb8ad75ff27176a9cd4268fa46/21w18a.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-05-05T15:24:35+00:00",
		"sha1": "2ce07891d7199eeb8ad75ff27176a9cd4268fa46",
		"complianceLevel": 1
	  },
	  {
		"id": "21w17a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a35d86fce90b5d6f00b45f2ff9123f592d3b3884/21w17a.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-04-28T13:54:05+00:00",
		"sha1": "a35d86fce90b5d6f00b45f2ff9123f592d3b3884",
		"complianceLevel": 1
	  },
	  {
		"id": "21w16a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/09cfaa8835ecdc440d8c335639638dd048d33855/21w16a.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-04-21T16:41:14+00:00",
		"sha1": "09cfaa8835ecdc440d8c335639638dd048d33855",
		"complianceLevel": 1
	  },
	  {
		"id": "21w15a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/cf518b3c640d5737d1ebddbf9dad022abba6ecf5/21w15a.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-04-14T13:41:34+00:00",
		"sha1": "cf518b3c640d5737d1ebddbf9dad022abba6ecf5",
		"complianceLevel": 1
	  },
	  {
		"id": "21w14a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0d17e06a6e7e112d963688defdd4fcae46b5531a/21w14a.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-04-07T14:04:09+00:00",
		"sha1": "0d17e06a6e7e112d963688defdd4fcae46b5531a",
		"complianceLevel": 1
	  },
	  {
		"id": "21w13a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e21dcc289d8e8e676b8005244137ec08c844feb4/21w13a.json",
		"time": "2021-07-06T11:54:23+00:00",
		"releaseTime": "2021-03-31T16:17:46+00:00",
		"sha1": "e21dcc289d8e8e676b8005244137ec08c844feb4",
		"complianceLevel": 1
	  },
	  {
		"id": "21w11a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5645a5ec22b8b05e0e0cd444bf73153a65fb541f/21w11a.json",
		"time": "2021-07-07T13:12:13+00:00",
		"releaseTime": "2021-03-17T15:05:50+00:00",
		"sha1": "5645a5ec22b8b05e0e0cd444bf73153a65fb541f",
		"complianceLevel": 1
	  },
	  {
		"id": "21w10a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1c9582fccc613c4c06e7ce2cd3a94b4ca61c9925/21w10a.json",
		"time": "2021-07-07T13:12:12+00:00",
		"releaseTime": "2021-03-10T15:24:38+00:00",
		"sha1": "1c9582fccc613c4c06e7ce2cd3a94b4ca61c9925",
		"complianceLevel": 1
	  },
	  {
		"id": "21w08b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/174f2274c10a7e5ad99ad12a646ff8411e533cdc/21w08b.json",
		"time": "2021-07-07T13:12:11+00:00",
		"releaseTime": "2021-02-25T11:46:34+00:00",
		"sha1": "174f2274c10a7e5ad99ad12a646ff8411e533cdc",
		"complianceLevel": 1
	  },
	  {
		"id": "21w08a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/78b17fe5d4c491a2ca144e9f5f241ac1e2aa651e/21w08a.json",
		"time": "2021-07-07T13:12:11+00:00",
		"releaseTime": "2021-02-24T14:38:51+00:00",
		"sha1": "78b17fe5d4c491a2ca144e9f5f241ac1e2aa651e",
		"complianceLevel": 1
	  },
	  {
		"id": "21w07a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/003d04040b963783d53bf1c6fbf382f26748fea4/21w07a.json",
		"time": "2021-07-07T13:12:10+00:00",
		"releaseTime": "2021-02-17T16:35:40+00:00",
		"sha1": "003d04040b963783d53bf1c6fbf382f26748fea4",
		"complianceLevel": 1
	  },
	  {
		"id": "21w06a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/785de8e3b0a6178b0990e8a218dda0366ed346df/21w06a.json",
		"time": "2021-07-07T13:12:09+00:00",
		"releaseTime": "2021-02-10T17:13:54+00:00",
		"sha1": "785de8e3b0a6178b0990e8a218dda0366ed346df",
		"complianceLevel": 1
	  },
	  {
		"id": "21w05b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/765d43907b09f2cddca6c8bd7a1e2941f2909fe6/21w05b.json",
		"time": "2021-07-07T13:12:09+00:00",
		"releaseTime": "2021-02-04T15:09:29+00:00",
		"sha1": "765d43907b09f2cddca6c8bd7a1e2941f2909fe6",
		"complianceLevel": 1
	  },
	  {
		"id": "21w05a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0b9febc2a12cf5e3439ac8e82748f4762138926c/21w05a.json",
		"time": "2021-07-07T13:12:08+00:00",
		"releaseTime": "2021-02-03T15:56:54+00:00",
		"sha1": "0b9febc2a12cf5e3439ac8e82748f4762138926c",
		"complianceLevel": 1
	  },
	  {
		"id": "21w03a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9deedf2397d863a6e0873d00d03bf79db169725b/21w03a.json",
		"time": "2021-07-07T13:12:07+00:00",
		"releaseTime": "2021-01-20T14:56:29+00:00",
		"sha1": "9deedf2397d863a6e0873d00d03bf79db169725b",
		"complianceLevel": 1
	  },
	  {
		"id": "1.16.5",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/66935fe3a8f602111a3f3cba9867d3fd6d80ef47/1.16.5.json",
		"time": "2021-07-07T13:08:46+00:00",
		"releaseTime": "2021-01-14T16:05:32+00:00",
		"sha1": "66935fe3a8f602111a3f3cba9867d3fd6d80ef47",
		"complianceLevel": 1
	  },
	  {
		"id": "1.16.5-rc1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b96e6ec33997b75edb8b7b9b56ad2307ac969a7d/1.16.5-rc1.json",
		"time": "2021-07-07T13:09:39+00:00",
		"releaseTime": "2021-01-13T15:58:55+00:00",
		"sha1": "b96e6ec33997b75edb8b7b9b56ad2307ac969a7d",
		"complianceLevel": 1
	  },
	  {
		"id": "20w51a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/48c4d7c0bd310a6614f7d8f7d77f6eccb4af679b/20w51a.json",
		"time": "2021-07-07T13:12:07+00:00",
		"releaseTime": "2020-12-16T16:27:57+00:00",
		"sha1": "48c4d7c0bd310a6614f7d8f7d77f6eccb4af679b",
		"complianceLevel": 1
	  },
	  {
		"id": "20w49a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/8778ca44e19b38ff58ef625f2cc4d13467153910/20w49a.json",
		"time": "2021-07-07T13:12:06+00:00",
		"releaseTime": "2020-12-02T16:47:20+00:00",
		"sha1": "8778ca44e19b38ff58ef625f2cc4d13467153910",
		"complianceLevel": 1
	  },
	  {
		"id": "20w48a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/111174f44d31c7a64012660584122d22fccfce26/20w48a.json",
		"time": "2021-07-07T13:12:05+00:00",
		"releaseTime": "2020-11-25T15:42:24+00:00",
		"sha1": "111174f44d31c7a64012660584122d22fccfce26",
		"complianceLevel": 1
	  },
	  {
		"id": "20w46a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b07744292f02468a3b3e7a1641cf401e9b99b3db/20w46a.json",
		"time": "2021-07-07T13:12:05+00:00",
		"releaseTime": "2020-11-11T15:30:32+00:00",
		"sha1": "b07744292f02468a3b3e7a1641cf401e9b99b3db",
		"complianceLevel": 1
	  },
	  {
		"id": "20w45a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4264ae5dd2187782a01d0d354b4ca5bec1e047de/20w45a.json",
		"time": "2021-07-07T13:12:04+00:00",
		"releaseTime": "2020-11-04T16:42:00+00:00",
		"sha1": "4264ae5dd2187782a01d0d354b4ca5bec1e047de",
		"complianceLevel": 1
	  },
	  {
		"id": "1.16.4",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/584f4f319f0acc2f8c9a7d992ae20b95f425c1c5/1.16.4.json",
		"time": "2021-07-07T13:08:45+00:00",
		"releaseTime": "2020-10-29T15:49:37+00:00",
		"sha1": "584f4f319f0acc2f8c9a7d992ae20b95f425c1c5",
		"complianceLevel": 1
	  },
	  {
		"id": "1.16.4-rc1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e263a7616eaf878b4c84d00cf053317dd5abffa1/1.16.4-rc1.json",
		"time": "2021-07-07T13:09:39+00:00",
		"releaseTime": "2020-10-27T16:31:08+00:00",
		"sha1": "e263a7616eaf878b4c84d00cf053317dd5abffa1",
		"complianceLevel": 1
	  },
	  {
		"id": "1.16.4-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/fe829be2c62b87d0d4e13aba0ece7786b564bcb3/1.16.4-pre2.json",
		"time": "2021-07-07T13:09:38+00:00",
		"releaseTime": "2020-10-22T15:32:17+00:00",
		"sha1": "fe829be2c62b87d0d4e13aba0ece7786b564bcb3",
		"complianceLevel": 1
	  },
	  {
		"id": "1.16.4-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/381243982348d85f9cd2c1500af1c728526d7bad/1.16.4-pre1.json",
		"time": "2021-07-07T13:09:38+00:00",
		"releaseTime": "2020-10-13T14:36:07+00:00",
		"sha1": "381243982348d85f9cd2c1500af1c728526d7bad",
		"complianceLevel": 1
	  },
	  {
		"id": "1.16.3",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/c0f2a74ac70b32fc1d0bf791c3f0b3f26b612fa0/1.16.3.json",
		"time": "2021-07-07T13:08:44+00:00",
		"releaseTime": "2020-09-10T13:42:37+00:00",
		"sha1": "c0f2a74ac70b32fc1d0bf791c3f0b3f26b612fa0",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16.3-rc1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0ca0f44ccc155120da044edba421eeb88d019af4/1.16.3-rc1.json",
		"time": "2021-07-07T13:09:37+00:00",
		"releaseTime": "2020-09-07T12:34:06+00:00",
		"sha1": "0ca0f44ccc155120da044edba421eeb88d019af4",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/361e7caa416f08d02b311115ee734d69cc92b94a/1.16.2.json",
		"time": "2021-07-07T13:08:44+00:00",
		"releaseTime": "2020-08-11T10:13:46+00:00",
		"sha1": "361e7caa416f08d02b311115ee734d69cc92b94a",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16.2-rc2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5c79553f36b4102f210deb396a3c7b58904bf42d/1.16.2-rc2.json",
		"time": "2021-07-07T13:09:36+00:00",
		"releaseTime": "2020-08-10T11:43:36+00:00",
		"sha1": "5c79553f36b4102f210deb396a3c7b58904bf42d",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16.2-rc1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/35dd547e508028bb23566b704245ba606d78e8bd/1.16.2-rc1.json",
		"time": "2021-07-07T13:09:36+00:00",
		"releaseTime": "2020-08-07T14:35:39+00:00",
		"sha1": "35dd547e508028bb23566b704245ba606d78e8bd",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16.2-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6654ec97ba19a05e72f7fce86139fda787d0d469/1.16.2-pre3.json",
		"time": "2021-07-07T13:09:35+00:00",
		"releaseTime": "2020-08-06T16:44:52+00:00",
		"sha1": "6654ec97ba19a05e72f7fce86139fda787d0d469",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16.2-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5b8d0e026e7b37f1dec188ab2bd956683cfdc969/1.16.2-pre2.json",
		"time": "2021-07-07T13:09:34+00:00",
		"releaseTime": "2020-08-05T15:30:50+00:00",
		"sha1": "5b8d0e026e7b37f1dec188ab2bd956683cfdc969",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16.2-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4012edf6898676810e61ad5e580aefb08d7d8e0e/1.16.2-pre1.json",
		"time": "2021-07-07T13:09:34+00:00",
		"releaseTime": "2020-07-29T13:19:05+00:00",
		"sha1": "4012edf6898676810e61ad5e580aefb08d7d8e0e",
		"complianceLevel": 0
	  },
	  {
		"id": "20w30a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/3d87282dec7603e8855d21ca254e9d134e1f3abf/20w30a.json",
		"time": "2021-07-07T13:12:03+00:00",
		"releaseTime": "2020-07-22T15:05:15+00:00",
		"sha1": "3d87282dec7603e8855d21ca254e9d134e1f3abf",
		"complianceLevel": 0
	  },
	  {
		"id": "20w29a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/40dcf13a8e1e6c59cb0ab6108b7c2cb2f15471f1/20w29a.json",
		"time": "2021-07-07T13:12:03+00:00",
		"releaseTime": "2020-07-15T14:13:47+00:00",
		"sha1": "40dcf13a8e1e6c59cb0ab6108b7c2cb2f15471f1",
		"complianceLevel": 0
	  },
	  {
		"id": "20w28a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/025e7a3bb518c5ba9908c94ce41fb4294025bbbc/20w28a.json",
		"time": "2021-07-07T13:12:02+00:00",
		"releaseTime": "2020-07-08T15:10:40+00:00",
		"sha1": "025e7a3bb518c5ba9908c94ce41fb4294025bbbc",
		"complianceLevel": 0
	  },
	  {
		"id": "20w27a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0cdda7643b9bcecf8a55da7d50aa1d2eef1a2024/20w27a.json",
		"time": "2021-07-07T13:12:01+00:00",
		"releaseTime": "2020-07-01T15:07:35+00:00",
		"sha1": "0cdda7643b9bcecf8a55da7d50aa1d2eef1a2024",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/a3622c3919f63be1e57af5803b1c9ac47f59297e/1.16.1.json",
		"time": "2021-07-07T13:08:43+00:00",
		"releaseTime": "2020-06-24T10:31:40+00:00",
		"sha1": "a3622c3919f63be1e57af5803b1c9ac47f59297e",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/5a157ca1ae150c3acb10afdfe714ba34bf322315/1.16.json",
		"time": "2021-07-07T13:08:46+00:00",
		"releaseTime": "2020-06-23T16:20:52+00:00",
		"sha1": "5a157ca1ae150c3acb10afdfe714ba34bf322315",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16-rc1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/09263e766c5880abb33bc5ae3b94b8b105848ade/1.16-rc1.json",
		"time": "2021-07-07T13:09:33+00:00",
		"releaseTime": "2020-06-18T12:49:28+00:00",
		"sha1": "09263e766c5880abb33bc5ae3b94b8b105848ade",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16-pre8",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/503cfde46f73b14e066e0c718c5cf9725651addd/1.16-pre8.json",
		"time": "2021-07-07T13:09:33+00:00",
		"releaseTime": "2020-06-17T14:45:23+00:00",
		"sha1": "503cfde46f73b14e066e0c718c5cf9725651addd",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16-pre7",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/3c76579ee82efca46472e2f08233e9544c739b00/1.16-pre7.json",
		"time": "2021-07-07T13:09:32+00:00",
		"releaseTime": "2020-06-16T15:31:35+00:00",
		"sha1": "3c76579ee82efca46472e2f08233e9544c739b00",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16-pre6",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/dac775b17fd2cf2a4669d1079242313adf2ec9e1/1.16-pre6.json",
		"time": "2021-07-07T13:09:31+00:00",
		"releaseTime": "2020-06-15T16:57:57+00:00",
		"sha1": "dac775b17fd2cf2a4669d1079242313adf2ec9e1",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16-pre5",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/54caf937f9c04dff3c46b60bb5a3c84e0e0e7f2b/1.16-pre5.json",
		"time": "2021-07-07T13:09:31+00:00",
		"releaseTime": "2020-06-12T14:33:59+00:00",
		"sha1": "54caf937f9c04dff3c46b60bb5a3c84e0e0e7f2b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1ceac7107add1a34674a0113053057655ff0cffb/1.16-pre4.json",
		"time": "2021-07-07T13:09:30+00:00",
		"releaseTime": "2020-06-11T15:45:55+00:00",
		"sha1": "1ceac7107add1a34674a0113053057655ff0cffb",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/60329b57af598ae647e508cb7a12a2a01225a3f7/1.16-pre3.json",
		"time": "2021-07-07T13:09:30+00:00",
		"releaseTime": "2020-06-10T14:57:43+00:00",
		"sha1": "60329b57af598ae647e508cb7a12a2a01225a3f7",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/bdd418cf5ab5725abc1a600760dea42d49e712ab/1.16-pre2.json",
		"time": "2021-07-07T13:09:29+00:00",
		"releaseTime": "2020-06-05T10:47:59+00:00",
		"sha1": "bdd418cf5ab5725abc1a600760dea42d49e712ab",
		"complianceLevel": 0
	  },
	  {
		"id": "1.16-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/eecc2c868c42d199394adb6b6fb7e7359b6962ff/1.16-pre1.json",
		"time": "2021-07-07T13:09:28+00:00",
		"releaseTime": "2020-06-04T18:17:51+00:00",
		"sha1": "eecc2c868c42d199394adb6b6fb7e7359b6962ff",
		"complianceLevel": 0
	  },
	  {
		"id": "20w22a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c31805fc7c1f7a359b16d2fe15e23fb8a678e2c1/20w22a.json",
		"time": "2021-07-07T13:12:01+00:00",
		"releaseTime": "2020-05-29T11:25:02+00:00",
		"sha1": "c31805fc7c1f7a359b16d2fe15e23fb8a678e2c1",
		"complianceLevel": 0
	  },
	  {
		"id": "20w21a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d87206332345014deb69b55503b3063a8422f6d9/20w21a.json",
		"time": "2021-07-07T13:12:00+00:00",
		"releaseTime": "2020-05-20T12:07:18+00:00",
		"sha1": "d87206332345014deb69b55503b3063a8422f6d9",
		"complianceLevel": 0
	  },
	  {
		"id": "20w20b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c8940dd5f3de496cc714761c4af5bac8018209b4/20w20b.json",
		"time": "2021-07-07T13:11:59+00:00",
		"releaseTime": "2020-05-14T08:16:26+00:00",
		"sha1": "c8940dd5f3de496cc714761c4af5bac8018209b4",
		"complianceLevel": 0
	  },
	  {
		"id": "20w20a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e9cc96ba4c0bf2e9fb89b40c53d109c87d471fad/20w20a.json",
		"time": "2021-07-07T13:11:59+00:00",
		"releaseTime": "2020-05-13T15:11:43+00:00",
		"sha1": "e9cc96ba4c0bf2e9fb89b40c53d109c87d471fad",
		"complianceLevel": 0
	  },
	  {
		"id": "20w19a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/212028103450cae6b3fc0c10827b2f854b6af54b/20w19a.json",
		"time": "2021-07-07T13:11:58+00:00",
		"releaseTime": "2020-05-06T16:23:24+00:00",
		"sha1": "212028103450cae6b3fc0c10827b2f854b6af54b",
		"complianceLevel": 0
	  },
	  {
		"id": "20w18a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b4a7ccd2675ad3a13cee80b471d12287f6c39abc/20w18a.json",
		"time": "2021-07-07T13:11:58+00:00",
		"releaseTime": "2020-04-29T15:16:34+00:00",
		"sha1": "b4a7ccd2675ad3a13cee80b471d12287f6c39abc",
		"complianceLevel": 0
	  },
	  {
		"id": "20w17a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/634970bcd5670e194bc37f49f56a7cd9b1472782/20w17a.json",
		"time": "2021-07-07T13:11:57+00:00",
		"releaseTime": "2020-04-22T13:47:50+00:00",
		"sha1": "634970bcd5670e194bc37f49f56a7cd9b1472782",
		"complianceLevel": 0
	  },
	  {
		"id": "20w16a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7d2845dae4a2f62abda54ec072f437e8b8b64e98/20w16a.json",
		"time": "2021-07-07T13:11:56+00:00",
		"releaseTime": "2020-04-15T14:13:01+00:00",
		"sha1": "7d2845dae4a2f62abda54ec072f437e8b8b64e98",
		"complianceLevel": 0
	  },
	  {
		"id": "20w15a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6b862346a07a40fab9663669e662656fa4517cda/20w15a.json",
		"time": "2021-07-07T13:11:56+00:00",
		"releaseTime": "2020-04-08T12:29:24+00:00",
		"sha1": "6b862346a07a40fab9663669e662656fa4517cda",
		"complianceLevel": 0
	  },
	  {
		"id": "20w14a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d231f86dcac021dc90fe53aa5e687ef916af491b/20w14a.json",
		"time": "2021-07-07T13:11:54+00:00",
		"releaseTime": "2020-04-02T14:28:06+00:00",
		"sha1": "d231f86dcac021dc90fe53aa5e687ef916af491b",
		"complianceLevel": 0
	  },
	  {
		"id": "20w14infinite",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d38463d42540d715198ab182daee5151d013dce6/20w14infinite.json",
		"time": "2021-07-07T13:11:55+00:00",
		"releaseTime": "2020-04-01T12:47:08+00:00",
		"sha1": "d38463d42540d715198ab182daee5151d013dce6",
		"complianceLevel": 0
	  },
	  {
		"id": "20w13b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c32dbc125f6a9d4744db79e8c3495fa9d897764d/20w13b.json",
		"time": "2021-07-07T13:11:54+00:00",
		"releaseTime": "2020-03-26T13:00:34+00:00",
		"sha1": "c32dbc125f6a9d4744db79e8c3495fa9d897764d",
		"complianceLevel": 0
	  },
	  {
		"id": "20w13a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/cc0b303de2c68d53cf44ec3192ef2578805df74c/20w13a.json",
		"time": "2021-07-07T13:11:53+00:00",
		"releaseTime": "2020-03-25T17:05:33+00:00",
		"sha1": "cc0b303de2c68d53cf44ec3192ef2578805df74c",
		"complianceLevel": 0
	  },
	  {
		"id": "20w12a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9607db8c5a24e208cce06b7ca6bcc637dec2a610/20w12a.json",
		"time": "2021-07-07T13:11:53+00:00",
		"releaseTime": "2020-03-18T16:42:06+00:00",
		"sha1": "9607db8c5a24e208cce06b7ca6bcc637dec2a610",
		"complianceLevel": 0
	  },
	  {
		"id": "20w11a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/23baac3a1323ba0fcc1d8b7d10108fda861faf24/20w11a.json",
		"time": "2021-07-07T13:11:52+00:00",
		"releaseTime": "2020-03-11T16:28:27+00:00",
		"sha1": "23baac3a1323ba0fcc1d8b7d10108fda861faf24",
		"complianceLevel": 0
	  },
	  {
		"id": "20w10a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/32f177b8763dfe1b7ae853e0d8725bddd16530ea/20w10a.json",
		"time": "2021-07-07T13:11:51+00:00",
		"releaseTime": "2020-03-04T16:21:41+00:00",
		"sha1": "32f177b8763dfe1b7ae853e0d8725bddd16530ea",
		"complianceLevel": 0
	  },
	  {
		"id": "20w09a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0a6993b1b9cd9552d1990a2c87eb6f1a221d62a2/20w09a.json",
		"time": "2021-07-07T13:11:51+00:00",
		"releaseTime": "2020-02-26T16:43:08+00:00",
		"sha1": "0a6993b1b9cd9552d1990a2c87eb6f1a221d62a2",
		"complianceLevel": 0
	  },
	  {
		"id": "20w08a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/8ccbfa3ac0b9094bff4385632f1589d363299eaa/20w08a.json",
		"time": "2021-07-07T13:11:50+00:00",
		"releaseTime": "2020-02-19T13:30:09+00:00",
		"sha1": "8ccbfa3ac0b9094bff4385632f1589d363299eaa",
		"complianceLevel": 0
	  },
	  {
		"id": "20w07a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f513f3b8331c13278e00ef8056fbe53870fd8834/20w07a.json",
		"time": "2021-07-07T13:11:49+00:00",
		"releaseTime": "2020-02-14T13:20:49+00:00",
		"sha1": "f513f3b8331c13278e00ef8056fbe53870fd8834",
		"complianceLevel": 0
	  },
	  {
		"id": "20w06a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/3f53be2e6c427bace9f7de6467cb85cf930e20bc/20w06a.json",
		"time": "2021-07-07T13:11:49+00:00",
		"releaseTime": "2020-02-05T16:05:22+00:00",
		"sha1": "3f53be2e6c427bace9f7de6467cb85cf930e20bc",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/52396828d64eae5bdd01a5ac787234a5f4c85e59/1.15.2.json",
		"time": "2021-07-07T13:08:42+00:00",
		"releaseTime": "2020-01-17T10:03:52+00:00",
		"sha1": "52396828d64eae5bdd01a5ac787234a5f4c85e59",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15.2-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/8a8cf04179742f5a439d820fcfaeb76369bf7303/1.15.2-pre2.json",
		"time": "2021-07-07T13:09:28+00:00",
		"releaseTime": "2020-01-16T12:35:57+00:00",
		"sha1": "8a8cf04179742f5a439d820fcfaeb76369bf7303",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15.2-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/66c95a724087cfd6cd01ad6845d9e4bfe1f5a341/1.15.2-pre1.json",
		"time": "2021-07-07T13:09:27+00:00",
		"releaseTime": "2020-01-14T16:19:31+00:00",
		"sha1": "66c95a724087cfd6cd01ad6845d9e4bfe1f5a341",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/6d5a4d2fece3c607d851226eabd499a6caa7cb75/1.15.1.json",
		"time": "2021-07-07T13:08:41+00:00",
		"releaseTime": "2019-12-16T10:29:47+00:00",
		"sha1": "6d5a4d2fece3c607d851226eabd499a6caa7cb75",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15.1-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4c7b181580f9f208236609d787baec59f3f690f5/1.15.1-pre1.json",
		"time": "2021-07-07T13:09:26+00:00",
		"releaseTime": "2019-12-12T14:02:30+00:00",
		"sha1": "4c7b181580f9f208236609d787baec59f3f690f5",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/dd52ed8176d69392eb889b54906356aeb668afbc/1.15.json",
		"time": "2021-07-07T13:08:42+00:00",
		"releaseTime": "2019-12-09T13:13:38+00:00",
		"sha1": "dd52ed8176d69392eb889b54906356aeb668afbc",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15-pre7",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d395e1efbfbd417f11e65dcea188b1ab4c06e830/1.15-pre7.json",
		"time": "2021-07-07T13:09:26+00:00",
		"releaseTime": "2019-12-09T12:14:11+00:00",
		"sha1": "d395e1efbfbd417f11e65dcea188b1ab4c06e830",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15-pre6",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/283e029b566a187aebec52114a27ee23dbe5126e/1.15-pre6.json",
		"time": "2021-07-07T13:09:25+00:00",
		"releaseTime": "2019-12-06T12:04:30+00:00",
		"sha1": "283e029b566a187aebec52114a27ee23dbe5126e",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15-pre5",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/28704c240e4b79da67bb6018cc816604df8ee4d1/1.15-pre5.json",
		"time": "2021-07-07T13:09:24+00:00",
		"releaseTime": "2019-12-05T13:20:00+00:00",
		"sha1": "28704c240e4b79da67bb6018cc816604df8ee4d1",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/113df5ae705f444465444359b9d0f650a7661e1e/1.15-pre4.json",
		"time": "2021-07-07T13:09:24+00:00",
		"releaseTime": "2019-12-03T12:24:24+00:00",
		"sha1": "113df5ae705f444465444359b9d0f650a7661e1e",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/239f0bc3df2a383c6a495e932c723be20aa2eaa4/1.15-pre3.json",
		"time": "2021-07-07T13:09:23+00:00",
		"releaseTime": "2019-11-28T17:17:50+00:00",
		"sha1": "239f0bc3df2a383c6a495e932c723be20aa2eaa4",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/cdd19150bbdff90ebbdd31dfc83742be3087c595/1.15-pre2.json",
		"time": "2021-07-07T13:09:22+00:00",
		"releaseTime": "2019-11-25T18:09:38+00:00",
		"sha1": "cdd19150bbdff90ebbdd31dfc83742be3087c595",
		"complianceLevel": 0
	  },
	  {
		"id": "1.15-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/be9d2c07b2cea26f99a55be6779e40325bc1741d/1.15-pre1.json",
		"time": "2021-07-07T13:09:22+00:00",
		"releaseTime": "2019-11-21T17:01:17+00:00",
		"sha1": "be9d2c07b2cea26f99a55be6779e40325bc1741d",
		"complianceLevel": 0
	  },
	  {
		"id": "19w46b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/296edef21650f132ab9a9a780a536b80bc8a8bf0/19w46b.json",
		"time": "2021-07-07T13:11:48+00:00",
		"releaseTime": "2019-11-14T13:29:24+00:00",
		"sha1": "296edef21650f132ab9a9a780a536b80bc8a8bf0",
		"complianceLevel": 0
	  },
	  {
		"id": "19w46a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/42c92811b04add3880d1f8ac98e69f3519ec7309/19w46a.json",
		"time": "2021-07-07T13:11:48+00:00",
		"releaseTime": "2019-11-13T16:37:46+00:00",
		"sha1": "42c92811b04add3880d1f8ac98e69f3519ec7309",
		"complianceLevel": 0
	  },
	  {
		"id": "19w45b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/fabcfa02d3ebbb430030a5c7973517471d1f8b5a/19w45b.json",
		"time": "2021-07-07T13:11:47+00:00",
		"releaseTime": "2019-11-08T12:42:44+00:00",
		"sha1": "fabcfa02d3ebbb430030a5c7973517471d1f8b5a",
		"complianceLevel": 0
	  },
	  {
		"id": "19w45a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f8615b2610d4f915332134ce7435e0d83fc95c46/19w45a.json",
		"time": "2021-07-07T13:11:46+00:00",
		"releaseTime": "2019-11-07T16:19:20+00:00",
		"sha1": "f8615b2610d4f915332134ce7435e0d83fc95c46",
		"complianceLevel": 0
	  },
	  {
		"id": "19w44a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0de8c8a841c7ede2e7e1b0e6d7a0ff0e0d52b85e/19w44a.json",
		"time": "2021-07-07T13:11:46+00:00",
		"releaseTime": "2019-10-30T15:31:44+00:00",
		"sha1": "0de8c8a841c7ede2e7e1b0e6d7a0ff0e0d52b85e",
		"complianceLevel": 0
	  },
	  {
		"id": "19w42a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/70ade3ff4eaae1c7c9308b77f979b090cddaf5ed/19w42a.json",
		"time": "2021-07-07T13:11:45+00:00",
		"releaseTime": "2019-10-16T15:30:39+00:00",
		"sha1": "70ade3ff4eaae1c7c9308b77f979b090cddaf5ed",
		"complianceLevel": 0
	  },
	  {
		"id": "19w41a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e423b844864daaf7f4075cf24bf049a7dd27b935/19w41a.json",
		"time": "2021-07-07T13:11:45+00:00",
		"releaseTime": "2019-10-09T15:21:35+00:00",
		"sha1": "e423b844864daaf7f4075cf24bf049a7dd27b935",
		"complianceLevel": 0
	  },
	  {
		"id": "19w40a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0148b539032fd5d4da5f72267357fd51cb93c0aa/19w40a.json",
		"time": "2021-07-07T13:11:44+00:00",
		"releaseTime": "2019-10-02T13:40:26+00:00",
		"sha1": "0148b539032fd5d4da5f72267357fd51cb93c0aa",
		"complianceLevel": 0
	  },
	  {
		"id": "19w39a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/19c811eac76700c48c8bf75e2d7242d6635182c7/19w39a.json",
		"time": "2021-07-07T13:11:44+00:00",
		"releaseTime": "2019-09-27T10:13:33+00:00",
		"sha1": "19c811eac76700c48c8bf75e2d7242d6635182c7",
		"complianceLevel": 0
	  },
	  {
		"id": "19w38b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/41c515ca21b6d4ff5c8adce117ca4188ff8d0e36/19w38b.json",
		"time": "2021-07-07T13:11:43+00:00",
		"releaseTime": "2019-09-18T14:59:13+00:00",
		"sha1": "41c515ca21b6d4ff5c8adce117ca4188ff8d0e36",
		"complianceLevel": 0
	  },
	  {
		"id": "19w38a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0dbff7626ca3175cce42f3c8b6d1c32166e2c958/19w38a.json",
		"time": "2021-07-07T13:11:42+00:00",
		"releaseTime": "2019-09-18T10:03:22+00:00",
		"sha1": "0dbff7626ca3175cce42f3c8b6d1c32166e2c958",
		"complianceLevel": 0
	  },
	  {
		"id": "19w37a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2bbbbe8bd55dab21504dbe09e8d8316a2314c0f4/19w37a.json",
		"time": "2021-07-07T13:11:42+00:00",
		"releaseTime": "2019-09-11T11:46:44+00:00",
		"sha1": "2bbbbe8bd55dab21504dbe09e8d8316a2314c0f4",
		"complianceLevel": 0
	  },
	  {
		"id": "19w36a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f9258be8d01d335f64b8def8d5bf93dfa09ad3a8/19w36a.json",
		"time": "2021-07-07T13:11:41+00:00",
		"releaseTime": "2019-09-04T11:19:34+00:00",
		"sha1": "f9258be8d01d335f64b8def8d5bf93dfa09ad3a8",
		"complianceLevel": 0
	  },
	  {
		"id": "19w35a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/30532ff135a5eafcad7cb989832f146e501c4dd0/19w35a.json",
		"time": "2021-07-07T13:11:41+00:00",
		"releaseTime": "2019-08-28T15:01:44+00:00",
		"sha1": "30532ff135a5eafcad7cb989832f146e501c4dd0",
		"complianceLevel": 0
	  },
	  {
		"id": "19w34a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9507147ec064ace6f5f1d513c96ebbee3f5db293/19w34a.json",
		"time": "2021-07-07T13:11:40+00:00",
		"releaseTime": "2019-08-22T12:06:21+00:00",
		"sha1": "9507147ec064ace6f5f1d513c96ebbee3f5db293",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.4",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/d73c3f908365863ebe6b01cf454990182f2652f4/1.14.4.json",
		"time": "2021-07-07T13:08:40+00:00",
		"releaseTime": "2019-07-19T09:25:47+00:00",
		"sha1": "d73c3f908365863ebe6b01cf454990182f2652f4",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.4-pre7",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4b319f52da9dbce8ef4b278867596356491d48ed/1.14.4-pre7.json",
		"time": "2021-07-07T13:09:21+00:00",
		"releaseTime": "2019-07-18T11:32:36+00:00",
		"sha1": "4b319f52da9dbce8ef4b278867596356491d48ed",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.4-pre6",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/70f428c9ccbbab91b7c89e3295d6072b36a6e8bc/1.14.4-pre6.json",
		"time": "2021-07-07T13:09:21+00:00",
		"releaseTime": "2019-07-15T12:39:49+00:00",
		"sha1": "70f428c9ccbbab91b7c89e3295d6072b36a6e8bc",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.4-pre5",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2076594973c79009e6349fb4554ef3afe412991a/1.14.4-pre5.json",
		"time": "2021-07-07T13:09:20+00:00",
		"releaseTime": "2019-07-11T10:52:33+00:00",
		"sha1": "2076594973c79009e6349fb4554ef3afe412991a",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.4-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f5ed795625ebdd8b59d5f5ee7d117e01fc5344a7/1.14.4-pre4.json",
		"time": "2021-07-07T13:09:20+00:00",
		"releaseTime": "2019-07-10T12:53:29+00:00",
		"sha1": "f5ed795625ebdd8b59d5f5ee7d117e01fc5344a7",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.4-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7f9624cbb4a81d42ab6a0b362815cf5f11171477/1.14.4-pre3.json",
		"time": "2021-07-07T13:09:19+00:00",
		"releaseTime": "2019-07-08T11:21:42+00:00",
		"sha1": "7f9624cbb4a81d42ab6a0b362815cf5f11171477",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.4-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7355190b45e31e5800308dcecc48e865d4833616/1.14.4-pre2.json",
		"time": "2021-07-07T13:09:19+00:00",
		"releaseTime": "2019-07-04T14:41:05+00:00",
		"sha1": "7355190b45e31e5800308dcecc48e865d4833616",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.4-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/98bab11790a1e731e3cfca0c44c177c4f5c30764/1.14.4-pre1.json",
		"time": "2021-07-07T13:09:18+00:00",
		"releaseTime": "2019-07-03T13:01:01+00:00",
		"sha1": "98bab11790a1e731e3cfca0c44c177c4f5c30764",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.3",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/7e1082895822b14f2749a0d8963453a3fb6fe512/1.14.3.json",
		"time": "2021-07-07T13:08:39+00:00",
		"releaseTime": "2019-06-24T12:52:52+00:00",
		"sha1": "7e1082895822b14f2749a0d8963453a3fb6fe512",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.3-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4d1f994dce63e26a8f92d034dc0451c4c4c32b65/1.14.3-pre4.json",
		"time": "2021-07-07T13:09:18+00:00",
		"releaseTime": "2019-06-19T11:44:29+00:00",
		"sha1": "4d1f994dce63e26a8f92d034dc0451c4c4c32b65",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.3-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d5b41ac670958ac48ca063963191a2f66db34b35/1.14.3-pre3.json",
		"time": "2021-07-07T13:09:17+00:00",
		"releaseTime": "2019-06-14T08:03:33+00:00",
		"sha1": "d5b41ac670958ac48ca063963191a2f66db34b35",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.3-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d601fea8b718b207b8ae1a9eafcc0218e4476720/1.14.3-pre2.json",
		"time": "2021-07-07T13:09:17+00:00",
		"releaseTime": "2019-06-07T09:11:29+00:00",
		"sha1": "d601fea8b718b207b8ae1a9eafcc0218e4476720",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.3-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ae25c95ffbbf1ccd0fc21b35a504685a3091c99a/1.14.3-pre1.json",
		"time": "2021-07-07T13:09:16+00:00",
		"releaseTime": "2019-06-03T14:34:20+00:00",
		"sha1": "ae25c95ffbbf1ccd0fc21b35a504685a3091c99a",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/e843715dc2d57d9d4f5ed3de1519df3c5534110c/1.14.2.json",
		"time": "2021-07-07T13:08:39+00:00",
		"releaseTime": "2019-05-27T11:48:25+00:00",
		"sha1": "e843715dc2d57d9d4f5ed3de1519df3c5534110c",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.2 Pre-Release 4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1d560cb1a100bcc14fe1be4a063add2e23783844/1.14.2%20Pre-Release%204.json",
		"time": "2021-07-07T13:09:16+00:00",
		"releaseTime": "2019-05-27T07:21:11+00:00",
		"sha1": "1d560cb1a100bcc14fe1be4a063add2e23783844",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.2 Pre-Release 3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7f33a966abc6dc701ca93cfeea1d7cd232ab6b0d/1.14.2%20Pre-Release%203.json",
		"time": "2021-07-07T13:09:15+00:00",
		"releaseTime": "2019-05-22T13:12:51+00:00",
		"sha1": "7f33a966abc6dc701ca93cfeea1d7cd232ab6b0d",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.2 Pre-Release 2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5b715998347860ed30045684167b5198a547eab1/1.14.2%20Pre-Release%202.json",
		"time": "2021-07-07T13:09:15+00:00",
		"releaseTime": "2019-05-17T12:21:03+00:00",
		"sha1": "5b715998347860ed30045684167b5198a547eab1",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.2 Pre-Release 1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/04800c34dc382fda1a96fe0e85f33acd41fd7156/1.14.2%20Pre-Release%201.json",
		"time": "2021-07-07T13:09:14+00:00",
		"releaseTime": "2019-05-16T15:40:25+00:00",
		"sha1": "04800c34dc382fda1a96fe0e85f33acd41fd7156",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/e31479f8f115a3da224494946786d8a8942822cc/1.14.1.json",
		"time": "2021-07-07T13:08:38+00:00",
		"releaseTime": "2019-05-13T11:10:12+00:00",
		"sha1": "e31479f8f115a3da224494946786d8a8942822cc",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.1 Pre-Release 2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4cca49c05d93dc2cba47db2a071129dc42c69493/1.14.1%20Pre-Release%202.json",
		"time": "2021-07-07T13:09:14+00:00",
		"releaseTime": "2019-05-09T14:01:04+00:00",
		"sha1": "4cca49c05d93dc2cba47db2a071129dc42c69493",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14.1 Pre-Release 1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/66e27aa238a8872ee13f80c2682303bfeeed65c0/1.14.1%20Pre-Release%201.json",
		"time": "2021-07-07T13:09:13+00:00",
		"releaseTime": "2019-05-07T14:44:42+00:00",
		"sha1": "66e27aa238a8872ee13f80c2682303bfeeed65c0",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/538fc13ebb63152dc0f417c9d54a09a2075a6f5d/1.14.json",
		"time": "2021-07-07T13:08:41+00:00",
		"releaseTime": "2019-04-23T14:52:44+00:00",
		"sha1": "538fc13ebb63152dc0f417c9d54a09a2075a6f5d",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14 Pre-Release 5",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/82a7b2c611ff3f2c6d607790d985d207fe3cf807/1.14%20Pre-Release%205.json",
		"time": "2021-07-07T13:09:13+00:00",
		"releaseTime": "2019-04-18T11:05:19+00:00",
		"sha1": "82a7b2c611ff3f2c6d607790d985d207fe3cf807",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14 Pre-Release 4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c245835e37a2f7ab449ecf29bdff193816b6cc5c/1.14%20Pre-Release%204.json",
		"time": "2021-07-07T13:09:12+00:00",
		"releaseTime": "2019-04-17T15:31:12+00:00",
		"sha1": "c245835e37a2f7ab449ecf29bdff193816b6cc5c",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14 Pre-Release 3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0305f465464bbe5c11d82ad4acc25a5025a403ff/1.14%20Pre-Release%203.json",
		"time": "2021-07-07T13:09:12+00:00",
		"releaseTime": "2019-04-16T13:57:10+00:00",
		"sha1": "0305f465464bbe5c11d82ad4acc25a5025a403ff",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14 Pre-Release 2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5cb67b8ff82181395d72784aad4813ce42a30f66/1.14%20Pre-Release%202.json",
		"time": "2021-07-07T13:09:11+00:00",
		"releaseTime": "2019-04-12T11:38:53+00:00",
		"sha1": "5cb67b8ff82181395d72784aad4813ce42a30f66",
		"complianceLevel": 0
	  },
	  {
		"id": "1.14 Pre-Release 1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d83c2b29cbf0b1fcc0ca5f0896656d7fbf4b8b37/1.14%20Pre-Release%201.json",
		"time": "2021-07-07T13:09:11+00:00",
		"releaseTime": "2019-04-10T14:24:16+00:00",
		"sha1": "d83c2b29cbf0b1fcc0ca5f0896656d7fbf4b8b37",
		"complianceLevel": 0
	  },
	  {
		"id": "19w14b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/98a89aafe6f0a52400cacb1e5402b3a57094a0f4/19w14b.json",
		"time": "2021-07-07T13:11:39+00:00",
		"releaseTime": "2019-04-05T10:33:58+00:00",
		"sha1": "98a89aafe6f0a52400cacb1e5402b3a57094a0f4",
		"complianceLevel": 0
	  },
	  {
		"id": "19w14a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f78deee6fa3004ade068f2573289dbe99370683f/19w14a.json",
		"time": "2021-07-07T13:11:39+00:00",
		"releaseTime": "2019-04-03T13:45:00+00:00",
		"sha1": "f78deee6fa3004ade068f2573289dbe99370683f",
		"complianceLevel": 0
	  },
	  {
		"id": "3D Shareware v1.34",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5c38ac67b627b0b57bd4dc38da74b4f428615acd/3D%20Shareware%20v1.34.json",
		"time": "2021-07-07T13:12:13+00:00",
		"releaseTime": "2019-04-01T11:18:08+00:00",
		"sha1": "5c38ac67b627b0b57bd4dc38da74b4f428615acd",
		"complianceLevel": 0
	  },
	  {
		"id": "19w13b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5ac5f0eb802947144e64c6834d1a746c0e18c72b/19w13b.json",
		"time": "2021-07-07T13:11:38+00:00",
		"releaseTime": "2019-03-29T16:53:22+00:00",
		"sha1": "5ac5f0eb802947144e64c6834d1a746c0e18c72b",
		"complianceLevel": 0
	  },
	  {
		"id": "19w13a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b132dc5c1a5d095cd6db5bab3a22981159b54ee3/19w13a.json",
		"time": "2021-07-07T13:11:38+00:00",
		"releaseTime": "2019-03-27T15:15:31+00:00",
		"sha1": "b132dc5c1a5d095cd6db5bab3a22981159b54ee3",
		"complianceLevel": 0
	  },
	  {
		"id": "19w12b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/705dd4dc770248678edf0d9acd31165bf782ec6d/19w12b.json",
		"time": "2021-07-07T13:11:37+00:00",
		"releaseTime": "2019-03-21T15:20:01+00:00",
		"sha1": "705dd4dc770248678edf0d9acd31165bf782ec6d",
		"complianceLevel": 0
	  },
	  {
		"id": "19w12a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/208229690acced4bf828fb251bbb3b4d89eb9ea2/19w12a.json",
		"time": "2021-07-07T13:11:37+00:00",
		"releaseTime": "2019-03-20T16:47:34+00:00",
		"sha1": "208229690acced4bf828fb251bbb3b4d89eb9ea2",
		"complianceLevel": 0
	  },
	  {
		"id": "19w11b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/fbfbe459e96ab501b56d68956f5c4dbd9814ed44/19w11b.json",
		"time": "2021-07-07T13:11:36+00:00",
		"releaseTime": "2019-03-14T14:26:23+00:00",
		"sha1": "fbfbe459e96ab501b56d68956f5c4dbd9814ed44",
		"complianceLevel": 0
	  },
	  {
		"id": "19w11a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4cc78e6b6c40f05aa80644d6120eeb0750c9318d/19w11a.json",
		"time": "2021-07-07T13:11:36+00:00",
		"releaseTime": "2019-03-13T13:59:29+00:00",
		"sha1": "4cc78e6b6c40f05aa80644d6120eeb0750c9318d",
		"complianceLevel": 0
	  },
	  {
		"id": "19w09a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5350317c1a1752eb5076a335d5911fad608c4f51/19w09a.json",
		"time": "2021-07-07T13:11:35+00:00",
		"releaseTime": "2019-02-27T14:44:30+00:00",
		"sha1": "5350317c1a1752eb5076a335d5911fad608c4f51",
		"complianceLevel": 0
	  },
	  {
		"id": "19w08b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e5db615cf56ec18d779e67b9cf3a70da5a4d4cce/19w08b.json",
		"time": "2021-07-07T13:11:35+00:00",
		"releaseTime": "2019-02-21T13:38:09+00:00",
		"sha1": "e5db615cf56ec18d779e67b9cf3a70da5a4d4cce",
		"complianceLevel": 0
	  },
	  {
		"id": "19w08a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a1e2c155d36855d233575f40153b25eb23bd4e06/19w08a.json",
		"time": "2021-07-07T13:11:34+00:00",
		"releaseTime": "2019-02-20T14:56:58+00:00",
		"sha1": "a1e2c155d36855d233575f40153b25eb23bd4e06",
		"complianceLevel": 0
	  },
	  {
		"id": "19w07a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/bfa3b3189eec5fd75f3b948ea25a0e87265f9f4b/19w07a.json",
		"time": "2021-07-07T13:11:34+00:00",
		"releaseTime": "2019-02-13T16:12:08+00:00",
		"sha1": "bfa3b3189eec5fd75f3b948ea25a0e87265f9f4b",
		"complianceLevel": 0
	  },
	  {
		"id": "19w06a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/375bbfbaaab10bf0f1289883c98654e606587b2e/19w06a.json",
		"time": "2021-07-07T13:11:33+00:00",
		"releaseTime": "2019-02-06T16:24:13+00:00",
		"sha1": "375bbfbaaab10bf0f1289883c98654e606587b2e",
		"complianceLevel": 0
	  },
	  {
		"id": "19w05a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/8f0dc7be1921138cca5c0f253193f63d79b31c04/19w05a.json",
		"time": "2021-07-07T13:11:33+00:00",
		"releaseTime": "2019-01-30T15:16:49+00:00",
		"sha1": "8f0dc7be1921138cca5c0f253193f63d79b31c04",
		"complianceLevel": 0
	  },
	  {
		"id": "19w04b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/bbf134b124fa2b9c935be8f85ed5bf5d6d0c4e5e/19w04b.json",
		"time": "2021-07-07T13:11:32+00:00",
		"releaseTime": "2019-01-25T12:20:15+00:00",
		"sha1": "bbf134b124fa2b9c935be8f85ed5bf5d6d0c4e5e",
		"complianceLevel": 0
	  },
	  {
		"id": "19w04a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/83de406b28e8e6dcef9c51d55cb5e11553f6e6bb/19w04a.json",
		"time": "2021-07-07T13:11:32+00:00",
		"releaseTime": "2019-01-24T15:31:52+00:00",
		"sha1": "83de406b28e8e6dcef9c51d55cb5e11553f6e6bb",
		"complianceLevel": 0
	  },
	  {
		"id": "19w03c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4268e7ef1776c6f897726020de2be4b12a436510/19w03c.json",
		"time": "2021-07-07T13:11:31+00:00",
		"releaseTime": "2019-01-18T11:27:13+00:00",
		"sha1": "4268e7ef1776c6f897726020de2be4b12a436510",
		"complianceLevel": 0
	  },
	  {
		"id": "19w03b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/3b58416b0415e5eaaf60f3fe7d043330a2cedb37/19w03b.json",
		"time": "2021-07-07T13:11:31+00:00",
		"releaseTime": "2019-01-17T16:43:27+00:00",
		"sha1": "3b58416b0415e5eaaf60f3fe7d043330a2cedb37",
		"complianceLevel": 0
	  },
	  {
		"id": "19w03a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4bb8ec5d7d12775db48283e8a2c59d7288b46745/19w03a.json",
		"time": "2021-07-07T13:11:30+00:00",
		"releaseTime": "2019-01-16T16:45:02+00:00",
		"sha1": "4bb8ec5d7d12775db48283e8a2c59d7288b46745",
		"complianceLevel": 0
	  },
	  {
		"id": "19w02a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/00bf58b0f625ea059ba8df1fe4604c22402faf4d/19w02a.json",
		"time": "2021-07-07T13:11:30+00:00",
		"releaseTime": "2019-01-09T15:52:07+00:00",
		"sha1": "00bf58b0f625ea059ba8df1fe4604c22402faf4d",
		"complianceLevel": 0
	  },
	  {
		"id": "18w50a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/42fc0fe030d081a9f93806a02724fce7473a6d76/18w50a.json",
		"time": "2021-07-07T13:11:29+00:00",
		"releaseTime": "2018-12-12T14:58:13+00:00",
		"sha1": "42fc0fe030d081a9f93806a02724fce7473a6d76",
		"complianceLevel": 0
	  },
	  {
		"id": "18w49a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9048273f9ed374bd8f3c81597f0905324fff3b36/18w49a.json",
		"time": "2021-07-07T13:11:29+00:00",
		"releaseTime": "2018-12-05T12:24:30+00:00",
		"sha1": "9048273f9ed374bd8f3c81597f0905324fff3b36",
		"complianceLevel": 0
	  },
	  {
		"id": "18w48b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/fefadb53447b45e1d6ae03b384f589cdd29a3c43/18w48b.json",
		"time": "2021-07-07T13:11:28+00:00",
		"releaseTime": "2018-11-30T10:37:31+00:00",
		"sha1": "fefadb53447b45e1d6ae03b384f589cdd29a3c43",
		"complianceLevel": 0
	  },
	  {
		"id": "18w48a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ef012cba0b843ada41086b674a304ec1b1a2bb7f/18w48a.json",
		"time": "2021-07-07T13:11:28+00:00",
		"releaseTime": "2018-11-29T13:11:38+00:00",
		"sha1": "ef012cba0b843ada41086b674a304ec1b1a2bb7f",
		"complianceLevel": 0
	  },
	  {
		"id": "18w47b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/36b346ee76bfede2f46928304d256526bea46bf9/18w47b.json",
		"time": "2021-07-07T13:11:27+00:00",
		"releaseTime": "2018-11-23T10:46:41+00:00",
		"sha1": "36b346ee76bfede2f46928304d256526bea46bf9",
		"complianceLevel": 0
	  },
	  {
		"id": "18w47a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1f3f376ff16032b41aad661a2d67bc69052e84e5/18w47a.json",
		"time": "2021-07-07T13:11:27+00:00",
		"releaseTime": "2018-11-21T15:45:22+00:00",
		"sha1": "1f3f376ff16032b41aad661a2d67bc69052e84e5",
		"complianceLevel": 0
	  },
	  {
		"id": "18w46a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6c689bd954ea03dd3562dac34e4d6605789dd61c/18w46a.json",
		"time": "2021-07-07T13:11:26+00:00",
		"releaseTime": "2018-11-15T13:43:14+00:00",
		"sha1": "6c689bd954ea03dd3562dac34e4d6605789dd61c",
		"complianceLevel": 0
	  },
	  {
		"id": "18w45a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c66f3d11f741f8d70dfdc13fedac40c3ac3e346a/18w45a.json",
		"time": "2021-07-07T13:11:26+00:00",
		"releaseTime": "2018-11-07T14:40:06+00:00",
		"sha1": "c66f3d11f741f8d70dfdc13fedac40c3ac3e346a",
		"complianceLevel": 0
	  },
	  {
		"id": "18w44a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e1d7ddfcffd550eadda8a3f3785758a696923489/18w44a.json",
		"time": "2021-07-07T13:11:25+00:00",
		"releaseTime": "2018-10-31T15:29:16+00:00",
		"sha1": "e1d7ddfcffd550eadda8a3f3785758a696923489",
		"complianceLevel": 0
	  },
	  {
		"id": "18w43c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9fb13023eac9cc512d9aec4caa54627622ed061a/18w43c.json",
		"time": "2021-07-07T13:11:24+00:00",
		"releaseTime": "2018-10-26T08:40:46+00:00",
		"sha1": "9fb13023eac9cc512d9aec4caa54627622ed061a",
		"complianceLevel": 0
	  },
	  {
		"id": "18w43b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/062a46b6a2f74d4cf5d8b71c2026b87e0c1b474b/18w43b.json",
		"time": "2021-07-07T13:11:24+00:00",
		"releaseTime": "2018-10-24T15:02:30+00:00",
		"sha1": "062a46b6a2f74d4cf5d8b71c2026b87e0c1b474b",
		"complianceLevel": 0
	  },
	  {
		"id": "18w43a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/279cc894b6e1d5f39f5628c5c60aaf6048ff9d88/18w43a.json",
		"time": "2021-07-07T13:11:23+00:00",
		"releaseTime": "2018-10-24T10:52:16+00:00",
		"sha1": "279cc894b6e1d5f39f5628c5c60aaf6048ff9d88",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/57f0285f5e6800233e3269a93ad11bfb631f2412/1.13.2.json",
		"time": "2021-07-07T13:08:37+00:00",
		"releaseTime": "2018-10-22T11:41:07+00:00",
		"sha1": "57f0285f5e6800233e3269a93ad11bfb631f2412",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13.2-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d098582233b0c5d1e985b78a13328191c7039ba3/1.13.2-pre2.json",
		"time": "2021-07-07T13:09:10+00:00",
		"releaseTime": "2018-10-18T14:46:12+00:00",
		"sha1": "d098582233b0c5d1e985b78a13328191c7039ba3",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13.2-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d04db44894a2bb44ec9d69ec4fb4e8e7d1280d20/1.13.2-pre1.json",
		"time": "2021-07-07T13:09:10+00:00",
		"releaseTime": "2018-10-16T13:40:58+00:00",
		"sha1": "d04db44894a2bb44ec9d69ec4fb4e8e7d1280d20",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/f63ffba80c4023981d9340392808d63d053f6172/1.13.1.json",
		"time": "2021-07-07T13:08:36+00:00",
		"releaseTime": "2018-08-22T14:03:42+00:00",
		"sha1": "f63ffba80c4023981d9340392808d63d053f6172",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13.1-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/fb40e3ada404ead18c6bbf1828de0e22121f8974/1.13.1-pre2.json",
		"time": "2021-07-07T13:09:09+00:00",
		"releaseTime": "2018-08-20T13:52:09+00:00",
		"sha1": "fb40e3ada404ead18c6bbf1828de0e22121f8974",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13.1-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/956b35cee969c3b532a0ca78d57fb8adedd40f17/1.13.1-pre1.json",
		"time": "2021-07-07T13:09:09+00:00",
		"releaseTime": "2018-08-16T13:08:44+00:00",
		"sha1": "956b35cee969c3b532a0ca78d57fb8adedd40f17",
		"complianceLevel": 0
	  },
	  {
		"id": "18w33a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/81af99b89230f35224e6e4fb0b36a5b07736e8dd/18w33a.json",
		"time": "2021-07-07T13:11:23+00:00",
		"releaseTime": "2018-08-15T14:28:56+00:00",
		"sha1": "81af99b89230f35224e6e4fb0b36a5b07736e8dd",
		"complianceLevel": 0
	  },
	  {
		"id": "18w32a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5af7fcbff19c21a1d15aea130ea4f540e5b62f96/18w32a.json",
		"time": "2021-07-07T13:11:22+00:00",
		"releaseTime": "2018-08-08T13:16:57+00:00",
		"sha1": "5af7fcbff19c21a1d15aea130ea4f540e5b62f96",
		"complianceLevel": 0
	  },
	  {
		"id": "18w31a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5a1b3ac24e2bb0abfed814dc74260e7250133e53/18w31a.json",
		"time": "2021-07-07T13:11:22+00:00",
		"releaseTime": "2018-08-01T12:54:44+00:00",
		"sha1": "5a1b3ac24e2bb0abfed814dc74260e7250133e53",
		"complianceLevel": 0
	  },
	  {
		"id": "18w30b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/93999fab24ebb3f53337623da6ec7d0b2a11577e/18w30b.json",
		"time": "2021-07-07T13:11:21+00:00",
		"releaseTime": "2018-07-26T16:06:57+00:00",
		"sha1": "93999fab24ebb3f53337623da6ec7d0b2a11577e",
		"complianceLevel": 0
	  },
	  {
		"id": "18w30a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9db7fd88150fafc42e08df36917f46700c0d39f3/18w30a.json",
		"time": "2021-07-07T13:11:21+00:00",
		"releaseTime": "2018-07-25T14:29:31+00:00",
		"sha1": "9db7fd88150fafc42e08df36917f46700c0d39f3",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/81db15e7c0e5e8f951b43f7bac9e4cdf711e983b/1.13.json",
		"time": "2021-07-07T13:08:38+00:00",
		"releaseTime": "2018-07-18T15:11:46+00:00",
		"sha1": "81db15e7c0e5e8f951b43f7bac9e4cdf711e983b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13-pre10",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/3c450152c554d6759157a9ce6ffadec9fb95e24c/1.13-pre10.json",
		"time": "2021-07-07T13:09:04+00:00",
		"releaseTime": "2018-07-17T14:48:06+00:00",
		"sha1": "3c450152c554d6759157a9ce6ffadec9fb95e24c",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13-pre9",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d1c7b6b70eec39512762ffb5073b69f7049bdfb1/1.13-pre9.json",
		"time": "2021-07-07T13:09:08+00:00",
		"releaseTime": "2018-07-16T14:17:42+00:00",
		"sha1": "d1c7b6b70eec39512762ffb5073b69f7049bdfb1",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13-pre8",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0ce1256e9dc58f664ebb7d75ef6f6232a784f1d0/1.13-pre8.json",
		"time": "2021-07-07T13:09:08+00:00",
		"releaseTime": "2018-07-13T11:45:00+00:00",
		"sha1": "0ce1256e9dc58f664ebb7d75ef6f6232a784f1d0",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13-pre7",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/3138bedca1829063c7544d11fba574f345ae8a37/1.13-pre7.json",
		"time": "2021-07-07T13:09:07+00:00",
		"releaseTime": "2018-07-10T14:21:42+00:00",
		"sha1": "3138bedca1829063c7544d11fba574f345ae8a37",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13-pre6",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/bf915899cc9446dbb0a215cf9270080617885107/1.13-pre6.json",
		"time": "2021-07-07T13:09:07+00:00",
		"releaseTime": "2018-07-04T12:36:00+00:00",
		"sha1": "bf915899cc9446dbb0a215cf9270080617885107",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13-pre5",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/df53ac0adaeea2456d22aa1af44d8610ff922b4d/1.13-pre5.json",
		"time": "2021-07-07T13:09:06+00:00",
		"releaseTime": "2018-06-28T13:58:53+00:00",
		"sha1": "df53ac0adaeea2456d22aa1af44d8610ff922b4d",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6f73cca640aa46104c89a57d8bd9e7ef719881a7/1.13-pre4.json",
		"time": "2021-07-07T13:09:06+00:00",
		"releaseTime": "2018-06-26T13:00:55+00:00",
		"sha1": "6f73cca640aa46104c89a57d8bd9e7ef719881a7",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/79560dc9500675e54582f977f658fb46e73b4fcf/1.13-pre3.json",
		"time": "2021-07-07T13:09:05+00:00",
		"releaseTime": "2018-06-21T12:57:11+00:00",
		"sha1": "79560dc9500675e54582f977f658fb46e73b4fcf",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/504adb6a0ae17c7dc4e85f8230f445d0af73729f/1.13-pre2.json",
		"time": "2021-07-07T13:09:05+00:00",
		"releaseTime": "2018-06-15T09:20:00+00:00",
		"sha1": "504adb6a0ae17c7dc4e85f8230f445d0af73729f",
		"complianceLevel": 0
	  },
	  {
		"id": "1.13-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/acb195539ae1d2a2ccddbbbb3382ca1f87735e3f/1.13-pre1.json",
		"time": "2021-07-07T13:09:04+00:00",
		"releaseTime": "2018-06-04T15:17:34+00:00",
		"sha1": "acb195539ae1d2a2ccddbbbb3382ca1f87735e3f",
		"complianceLevel": 0
	  },
	  {
		"id": "18w22c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/de1f6ed3defaf22bf1b27557f0c18dd1b2fbe55b/18w22c.json",
		"time": "2021-07-07T13:11:20+00:00",
		"releaseTime": "2018-05-31T13:53:15+00:00",
		"sha1": "de1f6ed3defaf22bf1b27557f0c18dd1b2fbe55b",
		"complianceLevel": 0
	  },
	  {
		"id": "18w22b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a1b9c5892cf052f3de6fb7b676c871c65c8b6d6d/18w22b.json",
		"time": "2021-07-07T13:11:20+00:00",
		"releaseTime": "2018-05-30T13:48:58+00:00",
		"sha1": "a1b9c5892cf052f3de6fb7b676c871c65c8b6d6d",
		"complianceLevel": 0
	  },
	  {
		"id": "18w22a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/95cd639784753702b0eb5dc9b1ee863f2ccf23b5/18w22a.json",
		"time": "2021-07-07T13:11:19+00:00",
		"releaseTime": "2018-05-29T13:23:55+00:00",
		"sha1": "95cd639784753702b0eb5dc9b1ee863f2ccf23b5",
		"complianceLevel": 0
	  },
	  {
		"id": "18w21b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7847b64c33796d23685ee143a7b647adbb2f285f/18w21b.json",
		"time": "2021-07-07T13:11:19+00:00",
		"releaseTime": "2018-05-25T10:09:09+00:00",
		"sha1": "7847b64c33796d23685ee143a7b647adbb2f285f",
		"complianceLevel": 0
	  },
	  {
		"id": "18w21a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/130e667a0ee5faa740f9c163945e9f460326ac79/18w21a.json",
		"time": "2021-07-07T13:11:18+00:00",
		"releaseTime": "2018-05-23T13:11:49+00:00",
		"sha1": "130e667a0ee5faa740f9c163945e9f460326ac79",
		"complianceLevel": 0
	  },
	  {
		"id": "18w20c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/773b53bf61ab96f7f587dc27ed172bf79bdf5760/18w20c.json",
		"time": "2021-07-07T13:11:18+00:00",
		"releaseTime": "2018-05-17T14:06:56+00:00",
		"sha1": "773b53bf61ab96f7f587dc27ed172bf79bdf5760",
		"complianceLevel": 0
	  },
	  {
		"id": "18w20b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/fa5b7a9613e6f300e0fa57b45e3762ad478c1caa/18w20b.json",
		"time": "2021-07-07T13:11:17+00:00",
		"releaseTime": "2018-05-16T14:35:35+00:00",
		"sha1": "fa5b7a9613e6f300e0fa57b45e3762ad478c1caa",
		"complianceLevel": 0
	  },
	  {
		"id": "18w20a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/acbd1c1ff90645f984ef42792a785d5f27fceccb/18w20a.json",
		"time": "2021-07-07T13:11:17+00:00",
		"releaseTime": "2018-05-15T14:02:25+00:00",
		"sha1": "acbd1c1ff90645f984ef42792a785d5f27fceccb",
		"complianceLevel": 0
	  },
	  {
		"id": "18w19b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/203e8329ba19a271ec5d6ab21aafaecd972b32ca/18w19b.json",
		"time": "2021-07-07T13:11:17+00:00",
		"releaseTime": "2018-05-09T10:00:51+00:00",
		"sha1": "203e8329ba19a271ec5d6ab21aafaecd972b32ca",
		"complianceLevel": 0
	  },
	  {
		"id": "18w19a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/3b5b82d46042f0d46d531e834e7d45e96388b3ce/18w19a.json",
		"time": "2021-07-07T13:11:16+00:00",
		"releaseTime": "2018-05-08T13:05:19+00:00",
		"sha1": "3b5b82d46042f0d46d531e834e7d45e96388b3ce",
		"complianceLevel": 0
	  },
	  {
		"id": "18w16a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c7bae03ea19a46a0f6eb5c9700b79b2b10c994b1/18w16a.json",
		"time": "2021-07-07T13:11:16+00:00",
		"releaseTime": "2018-04-19T14:46:35+00:00",
		"sha1": "c7bae03ea19a46a0f6eb5c9700b79b2b10c994b1",
		"complianceLevel": 0
	  },
	  {
		"id": "18w15a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f1b602e4d3bae10f219b7c3d553743b13c7de583/18w15a.json",
		"time": "2021-07-07T13:11:15+00:00",
		"releaseTime": "2018-04-11T14:54:22+00:00",
		"sha1": "f1b602e4d3bae10f219b7c3d553743b13c7de583",
		"complianceLevel": 0
	  },
	  {
		"id": "18w14b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b7ebaebacb95b90ab0eec0e9fb8519eb90d44440/18w14b.json",
		"time": "2021-07-07T13:11:15+00:00",
		"releaseTime": "2018-04-05T14:44:02+00:00",
		"sha1": "b7ebaebacb95b90ab0eec0e9fb8519eb90d44440",
		"complianceLevel": 0
	  },
	  {
		"id": "18w14a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6bbd59f49ff55ab4514a90f92728c4e68df4eb88/18w14a.json",
		"time": "2021-07-07T13:11:14+00:00",
		"releaseTime": "2018-04-04T14:36:14+00:00",
		"sha1": "6bbd59f49ff55ab4514a90f92728c4e68df4eb88",
		"complianceLevel": 0
	  },
	  {
		"id": "18w11a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/695d5fdb1c1d7890b2a01eeaadb3a07646744da3/18w11a.json",
		"time": "2021-07-07T13:11:14+00:00",
		"releaseTime": "2018-03-13T15:10:59+00:00",
		"sha1": "695d5fdb1c1d7890b2a01eeaadb3a07646744da3",
		"complianceLevel": 0
	  },
	  {
		"id": "18w10d",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/adfd6027307cb41a503c9f82cce5dfe4f2d5094a/18w10d.json",
		"time": "2021-07-07T13:11:13+00:00",
		"releaseTime": "2018-03-09T15:19:12+00:00",
		"sha1": "adfd6027307cb41a503c9f82cce5dfe4f2d5094a",
		"complianceLevel": 0
	  },
	  {
		"id": "18w10c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b0983f6f3ae01eaade43605470c64e5b9f1a23be/18w10c.json",
		"time": "2021-07-07T13:11:13+00:00",
		"releaseTime": "2018-03-08T15:29:23+00:00",
		"sha1": "b0983f6f3ae01eaade43605470c64e5b9f1a23be",
		"complianceLevel": 0
	  },
	  {
		"id": "18w10b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/669bcb31f3cccbeb15760bacb19943b02050153d/18w10b.json",
		"time": "2021-07-07T13:11:12+00:00",
		"releaseTime": "2018-03-07T15:56:01+00:00",
		"sha1": "669bcb31f3cccbeb15760bacb19943b02050153d",
		"complianceLevel": 0
	  },
	  {
		"id": "18w10a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/bcdd63501adab44435801cb2398a87a9513713e9/18w10a.json",
		"time": "2021-07-07T13:11:12+00:00",
		"releaseTime": "2018-03-06T15:54:24+00:00",
		"sha1": "bcdd63501adab44435801cb2398a87a9513713e9",
		"complianceLevel": 0
	  },
	  {
		"id": "18w09a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f36e20f54c8e60d2806c48b64c36e31004fa32a6/18w09a.json",
		"time": "2021-07-07T13:11:11+00:00",
		"releaseTime": "2018-03-01T14:15:10+00:00",
		"sha1": "f36e20f54c8e60d2806c48b64c36e31004fa32a6",
		"complianceLevel": 0
	  },
	  {
		"id": "18w08b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5e66f9ecad0bfd8dbb5e6a16b38355c755acbc63/18w08b.json",
		"time": "2021-07-07T13:11:11+00:00",
		"releaseTime": "2018-02-22T15:44:49+00:00",
		"sha1": "5e66f9ecad0bfd8dbb5e6a16b38355c755acbc63",
		"complianceLevel": 0
	  },
	  {
		"id": "18w08a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e0bb7cfb07d4cca7ed446952240ee3e0cc622e07/18w08a.json",
		"time": "2021-07-07T13:11:10+00:00",
		"releaseTime": "2018-02-21T14:59:00+00:00",
		"sha1": "e0bb7cfb07d4cca7ed446952240ee3e0cc622e07",
		"complianceLevel": 0
	  },
	  {
		"id": "18w07c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/164d5024816924443fe2c2bce9e44c20127f5b69/18w07c.json",
		"time": "2021-07-07T13:11:09+00:00",
		"releaseTime": "2018-02-16T13:23:32+00:00",
		"sha1": "164d5024816924443fe2c2bce9e44c20127f5b69",
		"complianceLevel": 0
	  },
	  {
		"id": "18w07b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e42afc54f3ee1e3d0dd02dab93e3c2ccc7296e45/18w07b.json",
		"time": "2021-07-07T13:11:09+00:00",
		"releaseTime": "2018-02-15T14:28:42+00:00",
		"sha1": "e42afc54f3ee1e3d0dd02dab93e3c2ccc7296e45",
		"complianceLevel": 0
	  },
	  {
		"id": "18w07a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/fe51edc1f13fec22e4602f057e1adeade9fac5e2/18w07a.json",
		"time": "2021-07-07T13:11:09+00:00",
		"releaseTime": "2018-02-14T17:34:13+00:00",
		"sha1": "fe51edc1f13fec22e4602f057e1adeade9fac5e2",
		"complianceLevel": 0
	  },
	  {
		"id": "18w06a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/8495721ade9b5d063c49d062d1f33846fc203d47/18w06a.json",
		"time": "2021-07-07T13:11:08+00:00",
		"releaseTime": "2018-02-09T12:09:55+00:00",
		"sha1": "8495721ade9b5d063c49d062d1f33846fc203d47",
		"complianceLevel": 0
	  },
	  {
		"id": "18w05a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/13874cddb59f225b24ef2cb1f7456a8fc7a17548/18w05a.json",
		"time": "2021-07-07T13:11:08+00:00",
		"releaseTime": "2018-01-31T13:32:09+00:00",
		"sha1": "13874cddb59f225b24ef2cb1f7456a8fc7a17548",
		"complianceLevel": 0
	  },
	  {
		"id": "18w03b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f8d78864b19b03d18224174f0be7c75873a3f3c0/18w03b.json",
		"time": "2021-07-07T13:11:07+00:00",
		"releaseTime": "2018-01-17T15:09:14+00:00",
		"sha1": "f8d78864b19b03d18224174f0be7c75873a3f3c0",
		"complianceLevel": 0
	  },
	  {
		"id": "18w03a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/dfed5fdfd653416520756fb2da6a79f0f7139c8b/18w03a.json",
		"time": "2021-07-07T13:11:07+00:00",
		"releaseTime": "2018-01-17T14:25:24+00:00",
		"sha1": "dfed5fdfd653416520756fb2da6a79f0f7139c8b",
		"complianceLevel": 0
	  },
	  {
		"id": "18w02a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b6e4df78e59b5f24c517d10b2cc54b5595b3f5ad/18w02a.json",
		"time": "2021-07-07T13:11:06+00:00",
		"releaseTime": "2018-01-10T11:54:55+00:00",
		"sha1": "b6e4df78e59b5f24c517d10b2cc54b5595b3f5ad",
		"complianceLevel": 0
	  },
	  {
		"id": "18w01a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7fc6a8d8a5fe83b29127bd26d4c67a62fbe42c5b/18w01a.json",
		"time": "2021-07-07T13:11:06+00:00",
		"releaseTime": "2018-01-03T13:29:30+00:00",
		"sha1": "7fc6a8d8a5fe83b29127bd26d4c67a62fbe42c5b",
		"complianceLevel": 0
	  },
	  {
		"id": "17w50a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b0eb6b51296f03bda24d56872870f729903aa5c6/17w50a.json",
		"time": "2021-07-07T13:11:05+00:00",
		"releaseTime": "2017-12-11T15:28:08+00:00",
		"sha1": "b0eb6b51296f03bda24d56872870f729903aa5c6",
		"complianceLevel": 0
	  },
	  {
		"id": "17w49b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7a64f441066143af72628f7175778d72b24c02a6/17w49b.json",
		"time": "2021-07-07T13:11:05+00:00",
		"releaseTime": "2017-12-07T15:29:54+00:00",
		"sha1": "7a64f441066143af72628f7175778d72b24c02a6",
		"complianceLevel": 0
	  },
	  {
		"id": "17w49a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2753f9a6861a6c70f4f9968e2e3f788437fe85f6/17w49a.json",
		"time": "2021-07-07T13:11:04+00:00",
		"releaseTime": "2017-12-06T14:24:30+00:00",
		"sha1": "2753f9a6861a6c70f4f9968e2e3f788437fe85f6",
		"complianceLevel": 0
	  },
	  {
		"id": "17w48a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/052c5d8edc42e5b5df4590f97d4d96030c31587b/17w48a.json",
		"time": "2021-07-07T13:11:04+00:00",
		"releaseTime": "2017-11-27T15:36:33+00:00",
		"sha1": "052c5d8edc42e5b5df4590f97d4d96030c31587b",
		"complianceLevel": 0
	  },
	  {
		"id": "17w47b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4320f9becf22167f57b00f277291739854c1132c/17w47b.json",
		"time": "2021-07-07T13:11:03+00:00",
		"releaseTime": "2017-11-23T15:30:12+00:00",
		"sha1": "4320f9becf22167f57b00f277291739854c1132c",
		"complianceLevel": 0
	  },
	  {
		"id": "17w47a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/18dbabe55967a3e7cdbe2796f2eb2603bb311bf0/17w47a.json",
		"time": "2021-07-07T13:11:03+00:00",
		"releaseTime": "2017-11-22T12:40:05+00:00",
		"sha1": "18dbabe55967a3e7cdbe2796f2eb2603bb311bf0",
		"complianceLevel": 0
	  },
	  {
		"id": "17w46a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/8f4d70f866bc675c372f76b03877629834b6066c/17w46a.json",
		"time": "2021-07-07T13:11:02+00:00",
		"releaseTime": "2017-11-15T15:21:55+00:00",
		"sha1": "8f4d70f866bc675c372f76b03877629834b6066c",
		"complianceLevel": 0
	  },
	  {
		"id": "17w45b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/54b19b33e106254f7be4050893a9d290b042d4a7/17w45b.json",
		"time": "2021-07-07T13:11:02+00:00",
		"releaseTime": "2017-11-10T10:07:02+00:00",
		"sha1": "54b19b33e106254f7be4050893a9d290b042d4a7",
		"complianceLevel": 0
	  },
	  {
		"id": "17w45a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c4a1c4186d3f1236b0d2431332976cc95ba8d936/17w45a.json",
		"time": "2021-07-07T13:11:01+00:00",
		"releaseTime": "2017-11-08T15:48:00+00:00",
		"sha1": "c4a1c4186d3f1236b0d2431332976cc95ba8d936",
		"complianceLevel": 0
	  },
	  {
		"id": "17w43b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5c3b54cc86bdfaaae819806dbd1f3caa73c7c4cd/17w43b.json",
		"time": "2021-07-07T13:11:00+00:00",
		"releaseTime": "2017-10-26T13:36:22+00:00",
		"sha1": "5c3b54cc86bdfaaae819806dbd1f3caa73c7c4cd",
		"complianceLevel": 0
	  },
	  {
		"id": "17w43a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ca84ae86884375ea172f6206687d932c1a9903e8/17w43a.json",
		"time": "2021-07-07T13:11:00+00:00",
		"releaseTime": "2017-10-25T14:43:50+00:00",
		"sha1": "ca84ae86884375ea172f6206687d932c1a9903e8",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/f07e0f1228f79b9b04313fc5640cd952474ba6f5/1.12.2.json",
		"time": "2021-07-07T13:08:36+00:00",
		"releaseTime": "2017-09-18T08:39:46+00:00",
		"sha1": "f07e0f1228f79b9b04313fc5640cd952474ba6f5",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12.2-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/afb7552dc02e4cf17ce32aff4714bca7e44f293e/1.12.2-pre2.json",
		"time": "2021-07-07T13:09:03+00:00",
		"releaseTime": "2017-09-15T08:21:17+00:00",
		"sha1": "afb7552dc02e4cf17ce32aff4714bca7e44f293e",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12.2-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/289c6d4e2e9f9c43e24ce47306d6780e2ad08f75/1.12.2-pre1.json",
		"time": "2021-07-07T13:09:03+00:00",
		"releaseTime": "2017-09-13T13:33:31+00:00",
		"sha1": "289c6d4e2e9f9c43e24ce47306d6780e2ad08f75",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/b089855afef020d39f0d1b56018188d93a9c3d90/1.12.1.json",
		"time": "2021-07-07T13:08:35+00:00",
		"releaseTime": "2017-08-03T12:40:39+00:00",
		"sha1": "b089855afef020d39f0d1b56018188d93a9c3d90",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12.1-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1c9f59b588131bde49faaef2dc09c363effa6dd8/1.12.1-pre1.json",
		"time": "2021-07-07T13:09:02+00:00",
		"releaseTime": "2017-08-02T10:53:55+00:00",
		"sha1": "1c9f59b588131bde49faaef2dc09c363effa6dd8",
		"complianceLevel": 0
	  },
	  {
		"id": "17w31a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/50f4593c20967d2090313b29c2e69adc44162feb/17w31a.json",
		"time": "2021-07-07T13:10:59+00:00",
		"releaseTime": "2017-08-01T09:41:23+00:00",
		"sha1": "50f4593c20967d2090313b29c2e69adc44162feb",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/7b4a07f16f9a309eac0e88425511d62ba52b5274/1.12.json",
		"time": "2021-07-07T13:08:36+00:00",
		"releaseTime": "2017-06-02T13:50:27+00:00",
		"sha1": "7b4a07f16f9a309eac0e88425511d62ba52b5274",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12-pre7",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c9c584f0ccaf955ce28547ca4ebfe53e10d1fe61/1.12-pre7.json",
		"time": "2021-07-07T13:09:01+00:00",
		"releaseTime": "2017-05-31T10:56:41+00:00",
		"sha1": "c9c584f0ccaf955ce28547ca4ebfe53e10d1fe61",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12-pre6",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/76aad60a521cc0717f9dfd165b8268f724e403e8/1.12-pre6.json",
		"time": "2021-07-07T13:09:01+00:00",
		"releaseTime": "2017-05-29T11:45:12+00:00",
		"sha1": "76aad60a521cc0717f9dfd165b8268f724e403e8",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12-pre5",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/58eb99b5d4b7ebb12f09bde04a87b61167b9e789/1.12-pre5.json",
		"time": "2021-07-07T13:09:00+00:00",
		"releaseTime": "2017-05-19T07:43:28+00:00",
		"sha1": "58eb99b5d4b7ebb12f09bde04a87b61167b9e789",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7a6d417d018c16ec07da37af4cbee61cf940d438/1.12-pre4.json",
		"time": "2021-07-07T13:09:00+00:00",
		"releaseTime": "2017-05-18T12:28:16+00:00",
		"sha1": "7a6d417d018c16ec07da37af4cbee61cf940d438",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5d87f11bc9988336b1f41d5417f55a2e154a8a55/1.12-pre3.json",
		"time": "2021-07-07T13:09:00+00:00",
		"releaseTime": "2017-05-17T14:09:18+00:00",
		"sha1": "5d87f11bc9988336b1f41d5417f55a2e154a8a55",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c9c7a2f0d284a38c95bcd1a5995f7561bf42f093/1.12-pre2.json",
		"time": "2021-07-07T13:08:59+00:00",
		"releaseTime": "2017-05-11T12:11:12+00:00",
		"sha1": "c9c7a2f0d284a38c95bcd1a5995f7561bf42f093",
		"complianceLevel": 0
	  },
	  {
		"id": "1.12-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2c855db043af01f275ff482b5a856ac31649b257/1.12-pre1.json",
		"time": "2021-07-07T13:08:58+00:00",
		"releaseTime": "2017-05-10T11:37:17+00:00",
		"sha1": "2c855db043af01f275ff482b5a856ac31649b257",
		"complianceLevel": 0
	  },
	  {
		"id": "17w18b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2a826d7dfb3b020d5de992a6fad9910cc8925767/17w18b.json",
		"time": "2021-07-07T13:10:59+00:00",
		"releaseTime": "2017-05-04T13:40:22+00:00",
		"sha1": "2a826d7dfb3b020d5de992a6fad9910cc8925767",
		"complianceLevel": 0
	  },
	  {
		"id": "17w18a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e8f76067aa1ce41332fafc488328803361676714/17w18a.json",
		"time": "2021-07-07T13:10:59+00:00",
		"releaseTime": "2017-05-03T14:50:23+00:00",
		"sha1": "e8f76067aa1ce41332fafc488328803361676714",
		"complianceLevel": 0
	  },
	  {
		"id": "17w17b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1e608b422935002b132565bec45013cbb66675a1/17w17b.json",
		"time": "2021-07-07T13:10:58+00:00",
		"releaseTime": "2017-04-27T13:24:23+00:00",
		"sha1": "1e608b422935002b132565bec45013cbb66675a1",
		"complianceLevel": 0
	  },
	  {
		"id": "17w17a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4c254f395fc9717fcc253d11af2fb94b94b26e9d/17w17a.json",
		"time": "2021-07-07T13:10:58+00:00",
		"releaseTime": "2017-04-26T13:48:23+00:00",
		"sha1": "4c254f395fc9717fcc253d11af2fb94b94b26e9d",
		"complianceLevel": 0
	  },
	  {
		"id": "17w16b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b1c8bda5b248f7ede3b7d49c9cf1537c74cd1533/17w16b.json",
		"time": "2021-07-07T13:10:57+00:00",
		"releaseTime": "2017-04-21T12:02:59+00:00",
		"sha1": "b1c8bda5b248f7ede3b7d49c9cf1537c74cd1533",
		"complianceLevel": 0
	  },
	  {
		"id": "17w16a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/eb622a568b68e47622e116baa4598525ed7b148a/17w16a.json",
		"time": "2021-07-07T13:10:57+00:00",
		"releaseTime": "2017-04-20T13:58:35+00:00",
		"sha1": "eb622a568b68e47622e116baa4598525ed7b148a",
		"complianceLevel": 0
	  },
	  {
		"id": "17w15a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a3d8c45522c6fb8db3113dd669aec0231eeabeb5/17w15a.json",
		"time": "2021-07-07T13:10:56+00:00",
		"releaseTime": "2017-04-12T09:30:50+00:00",
		"sha1": "a3d8c45522c6fb8db3113dd669aec0231eeabeb5",
		"complianceLevel": 0
	  },
	  {
		"id": "17w14a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/51f1acd9b0e5f8dff45dc80b4d73eb8e8317e733/17w14a.json",
		"time": "2019-06-28T07:09:15+00:00",
		"releaseTime": "2017-04-05T13:58:01+00:00",
		"sha1": "51f1acd9b0e5f8dff45dc80b4d73eb8e8317e733",
		"complianceLevel": 0
	  },
	  {
		"id": "17w13b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b730c949a3fc6daf051b68e55d0730fbe7334968/17w13b.json",
		"time": "2021-07-07T13:10:56+00:00",
		"releaseTime": "2017-03-31T11:06:35+00:00",
		"sha1": "b730c949a3fc6daf051b68e55d0730fbe7334968",
		"complianceLevel": 0
	  },
	  {
		"id": "17w13a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/65c4f9564999cb0d00152560ad40a85c02b43a7e/17w13a.json",
		"time": "2021-07-07T13:10:55+00:00",
		"releaseTime": "2017-03-30T09:32:19+00:00",
		"sha1": "65c4f9564999cb0d00152560ad40a85c02b43a7e",
		"complianceLevel": 0
	  },
	  {
		"id": "17w06a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/8e8837a81358fec416e6bda2a665345b60b5b30b/17w06a.json",
		"time": "2021-07-07T13:10:55+00:00",
		"releaseTime": "2017-02-08T13:16:29+00:00",
		"sha1": "8e8837a81358fec416e6bda2a665345b60b5b30b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.11.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/4114e19b0bb62ea8f12e9a6f2762c3a249f89b91/1.11.2.json",
		"time": "2021-07-07T13:08:34+00:00",
		"releaseTime": "2016-12-21T09:29:12+00:00",
		"sha1": "4114e19b0bb62ea8f12e9a6f2762c3a249f89b91",
		"complianceLevel": 0
	  },
	  {
		"id": "1.11.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/e9ab403f6ddc54ee6292da63889768cdd611a086/1.11.1.json",
		"time": "2021-07-07T13:08:34+00:00",
		"releaseTime": "2016-12-20T14:05:34+00:00",
		"sha1": "e9ab403f6ddc54ee6292da63889768cdd611a086",
		"complianceLevel": 0
	  },
	  {
		"id": "16w50a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ff04f22c90c2f8e756a3fbb27deb33cc0ed0b7eb/16w50a.json",
		"time": "2021-07-07T13:10:54+00:00",
		"releaseTime": "2016-12-15T14:38:52+00:00",
		"sha1": "ff04f22c90c2f8e756a3fbb27deb33cc0ed0b7eb",
		"complianceLevel": 0
	  },
	  {
		"id": "1.11",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/2ec7a29f402bf399f586397a041311354cc2da43/1.11.json",
		"time": "2021-07-07T13:08:35+00:00",
		"releaseTime": "2016-11-14T14:34:40+00:00",
		"sha1": "2ec7a29f402bf399f586397a041311354cc2da43",
		"complianceLevel": 0
	  },
	  {
		"id": "1.11-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/03d3c9bbb5859f7e3cd91eb77255afa5efd1b42c/1.11-pre1.json",
		"time": "2021-07-07T13:08:58+00:00",
		"releaseTime": "2016-11-08T13:42:50+00:00",
		"sha1": "03d3c9bbb5859f7e3cd91eb77255afa5efd1b42c",
		"complianceLevel": 0
	  },
	  {
		"id": "16w44a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2be7bde8f32b69e859c0de79d44bcd7f2587bf4f/16w44a.json",
		"time": "2021-07-07T13:10:54+00:00",
		"releaseTime": "2016-11-03T14:17:11+00:00",
		"sha1": "2be7bde8f32b69e859c0de79d44bcd7f2587bf4f",
		"complianceLevel": 0
	  },
	  {
		"id": "16w43a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5be866c37b82db0d03bd00fcfff06460b3caeadb/16w43a.json",
		"time": "2021-07-07T13:10:54+00:00",
		"releaseTime": "2016-10-27T09:00:51+00:00",
		"sha1": "5be866c37b82db0d03bd00fcfff06460b3caeadb",
		"complianceLevel": 0
	  },
	  {
		"id": "16w42a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/45eb35c773873abf935d4582524e98ed8730e7af/16w42a.json",
		"time": "2021-07-07T13:10:53+00:00",
		"releaseTime": "2016-10-19T11:17:47+00:00",
		"sha1": "45eb35c773873abf935d4582524e98ed8730e7af",
		"complianceLevel": 0
	  },
	  {
		"id": "16w41a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/fb060cbbe0e995b2bbf83173477079ba8f7ab142/16w41a.json",
		"time": "2021-07-07T13:10:53+00:00",
		"releaseTime": "2016-10-13T14:28:35+00:00",
		"sha1": "fb060cbbe0e995b2bbf83173477079ba8f7ab142",
		"complianceLevel": 0
	  },
	  {
		"id": "16w40a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ed32f1c827a8529183794cc116678a284f6d1a8e/16w40a.json",
		"time": "2021-07-07T13:10:52+00:00",
		"releaseTime": "2016-10-06T13:57:59+00:00",
		"sha1": "ed32f1c827a8529183794cc116678a284f6d1a8e",
		"complianceLevel": 0
	  },
	  {
		"id": "16w39c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7dd1bc6ec9139a4b5f97dc0daa4a6df2d5465056/16w39c.json",
		"time": "2021-07-07T13:10:52+00:00",
		"releaseTime": "2016-09-30T14:11:48+00:00",
		"sha1": "7dd1bc6ec9139a4b5f97dc0daa4a6df2d5465056",
		"complianceLevel": 0
	  },
	  {
		"id": "16w39b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4dd8c007b63d2c3d1f338e8fc48ddaa514bdd319/16w39b.json",
		"time": "2021-07-07T13:10:51+00:00",
		"releaseTime": "2016-09-29T14:39:39+00:00",
		"sha1": "4dd8c007b63d2c3d1f338e8fc48ddaa514bdd319",
		"complianceLevel": 0
	  },
	  {
		"id": "16w39a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/11fa3a9a3db9d48e28ef410b4274feae314d5478/16w39a.json",
		"time": "2021-07-07T13:10:51+00:00",
		"releaseTime": "2016-09-28T13:32:06+00:00",
		"sha1": "11fa3a9a3db9d48e28ef410b4274feae314d5478",
		"complianceLevel": 0
	  },
	  {
		"id": "16w38a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a1102beb99bd587c1c3875e93dc352d51d4b916f/16w38a.json",
		"time": "2021-07-07T13:10:51+00:00",
		"releaseTime": "2016-09-20T12:40:49+00:00",
		"sha1": "a1102beb99bd587c1c3875e93dc352d51d4b916f",
		"complianceLevel": 0
	  },
	  {
		"id": "16w36a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7ab70aaa96241401de5124ac3309ca5070511206/16w36a.json",
		"time": "2021-07-07T13:10:50+00:00",
		"releaseTime": "2016-09-08T14:55:10+00:00",
		"sha1": "7ab70aaa96241401de5124ac3309ca5070511206",
		"complianceLevel": 0
	  },
	  {
		"id": "16w35a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/70500461907de66d62b213ad55a77772a9bdf855/16w35a.json",
		"time": "2021-07-07T13:10:50+00:00",
		"releaseTime": "2016-09-01T13:13:38+00:00",
		"sha1": "70500461907de66d62b213ad55a77772a9bdf855",
		"complianceLevel": 0
	  },
	  {
		"id": "16w33a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0930aa6863c23406297f5d707cf9681cd09b907c/16w33a.json",
		"time": "2021-07-07T13:10:49+00:00",
		"releaseTime": "2016-08-17T12:48:57+00:00",
		"sha1": "0930aa6863c23406297f5d707cf9681cd09b907c",
		"complianceLevel": 0
	  },
	  {
		"id": "16w32b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b63ddc506f6ce63b0bf6cbfc0124294a05cc21a5/16w32b.json",
		"time": "2021-07-07T13:10:49+00:00",
		"releaseTime": "2016-08-11T14:34:29+00:00",
		"sha1": "b63ddc506f6ce63b0bf6cbfc0124294a05cc21a5",
		"complianceLevel": 0
	  },
	  {
		"id": "16w32a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/90fcf7b32d1f637ec126b14815ec510cdc1767d6/16w32a.json",
		"time": "2021-07-07T13:10:48+00:00",
		"releaseTime": "2016-08-10T12:30:10+00:00",
		"sha1": "90fcf7b32d1f637ec126b14815ec510cdc1767d6",
		"complianceLevel": 0
	  },
	  {
		"id": "1.10.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/a86a4eaacfee738c8d609baf6d414175f94c26f6/1.10.2.json",
		"time": "2021-07-07T13:08:33+00:00",
		"releaseTime": "2016-06-23T09:17:32+00:00",
		"sha1": "a86a4eaacfee738c8d609baf6d414175f94c26f6",
		"complianceLevel": 0
	  },
	  {
		"id": "1.10.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/37e79563a008ee1c8b478fd46b8cc8c0d6b2b777/1.10.1.json",
		"time": "2021-07-07T13:08:32+00:00",
		"releaseTime": "2016-06-22T10:13:22+00:00",
		"sha1": "37e79563a008ee1c8b478fd46b8cc8c0d6b2b777",
		"complianceLevel": 0
	  },
	  {
		"id": "1.10",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/1548866cad074c15272ed916d04f35e34d81fb2d/1.10.json",
		"time": "2021-07-07T13:08:33+00:00",
		"releaseTime": "2016-06-08T13:06:18+00:00",
		"sha1": "1548866cad074c15272ed916d04f35e34d81fb2d",
		"complianceLevel": 0
	  },
	  {
		"id": "1.10-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/894de648fcdbd11861d29e232934a85777fb0eb4/1.10-pre2.json",
		"time": "2021-07-07T13:08:58+00:00",
		"releaseTime": "2016-06-07T14:56:34+00:00",
		"sha1": "894de648fcdbd11861d29e232934a85777fb0eb4",
		"complianceLevel": 0
	  },
	  {
		"id": "1.10-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2b852056dbcb0e9143822710fba883720de5cd9e/1.10-pre1.json",
		"time": "2021-07-07T13:08:57+00:00",
		"releaseTime": "2016-06-02T14:45:16+00:00",
		"sha1": "2b852056dbcb0e9143822710fba883720de5cd9e",
		"complianceLevel": 0
	  },
	  {
		"id": "16w21b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5dfa6f81866a0d1eafc60149d5c7b386eca75859/16w21b.json",
		"time": "2021-07-07T13:10:48+00:00",
		"releaseTime": "2016-05-26T12:47:22+00:00",
		"sha1": "5dfa6f81866a0d1eafc60149d5c7b386eca75859",
		"complianceLevel": 0
	  },
	  {
		"id": "16w21a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e18045ac8c1489919d6d3ffa84062a2ad6e88469/16w21a.json",
		"time": "2021-07-07T13:10:47+00:00",
		"releaseTime": "2016-05-25T13:12:09+00:00",
		"sha1": "e18045ac8c1489919d6d3ffa84062a2ad6e88469",
		"complianceLevel": 0
	  },
	  {
		"id": "16w20a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1354a7bab18bd0b0641238bfd763b10e8e4c880c/16w20a.json",
		"time": "2021-07-07T13:10:47+00:00",
		"releaseTime": "2016-05-18T12:45:14+00:00",
		"sha1": "1354a7bab18bd0b0641238bfd763b10e8e4c880c",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9.4",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/fd5c3c2e5f5ec402af673e48ac8604421d97c929/1.9.4.json",
		"time": "2021-07-07T13:08:56+00:00",
		"releaseTime": "2016-05-10T10:17:16+00:00",
		"sha1": "fd5c3c2e5f5ec402af673e48ac8604421d97c929",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9.3",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/ba0ec1ec92c6d08999e177d288a634955958f6cb/1.9.3.json",
		"time": "2021-07-07T13:08:56+00:00",
		"releaseTime": "2016-05-10T08:33:35+00:00",
		"sha1": "ba0ec1ec92c6d08999e177d288a634955958f6cb",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9.3-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1903f3d28b272ed9756b14fc1764ef8d7236050b/1.9.3-pre3.json",
		"time": "2021-07-07T13:09:54+00:00",
		"releaseTime": "2016-05-03T09:28:11+00:00",
		"sha1": "1903f3d28b272ed9756b14fc1764ef8d7236050b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9.3-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7fadcf97db459a6611670ab3ecc2416adfd954b2/1.9.3-pre2.json",
		"time": "2021-07-07T13:09:53+00:00",
		"releaseTime": "2016-04-27T13:33:20+00:00",
		"sha1": "7fadcf97db459a6611670ab3ecc2416adfd954b2",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9.3-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f33b0474c1b3b71b104c6067e0615a4d643704fd/1.9.3-pre1.json",
		"time": "2021-07-07T13:09:53+00:00",
		"releaseTime": "2016-04-21T12:41:42+00:00",
		"sha1": "f33b0474c1b3b71b104c6067e0615a4d643704fd",
		"complianceLevel": 0
	  },
	  {
		"id": "16w15b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/03dbaf9fefec8380f1e5810982c58b02be223e95/16w15b.json",
		"time": "2021-07-07T13:10:46+00:00",
		"releaseTime": "2016-04-13T13:56:41+00:00",
		"sha1": "03dbaf9fefec8380f1e5810982c58b02be223e95",
		"complianceLevel": 0
	  },
	  {
		"id": "16w15a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a372843057e685d067a6b762703ad29f3d3fe891/16w15a.json",
		"time": "2021-07-07T13:10:46+00:00",
		"releaseTime": "2016-04-11T14:38:28+00:00",
		"sha1": "a372843057e685d067a6b762703ad29f3d3fe891",
		"complianceLevel": 0
	  },
	  {
		"id": "16w14a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/650778c7f096f2a265138fa84c325a251859ea58/16w14a.json",
		"time": "2021-07-07T13:10:45+00:00",
		"releaseTime": "2016-04-07T12:47:51+00:00",
		"sha1": "650778c7f096f2a265138fa84c325a251859ea58",
		"complianceLevel": 0
	  },
	  {
		"id": "1.RV-Pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/cc13c3c7e0648fb9ef096588abea984135e82d8d/1.RV-Pre1.json",
		"time": "2021-07-07T13:09:54+00:00",
		"releaseTime": "2016-03-31T16:18:53+00:00",
		"sha1": "cc13c3c7e0648fb9ef096588abea984135e82d8d",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/78c4500478c68e9c14245960ef1ed9c05650b2e0/1.9.2.json",
		"time": "2021-07-07T13:08:55+00:00",
		"releaseTime": "2016-03-30T15:23:55+00:00",
		"sha1": "78c4500478c68e9c14245960ef1ed9c05650b2e0",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/e50b714eff37d59e229dca41650a01f41164388a/1.9.1.json",
		"time": "2021-07-07T13:08:55+00:00",
		"releaseTime": "2016-03-30T13:43:07+00:00",
		"sha1": "e50b714eff37d59e229dca41650a01f41164388a",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9.1-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/bbd2e10c8be7de601c84df161c0954715b30fa5b/1.9.1-pre3.json",
		"time": "2021-07-07T13:09:53+00:00",
		"releaseTime": "2016-03-11T09:20:36+00:00",
		"sha1": "bbd2e10c8be7de601c84df161c0954715b30fa5b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9.1-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/04e4e3d74c09cd2ea0db5bbb3ef28e6abefc4ad8/1.9.1-pre2.json",
		"time": "2021-07-07T13:09:52+00:00",
		"releaseTime": "2016-03-10T15:06:03+00:00",
		"sha1": "04e4e3d74c09cd2ea0db5bbb3ef28e6abefc4ad8",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9.1-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ce2698859c57de9f2411e52ec5b0811a51f0b1c7/1.9.1-pre1.json",
		"time": "2021-07-07T13:09:52+00:00",
		"releaseTime": "2016-03-09T16:27:29+00:00",
		"sha1": "ce2698859c57de9f2411e52ec5b0811a51f0b1c7",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/d4790da414bdcc3ce4d15c45ba0119a818813c59/1.9.json",
		"time": "2021-07-07T13:08:57+00:00",
		"releaseTime": "2016-02-29T13:49:54+00:00",
		"sha1": "d4790da414bdcc3ce4d15c45ba0119a818813c59",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a73d890c66cf9e69e3e16454fc2e2bb4b38bc082/1.9-pre4.json",
		"time": "2021-07-07T13:09:51+00:00",
		"releaseTime": "2016-02-26T15:21:11+00:00",
		"sha1": "a73d890c66cf9e69e3e16454fc2e2bb4b38bc082",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4597a0e7b49b6bbf6d299c57dc34d38f163d02c8/1.9-pre3.json",
		"time": "2021-07-07T13:09:51+00:00",
		"releaseTime": "2016-02-24T15:52:36+00:00",
		"sha1": "4597a0e7b49b6bbf6d299c57dc34d38f163d02c8",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d694c9a3c6e287e7d5bc713af20627480178977a/1.9-pre2.json",
		"time": "2021-07-07T13:09:50+00:00",
		"releaseTime": "2016-02-18T17:41:00+00:00",
		"sha1": "d694c9a3c6e287e7d5bc713af20627480178977a",
		"complianceLevel": 0
	  },
	  {
		"id": "1.9-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0aa6f09304db44f93a8b9d9336cc39e48edf4e18/1.9-pre1.json",
		"time": "2021-07-07T13:09:50+00:00",
		"releaseTime": "2016-02-17T15:23:19+00:00",
		"sha1": "0aa6f09304db44f93a8b9d9336cc39e48edf4e18",
		"complianceLevel": 0
	  },
	  {
		"id": "16w07b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6bdb531e36ca2b123604fe67e3ad1230ae94ef84/16w07b.json",
		"time": "2021-07-07T13:10:45+00:00",
		"releaseTime": "2016-02-16T15:22:39+00:00",
		"sha1": "6bdb531e36ca2b123604fe67e3ad1230ae94ef84",
		"complianceLevel": 0
	  },
	  {
		"id": "16w07a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/8d74132aa41c55e9f35326bf24d6843721949a45/16w07a.json",
		"time": "2021-07-07T13:10:44+00:00",
		"releaseTime": "2016-02-15T15:48:46+00:00",
		"sha1": "8d74132aa41c55e9f35326bf24d6843721949a45",
		"complianceLevel": 0
	  },
	  {
		"id": "16w06a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/364f015778b596481506cc8fe8cc18a15b33079b/16w06a.json",
		"time": "2021-07-07T13:10:44+00:00",
		"releaseTime": "2016-02-10T15:06:41+00:00",
		"sha1": "364f015778b596481506cc8fe8cc18a15b33079b",
		"complianceLevel": 0
	  },
	  {
		"id": "16w05b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e4c1cdcd3f697e217bbded74a7c36be8941fcc1b/16w05b.json",
		"time": "2021-07-07T13:10:43+00:00",
		"releaseTime": "2016-02-04T15:28:02+00:00",
		"sha1": "e4c1cdcd3f697e217bbded74a7c36be8941fcc1b",
		"complianceLevel": 0
	  },
	  {
		"id": "16w05a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2ae800b28571760e5282d43d04185c8c7bb5e143/16w05a.json",
		"time": "2021-07-07T13:10:43+00:00",
		"releaseTime": "2016-02-03T15:48:38+00:00",
		"sha1": "2ae800b28571760e5282d43d04185c8c7bb5e143",
		"complianceLevel": 0
	  },
	  {
		"id": "16w04a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7cc40a3f02fd634f30b439b35d9a1f99d8bcea67/16w04a.json",
		"time": "2021-07-07T13:10:43+00:00",
		"releaseTime": "2016-01-28T15:37:24+00:00",
		"sha1": "7cc40a3f02fd634f30b439b35d9a1f99d8bcea67",
		"complianceLevel": 0
	  },
	  {
		"id": "16w03a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/501c6a8494456fd0afa82e447752992b742d3aa4/16w03a.json",
		"time": "2021-07-07T13:10:42+00:00",
		"releaseTime": "2016-01-20T14:29:24+00:00",
		"sha1": "501c6a8494456fd0afa82e447752992b742d3aa4",
		"complianceLevel": 0
	  },
	  {
		"id": "16w02a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4ffbef6451aea4f6128f54f8683fbf131070b44c/16w02a.json",
		"time": "2021-07-07T13:10:41+00:00",
		"releaseTime": "2016-01-13T15:15:16+00:00",
		"sha1": "4ffbef6451aea4f6128f54f8683fbf131070b44c",
		"complianceLevel": 0
	  },
	  {
		"id": "15w51b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/04462230e67acddbcc3caf949ab31822facb34b8/15w51b.json",
		"time": "2021-07-07T13:10:41+00:00",
		"releaseTime": "2015-12-17T15:30:41+00:00",
		"sha1": "04462230e67acddbcc3caf949ab31822facb34b8",
		"complianceLevel": 0
	  },
	  {
		"id": "15w51a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a8a688ad68f561fbf844b664d1c282710f9fc0d8/15w51a.json",
		"time": "2021-07-07T13:10:41+00:00",
		"releaseTime": "2015-12-17T14:02:37+00:00",
		"sha1": "a8a688ad68f561fbf844b664d1c282710f9fc0d8",
		"complianceLevel": 0
	  },
	  {
		"id": "15w50a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f024cadb981e397e6701c47044d0057c16a58f44/15w50a.json",
		"time": "2021-07-07T13:10:40+00:00",
		"releaseTime": "2015-12-09T15:35:57+00:00",
		"sha1": "f024cadb981e397e6701c47044d0057c16a58f44",
		"complianceLevel": 0
	  },
	  {
		"id": "15w49b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1dbe6cd541c43d6fa1643f3e50ee1e874612958c/15w49b.json",
		"time": "2021-07-07T13:10:40+00:00",
		"releaseTime": "2015-12-03T15:23:22+00:00",
		"sha1": "1dbe6cd541c43d6fa1643f3e50ee1e874612958c",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.9",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/0a5a761091458d69e3dea629a018eff7d74eb534/1.8.9.json",
		"time": "2021-07-07T13:08:54+00:00",
		"releaseTime": "2015-12-03T09:24:39+00:00",
		"sha1": "0a5a761091458d69e3dea629a018eff7d74eb534",
		"complianceLevel": 0
	  },
	  {
		"id": "15w49a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/940efd3592363dcd0a45b63ed159f5fa6e7ddfce/15w49a.json",
		"time": "2021-07-07T13:10:39+00:00",
		"releaseTime": "2015-12-02T15:09:37+00:00",
		"sha1": "940efd3592363dcd0a45b63ed159f5fa6e7ddfce",
		"complianceLevel": 0
	  },
	  {
		"id": "15w47c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ba9ee914993016dd7a54a4364081459c63fddfce/15w47c.json",
		"time": "2021-07-07T13:10:39+00:00",
		"releaseTime": "2015-11-20T12:46:56+00:00",
		"sha1": "ba9ee914993016dd7a54a4364081459c63fddfce",
		"complianceLevel": 0
	  },
	  {
		"id": "15w47b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/af91931871af0f789b7db0bd2b56398601512c72/15w47b.json",
		"time": "2021-07-07T13:10:39+00:00",
		"releaseTime": "2015-11-19T14:48:03+00:00",
		"sha1": "af91931871af0f789b7db0bd2b56398601512c72",
		"complianceLevel": 0
	  },
	  {
		"id": "15w47a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f4877328436495df20bae5a48c94c0f341d57504/15w47a.json",
		"time": "2021-07-07T13:10:38+00:00",
		"releaseTime": "2015-11-18T15:53:41+00:00",
		"sha1": "f4877328436495df20bae5a48c94c0f341d57504",
		"complianceLevel": 0
	  },
	  {
		"id": "15w46a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4367245c01603ff6df73bc97b53ecec3a1804f01/15w46a.json",
		"time": "2021-07-07T13:10:38+00:00",
		"releaseTime": "2015-11-12T12:11:47+00:00",
		"sha1": "4367245c01603ff6df73bc97b53ecec3a1804f01",
		"complianceLevel": 0
	  },
	  {
		"id": "15w45a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1ebbad74c0af351340f64ca2b3cb913fc027ad30/15w45a.json",
		"time": "2021-07-07T13:10:37+00:00",
		"releaseTime": "2015-11-05T13:04:07+00:00",
		"sha1": "1ebbad74c0af351340f64ca2b3cb913fc027ad30",
		"complianceLevel": 0
	  },
	  {
		"id": "15w44b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0501e9eaf57c5ff6dfd56b70ad8f61ec789fa5bb/15w44b.json",
		"time": "2021-07-07T13:10:37+00:00",
		"releaseTime": "2015-10-30T11:23:17+00:00",
		"sha1": "0501e9eaf57c5ff6dfd56b70ad8f61ec789fa5bb",
		"complianceLevel": 0
	  },
	  {
		"id": "15w44a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9fbcc6db118c2091db755f3cd3683b07554e4eb8/15w44a.json",
		"time": "2021-07-07T13:10:37+00:00",
		"releaseTime": "2015-10-28T15:09:36+00:00",
		"sha1": "9fbcc6db118c2091db755f3cd3683b07554e4eb8",
		"complianceLevel": 0
	  },
	  {
		"id": "15w43c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/95d5b716b53a547bda2baf65edf5e0352f573ada/15w43c.json",
		"time": "2021-07-07T13:10:36+00:00",
		"releaseTime": "2015-10-23T15:35:55+00:00",
		"sha1": "95d5b716b53a547bda2baf65edf5e0352f573ada",
		"complianceLevel": 0
	  },
	  {
		"id": "15w43b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/305d0dfaa25096ff548e894555d31614c9e79de4/15w43b.json",
		"time": "2021-07-07T13:10:36+00:00",
		"releaseTime": "2015-10-22T14:11:58+00:00",
		"sha1": "305d0dfaa25096ff548e894555d31614c9e79de4",
		"complianceLevel": 0
	  },
	  {
		"id": "15w43a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b1029b6f3a62d663045a0b97126eda2ba27d1bc4/15w43a.json",
		"time": "2021-07-07T13:10:35+00:00",
		"releaseTime": "2015-10-21T15:28:52+00:00",
		"sha1": "b1029b6f3a62d663045a0b97126eda2ba27d1bc4",
		"complianceLevel": 0
	  },
	  {
		"id": "15w42a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/fbd7577f54df809acec14dc2e439d92f55b3301e/15w42a.json",
		"time": "2021-07-07T13:10:35+00:00",
		"releaseTime": "2015-10-14T13:25:14+00:00",
		"sha1": "fbd7577f54df809acec14dc2e439d92f55b3301e",
		"complianceLevel": 0
	  },
	  {
		"id": "15w41b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e83e95be6ed251ad548b7658d5b9226a0150ba09/15w41b.json",
		"time": "2021-07-07T13:10:34+00:00",
		"releaseTime": "2015-10-07T14:07:26+00:00",
		"sha1": "e83e95be6ed251ad548b7658d5b9226a0150ba09",
		"complianceLevel": 0
	  },
	  {
		"id": "15w41a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9b997fb40c62f4c16000b1d4f358fe86ee7f8fb3/15w41a.json",
		"time": "2021-07-07T13:10:34+00:00",
		"releaseTime": "2015-10-07T13:19:53+00:00",
		"sha1": "9b997fb40c62f4c16000b1d4f358fe86ee7f8fb3",
		"complianceLevel": 0
	  },
	  {
		"id": "15w40b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/894e7146b7785b3f4f9f50ac991d189d1eb5516f/15w40b.json",
		"time": "2021-07-07T13:10:33+00:00",
		"releaseTime": "2015-09-30T14:13:54+00:00",
		"sha1": "894e7146b7785b3f4f9f50ac991d189d1eb5516f",
		"complianceLevel": 0
	  },
	  {
		"id": "15w40a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4d6ef2dbdcf0a34be842c62cf66c86f308159080/15w40a.json",
		"time": "2021-07-07T13:10:33+00:00",
		"releaseTime": "2015-09-30T13:13:54+00:00",
		"sha1": "4d6ef2dbdcf0a34be842c62cf66c86f308159080",
		"complianceLevel": 0
	  },
	  {
		"id": "15w39c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/89d2ddd11a60419122b24546ea6f9ada8138dadb/15w39c.json",
		"time": "2021-07-07T13:10:32+00:00",
		"releaseTime": "2015-09-23T13:13:54+00:00",
		"sha1": "89d2ddd11a60419122b24546ea6f9ada8138dadb",
		"complianceLevel": 0
	  },
	  {
		"id": "15w39b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/02b022f07c10e308306fd9e365448e6011cac62d/15w39b.json",
		"time": "2021-07-07T13:10:32+00:00",
		"releaseTime": "2015-09-21T15:09:52+00:00",
		"sha1": "02b022f07c10e308306fd9e365448e6011cac62d",
		"complianceLevel": 0
	  },
	  {
		"id": "15w39a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/44b5ca60acc08636a7c405c7514ba898a0756b51/15w39a.json",
		"time": "2021-07-07T13:10:31+00:00",
		"releaseTime": "2015-09-21T13:16:32+00:00",
		"sha1": "44b5ca60acc08636a7c405c7514ba898a0756b51",
		"complianceLevel": 0
	  },
	  {
		"id": "15w38b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a7ae63b9909c0cdd7833137a45ea6cae4f82a55e/15w38b.json",
		"time": "2021-07-07T13:10:31+00:00",
		"releaseTime": "2015-09-17T14:22:31+00:00",
		"sha1": "a7ae63b9909c0cdd7833137a45ea6cae4f82a55e",
		"complianceLevel": 0
	  },
	  {
		"id": "15w38a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/76027a2788d0af51508b88bef8d34fa910d9e407/15w38a.json",
		"time": "2021-07-07T13:10:30+00:00",
		"releaseTime": "2015-09-16T14:22:31+00:00",
		"sha1": "76027a2788d0af51508b88bef8d34fa910d9e407",
		"complianceLevel": 0
	  },
	  {
		"id": "15w37a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/10ee5253b9abb7f1416d8dfa732905812791f30a/15w37a.json",
		"time": "2021-07-07T13:10:30+00:00",
		"releaseTime": "2015-09-10T14:22:31+00:00",
		"sha1": "10ee5253b9abb7f1416d8dfa732905812791f30a",
		"complianceLevel": 0
	  },
	  {
		"id": "15w36d",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e9b437bb78f77d2f4b4049fde64883b637666dc1/15w36d.json",
		"time": "2021-07-07T13:10:29+00:00",
		"releaseTime": "2015-09-04T14:22:31+00:00",
		"sha1": "e9b437bb78f77d2f4b4049fde64883b637666dc1",
		"complianceLevel": 0
	  },
	  {
		"id": "15w36c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9e83553834966f7a99f2e92b8d2c388980e537a7/15w36c.json",
		"time": "2021-07-07T13:10:29+00:00",
		"releaseTime": "2015-09-02T16:07:22+00:00",
		"sha1": "9e83553834966f7a99f2e92b8d2c388980e537a7",
		"complianceLevel": 0
	  },
	  {
		"id": "15w36b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6e0175c91b3f9142fb4ab5173893d7236dacdc6b/15w36b.json",
		"time": "2021-07-07T13:10:29+00:00",
		"releaseTime": "2015-09-02T15:36:25+00:00",
		"sha1": "6e0175c91b3f9142fb4ab5173893d7236dacdc6b",
		"complianceLevel": 0
	  },
	  {
		"id": "15w36a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/145fe914deeeb5a51b01c530a4c33eae7b1d758a/15w36a.json",
		"time": "2021-07-07T13:10:28+00:00",
		"releaseTime": "2015-09-02T14:46:40+00:00",
		"sha1": "145fe914deeeb5a51b01c530a4c33eae7b1d758a",
		"complianceLevel": 0
	  },
	  {
		"id": "15w35e",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/5bc4dbb1921bf4270f674d1bfad3eb064197f7a6/15w35e.json",
		"time": "2021-07-07T13:10:28+00:00",
		"releaseTime": "2015-08-28T18:14:02+00:00",
		"sha1": "5bc4dbb1921bf4270f674d1bfad3eb064197f7a6",
		"complianceLevel": 0
	  },
	  {
		"id": "15w35d",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/388c752feba74d5b1113fce30d555069c25f067f/15w35d.json",
		"time": "2021-07-07T13:10:27+00:00",
		"releaseTime": "2015-08-28T16:25:35+00:00",
		"sha1": "388c752feba74d5b1113fce30d555069c25f067f",
		"complianceLevel": 0
	  },
	  {
		"id": "15w35c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f7ce69df26bd6a363dc3f037e4640ff6ba1707b9/15w35c.json",
		"time": "2021-07-07T13:10:27+00:00",
		"releaseTime": "2015-08-28T11:21:00+00:00",
		"sha1": "f7ce69df26bd6a363dc3f037e4640ff6ba1707b9",
		"complianceLevel": 0
	  },
	  {
		"id": "15w35b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0087ef488703747d72296bf1f64871574f2213ee/15w35b.json",
		"time": "2021-07-07T13:10:26+00:00",
		"releaseTime": "2015-08-24T15:39:18+00:00",
		"sha1": "0087ef488703747d72296bf1f64871574f2213ee",
		"complianceLevel": 0
	  },
	  {
		"id": "15w35a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/03307c22a5cbe2b1a21348bffc7d45ab3d99762d/15w35a.json",
		"time": "2021-07-07T13:10:26+00:00",
		"releaseTime": "2015-08-24T14:19:31+00:00",
		"sha1": "03307c22a5cbe2b1a21348bffc7d45ab3d99762d",
		"complianceLevel": 0
	  },
	  {
		"id": "15w34d",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e2b57351a63417b92b13a32e9d6ce0f5ae6ade49/15w34d.json",
		"time": "2021-07-07T13:10:26+00:00",
		"releaseTime": "2015-08-21T15:27:55+00:00",
		"sha1": "e2b57351a63417b92b13a32e9d6ce0f5ae6ade49",
		"complianceLevel": 0
	  },
	  {
		"id": "15w34c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e1d484327695c136c701daaeb9eec46984e7a101/15w34c.json",
		"time": "2021-07-07T13:10:25+00:00",
		"releaseTime": "2015-08-21T12:45:20+00:00",
		"sha1": "e1d484327695c136c701daaeb9eec46984e7a101",
		"complianceLevel": 0
	  },
	  {
		"id": "15w34b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/eb07ca235a8ef1bee67a43b19887dc0d3387a6f7/15w34b.json",
		"time": "2021-07-07T13:10:25+00:00",
		"releaseTime": "2015-08-20T14:00:03+00:00",
		"sha1": "eb07ca235a8ef1bee67a43b19887dc0d3387a6f7",
		"complianceLevel": 0
	  },
	  {
		"id": "15w34a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/dabaa043a724f0848f1565b8d0bb479227cc2de4/15w34a.json",
		"time": "2021-07-07T13:10:24+00:00",
		"releaseTime": "2015-08-19T12:56:01+00:00",
		"sha1": "dabaa043a724f0848f1565b8d0bb479227cc2de4",
		"complianceLevel": 0
	  },
	  {
		"id": "15w33c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/aa5694a6d5aa610eebbb920bbea5ff9815814add/15w33c.json",
		"time": "2021-07-07T13:10:24+00:00",
		"releaseTime": "2015-08-14T13:10:46+00:00",
		"sha1": "aa5694a6d5aa610eebbb920bbea5ff9815814add",
		"complianceLevel": 0
	  },
	  {
		"id": "15w33b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/32ce85d9302438348acca056c20e2b54e7803311/15w33b.json",
		"time": "2021-07-07T13:10:23+00:00",
		"releaseTime": "2015-08-12T15:29:11+00:00",
		"sha1": "32ce85d9302438348acca056c20e2b54e7803311",
		"complianceLevel": 0
	  },
	  {
		"id": "15w33a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/88adbb1616ab2dcf394926cbdacf9d24bce8257c/15w33a.json",
		"time": "2021-07-07T13:10:23+00:00",
		"releaseTime": "2015-08-12T14:05:07+00:00",
		"sha1": "88adbb1616ab2dcf394926cbdacf9d24bce8257c",
		"complianceLevel": 0
	  },
	  {
		"id": "15w32c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7220309b2ed58abb816ed5898ff361be44d44c2d/15w32c.json",
		"time": "2021-07-07T13:10:23+00:00",
		"releaseTime": "2015-08-07T14:08:17+00:00",
		"sha1": "7220309b2ed58abb816ed5898ff361be44d44c2d",
		"complianceLevel": 0
	  },
	  {
		"id": "15w32b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f0a5d8cde1b51810a7bba49135b17bb183513f75/15w32b.json",
		"time": "2021-07-07T13:10:22+00:00",
		"releaseTime": "2015-08-06T13:51:47+00:00",
		"sha1": "f0a5d8cde1b51810a7bba49135b17bb183513f75",
		"complianceLevel": 0
	  },
	  {
		"id": "15w32a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/44138639b2b4d415ce8f5a4b711378abd981a74e/15w32a.json",
		"time": "2021-07-07T13:10:22+00:00",
		"releaseTime": "2015-08-05T12:22:42+00:00",
		"sha1": "44138639b2b4d415ce8f5a4b711378abd981a74e",
		"complianceLevel": 0
	  },
	  {
		"id": "15w31c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/088767c91ec7e570e8ba8a113cbe07eb540a2b61/15w31c.json",
		"time": "2021-07-07T13:10:21+00:00",
		"releaseTime": "2015-07-31T13:45:08+00:00",
		"sha1": "088767c91ec7e570e8ba8a113cbe07eb540a2b61",
		"complianceLevel": 0
	  },
	  {
		"id": "15w31b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6993751373d0c8aedf6d16efc463b14c897c7bcd/15w31b.json",
		"time": "2021-07-07T13:10:21+00:00",
		"releaseTime": "2015-07-30T13:38:32+00:00",
		"sha1": "6993751373d0c8aedf6d16efc463b14c897c7bcd",
		"complianceLevel": 0
	  },
	  {
		"id": "15w31a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7a917ea3d54bf69e00b275675c964e08a43f9ec3/15w31a.json",
		"time": "2021-07-07T13:10:21+00:00",
		"releaseTime": "2015-07-29T13:24:33+00:00",
		"sha1": "7a917ea3d54bf69e00b275675c964e08a43f9ec3",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.8",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/840671701bed1ac92694ae0f9ceded84bd104263/1.8.8.json",
		"time": "2021-07-07T13:08:54+00:00",
		"releaseTime": "2015-07-27T10:31:28+00:00",
		"sha1": "840671701bed1ac92694ae0f9ceded84bd104263",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.7",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/2c32fd36779ef294d8f0f64e695cd7e0e6bccc90/1.8.7.json",
		"time": "2021-07-07T13:08:53+00:00",
		"releaseTime": "2015-06-05T10:10:44+00:00",
		"sha1": "2c32fd36779ef294d8f0f64e695cd7e0e6bccc90",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.6",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/17546e9973de7ff7966942939ac84c5ce7bf5fc9/1.8.6.json",
		"time": "2021-07-07T13:08:53+00:00",
		"releaseTime": "2015-05-25T10:31:19+00:00",
		"sha1": "17546e9973de7ff7966942939ac84c5ce7bf5fc9",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.5",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/ea7a82250794aaf9b10e2b6aa57f183ef4caa8d3/1.8.5.json",
		"time": "2021-07-07T13:08:52+00:00",
		"releaseTime": "2015-05-22T11:15:28+00:00",
		"sha1": "ea7a82250794aaf9b10e2b6aa57f183ef4caa8d3",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.4",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/f0820dbf72b95c9ce91d7ef25c8197435357f8bd/1.8.4.json",
		"time": "2021-07-07T13:08:52+00:00",
		"releaseTime": "2015-04-17T11:37:50+00:00",
		"sha1": "f0820dbf72b95c9ce91d7ef25c8197435357f8bd",
		"complianceLevel": 0
	  },
	  {
		"id": "15w14a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/24222926b78e0efbd1e552b42a7b3e8f499cb08a/15w14a.json",
		"time": "2021-07-07T13:10:20+00:00",
		"releaseTime": "2015-04-01T07:08:00+00:00",
		"sha1": "24222926b78e0efbd1e552b42a7b3e8f499cb08a",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.3",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/47bd60baa69eeb66a40ad415b91ea75874393f28/1.8.3.json",
		"time": "2021-07-07T13:08:52+00:00",
		"releaseTime": "2015-02-20T14:00:09+00:00",
		"sha1": "47bd60baa69eeb66a40ad415b91ea75874393f28",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/8ad0750111152dd36fab753b6967ea8f232ff9a5/1.8.2.json",
		"time": "2021-07-07T13:08:51+00:00",
		"releaseTime": "2015-02-19T15:47:29+00:00",
		"sha1": "8ad0750111152dd36fab753b6967ea8f232ff9a5",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.2-pre7",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d42c640ca64a9f76c93013c33add28fa614fcb5b/1.8.2-pre7.json",
		"time": "2021-07-07T13:09:49+00:00",
		"releaseTime": "2015-02-16T13:01:35+00:00",
		"sha1": "d42c640ca64a9f76c93013c33add28fa614fcb5b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.2-pre6",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ee9bf993471e18b9b8c94ac710495f8988e27618/1.8.2-pre6.json",
		"time": "2021-07-07T13:09:49+00:00",
		"releaseTime": "2015-01-30T11:58:24+00:00",
		"sha1": "ee9bf993471e18b9b8c94ac710495f8988e27618",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.2-pre5",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a80a30bca07deab8b6f4ac10eaf37dc3eff7a13b/1.8.2-pre5.json",
		"time": "2021-07-07T13:09:48+00:00",
		"releaseTime": "2015-01-26T15:03:24+00:00",
		"sha1": "a80a30bca07deab8b6f4ac10eaf37dc3eff7a13b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.2-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/71246befca940f6c1f39ca581496fafd635a377b/1.8.2-pre4.json",
		"time": "2021-07-07T13:09:48+00:00",
		"releaseTime": "2015-01-16T14:19:59+00:00",
		"sha1": "71246befca940f6c1f39ca581496fafd635a377b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.2-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/efd5a5d1205500ac52fb24e5dab9493933380712/1.8.2-pre3.json",
		"time": "2021-07-07T13:09:48+00:00",
		"releaseTime": "2015-01-15T16:44:33+00:00",
		"sha1": "efd5a5d1205500ac52fb24e5dab9493933380712",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.2-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/fa152a8676f821f054bd70480d01682d6bd38a84/1.8.2-pre2.json",
		"time": "2021-07-07T13:09:47+00:00",
		"releaseTime": "2015-01-15T15:07:31+00:00",
		"sha1": "fa152a8676f821f054bd70480d01682d6bd38a84",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.2-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6209c1806912a044434ea1b16d0d6923cca7812c/1.8.2-pre1.json",
		"time": "2021-07-07T13:09:47+00:00",
		"releaseTime": "2014-12-18T11:29:41+00:00",
		"sha1": "6209c1806912a044434ea1b16d0d6923cca7812c",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/ebd569c4cbf569133a4f82610b7d55ee6c8b112a/1.8.1.json",
		"time": "2021-07-07T13:08:51+00:00",
		"releaseTime": "2014-11-24T14:13:31+00:00",
		"sha1": "ebd569c4cbf569133a4f82610b7d55ee6c8b112a",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.1-pre5",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9ef79b52618e178cd91aaff1a6d8a4b17aa396f6/1.8.1-pre5.json",
		"time": "2021-07-07T13:09:46+00:00",
		"releaseTime": "2014-11-19T14:30:48+00:00",
		"sha1": "9ef79b52618e178cd91aaff1a6d8a4b17aa396f6",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.1-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7cac56161bcbd0b10dd1b341dec65f9b678159f4/1.8.1-pre4.json",
		"time": "2021-07-07T13:09:46+00:00",
		"releaseTime": "2014-11-06T14:10:50+00:00",
		"sha1": "7cac56161bcbd0b10dd1b341dec65f9b678159f4",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.1-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f0dd250b0e4ad98c2a7bdb501b548da466a8f6e7/1.8.1-pre3.json",
		"time": "2021-07-07T13:09:45+00:00",
		"releaseTime": "2014-10-23T12:59:42+00:00",
		"sha1": "f0dd250b0e4ad98c2a7bdb501b548da466a8f6e7",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.1-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ab764f8dc1992a9974990e80c626410560f1bef0/1.8.1-pre2.json",
		"time": "2021-07-07T13:09:45+00:00",
		"releaseTime": "2014-10-16T14:19:27+00:00",
		"sha1": "ab764f8dc1992a9974990e80c626410560f1bef0",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8.1-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/06eae8232dbb549a84e733af8e4ebffa32c0436c/1.8.1-pre1.json",
		"time": "2021-07-07T13:09:44+00:00",
		"releaseTime": "2014-10-15T13:25:11+00:00",
		"sha1": "06eae8232dbb549a84e733af8e4ebffa32c0436c",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/32a929f09f3161886ecdf814b7d458c6cc749398/1.8.json",
		"time": "2021-07-07T13:08:55+00:00",
		"releaseTime": "2014-09-02T08:24:35+00:00",
		"sha1": "32a929f09f3161886ecdf814b7d458c6cc749398",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/58e6445cb7a1d32560856c60583a4540e1703dca/1.8-pre3.json",
		"time": "2021-07-07T13:09:44+00:00",
		"releaseTime": "2014-08-28T09:40:54+00:00",
		"sha1": "58e6445cb7a1d32560856c60583a4540e1703dca",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/37985cd99d6dabbb707cb58e450f2a1ff387cd68/1.8-pre2.json",
		"time": "2021-07-07T13:09:44+00:00",
		"releaseTime": "2014-08-25T14:52:18+00:00",
		"sha1": "37985cd99d6dabbb707cb58e450f2a1ff387cd68",
		"complianceLevel": 0
	  },
	  {
		"id": "1.8-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/69e33f93b5fa6c097dc6ed36781021fe809a4b41/1.8-pre1.json",
		"time": "2021-07-07T13:09:43+00:00",
		"releaseTime": "2014-08-21T13:56:26+00:00",
		"sha1": "69e33f93b5fa6c097dc6ed36781021fe809a4b41",
		"complianceLevel": 0
	  },
	  {
		"id": "14w34d",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/16394df985ca6d9a8ee789a9e89fe95cb9a23be3/14w34d.json",
		"time": "2021-07-07T13:10:20+00:00",
		"releaseTime": "2014-08-20T12:46:59+00:00",
		"sha1": "16394df985ca6d9a8ee789a9e89fe95cb9a23be3",
		"complianceLevel": 0
	  },
	  {
		"id": "14w34c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0c5a19458e0b553fa240dfc5a61289c22e111cb1/14w34c.json",
		"time": "2021-07-07T13:10:19+00:00",
		"releaseTime": "2014-08-19T15:31:24+00:00",
		"sha1": "0c5a19458e0b553fa240dfc5a61289c22e111cb1",
		"complianceLevel": 0
	  },
	  {
		"id": "14w34b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6c89f3a7de55963789593ba4d2106a9c8e949bab/14w34b.json",
		"time": "2021-07-07T13:10:19+00:00",
		"releaseTime": "2014-08-18T15:14:28+00:00",
		"sha1": "6c89f3a7de55963789593ba4d2106a9c8e949bab",
		"complianceLevel": 0
	  },
	  {
		"id": "14w34a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/21c3aed8e37e4134aec4530f221f94edea70899c/14w34a.json",
		"time": "2021-07-07T13:10:18+00:00",
		"releaseTime": "2014-08-18T14:14:11+00:00",
		"sha1": "21c3aed8e37e4134aec4530f221f94edea70899c",
		"complianceLevel": 0
	  },
	  {
		"id": "14w33c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/42ed6488b8e3e71aa083e6eb735921eb4c75e81c/14w33c.json",
		"time": "2021-07-07T13:10:18+00:00",
		"releaseTime": "2014-08-15T18:00:26+00:00",
		"sha1": "42ed6488b8e3e71aa083e6eb735921eb4c75e81c",
		"complianceLevel": 0
	  },
	  {
		"id": "14w33b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c13a49d55ca676a894b88bbf46a42e86a092a4b4/14w33b.json",
		"time": "2021-07-07T13:10:17+00:00",
		"releaseTime": "2014-08-15T16:23:51+00:00",
		"sha1": "c13a49d55ca676a894b88bbf46a42e86a092a4b4",
		"complianceLevel": 0
	  },
	  {
		"id": "14w33a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a28e406554328a2c046b215d603a4e7f0a9d7393/14w33a.json",
		"time": "2021-07-07T13:10:17+00:00",
		"releaseTime": "2014-08-13T15:08:14+00:00",
		"sha1": "a28e406554328a2c046b215d603a4e7f0a9d7393",
		"complianceLevel": 0
	  },
	  {
		"id": "14w32d",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f621a856544f9901ae6f47848c2e2376d70b00ed/14w32d.json",
		"time": "2021-07-07T13:10:17+00:00",
		"releaseTime": "2014-08-08T15:13:41+00:00",
		"sha1": "f621a856544f9901ae6f47848c2e2376d70b00ed",
		"complianceLevel": 0
	  },
	  {
		"id": "14w32c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/aaf6f41f3f7cb1e09ab1cd89a5943c5086ce5168/14w32c.json",
		"time": "2021-07-07T13:10:16+00:00",
		"releaseTime": "2014-08-08T14:11:20+00:00",
		"sha1": "aaf6f41f3f7cb1e09ab1cd89a5943c5086ce5168",
		"complianceLevel": 0
	  },
	  {
		"id": "14w32b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/20cbbc7264a089e913891bf45097a138ddf36158/14w32b.json",
		"time": "2021-07-07T13:10:16+00:00",
		"releaseTime": "2014-08-07T14:45:17+00:00",
		"sha1": "20cbbc7264a089e913891bf45097a138ddf36158",
		"complianceLevel": 0
	  },
	  {
		"id": "14w32a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7a78f68530df526855811fda58bedb4dce4bdf44/14w32a.json",
		"time": "2021-07-07T13:10:16+00:00",
		"releaseTime": "2014-08-06T14:01:16+00:00",
		"sha1": "7a78f68530df526855811fda58bedb4dce4bdf44",
		"complianceLevel": 0
	  },
	  {
		"id": "14w31a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/16d78de0421d6cb0ef46172826e53bd73cb3c76b/14w31a.json",
		"time": "2021-07-07T13:10:15+00:00",
		"releaseTime": "2014-07-30T15:38:05+00:00",
		"sha1": "16d78de0421d6cb0ef46172826e53bd73cb3c76b",
		"complianceLevel": 0
	  },
	  {
		"id": "14w30c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/bf6acdcd84cd8bcd67c12f7ed0ed897e28fbef2e/14w30c.json",
		"time": "2021-07-07T13:10:15+00:00",
		"releaseTime": "2014-07-24T14:39:09+00:00",
		"sha1": "bf6acdcd84cd8bcd67c12f7ed0ed897e28fbef2e",
		"complianceLevel": 0
	  },
	  {
		"id": "14w30b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/bbe511406ce99d79c8659f495cdb8060fffd8718/14w30b.json",
		"time": "2021-07-07T13:10:14+00:00",
		"releaseTime": "2014-07-23T15:03:03+00:00",
		"sha1": "bbe511406ce99d79c8659f495cdb8060fffd8718",
		"complianceLevel": 0
	  },
	  {
		"id": "14w30a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7b7ff6da16b931f20e5bff1d527c42fcb9506269/14w30a.json",
		"time": "2021-07-07T13:10:14+00:00",
		"releaseTime": "2014-07-23T13:15:42+00:00",
		"sha1": "7b7ff6da16b931f20e5bff1d527c42fcb9506269",
		"complianceLevel": 0
	  },
	  {
		"id": "14w29b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/961748fc6fb51b3c7cf57060d4f5a063cc41502c/14w29b.json",
		"time": "2021-07-07T13:10:14+00:00",
		"releaseTime": "2014-07-16T17:27:40+00:00",
		"sha1": "961748fc6fb51b3c7cf57060d4f5a063cc41502c",
		"complianceLevel": 0
	  },
	  {
		"id": "14w29a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7859849cc97b7ac84976cf0ba1ab19d6d85e741a/14w29a.json",
		"time": "2021-07-07T13:10:13+00:00",
		"releaseTime": "2014-07-16T15:18:17+00:00",
		"sha1": "7859849cc97b7ac84976cf0ba1ab19d6d85e741a",
		"complianceLevel": 0
	  },
	  {
		"id": "14w28b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c1abec08d053046e23934be21e3475968fef5aba/14w28b.json",
		"time": "2021-07-07T13:10:13+00:00",
		"releaseTime": "2014-07-10T14:28:48+00:00",
		"sha1": "c1abec08d053046e23934be21e3475968fef5aba",
		"complianceLevel": 0
	  },
	  {
		"id": "14w28a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4e8bb577176355b0d7f4f5adabca6b6b2e160211/14w28a.json",
		"time": "2021-07-07T13:10:12+00:00",
		"releaseTime": "2014-07-09T15:42:36+00:00",
		"sha1": "4e8bb577176355b0d7f4f5adabca6b6b2e160211",
		"complianceLevel": 0
	  },
	  {
		"id": "14w27b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/60f220693f30a7c92b53631e6f3a48e2d5e3ce13/14w27b.json",
		"time": "2021-07-07T13:10:12+00:00",
		"releaseTime": "2014-07-02T18:34:56+00:00",
		"sha1": "60f220693f30a7c92b53631e6f3a48e2d5e3ce13",
		"complianceLevel": 0
	  },
	  {
		"id": "14w27a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d4ed459f98feab3d0e787235bebe0aa70c27eda3/14w27a.json",
		"time": "2021-07-07T13:10:12+00:00",
		"releaseTime": "2014-07-02T16:07:20+00:00",
		"sha1": "d4ed459f98feab3d0e787235bebe0aa70c27eda3",
		"complianceLevel": 0
	  },
	  {
		"id": "14w26c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e3d5c17789ae8283457dfe69c8f67c50e948f98c/14w26c.json",
		"time": "2021-07-07T13:10:11+00:00",
		"releaseTime": "2014-06-26T15:05:03+00:00",
		"sha1": "e3d5c17789ae8283457dfe69c8f67c50e948f98c",
		"complianceLevel": 0
	  },
	  {
		"id": "14w26b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0c793c15ccf37c051dfc8189e34febac7b84add0/14w26b.json",
		"time": "2021-07-07T13:10:11+00:00",
		"releaseTime": "2014-06-25T15:08:39+00:00",
		"sha1": "0c793c15ccf37c051dfc8189e34febac7b84add0",
		"complianceLevel": 0
	  },
	  {
		"id": "14w26a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4bd91bc2016638c36e5f75fbddc76f799248b524/14w26a.json",
		"time": "2021-07-07T13:10:10+00:00",
		"releaseTime": "2014-06-25T13:59:27+00:00",
		"sha1": "4bd91bc2016638c36e5f75fbddc76f799248b524",
		"complianceLevel": 0
	  },
	  {
		"id": "14w25b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d5a45bde6dc849ab57b5a071db8c464cde175624/14w25b.json",
		"time": "2021-07-07T13:10:10+00:00",
		"releaseTime": "2014-06-19T12:29:48+00:00",
		"sha1": "d5a45bde6dc849ab57b5a071db8c464cde175624",
		"complianceLevel": 0
	  },
	  {
		"id": "14w25a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a640a289c75b3e8ce5c005a5d5b6b3cddc129add/14w25a.json",
		"time": "2021-07-07T13:10:10+00:00",
		"releaseTime": "2014-06-18T15:52:28+00:00",
		"sha1": "a640a289c75b3e8ce5c005a5d5b6b3cddc129add",
		"complianceLevel": 0
	  },
	  {
		"id": "14w21b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4587f874b8ff6cf22118cab83ecaac53c364dbbb/14w21b.json",
		"time": "2021-07-07T13:10:09+00:00",
		"releaseTime": "2014-05-22T15:17:55+00:00",
		"sha1": "4587f874b8ff6cf22118cab83ecaac53c364dbbb",
		"complianceLevel": 0
	  },
	  {
		"id": "14w21a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/edf0e7281e24fd174355902e92f1389d5a864c8d/14w21a.json",
		"time": "2021-07-07T13:10:09+00:00",
		"releaseTime": "2014-05-22T14:44:33+00:00",
		"sha1": "edf0e7281e24fd174355902e92f1389d5a864c8d",
		"complianceLevel": 0
	  },
	  {
		"id": "14w20b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/15fc0d317ed8e3f245db937df7d19dfa0602bc33/14w20b.json",
		"time": "2021-07-07T13:10:08+00:00",
		"releaseTime": "2014-05-15T16:47:21+00:00",
		"sha1": "15fc0d317ed8e3f245db937df7d19dfa0602bc33",
		"complianceLevel": 0
	  },
	  {
		"id": "14w20a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2cc1b54cf0bfd9aca8f08f04c6d972749ebd2a2a/14w20a.json",
		"time": "2021-07-07T13:10:08+00:00",
		"releaseTime": "2014-05-15T14:01:20+00:00",
		"sha1": "2cc1b54cf0bfd9aca8f08f04c6d972749ebd2a2a",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.10",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/4e4b8a264e14dd35a07ffd67dec548465fb94d81/1.7.10.json",
		"time": "2021-07-07T13:08:47+00:00",
		"releaseTime": "2014-05-14T17:29:23+00:00",
		"sha1": "4e4b8a264e14dd35a07ffd67dec548465fb94d81",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.10-pre4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2219e064bb727bc298de18e19985c030b77eebb0/1.7.10-pre4.json",
		"time": "2021-07-07T13:09:42+00:00",
		"releaseTime": "2014-05-14T16:29:23+00:00",
		"sha1": "2219e064bb727bc298de18e19985c030b77eebb0",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.10-pre3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e92519b17152bf674d8d8138b04c1948b1984d0c/1.7.10-pre3.json",
		"time": "2021-07-07T13:09:41+00:00",
		"releaseTime": "2014-05-14T15:29:23+00:00",
		"sha1": "e92519b17152bf674d8d8138b04c1948b1984d0c",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.10-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9292da53e7e326fb1a6628fef6e9bbca10279fbc/1.7.10-pre2.json",
		"time": "2021-07-07T13:09:41+00:00",
		"releaseTime": "2014-05-14T14:29:23+00:00",
		"sha1": "9292da53e7e326fb1a6628fef6e9bbca10279fbc",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.10-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/93ad94821119415428af9622c4593b56efc4dc38/1.7.10-pre1.json",
		"time": "2021-07-07T13:09:40+00:00",
		"releaseTime": "2014-05-14T13:29:23+00:00",
		"sha1": "93ad94821119415428af9622c4593b56efc4dc38",
		"complianceLevel": 0
	  },
	  {
		"id": "14w19a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/30b78edaec0fc1dd1ce6d111a8bd068e9f263fb8/14w19a.json",
		"time": "2021-07-07T13:10:08+00:00",
		"releaseTime": "2014-05-08T14:24:19+00:00",
		"sha1": "30b78edaec0fc1dd1ce6d111a8bd068e9f263fb8",
		"complianceLevel": 0
	  },
	  {
		"id": "14w18b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/040ab736abf398bfe8c515fd088d01dda2087939/14w18b.json",
		"time": "2021-07-07T13:10:07+00:00",
		"releaseTime": "2014-05-02T11:38:17+00:00",
		"sha1": "040ab736abf398bfe8c515fd088d01dda2087939",
		"complianceLevel": 0
	  },
	  {
		"id": "14w18a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/50ac0ae24be3249c505714d3a5921f7b4b73c7a3/14w18a.json",
		"time": "2021-07-07T13:10:07+00:00",
		"releaseTime": "2014-04-30T10:25:35+00:00",
		"sha1": "50ac0ae24be3249c505714d3a5921f7b4b73c7a3",
		"complianceLevel": 0
	  },
	  {
		"id": "14w17a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/50c980088ef84a5c27c6db93231f0fa7ea4e00e6/14w17a.json",
		"time": "2021-07-07T13:10:07+00:00",
		"releaseTime": "2014-04-24T15:44:49+00:00",
		"sha1": "50c980088ef84a5c27c6db93231f0fa7ea4e00e6",
		"complianceLevel": 0
	  },
	  {
		"id": "14w11b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e5a5ca85e88830dc8182a5991a1c61d1d7d9c6c0/14w11b.json",
		"time": "2021-07-07T13:10:06+00:00",
		"releaseTime": "2014-04-14T14:36:19+00:00",
		"sha1": "e5a5ca85e88830dc8182a5991a1c61d1d7d9c6c0",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.9",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/c886e6096a1befb7680d95654fd3f1f0cb12fff2/1.7.9.json",
		"time": "2021-07-07T13:08:50+00:00",
		"releaseTime": "2014-04-14T13:29:23+00:00",
		"sha1": "c886e6096a1befb7680d95654fd3f1f0cb12fff2",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.8",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/b633b9168e3732b7900efb87b534c9e0003d2462/1.7.8.json",
		"time": "2021-07-07T13:08:50+00:00",
		"releaseTime": "2014-04-09T07:58:16+00:00",
		"sha1": "b633b9168e3732b7900efb87b534c9e0003d2462",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.7",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/40460e82f921e3a68d4e80d1ea4f74fe9ce4b6b9/1.7.7.json",
		"time": "2021-07-07T13:08:50+00:00",
		"releaseTime": "2014-04-09T07:52:16+00:00",
		"sha1": "40460e82f921e3a68d4e80d1ea4f74fe9ce4b6b9",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.6",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/8156f8ea26310e8b53b4f395e35dd3bf06994b0a/1.7.6.json",
		"time": "2021-07-07T13:08:49+00:00",
		"releaseTime": "2014-04-09T07:52:06+00:00",
		"sha1": "8156f8ea26310e8b53b4f395e35dd3bf06994b0a",
		"complianceLevel": 0
	  },
	  {
		"id": "14w11a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e8aa5f7790b65a1e9e0ca3361a209c79e4c22bbc/14w11a.json",
		"time": "2021-07-07T13:10:06+00:00",
		"releaseTime": "2014-03-13T14:02:50+00:00",
		"sha1": "e8aa5f7790b65a1e9e0ca3361a209c79e4c22bbc",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.6-pre2",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/69c658d31cc01286b0694d37c10cfc566424a610/1.7.6-pre2.json",
		"time": "2021-07-07T13:09:42+00:00",
		"releaseTime": "2014-03-08T11:00:01+00:00",
		"sha1": "69c658d31cc01286b0694d37c10cfc566424a610",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.6-pre1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7c7b2c16cd38b79391c66fefb16d07421876003e/1.7.6-pre1.json",
		"time": "2021-07-07T13:09:42+00:00",
		"releaseTime": "2014-03-08T11:00:00+00:00",
		"sha1": "7c7b2c16cd38b79391c66fefb16d07421876003e",
		"complianceLevel": 0
	  },
	  {
		"id": "14w10c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d67767e03b8cc63ddf2a230a0667cd6668c23059/14w10c.json",
		"time": "2021-07-07T13:10:05+00:00",
		"releaseTime": "2014-03-07T13:49:55+00:00",
		"sha1": "d67767e03b8cc63ddf2a230a0667cd6668c23059",
		"complianceLevel": 0
	  },
	  {
		"id": "14w10b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ca975bfeea3fc0e74278ad0a6bf7634344940c89/14w10b.json",
		"time": "2021-07-07T13:10:05+00:00",
		"releaseTime": "2014-03-06T16:25:39+00:00",
		"sha1": "ca975bfeea3fc0e74278ad0a6bf7634344940c89",
		"complianceLevel": 0
	  },
	  {
		"id": "14w10a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/79c04521183b50308b630d543fb5490bb57f3d17/14w10a.json",
		"time": "2021-07-07T13:10:05+00:00",
		"releaseTime": "2014-03-06T14:23:04+00:00",
		"sha1": "79c04521183b50308b630d543fb5490bb57f3d17",
		"complianceLevel": 0
	  },
	  {
		"id": "14w08a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/829aec4bab1531a1b5176e46789c789411573b68/14w08a.json",
		"time": "2021-07-07T13:10:04+00:00",
		"releaseTime": "2014-02-26T17:00:00+00:00",
		"sha1": "829aec4bab1531a1b5176e46789c789411573b68",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.5",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/c66f7c06e791075a27664efe46a540c64245a46f/1.7.5.json",
		"time": "2021-07-07T13:08:49+00:00",
		"releaseTime": "2014-02-26T09:22:17+00:00",
		"sha1": "c66f7c06e791075a27664efe46a540c64245a46f",
		"complianceLevel": 0
	  },
	  {
		"id": "14w07a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0f07842faaa11ae2a7f3713aef30cc25f24e5f51/14w07a.json",
		"time": "2021-07-07T13:10:04+00:00",
		"releaseTime": "2014-02-14T11:05:07+00:00",
		"sha1": "0f07842faaa11ae2a7f3713aef30cc25f24e5f51",
		"complianceLevel": 0
	  },
	  {
		"id": "14w06b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/10daefe135cbcaffdd57ce5b858a5fe198ba133d/14w06b.json",
		"time": "2021-07-07T13:10:03+00:00",
		"releaseTime": "2014-02-06T17:30:42+00:00",
		"sha1": "10daefe135cbcaffdd57ce5b858a5fe198ba133d",
		"complianceLevel": 0
	  },
	  {
		"id": "14w06a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/adcac264b8fa163f70fbba49b9e36bc880bd1769/14w06a.json",
		"time": "2021-07-07T13:10:03+00:00",
		"releaseTime": "2014-02-06T14:30:17+00:00",
		"sha1": "adcac264b8fa163f70fbba49b9e36bc880bd1769",
		"complianceLevel": 0
	  },
	  {
		"id": "14w05b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4087f88bfec46db1dc49d6eecd8fd771e007b15f/14w05b.json",
		"time": "2021-07-07T13:10:02+00:00",
		"releaseTime": "2014-01-31T14:05:50+00:00",
		"sha1": "4087f88bfec46db1dc49d6eecd8fd771e007b15f",
		"complianceLevel": 0
	  },
	  {
		"id": "14w05a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2dbf1faf1aa7bcd0ef725957b86feb9c606c7867/14w05a.json",
		"time": "2021-07-07T13:10:02+00:00",
		"releaseTime": "2014-01-30T15:32:41+00:00",
		"sha1": "2dbf1faf1aa7bcd0ef725957b86feb9c606c7867",
		"complianceLevel": 0
	  },
	  {
		"id": "14w04b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ac0d995307e19635b6f2b2ca6b20ce28d94f1add/14w04b.json",
		"time": "2021-07-07T13:10:02+00:00",
		"releaseTime": "2014-01-24T15:48:46+00:00",
		"sha1": "ac0d995307e19635b6f2b2ca6b20ce28d94f1add",
		"complianceLevel": 0
	  },
	  {
		"id": "14w04a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0a14ebaacbf5974c70a4fbcf4eaefd47ae20b41e/14w04a.json",
		"time": "2021-07-07T13:10:01+00:00",
		"releaseTime": "2014-01-23T15:26:13+00:00",
		"sha1": "0a14ebaacbf5974c70a4fbcf4eaefd47ae20b41e",
		"complianceLevel": 0
	  },
	  {
		"id": "14w03b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/049235743e2df9880c9dd9d08784425a06b51f19/14w03b.json",
		"time": "2021-07-07T13:10:01+00:00",
		"releaseTime": "2014-01-16T16:36:19+00:00",
		"sha1": "049235743e2df9880c9dd9d08784425a06b51f19",
		"complianceLevel": 0
	  },
	  {
		"id": "14w03a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/d62cfce364aa87f508afd293ce3f374086b99828/14w03a.json",
		"time": "2021-07-07T13:10:00+00:00",
		"releaseTime": "2014-01-16T14:45:13+00:00",
		"sha1": "d62cfce364aa87f508afd293ce3f374086b99828",
		"complianceLevel": 0
	  },
	  {
		"id": "14w02c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/881525eefd49df390f825f5066a01b3010434f42/14w02c.json",
		"time": "2021-07-07T13:10:00+00:00",
		"releaseTime": "2014-01-10T15:42:36+00:00",
		"sha1": "881525eefd49df390f825f5066a01b3010434f42",
		"complianceLevel": 0
	  },
	  {
		"id": "14w02b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/85e1190bbbbee9459d32600b1e911c22bc00b6eb/14w02b.json",
		"time": "2021-07-07T13:10:00+00:00",
		"releaseTime": "2014-01-09T15:45:55+00:00",
		"sha1": "85e1190bbbbee9459d32600b1e911c22bc00b6eb",
		"complianceLevel": 0
	  },
	  {
		"id": "14w02a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4edeafe114d65c004cb70477923d37a2b9e8df4e/14w02a.json",
		"time": "2021-07-07T13:09:59+00:00",
		"releaseTime": "2014-01-09T14:44:41+00:00",
		"sha1": "4edeafe114d65c004cb70477923d37a2b9e8df4e",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.4",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/5cd47c02a53127b0c897c677efed176c1ac7e292/1.7.4.json",
		"time": "2021-07-07T13:08:48+00:00",
		"releaseTime": "2013-12-09T12:28:10+00:00",
		"sha1": "5cd47c02a53127b0c897c677efed176c1ac7e292",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.3",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/e6a8936b3236c8726e400d97ae3580e0fe6feb35/1.7.3.json",
		"time": "2021-07-07T13:08:48+00:00",
		"releaseTime": "2013-12-06T13:55:34+00:00",
		"sha1": "e6a8936b3236c8726e400d97ae3580e0fe6feb35",
		"complianceLevel": 0
	  },
	  {
		"id": "13w49a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7387830c28e2367703180f18b9dde508a7f7854f/13w49a.json",
		"time": "2021-07-07T13:09:59+00:00",
		"releaseTime": "2013-12-05T14:34:41+00:00",
		"sha1": "7387830c28e2367703180f18b9dde508a7f7854f",
		"complianceLevel": 0
	  },
	  {
		"id": "13w48b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f363a50b9e624b752d5e7563b1bdc9934d43a307/13w48b.json",
		"time": "2021-07-07T13:09:59+00:00",
		"releaseTime": "2013-11-26T18:36:08+00:00",
		"sha1": "f363a50b9e624b752d5e7563b1bdc9934d43a307",
		"complianceLevel": 0
	  },
	  {
		"id": "13w48a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b076cd1b383b6bbdaf546f367c3b1cd6f9989e23/13w48a.json",
		"time": "2021-07-07T13:09:58+00:00",
		"releaseTime": "2013-11-25T16:53:39+00:00",
		"sha1": "b076cd1b383b6bbdaf546f367c3b1cd6f9989e23",
		"complianceLevel": 0
	  },
	  {
		"id": "13w47e",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2ab160e4a983a2b15f332a0059f282e1072faae7/13w47e.json",
		"time": "2021-07-07T13:09:58+00:00",
		"releaseTime": "2013-11-22T15:16:38+00:00",
		"sha1": "2ab160e4a983a2b15f332a0059f282e1072faae7",
		"complianceLevel": 0
	  },
	  {
		"id": "13w47d",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ab6feecefeb0635fce4f3aed65f8d6cf6c4c8587/13w47d.json",
		"time": "2021-07-07T13:09:58+00:00",
		"releaseTime": "2013-11-22T13:51:15+00:00",
		"sha1": "ab6feecefeb0635fce4f3aed65f8d6cf6c4c8587",
		"complianceLevel": 0
	  },
	  {
		"id": "13w47c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2e01ed51340bc86e3188a81103587163a3f92a0d/13w47c.json",
		"time": "2021-07-07T13:09:57+00:00",
		"releaseTime": "2013-11-21T17:10:33+00:00",
		"sha1": "2e01ed51340bc86e3188a81103587163a3f92a0d",
		"complianceLevel": 0
	  },
	  {
		"id": "13w47b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/77e46f9404acbecba9722f007b8387185b86fd49/13w47b.json",
		"time": "2021-07-07T13:09:57+00:00",
		"releaseTime": "2013-11-21T16:57:41+00:00",
		"sha1": "77e46f9404acbecba9722f007b8387185b86fd49",
		"complianceLevel": 0
	  },
	  {
		"id": "13w47a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c43ec3622b72ba59322fc5b15d3baa2de37c1d2b/13w47a.json",
		"time": "2021-07-07T13:09:56+00:00",
		"releaseTime": "2013-11-21T15:59:58+00:00",
		"sha1": "c43ec3622b72ba59322fc5b15d3baa2de37c1d2b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/9a39eb6c9bcc7240fe1066527bf76bd5db37020a/1.7.2.json",
		"time": "2021-07-07T13:08:48+00:00",
		"releaseTime": "2013-10-25T13:00:00+00:00",
		"sha1": "9a39eb6c9bcc7240fe1066527bf76bd5db37020a",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7.1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7b07fc0656e54d3eddc796acfa9360a4af119fef/1.7.1.json",
		"time": "2021-07-07T13:09:40+00:00",
		"releaseTime": "2013-10-23T12:01:07+00:00",
		"sha1": "7b07fc0656e54d3eddc796acfa9360a4af119fef",
		"complianceLevel": 0
	  },
	  {
		"id": "1.7",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9a5a70f7117d8ff7d3706c0cd90f598220b656b3/1.7.json",
		"time": "2021-07-07T13:09:43+00:00",
		"releaseTime": "2013-10-22T15:04:05+00:00",
		"sha1": "9a5a70f7117d8ff7d3706c0cd90f598220b656b3",
		"complianceLevel": 0
	  },
	  {
		"id": "13w43a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/3d92a0d48abe6b24f70bda1a2c05a0fd6806a90c/13w43a.json",
		"time": "2021-07-07T13:09:56+00:00",
		"releaseTime": "2013-10-21T16:34:47+00:00",
		"sha1": "3d92a0d48abe6b24f70bda1a2c05a0fd6806a90c",
		"complianceLevel": 0
	  },
	  {
		"id": "13w42b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/1df97c4600ec63d977508c9f144710bd202a7a4d/13w42b.json",
		"time": "2021-07-07T13:09:56+00:00",
		"releaseTime": "2013-10-18T16:34:08+00:00",
		"sha1": "1df97c4600ec63d977508c9f144710bd202a7a4d",
		"complianceLevel": 0
	  },
	  {
		"id": "13w42a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/c9f75a0a8e50d04e51e10f017d95af64abd52152/13w42a.json",
		"time": "2021-07-07T13:09:55+00:00",
		"releaseTime": "2013-10-17T18:33:05+00:00",
		"sha1": "c9f75a0a8e50d04e51e10f017d95af64abd52152",
		"complianceLevel": 0
	  },
	  {
		"id": "13w41b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e06d50f70e2cc90bc7ad53f5031f36f41e264bd9/13w41b.json",
		"time": "2021-07-07T13:09:55+00:00",
		"releaseTime": "2013-10-11T15:09:17+00:00",
		"sha1": "e06d50f70e2cc90bc7ad53f5031f36f41e264bd9",
		"complianceLevel": 0
	  },
	  {
		"id": "13w41a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/47d2d9842b98dfbb03505143801651061c09652b/13w41a.json",
		"time": "2021-07-07T13:09:55+00:00",
		"releaseTime": "2013-10-10T14:21:43+00:00",
		"sha1": "47d2d9842b98dfbb03505143801651061c09652b",
		"complianceLevel": 0
	  },
	  {
		"id": "13w39b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f2bb7d91452ebcd98bc77b0e69d279b1146c02a4/13w39b.json",
		"time": "2019-06-28T07:08:09+00:00",
		"releaseTime": "2013-09-27T12:15:58+00:00",
		"sha1": "f2bb7d91452ebcd98bc77b0e69d279b1146c02a4",
		"complianceLevel": 0
	  },
	  {
		"id": "13w39a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2f4e00c7f6e3bc4a707c027ae2b527bcede4adda/13w39a.json",
		"time": "2019-06-28T07:08:09+00:00",
		"releaseTime": "2013-09-26T15:11:19+00:00",
		"sha1": "2f4e00c7f6e3bc4a707c027ae2b527bcede4adda",
		"complianceLevel": 0
	  },
	  {
		"id": "13w38c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/a9e87e0699f19fea280878f5deb744c5d5d3ccb1/13w38c.json",
		"time": "2019-06-28T07:08:09+00:00",
		"releaseTime": "2013-09-20T15:11:34+00:00",
		"sha1": "a9e87e0699f19fea280878f5deb744c5d5d3ccb1",
		"complianceLevel": 0
	  },
	  {
		"id": "13w38b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/6f426be1993b140ab5d10459c91eb1f542d58c82/13w38b.json",
		"time": "2019-06-28T07:08:09+00:00",
		"releaseTime": "2013-09-20T13:45:40+00:00",
		"sha1": "6f426be1993b140ab5d10459c91eb1f542d58c82",
		"complianceLevel": 0
	  },
	  {
		"id": "13w38a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e6dc1d9f9c8efeec67af438d5bf61be082f6e8a4/13w38a.json",
		"time": "2019-06-28T07:08:09+00:00",
		"releaseTime": "2013-09-19T16:34:21+00:00",
		"sha1": "e6dc1d9f9c8efeec67af438d5bf61be082f6e8a4",
		"complianceLevel": 0
	  },
	  {
		"id": "1.6.4",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/b71bae449192fbbe1582ff32fb3765edf0b9b0a8/1.6.4.json",
		"time": "2019-06-28T07:06:16+00:00",
		"releaseTime": "2013-09-19T15:52:37+00:00",
		"sha1": "b71bae449192fbbe1582ff32fb3765edf0b9b0a8",
		"complianceLevel": 0
	  },
	  {
		"id": "1.6.3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/903d6ba1bc87c301d88fa418f8b33446201c7d4e/1.6.3.json",
		"time": "2019-06-28T07:07:47+00:00",
		"releaseTime": "2013-09-13T10:54:41+00:00",
		"sha1": "903d6ba1bc87c301d88fa418f8b33446201c7d4e",
		"complianceLevel": 0
	  },
	  {
		"id": "13w37b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b8d28154ee056af6af3c8c37815418fe0e9f34f8/13w37b.json",
		"time": "2019-06-28T07:08:08+00:00",
		"releaseTime": "2013-09-13T10:54:41+00:00",
		"sha1": "b8d28154ee056af6af3c8c37815418fe0e9f34f8",
		"complianceLevel": 0
	  },
	  {
		"id": "13w37a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2f33c613a4bb81ef5f56be03a8f578208ada382a/13w37a.json",
		"time": "2019-06-28T07:08:08+00:00",
		"releaseTime": "2013-09-12T14:23:14+00:00",
		"sha1": "2f33c613a4bb81ef5f56be03a8f578208ada382a",
		"complianceLevel": 0
	  },
	  {
		"id": "13w36b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/4a538e23057a596fc8c7e04d8a7738d866467f51/13w36b.json",
		"time": "2019-06-28T07:08:08+00:00",
		"releaseTime": "2013-09-06T12:31:58+00:00",
		"sha1": "4a538e23057a596fc8c7e04d8a7738d866467f51",
		"complianceLevel": 0
	  },
	  {
		"id": "13w36a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/bc915c4dc167dfba92fcc0ae3aa051ae0f9f089b/13w36a.json",
		"time": "2019-06-28T07:08:08+00:00",
		"releaseTime": "2013-09-05T13:05:40+00:00",
		"sha1": "bc915c4dc167dfba92fcc0ae3aa051ae0f9f089b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.6.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/c0729761bf65dc58138ce508645dba1442fa78b8/1.6.2.json",
		"time": "2019-06-28T07:06:16+00:00",
		"releaseTime": "2013-07-05T13:09:02+00:00",
		"sha1": "c0729761bf65dc58138ce508645dba1442fa78b8",
		"complianceLevel": 0
	  },
	  {
		"id": "1.6.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/7fd8e0c76f62813eb0465e31bb74b160c01472d6/1.6.1.json",
		"time": "2019-06-28T07:06:16+00:00",
		"releaseTime": "2013-06-28T14:48:41+00:00",
		"sha1": "7fd8e0c76f62813eb0465e31bb74b160c01472d6",
		"complianceLevel": 0
	  },
	  {
		"id": "1.6",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/20116297638f7c70cd046e25a6ac90fee4cae61a/1.6.json",
		"time": "2019-06-28T07:07:47+00:00",
		"releaseTime": "2013-06-25T13:08:56+00:00",
		"sha1": "20116297638f7c70cd046e25a6ac90fee4cae61a",
		"complianceLevel": 0
	  },
	  {
		"id": "13w26a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b349702aef5e3adaebec30c79338300423943930/13w26a.json",
		"time": "2019-06-28T07:08:08+00:00",
		"releaseTime": "2013-06-24T16:06:06+00:00",
		"sha1": "b349702aef5e3adaebec30c79338300423943930",
		"complianceLevel": 0
	  },
	  {
		"id": "13w25c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/934788bc580ef0a19725ee5bd31f02a0b866e0bf/13w25c.json",
		"time": "2019-06-28T07:08:08+00:00",
		"releaseTime": "2013-06-20T15:23:37+00:00",
		"sha1": "934788bc580ef0a19725ee5bd31f02a0b866e0bf",
		"complianceLevel": 0
	  },
	  {
		"id": "13w25b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/8b7870ddd0d0b38779479ad782d65ad80e688cf7/13w25b.json",
		"time": "2019-06-28T07:08:08+00:00",
		"releaseTime": "2013-06-18T15:13:27+00:00",
		"sha1": "8b7870ddd0d0b38779479ad782d65ad80e688cf7",
		"complianceLevel": 0
	  },
	  {
		"id": "13w25a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/65c0e5fff89b477ac6f8ddb336f0e718d525d311/13w25a.json",
		"time": "2019-06-28T07:08:08+00:00",
		"releaseTime": "2013-06-17T14:08:06+00:00",
		"sha1": "65c0e5fff89b477ac6f8ddb336f0e718d525d311",
		"complianceLevel": 0
	  },
	  {
		"id": "13w24b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/e1294b52803771cfb06767c4c40dced70475cb25/13w24b.json",
		"time": "2019-06-28T07:08:08+00:00",
		"releaseTime": "2013-06-14T12:19:13+00:00",
		"sha1": "e1294b52803771cfb06767c4c40dced70475cb25",
		"complianceLevel": 0
	  },
	  {
		"id": "13w24a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/74666ab85cc5539f08aec638eabd63a552ed4125/13w24a.json",
		"time": "2019-06-28T07:08:07+00:00",
		"releaseTime": "2013-06-13T15:32:23+00:00",
		"sha1": "74666ab85cc5539f08aec638eabd63a552ed4125",
		"complianceLevel": 0
	  },
	  {
		"id": "13w23b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b19c87c7bfe0c0b202f79fc4b870b0ae97b00e53/13w23b.json",
		"time": "2019-06-28T07:08:07+00:00",
		"releaseTime": "2013-06-08T00:32:01+00:00",
		"sha1": "b19c87c7bfe0c0b202f79fc4b870b0ae97b00e53",
		"complianceLevel": 0
	  },
	  {
		"id": "13w23a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/400a136ff102882dfa9bb8990aef32b81309d46a/13w23a.json",
		"time": "2019-06-28T07:08:07+00:00",
		"releaseTime": "2013-06-07T16:04:20+00:00",
		"sha1": "400a136ff102882dfa9bb8990aef32b81309d46a",
		"complianceLevel": 0
	  },
	  {
		"id": "13w22a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f8277e55fae1555de47b44f5e6620f13b79fbe4e/13w22a.json",
		"time": "2019-06-28T07:08:07+00:00",
		"releaseTime": "2013-05-30T14:38:40+00:00",
		"sha1": "f8277e55fae1555de47b44f5e6620f13b79fbe4e",
		"complianceLevel": 0
	  },
	  {
		"id": "13w21b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/eab4bc12a78d862fadb36192c5351e35888eab15/13w21b.json",
		"time": "2019-06-28T07:08:07+00:00",
		"releaseTime": "2013-05-27T08:50:42+00:00",
		"sha1": "eab4bc12a78d862fadb36192c5351e35888eab15",
		"complianceLevel": 0
	  },
	  {
		"id": "13w21a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/ff59021171bd73aa155e40f84a924e1ab3f5307d/13w21a.json",
		"time": "2019-06-28T07:08:07+00:00",
		"releaseTime": "2013-05-23T15:38:28+00:00",
		"sha1": "ff59021171bd73aa155e40f84a924e1ab3f5307d",
		"complianceLevel": 0
	  },
	  {
		"id": "13w19a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/9360f33fa1391cbbfead0e0033c5a1e763f28d19/13w19a.json",
		"time": "2019-06-28T07:08:07+00:00",
		"releaseTime": "2013-05-10T14:48:02+00:00",
		"sha1": "9360f33fa1391cbbfead0e0033c5a1e763f28d19",
		"complianceLevel": 0
	  },
	  {
		"id": "13w18c",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/79ab628c68cefd8daa25d186f7961299ef4e63a9/13w18c.json",
		"time": "2019-06-28T07:08:07+00:00",
		"releaseTime": "2013-05-03T09:19:35+00:00",
		"sha1": "79ab628c68cefd8daa25d186f7961299ef4e63a9",
		"complianceLevel": 0
	  },
	  {
		"id": "13w18b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/2eea88342657ab8327af70ac2e1416859e0ef02a/13w18b.json",
		"time": "2019-06-28T07:08:07+00:00",
		"releaseTime": "2013-05-02T17:12:25+00:00",
		"sha1": "2eea88342657ab8327af70ac2e1416859e0ef02a",
		"complianceLevel": 0
	  },
	  {
		"id": "13w18a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/0fdbf2d4561027558ba9475c5d318540c743dec9/13w18a.json",
		"time": "2019-06-28T07:08:06+00:00",
		"releaseTime": "2013-05-02T15:45:59+00:00",
		"sha1": "0fdbf2d4561027558ba9475c5d318540c743dec9",
		"complianceLevel": 0
	  },
	  {
		"id": "13w17a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/f65b6bd3c813d67b026a3c9ec12e3280c495cf87/13w17a.json",
		"time": "2019-06-28T07:08:06+00:00",
		"releaseTime": "2013-04-25T15:50:00+00:00",
		"sha1": "f65b6bd3c813d67b026a3c9ec12e3280c495cf87",
		"complianceLevel": 0
	  },
	  {
		"id": "1.5.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/a611948faa091ef3a0af43d39f77589c5e402170/1.5.2.json",
		"time": "2019-06-28T07:06:16+00:00",
		"releaseTime": "2013-04-25T15:45:00+00:00",
		"sha1": "a611948faa091ef3a0af43d39f77589c5e402170",
		"complianceLevel": 0
	  },
	  {
		"id": "13w16b",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/b2c1241c6ca9d119a9f2df431cdf18a56f883d0d/13w16b.json",
		"time": "2019-06-28T07:08:06+00:00",
		"releaseTime": "2013-04-23T21:51:22+00:00",
		"sha1": "b2c1241c6ca9d119a9f2df431cdf18a56f883d0d",
		"complianceLevel": 0
	  },
	  {
		"id": "13w16a",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/7efbf7a86354b08f0a5011c96752c92cabb2e67c/13w16a.json",
		"time": "2019-06-28T07:08:06+00:00",
		"releaseTime": "2013-04-21T12:49:30+00:00",
		"sha1": "7efbf7a86354b08f0a5011c96752c92cabb2e67c",
		"complianceLevel": 0
	  },
	  {
		"id": "1.5.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/46b50f586db97821f22d2c94914c3a31f733a264/1.5.1.json",
		"time": "2019-06-28T07:06:16+00:00",
		"releaseTime": "2013-03-20T10:00:00+00:00",
		"sha1": "46b50f586db97821f22d2c94914c3a31f733a264",
		"complianceLevel": 0
	  },
	  {
		"id": "1.5",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/00fa1c125d5ad5ccdbc55f1d1b53b6f52c85bd22/1.5.json",
		"time": "2019-06-28T07:07:47+00:00",
		"releaseTime": "2013-03-06T22:00:00+00:00",
		"sha1": "00fa1c125d5ad5ccdbc55f1d1b53b6f52c85bd22",
		"complianceLevel": 0
	  },
	  {
		"id": "1.4.7",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/89802d57ccee3d03ec59d2ab1f44386234adb399/1.4.7.json",
		"time": "2019-06-28T07:06:15+00:00",
		"releaseTime": "2012-12-27T22:00:00+00:00",
		"sha1": "89802d57ccee3d03ec59d2ab1f44386234adb399",
		"complianceLevel": 0
	  },
	  {
		"id": "1.4.5",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/19555cce1f6d04e2b417ac2e9e06b6b752d5a2de/1.4.5.json",
		"time": "2019-06-28T07:06:15+00:00",
		"releaseTime": "2012-12-19T22:00:00+00:00",
		"sha1": "19555cce1f6d04e2b417ac2e9e06b6b752d5a2de",
		"complianceLevel": 0
	  },
	  {
		"id": "1.4.6",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/a7d02971582fcdf14ea275cc549cb57604a37079/1.4.6.json",
		"time": "2019-06-28T07:06:15+00:00",
		"releaseTime": "2012-12-19T22:00:00+00:00",
		"sha1": "a7d02971582fcdf14ea275cc549cb57604a37079",
		"complianceLevel": 0
	  },
	  {
		"id": "1.4.4",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/578d09d5753e590ea043a68bdaaea45a5b9711ab/1.4.4.json",
		"time": "2019-06-28T07:06:15+00:00",
		"releaseTime": "2012-12-13T22:00:00+00:00",
		"sha1": "578d09d5753e590ea043a68bdaaea45a5b9711ab",
		"complianceLevel": 0
	  },
	  {
		"id": "1.4.3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/180deba5c263367e914217be0701bd9e1a44f13b/1.4.3.json",
		"time": "2019-06-28T07:07:47+00:00",
		"releaseTime": "2012-11-30T22:00:00+00:00",
		"sha1": "180deba5c263367e914217be0701bd9e1a44f13b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.4.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/aec7405cf74ae5f79b13f6e8c88f66920eac0e1b/1.4.2.json",
		"time": "2019-06-28T07:06:15+00:00",
		"releaseTime": "2012-11-24T22:00:00+00:00",
		"sha1": "aec7405cf74ae5f79b13f6e8c88f66920eac0e1b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.4.1",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/859f584890b4db227f267026510c6ac2a5076d94/1.4.1.json",
		"time": "2019-06-28T07:07:47+00:00",
		"releaseTime": "2012-11-22T22:00:00+00:00",
		"sha1": "859f584890b4db227f267026510c6ac2a5076d94",
		"complianceLevel": 0
	  },
	  {
		"id": "1.4",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/070986801bc1d42eac872758cf12f00afa7b5f35/1.4.json",
		"time": "2019-06-28T07:07:47+00:00",
		"releaseTime": "2012-11-18T22:00:00+00:00",
		"sha1": "070986801bc1d42eac872758cf12f00afa7b5f35",
		"complianceLevel": 0
	  },
	  {
		"id": "1.3.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/2bd0ca9b77465a29df4b9449772d008f8724dd19/1.3.2.json",
		"time": "2019-06-28T07:06:15+00:00",
		"releaseTime": "2012-08-15T22:00:00+00:00",
		"sha1": "2bd0ca9b77465a29df4b9449772d008f8724dd19",
		"complianceLevel": 0
	  },
	  {
		"id": "1.3.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/be773d2fff5c6e4db9929ae4ea780f8837323b6b/1.3.1.json",
		"time": "2019-06-28T07:06:15+00:00",
		"releaseTime": "2012-07-31T22:00:00+00:00",
		"sha1": "be773d2fff5c6e4db9929ae4ea780f8837323b6b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.3",
		"type": "snapshot",
		"url": "https://launchermeta.mojang.com/v1/packages/aeb1bb40dc59420433ae46a0f133392508218bbe/1.3.json",
		"time": "2019-06-28T07:07:47+00:00",
		"releaseTime": "2012-07-25T22:00:00+00:00",
		"sha1": "aeb1bb40dc59420433ae46a0f133392508218bbe",
		"complianceLevel": 0
	  },
	  {
		"id": "1.2.5",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/886723c6385e62bb8dbe8026abf685140602404b/1.2.5.json",
		"time": "2019-06-28T07:06:15+00:00",
		"releaseTime": "2012-03-29T22:00:00+00:00",
		"sha1": "886723c6385e62bb8dbe8026abf685140602404b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.2.4",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/3cc81a28957e8c7d94bdd4f6b2c51ef8cae54f9b/1.2.4.json",
		"time": "2019-06-28T07:06:15+00:00",
		"releaseTime": "2012-03-21T22:00:00+00:00",
		"sha1": "3cc81a28957e8c7d94bdd4f6b2c51ef8cae54f9b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.2.3",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/4100a29330f76c93ef3bdc47d11fa016fd73de99/1.2.3.json",
		"time": "2019-06-28T07:06:14+00:00",
		"releaseTime": "2012-03-01T22:00:00+00:00",
		"sha1": "4100a29330f76c93ef3bdc47d11fa016fd73de99",
		"complianceLevel": 0
	  },
	  {
		"id": "1.2.2",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/897ce0e56d1cecc9e720f1934c09ea395008aa9e/1.2.2.json",
		"time": "2019-06-28T07:06:14+00:00",
		"releaseTime": "2012-02-29T22:00:01+00:00",
		"sha1": "897ce0e56d1cecc9e720f1934c09ea395008aa9e",
		"complianceLevel": 0
	  },
	  {
		"id": "1.2.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/dc5f9e134da9e18a2db42ddc246aba5cdfe28d3c/1.2.1.json",
		"time": "2019-06-28T07:06:14+00:00",
		"releaseTime": "2012-02-29T22:00:00+00:00",
		"sha1": "dc5f9e134da9e18a2db42ddc246aba5cdfe28d3c",
		"complianceLevel": 0
	  },
	  {
		"id": "1.1",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/1c1aaa3dea5155b549d4baf3491e3f5f564c3b8b/1.1.json",
		"time": "2019-06-28T07:05:45+00:00",
		"releaseTime": "2012-01-11T22:00:00+00:00",
		"sha1": "1c1aaa3dea5155b549d4baf3491e3f5f564c3b8b",
		"complianceLevel": 0
	  },
	  {
		"id": "1.0",
		"type": "release",
		"url": "https://launchermeta.mojang.com/v1/packages/edfc56a64dfc6430665d745264732db53b0d1b41/1.0.json",
		"time": "2019-06-28T07:05:45+00:00",
		"releaseTime": "2011-11-17T22:00:00+00:00",
		"sha1": "edfc56a64dfc6430665d745264732db53b0d1b41",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.8.1",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/5b5a652a385b3d04af8e05bb773696a227ebc300/b1.8.1.json",
		"time": "2019-06-28T07:05:45+00:00",
		"releaseTime": "2011-09-18T22:00:00+00:00",
		"sha1": "5b5a652a385b3d04af8e05bb773696a227ebc300",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.8",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/c2ce53a6d847e3c3efb640dc7d32e47e72e0a24f/b1.8.json",
		"time": "2019-06-28T07:05:45+00:00",
		"releaseTime": "2011-09-14T22:00:00+00:00",
		"sha1": "c2ce53a6d847e3c3efb640dc7d32e47e72e0a24f",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.7.3",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/0de35233bea254fd0011cceb4aa96e0a32b7efd2/b1.7.3.json",
		"time": "2019-06-28T07:05:45+00:00",
		"releaseTime": "2011-07-07T22:00:00+00:00",
		"sha1": "0de35233bea254fd0011cceb4aa96e0a32b7efd2",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.7.2",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/b9ac061f45d78c16739c95871d1540c681a187c8/b1.7.2.json",
		"time": "2019-06-28T07:05:45+00:00",
		"releaseTime": "2011-06-30T22:00:00+00:00",
		"sha1": "b9ac061f45d78c16739c95871d1540c681a187c8",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.7",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/3b7900fbef699471069015beef2a0d62cb2efabf/b1.7.json",
		"time": "2019-06-28T07:05:45+00:00",
		"releaseTime": "2011-06-29T22:00:00+00:00",
		"sha1": "3b7900fbef699471069015beef2a0d62cb2efabf",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.6.6",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/2c39299d7841e5273f5a6347373ba89eb48149d2/b1.6.6.json",
		"time": "2019-06-28T07:05:44+00:00",
		"releaseTime": "2011-05-30T22:00:00+00:00",
		"sha1": "2c39299d7841e5273f5a6347373ba89eb48149d2",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.6.5",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/f52059e5de7e06cbceba3642895005a36c40a6ef/b1.6.5.json",
		"time": "2019-06-28T07:05:44+00:00",
		"releaseTime": "2011-05-27T22:00:00+00:00",
		"sha1": "f52059e5de7e06cbceba3642895005a36c40a6ef",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.6.4",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/3a6546e2e4be5492c974359880f5cd6f1c513478/b1.6.4.json",
		"time": "2019-06-28T07:05:44+00:00",
		"releaseTime": "2011-05-25T22:00:04+00:00",
		"sha1": "3a6546e2e4be5492c974359880f5cd6f1c513478",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.6.3",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/4aa947cc55a28139e06ba6cde47b2c5aa3d9941c/b1.6.3.json",
		"time": "2019-06-28T07:05:44+00:00",
		"releaseTime": "2011-05-25T22:00:03+00:00",
		"sha1": "4aa947cc55a28139e06ba6cde47b2c5aa3d9941c",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.6.2",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/299360b97a57c73c0a6f258313200f40a37ab758/b1.6.2.json",
		"time": "2019-06-28T07:05:44+00:00",
		"releaseTime": "2011-05-25T22:00:02+00:00",
		"sha1": "299360b97a57c73c0a6f258313200f40a37ab758",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.6.1",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/72fc5a50e6a2fc9cdb5d77dcaa1f4cd4391536d3/b1.6.1.json",
		"time": "2019-06-28T07:05:44+00:00",
		"releaseTime": "2011-05-25T22:00:01+00:00",
		"sha1": "72fc5a50e6a2fc9cdb5d77dcaa1f4cd4391536d3",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.6",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/70c6be33812e5a030db15b28ccd72d30a5dcbca7/b1.6.json",
		"time": "2019-06-28T07:05:44+00:00",
		"releaseTime": "2011-05-25T22:00:00+00:00",
		"sha1": "70c6be33812e5a030db15b28ccd72d30a5dcbca7",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.5_01",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/cda4fc0e6f35c8645a013045be7c746e1a5b63d3/b1.5_01.json",
		"time": "2019-06-28T07:05:44+00:00",
		"releaseTime": "2011-04-19T22:00:00+00:00",
		"sha1": "cda4fc0e6f35c8645a013045be7c746e1a5b63d3",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.5",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/cd33026e81d260495f859a446a1f17e3071a17ea/b1.5.json",
		"time": "2019-06-28T07:05:44+00:00",
		"releaseTime": "2011-04-18T22:00:00+00:00",
		"sha1": "cd33026e81d260495f859a446a1f17e3071a17ea",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.4_01",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/7a7fbff959b9576845d83b215cc82ecf6bca5bcf/b1.4_01.json",
		"time": "2019-06-28T07:05:43+00:00",
		"releaseTime": "2011-04-04T22:00:00+00:00",
		"sha1": "7a7fbff959b9576845d83b215cc82ecf6bca5bcf",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.4",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/ce8cd5c165f2941d343cd98ec76a5b93d9a95c8c/b1.4.json",
		"time": "2019-06-28T07:05:43+00:00",
		"releaseTime": "2011-03-30T22:00:00+00:00",
		"sha1": "ce8cd5c165f2941d343cd98ec76a5b93d9a95c8c",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.3_01",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/f551611d75278e2847cca77421a51b99f4bb32de/b1.3_01.json",
		"time": "2019-06-28T07:05:43+00:00",
		"releaseTime": "2011-02-22T22:00:00+00:00",
		"sha1": "f551611d75278e2847cca77421a51b99f4bb32de",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.3b",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/a803c5a24b04ccab17269183b872cdaa38b32ae3/b1.3b.json",
		"time": "2019-06-28T07:05:43+00:00",
		"releaseTime": "2011-02-21T22:00:00+00:00",
		"sha1": "a803c5a24b04ccab17269183b872cdaa38b32ae3",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.2_02",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/afeabd62ef6bab19c8d570ab39a5e4161b27b85d/b1.2_02.json",
		"time": "2019-06-28T07:05:43+00:00",
		"releaseTime": "2011-01-20T22:00:00+00:00",
		"sha1": "afeabd62ef6bab19c8d570ab39a5e4161b27b85d",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.2_01",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/e9ea31d80bb8979bf3042165094d2e10b2899eaf/b1.2_01.json",
		"time": "2019-06-28T07:05:43+00:00",
		"releaseTime": "2011-01-13T22:00:00+00:00",
		"sha1": "e9ea31d80bb8979bf3042165094d2e10b2899eaf",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.2",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/703f0c0797d0b83dac8304b1c35ee4dcee6a2413/b1.2.json",
		"time": "2019-06-28T07:05:43+00:00",
		"releaseTime": "2011-01-12T22:00:00+00:00",
		"sha1": "703f0c0797d0b83dac8304b1c35ee4dcee6a2413",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.1_02",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/219e087b54dc2ca9a51a0204238cccaa035b5cef/b1.1_02.json",
		"time": "2019-06-28T07:05:43+00:00",
		"releaseTime": "2010-12-21T22:00:01+00:00",
		"sha1": "219e087b54dc2ca9a51a0204238cccaa035b5cef",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.1_01",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/d0d10f21a0f8115bdfbc2f0ddef2d93833af2e9d/b1.1_01.json",
		"time": "2019-06-28T07:05:43+00:00",
		"releaseTime": "2010-12-21T22:00:00+00:00",
		"sha1": "d0d10f21a0f8115bdfbc2f0ddef2d93833af2e9d",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.0.2",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/f8d933ed690495b66f76d0a5045e40c18881a26a/b1.0.2.json",
		"time": "2019-06-28T07:05:42+00:00",
		"releaseTime": "2010-12-20T22:00:00+00:00",
		"sha1": "f8d933ed690495b66f76d0a5045e40c18881a26a",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.0_01",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/c7559faa3f520837456b2a3736c663a856812049/b1.0_01.json",
		"time": "2019-06-28T07:05:42+00:00",
		"releaseTime": "2010-12-19T22:00:01+00:00",
		"sha1": "c7559faa3f520837456b2a3736c663a856812049",
		"complianceLevel": 0
	  },
	  {
		"id": "b1.0",
		"type": "old_beta",
		"url": "https://launchermeta.mojang.com/v1/packages/5be70efae7eea5f18ce164fb264eaa9563a054ee/b1.0.json",
		"time": "2019-06-28T07:05:42+00:00",
		"releaseTime": "2010-12-19T22:00:00+00:00",
		"sha1": "5be70efae7eea5f18ce164fb264eaa9563a054ee",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.6",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/d385e176aa7d3d3702bac78ad1ba906a77de13df/a1.2.6.json",
		"time": "2019-06-28T07:05:41+00:00",
		"releaseTime": "2010-12-02T22:00:00+00:00",
		"sha1": "d385e176aa7d3d3702bac78ad1ba906a77de13df",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.5",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/491a4961f00770bd130206c013795f35af948493/a1.2.5.json",
		"time": "2019-06-28T07:05:41+00:00",
		"releaseTime": "2010-11-30T22:00:00+00:00",
		"sha1": "491a4961f00770bd130206c013795f35af948493",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.4_01",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/e802a257031c5b9297c971599cc2573c2efece2c/a1.2.4_01.json",
		"time": "2019-06-28T07:05:41+00:00",
		"releaseTime": "2010-11-29T22:00:00+00:00",
		"sha1": "e802a257031c5b9297c971599cc2573c2efece2c",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.3_04",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/467403a159661d486fbe49480faf0909ea533759/a1.2.3_04.json",
		"time": "2019-06-28T07:05:41+00:00",
		"releaseTime": "2010-11-25T22:00:00+00:00",
		"sha1": "467403a159661d486fbe49480faf0909ea533759",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.3_02",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/31fa028661857f2e3d3732d07a6d36ec21d6dbdc/a1.2.3_02.json",
		"time": "2019-06-28T07:05:41+00:00",
		"releaseTime": "2010-11-24T22:00:00+00:00",
		"sha1": "31fa028661857f2e3d3732d07a6d36ec21d6dbdc",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.3_01",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/2dbccc4579a4481dc8d72a962d396de044648522/a1.2.3_01.json",
		"time": "2019-06-28T07:05:40+00:00",
		"releaseTime": "2010-11-23T22:00:01+00:00",
		"sha1": "2dbccc4579a4481dc8d72a962d396de044648522",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.3",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/48f077bf27e0a01a0bb2051e0ac17a96693cb730/a1.2.3.json",
		"time": "2019-06-28T07:05:40+00:00",
		"releaseTime": "2010-11-23T22:00:00+00:00",
		"sha1": "48f077bf27e0a01a0bb2051e0ac17a96693cb730",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.2b",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/0fbd340ae9087db32535ed0fb2d119240e7e0aaa/a1.2.2b.json",
		"time": "2019-06-28T07:05:40+00:00",
		"releaseTime": "2010-11-09T22:00:01+00:00",
		"sha1": "0fbd340ae9087db32535ed0fb2d119240e7e0aaa",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.2a",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/6679a974769ad2c6b88025ccbb60c72612dcee1f/a1.2.2a.json",
		"time": "2019-06-28T07:05:40+00:00",
		"releaseTime": "2010-11-09T22:00:00+00:00",
		"sha1": "6679a974769ad2c6b88025ccbb60c72612dcee1f",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.1_01",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/79e993b71f867e777a0ac9e2816bfd3df5c1aaed/a1.2.1_01.json",
		"time": "2019-06-28T07:05:40+00:00",
		"releaseTime": "2010-11-04T22:00:01+00:00",
		"sha1": "79e993b71f867e777a0ac9e2816bfd3df5c1aaed",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.1",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/f601478f5b0bf55ba0a302a4d1a8ce402c9311c1/a1.2.1.json",
		"time": "2019-06-28T07:05:40+00:00",
		"releaseTime": "2010-11-04T22:00:00+00:00",
		"sha1": "f601478f5b0bf55ba0a302a4d1a8ce402c9311c1",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.0_02",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/48465dfa2e2066e3b06e8fca934f2bbd7d03e65c/a1.2.0_02.json",
		"time": "2019-06-28T07:05:40+00:00",
		"releaseTime": "2010-11-03T22:00:00+00:00",
		"sha1": "48465dfa2e2066e3b06e8fca934f2bbd7d03e65c",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.0_01",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/e31556ae6f925d4dc7a06a5110cce731eca6eaa8/a1.2.0_01.json",
		"time": "2019-06-28T07:05:40+00:00",
		"releaseTime": "2010-10-30T22:00:00+00:00",
		"sha1": "e31556ae6f925d4dc7a06a5110cce731eca6eaa8",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.2.0",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/a0e7c34fba8c11fa8b84a56a01ed8505ad34b52a/a1.2.0.json",
		"time": "2019-06-28T07:05:39+00:00",
		"releaseTime": "2010-10-29T22:00:00+00:00",
		"sha1": "a0e7c34fba8c11fa8b84a56a01ed8505ad34b52a",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.1.2_01",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/c28a71888f39178cfc6f882fd62f49de920d138e/a1.1.2_01.json",
		"time": "2019-06-28T07:05:39+00:00",
		"releaseTime": "2010-09-22T22:00:00+00:00",
		"sha1": "c28a71888f39178cfc6f882fd62f49de920d138e",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.1.2",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/1231fefa206b10993e4b77b13f86fc72759db038/a1.1.2.json",
		"time": "2019-06-28T07:05:39+00:00",
		"releaseTime": "2010-09-19T22:00:00+00:00",
		"sha1": "1231fefa206b10993e4b77b13f86fc72759db038",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.1.0",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/a30685ddf58121f25dd845e5361a3d262409489d/a1.1.0.json",
		"time": "2019-06-28T07:05:39+00:00",
		"releaseTime": "2010-09-12T22:00:00+00:00",
		"sha1": "a30685ddf58121f25dd845e5361a3d262409489d",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.0.17_04",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/7c85d4575c9cc951135c3e6e958e1fcb2ed4e4ae/a1.0.17_04.json",
		"time": "2019-06-28T07:05:39+00:00",
		"releaseTime": "2010-08-22T22:00:00+00:00",
		"sha1": "7c85d4575c9cc951135c3e6e958e1fcb2ed4e4ae",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.0.17_02",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/54d1b6731a61a0b513ec0f677b58249be6c94086/a1.0.17_02.json",
		"time": "2019-06-28T07:05:39+00:00",
		"releaseTime": "2010-08-19T22:00:00+00:00",
		"sha1": "54d1b6731a61a0b513ec0f677b58249be6c94086",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.0.16",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/139fcc1a590a1166b872046bf4b82cfb04452061/a1.0.16.json",
		"time": "2019-06-28T07:05:38+00:00",
		"releaseTime": "2010-08-11T22:00:00+00:00",
		"sha1": "139fcc1a590a1166b872046bf4b82cfb04452061",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.0.15",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/786cd66bd91d3bd8ebd5d7fb10ec63daf671b646/a1.0.15.json",
		"time": "2019-06-28T07:05:38+00:00",
		"releaseTime": "2010-08-03T22:00:00+00:00",
		"sha1": "786cd66bd91d3bd8ebd5d7fb10ec63daf671b646",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.0.14",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/eff58a647f38df437241e769a9703baa40e0efd8/a1.0.14.json",
		"time": "2019-06-28T07:05:38+00:00",
		"releaseTime": "2010-07-29T22:00:00+00:00",
		"sha1": "eff58a647f38df437241e769a9703baa40e0efd8",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.0.11",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/3467dbde729f4e43047e54627ca20c82c9ee9a66/a1.0.11.json",
		"time": "2019-06-28T07:05:38+00:00",
		"releaseTime": "2010-07-22T22:00:00+00:00",
		"sha1": "3467dbde729f4e43047e54627ca20c82c9ee9a66",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.0.5_01",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/269ea3fdb671777a7bd0e97d4046cc27dc909694/a1.0.5_01.json",
		"time": "2019-06-28T07:05:39+00:00",
		"releaseTime": "2010-07-12T22:00:00+00:00",
		"sha1": "269ea3fdb671777a7bd0e97d4046cc27dc909694",
		"complianceLevel": 0
	  },
	  {
		"id": "a1.0.4",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/a1f71c0a68d59c0b6570073b440fb55ff889ba5a/a1.0.4.json",
		"time": "2019-06-28T07:05:39+00:00",
		"releaseTime": "2010-07-08T22:00:00+00:00",
		"sha1": "a1f71c0a68d59c0b6570073b440fb55ff889ba5a",
		"complianceLevel": 0
	  },
	  {
		"id": "inf-20100618",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/065ce5795aaf172080a4975cefac0d248bee7a3b/inf-20100618.json",
		"time": "2019-06-28T07:05:42+00:00",
		"releaseTime": "2010-06-15T22:00:00+00:00",
		"sha1": "065ce5795aaf172080a4975cefac0d248bee7a3b",
		"complianceLevel": 0
	  },
	  {
		"id": "c0.30_01c",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/0bb9bdebc3e124818fd31779a4bb394283050a02/c0.30_01c.json",
		"time": "2019-06-28T07:05:42+00:00",
		"releaseTime": "2009-12-21T22:00:00+00:00",
		"sha1": "0bb9bdebc3e124818fd31779a4bb394283050a02",
		"complianceLevel": 0
	  },
	  {
		"id": "c0.0.13a",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/5d77dd27f9684b4cbfbecc29a2c8c28d36a5e9c2/c0.0.13a.json",
		"time": "2019-06-28T07:05:41+00:00",
		"releaseTime": "2009-05-30T22:00:00+00:00",
		"sha1": "5d77dd27f9684b4cbfbecc29a2c8c28d36a5e9c2",
		"complianceLevel": 0
	  },
	  {
		"id": "c0.0.13a_03",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/3e038a3d4ce26771a8019c8c348a18844b950fdc/c0.0.13a_03.json",
		"time": "2019-06-28T07:05:41+00:00",
		"releaseTime": "2009-05-21T22:00:00+00:00",
		"sha1": "3e038a3d4ce26771a8019c8c348a18844b950fdc",
		"complianceLevel": 0
	  },
	  {
		"id": "c0.0.11a",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/d700384628b420acbe14388efc3b563e85ff7961/c0.0.11a.json",
		"time": "2019-06-28T07:05:41+00:00",
		"releaseTime": "2009-05-16T22:00:00+00:00",
		"sha1": "d700384628b420acbe14388efc3b563e85ff7961",
		"complianceLevel": 0
	  },
	  {
		"id": "rd-161348",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/a937d17cca60af0a7d45d04b49a849af16b08a28/rd-161348.json",
		"time": "2019-06-28T07:05:42+00:00",
		"releaseTime": "2009-05-16T11:48:00+00:00",
		"sha1": "a937d17cca60af0a7d45d04b49a849af16b08a28",
		"complianceLevel": 0
	  },
	  {
		"id": "rd-160052",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/c33dd04acfdbf34dcdfcca64db8545339ea24f02/rd-160052.json",
		"time": "2019-06-28T07:05:42+00:00",
		"releaseTime": "2009-05-15T22:52:00+00:00",
		"sha1": "c33dd04acfdbf34dcdfcca64db8545339ea24f02",
		"complianceLevel": 0
	  },
	  {
		"id": "rd-20090515",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/1bcd01f323df5c5092e9f0967b3d310d5bc0013a/rd-20090515.json",
		"time": "2019-06-28T07:05:42+00:00",
		"releaseTime": "2009-05-14T22:00:00+00:00",
		"sha1": "1bcd01f323df5c5092e9f0967b3d310d5bc0013a",
		"complianceLevel": 0
	  },
	  {
		"id": "rd-132328",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/77baa48d9cbbc6c3165c294e5bcdab2ca6903d57/rd-132328.json",
		"time": "2019-06-28T07:05:42+00:00",
		"releaseTime": "2009-05-13T21:28:00+00:00",
		"sha1": "77baa48d9cbbc6c3165c294e5bcdab2ca6903d57",
		"complianceLevel": 0
	  },
	  {
		"id": "rd-132211",
		"type": "old_alpha",
		"url": "https://launchermeta.mojang.com/v1/packages/0f2a46082313d0ec67972f9f63c3fa6591f9bb85/rd-132211.json",
		"time": "2019-06-28T07:05:42+00:00",
		"releaseTime": "2009-05-13T20:11:00+00:00",
		"sha1": "0f2a46082313d0ec67972f9f63c3fa6591f9bb85",
		"complianceLevel": 0
	  }
	]
  }`

func TestParseManifests(t *testing.T) {
	pd := NewManifestsProvider()
	pd.ParseManifests([]byte(p))
	r := pd.DownloadInfo
	t.Log(len(r))
	t.FailNow()
}