# mockhook
Mockhook is a mock webhook for Facebook. 
It handles verification and prints incoming requests in prety JSON format.

## Get and run
```
go get github.com/nguyendangminh/mockhook
mockhook --help
mockhook -p 1203 -h webhook -t my_secret_token
```

## Sample of output
```
INFO[0000] The webhook is ready at port 1203/webhook    
INFO[0004] New message: 
{
  "entry": [
    {
      "id": "227709611044112",
      "messaging": [
        {
          "message": {
            "mid": "mid.$cAADPGZhswHhj_nEXBFdzSvOcIuDY",
            "nlp": {
              "entities": {
                "greetings": [
                  {
                    "confidence": 0.56252150963886,
                    "value": "true"
                  }
                ]
              }
            },
            "seq": 184630,
            "text": "fdss"
          },
          "recipient": {
            "id": "227709611044112"
          },
          "sender": {
            "id": "1260469204062102"
          },
          "timestamp": 1502385786628
        }
      ],
      "time": 1502385787047
    }
  ],
  "object": "page"
}
```
