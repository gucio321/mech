# NBC

https://github.com/89z/mech/issues/83

Fail:

~~~
POST /access/vod/nbcuniversal/9000221348 HTTP/1.1
Host: access-cloudpath.media.nbcuni.com
Authorization: NBC-Security key=android_nbcuniversal,version=2.4,time=1655574989216,hash=bd00d39f14eb132453a85d3ecc3cfc7dd6abe8ce2e3962754f04192fe520cbfd
Content-Type: application/json

{
 "device": "android",
 "deviceId": "android",
 "externalAdvertiserId": "NBC",
 "mpx": {
  "accountId": 2410887629
 }
}
~~~