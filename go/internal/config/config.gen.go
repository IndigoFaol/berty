// file generated. see /config.
package config

import (
	"encoding/json"

	"berty.tech/berty/v2/go/pkg/protocoltypes"
)

var Config protocoltypes.Config

// FIXME: make it more nicely
func init() {
	input := `
{
  "berty": {
    "contacts": {
      "betabot": {
        "link": "https://berty.tech/id#contact/oZBLFkhRFWg9NXmy8wtjzZKeeUwfaWQVkUXfNY6h9z4QnGXuXSox8zs2yXeFHzt9DKJ8B4WHy7Zjy1g8tSat6ee2Jh7vp4J/name=BetaBot",
        "name": "BetaBot",
        "description": "Official BetaBot",
        "kind": "Bot"
      },
      "betabot-dev": {
        "link": "https://berty.tech/id#contact/oZBLFQQTqmZ25yvcRRUVgzzZCpDY7YAVXDEz3g8kmwcSPCmnQrqiMbLrfSp1GRLaNBX6ZRarid8jkSUX4yQYL4oBZcLDtFY/name=BetaBot-Dev",
        "name": "BetaBot Dev",
        "description": "BetaBot-Dev",
        "kind": "Bot"
      },
      "testbot": {
        "link": "https://berty.tech/id#contact/oZBLG8gg1RTwCW2u1AxA44dT1Lum2PAr2nBra6WdmTpozq7vdDkae9FjyCW3QZ86AE9pMAudh4GKTfNj1jELcdvtJy44Rbp/name=TestBot",
        "name": "TestBot",
        "description": "Official TestBot",
        "kind": "Bot"
      },
      "testbot-dev": {
        "link": "https://berty.tech/id#contact/oZBLFG6fwy21Dn4o511HC6ghxEghB5apq2QhjFepXCkbhfCWjzvRzhetVCGYwWhCtrWVL1TmExtHvdgGBvx1WK1RnXTdato/name=TestBot-Dev",
        "name": "TestBot Dev",
        "description": "TestBot-Dev",
        "kind": "Bot"
      }
    },
    "conversations": {
      "dev-test": {
        "link": "https://berty.tech/id#group/5QdUv6Fn3uvi3AchNcqFECaZvngDzWGUJV4wntYCWuYyjXrNi4ykvnP2ZCeWA1YLTmCSFbZaimXzp8rZtkKqby8nGv7S2AXyJPggo3aMghJS3rpeuDx6pbbNqEoNXJK67pWVpcLB6VrXM54zRNbi4zhEsLdATGzguWoruShcRHzFpuP75nyvRnXnfRaHPjdbyC/name=Dev+Test",
        "name": "Dev Test",
        "description": "Simple Conversation with Replication"
      },
      "dev-test-2": {
        "link": "https://berty.tech/id#group/5QdUv6Fn3uwSPQZpLUXcNpQ24bR9y1TUK6xJfE9khJo4jrP8jF4QfdgpTT73Us58y6XZhtvJHLDCouXfgsKDxXqZAsfLDhUZLb48PDcBYYVUk7nzxh1MKh5a6Ug1bP4KpdrVDYQx1gCZA4HZavEUgUBbC1pYBZ2DaY27sCMfqxpt79RjZCBwmVqW1DbbuYCUoc/name=Dev+Test+2",
        "name": "Dev Test 2",
        "description": "Simple Conversation with Replication"
      }
    }
  },
  "p2p": {
    "rdvp": [
      {
        "maddr": "/ip4/51.159.21.214/tcp/4040/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p"
      },
      {
        "maddr": "/ip4/51.159.21.214/udp/4040/quic/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p"
      },
      {
        "maddr": "/ip4/51.15.25.224/tcp/4040/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ"
      },
      {
        "maddr": "/ip4/51.15.25.224/udp/4040/quic/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ"
      },
      {
        "maddr": "/ip4/51.75.127.200/tcp/4141/p2p/12D3KooWPwRwwKatdy5yzRVCYPHib3fntYgbFB4nqrJPHWAqXD7z"
      },
      {
        "maddr": "/ip4/51.75.127.200/udp/4141/quic/p2p/12D3KooWPwRwwKatdy5yzRVCYPHib3fntYgbFB4nqrJPHWAqXD7z"
      }
    ],
    "relayHack": [
      "/ip4/51.159.21.214/udp/4040/quic/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p",
      "/ip4/51.15.25.224/udp/4040/quic/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ",
      "/ip4/51.75.127.200/udp/4141/quic/p2p/12D3KooWPwRwwKatdy5yzRVCYPHib3fntYgbFB4nqrJPHWAqXD7z"
    ]
  }
}`
	err := json.Unmarshal([]byte(input), &Config)
	if err != nil {
		panic(err)
	}
}
