package crypto

import "golang.org/x/crypto/acme/autocert"

//Manager is a stateful certificate manager built on top of acme.Client.
// It obtains and refreshes certificates automatically using "tls-alpn-01", "tls-sni-01", "tls-sni-02" and "http-01" challenge types,
// as well as providing them to a TLS server via tls.Config.
//You must specify a cache implementation, such as DirCache, to reuse obtained certificates across program restarts.
// Otherwise your server is very likely to exceed the certificate issuer's request rate limits.
// Config Variables:
// autocert.cache =  Is a cache optionally stores and retrieves previously-obtained certificates and other state
// autocert.whitlist = Is a stringslice of which domains the Manager will attempt to retrieve new certificates for. It does not affect cached certs.
func NewCertManager(dir string, whitelist ...string) *autocert.Manager {
	return &autocert.Manager{
		Prompt:          autocert.AcceptTOS,
		Cache:           autocert.DirCache(dir),
		HostPolicy:      autocert.HostWhitelist(whitelist...),
		RenewBefore:     0,
		Client:          nil,
		Email:           "",
		ExtraExtensions: nil,
	}
}