package gosoap

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"time"
)

/*
************************

	WS-Security types

************************
*/
const (
	passwordType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest"
	encodingType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary"
)

// Security type :XMLName xml.Name `xml:"http://purl.org/rss/1.0/modules/content/ encoded"`
type Security struct {
	//XMLName xml.Name  `xml:"wsse:Security"`
	MustUnderstand string   `xml:"s:mustUnderstand,attr"`
	XMLName        xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security"`
	Auth           wsAuth
}

type password struct {
	//XMLName xml.Name `xml:"wsse:Password"`
	Type     string `xml:"Type,attr"`
	Password string `xml:",chardata"`
}

type nonce struct {
	//XMLName xml.Name `xml:"wsse:Nonce"`
	Type  string `xml:"EncodingType,attr"`
	Nonce string `xml:",chardata"`
}

type wsAuth struct {
	XMLName  xml.Name `xml:"UsernameToken"`
	Username string   `xml:"Username"`
	Password password `xml:"Password"`
	Nonce    nonce    `xml:"Nonce"`
	Created  string   `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd Created"`
}

/*
   <Security s:mustUnderstand="1" xmlns="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd">
       <UsernameToken>
           <Username>admin</Username>
           <Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest">edBuG+qVavQKLoWuGWQdPab4IBE=</Password>
           <Nonce EncodingType="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary">S7wO1ZFTh0KXv2CR7bd2ZXkLAAAAAA==</Nonce>
           <Created xmlns="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">2018-04-10T18:04:25.836Z</Created>
       </UsernameToken>
   </Security>
*/

// NewSecurity get a new security
func NewSecurity(username, passwd string, timeDiff time.Duration) Security {
	nonceSeq := generateNonce()                                     // Generate proper Base64-encoded Nonce
	created := time.Now().Add(-timeDiff).UTC().Format(time.RFC3339) // Ensure no milliseconds

	// wsPassword := ""
	// if passwd != "" {
	wsPassword := generateTokenBase64(nonceSeq, created, passwd)
	// }

	security := Security{
		MustUnderstand: "1",
		Auth: wsAuth{
			Username: username,
			Password: password{
				Type:     passwordType,
				Password: wsPassword,
			},
			Nonce: nonce{
				Type:  encodingType,
				Nonce: nonceSeq,
			},
			Created: created,
		},
	}

	return security
}

func generateNonce() string {
	nonceBytes := make([]byte, 16)                       // Generate 16 random bytes
	rand.Read(nonceBytes)                                // Fill with random binary data
	return base64.StdEncoding.EncodeToString(nonceBytes) // Encode as Base64
}

// Digest = B64ENCODE( SHA1( B64DECODE( Nonce ) + Date + Password ) )
func generateTokenBase64(Nonce string, Created string, Password string) string {
	sDec, _ := base64.StdEncoding.DecodeString(Nonce) // Decode Nonce properly

	hasher := sha1.New()
	hasher.Write(sDec)             // Use raw binary Nonce
	hasher.Write([]byte(Created))  // Use timestamp as raw bytes
	hasher.Write([]byte(Password)) // Use password as raw bytes

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil)) // Convert SHA1 hash to Base64
}
