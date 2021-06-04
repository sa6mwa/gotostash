# decodeauthzmessages (package)

Decodes AWS Authorization Messages. You need something like the following policy statement for this to work:
```yaml
- Action:
  - "sts:DecodeAuthorizationMessage"
  Resource:
  - "*"
  Effect: Allow
```

# Usage

This is just an example...

```go
package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/sa6mwa/gotostash/pkg/decodeauthzmessages"
)

var decodeAuthZMessages bool = true

func main() {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		log.Fatal(err)
	}

	// if we are to decode authorization messages, add handler to session
	if decodeAuthZMessages {
		decodeauthzmessages.SetHandler(s)
	}

	// any AWS error you print containing an encoded authorization message will
	// try to be decoded
}
```
